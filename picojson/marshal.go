// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"bytes"
	"cmp"
	"encoding/base64"
	"math"
	"slices"
	"strconv"

	"storj.io/picobuf/internal/protowire"
	"storj.io/picobuf/picojson/internal/jsontext"
)

type marshaler struct {
	schema Schema
	out    []byte
	// irregular is set once some wire data needed grouped.
	irregular bool
}

// wireValue is one occurrence of a field in wire data.
type wireValue struct {
	typ protowire.Type
	num uint64 // varint and fixed values
	buf []byte // length-delimited values
}

// collect groups the wire data of msg by field, in the order of the fields.
// Unknown fields are dropped, as in protojson.
func collect(msg *message, wire []byte) ([][]wireValue, error) {
	fields := make([]field, msg.count)
	byNumber := make(map[int32]int, msg.count)
	it := msg.iter()
	for i := range fields {
		fields[i], _ = it.next()
		byNumber[fields[i].Number] = i
	}
	// Count the values first, so that they fit in one allocation.
	counts := make([]int, msg.count)
	total := 0
	for rest := wire; len(rest) > 0; {
		num, v, n := consumeValue(rest)
		if n < 0 {
			return nil, errSyntax
		}
		rest = rest[n:]
		if i, ok := byNumber[int32(num)]; ok && matchesWire(&fields[i], v) {
			counts[i]++
			total++
		}
	}
	all := make([]wireValue, total)
	values := make([][]wireValue, msg.count)
	for i, c := range counts {
		values[i], all = all[:0:c], all[c:]
	}
	var oneofs map[int32]int
	for rest := wire; len(rest) > 0; {
		num, v, n := consumeValue(rest)
		rest = rest[n:]
		if i, ok := byNumber[int32(num)]; ok && matchesWire(&fields[i], v) {
			if oneof := fields[i].Oneof; oneof != 0 {
				if previous, set := oneofs[oneof]; set && previous != i {
					// Switching members clears earlier values, including messages
					// that must not merge if selected again later.
					values[previous] = values[previous][:0]
				}
				if oneofs == nil {
					oneofs = make(map[int32]int)
				}
				oneofs[oneof] = i
			}
			values[i] = append(values[i], v)
		}
	}
	return values, nil
}

// consumeValue parses one field from wire and returns its length, or a
// negative length on error.
func consumeValue(wire []byte) (protowire.Number, wireValue, int) {
	num, typ, tagLen := protowire.ConsumeTag(wire)
	if tagLen < 0 {
		return 0, wireValue{}, tagLen
	}
	wire = wire[tagLen:]
	v := wireValue{typ: typ}
	var n int
	switch typ {
	case protowire.VarintType:
		v.num, n = protowire.ConsumeVarint(wire)
	case protowire.Fixed32Type:
		var x uint32
		x, n = protowire.ConsumeFixed32(wire)
		v.num = uint64(x)
	case protowire.Fixed64Type:
		v.num, n = protowire.ConsumeFixed64(wire)
	case protowire.BytesType:
		v.buf, n = protowire.ConsumeBytes(wire)
	default:
		n = protowire.ConsumeFieldValue(num, typ, wire)
	}
	if n < 0 {
		return 0, wireValue{}, n
	}
	return num, v, tagLen + n
}

func (m *marshaler) message(offset int32, wire []byte) error {
	msg := m.schema.message(offset)
	if handled, err := m.wellKnown(&msg, wire); handled {
		return err
	}
	if m.irregular {
		// Check first, so that a subtree is not rendered again at every
		// level that needs grouped.
		if ok, _ := m.walk(&msg, wire, false); !ok {
			return m.grouped(&msg, wire)
		}
		_, err := m.walk(&msg, wire, true)
		return err
	}
	start := len(m.out)
	ok, err := m.walk(&msg, wire, true)
	if ok {
		return nil
	}
	// With regular wire data, no later occurrence of a field can replace
	// the value that failed, so the error is final.
	if err != nil {
		if regular, _ := m.walk(&msg, wire, false); regular {
			return err
		}
	}
	m.irregular = true
	m.out = m.out[:start]
	return m.grouped(&msg, wire)
}

// walk handles msg in a single pass over the wire data. That works for
// regular wire data, as picobuf encodes it: fields in declaration order,
// each singular field and oneof at most once, the values of a repeated
// field together, and map entries sorted by key. Otherwise, such as with
// unrecognized fields that picobuf appends, it reports false, and grouped
// renders msg.
//
// With render, walk renders msg and stops at the first error. Without, it
// only reports whether the wire data is regular.
func (m *marshaler) walk(msg *message, wire []byte, render bool) (bool, error) {
	if render {
		m.out = append(m.out, '{')
	}
	it := msg.iter()
	var (
		f      field // the field of the previous value
		entry  mapEntry
		open   bool // whether the list or map of f is open
		first  = true
		oneofs uint64
	)
	for len(wire) > 0 {
		num, v, n := consumeValue(wire)
		if n < 0 {
			return false, errSyntax
		}
		wire = wire[n:]

		if open && int32(num) == f.Number {
			if !matchesWire(&f, v) || emptyPacked(&f, v) {
				continue
			}
			var err error
			if f.Map {
				var ok bool
				if ok, err = m.mapEntry(&entry, v, render); !ok {
					return false, err
				}
			} else if render {
				err = m.elements(&f, v, false)
			}
			if err != nil {
				return false, fieldError(msg, &f, err)
			}
			continue
		}

		probe := it
		next, ok := probe.find(int32(num))
		if !ok {
			if msg.indexByNumber(int32(num)) >= 0 {
				return false, nil // out of order
			}
			continue // unknown
		}
		if !matchesWire(&next, v) || emptyPacked(&next, v) {
			continue
		}
		it = probe
		if render && open {
			m.out = append(m.out, closing(&f))
		}
		f, open = next, false
		if !f.List && !f.Map && !f.Presence && f.Kind != messageKind && isDefault(f.Kind, v) {
			continue
		}
		if f.Oneof != 0 {
			// ponytail: messages with more than 64 oneofs use grouped.
			bit := uint64(1) << (f.Oneof - 1)
			if f.Oneof > 64 || oneofs&bit != 0 {
				return false, nil
			}
			oneofs |= bit
		}
		if render {
			if !first {
				m.out = append(m.out, ',')
			}
			first = false
			m.out, _ = jsontext.AppendQuote(m.out, msg.text(f.jsonName))
			m.out = append(m.out, ':')
		}

		var err error
		switch {
		case f.Map:
			open = true
			if render {
				m.out = append(m.out, '{')
			}
			entry = m.newMapEntry(&f)
			if ok, err = m.mapEntry(&entry, v, render); !ok {
				return false, err
			}
		case f.List:
			open = true
			if render {
				m.out = append(m.out, '[')
				err = m.elements(&f, v, true)
			}
		case render:
			err = m.value(&f, v)
		}
		if err != nil {
			return false, fieldError(msg, &f, err)
		}
	}
	if render {
		if open {
			m.out = append(m.out, closing(&f))
		}
		m.out = append(m.out, '}')
	}
	return true, nil
}

// closing returns the character that closes the list or map of f.
func closing(f *field) byte {
	if f.Map {
		return '}'
	}
	return ']'
}

// grouped renders msg after grouping the wire data by field.
func (m *marshaler) grouped(msg *message, wire []byte) error {
	values, err := collect(msg, wire)
	if err != nil {
		return err
	}
	m.out = append(m.out, '{')
	first := true
	it := msg.iter()
	for _, vs := range values {
		f, _ := it.next()
		vs = slices.DeleteFunc(vs, func(v wireValue) bool { return emptyPacked(&f, v) })
		if len(vs) == 0 {
			continue
		}
		var single wireValue
		if !f.List && !f.Map {
			single = lastValue(&f, vs)
			if !f.Presence && f.Kind != messageKind && isDefault(f.Kind, single) {
				continue
			}
		}
		if !first {
			m.out = append(m.out, ',')
		}
		first = false
		m.out, _ = jsontext.AppendQuote(m.out, msg.text(f.jsonName))
		m.out = append(m.out, ':')

		switch {
		case f.Map:
			err = m.mapValue(&f, vs)
		case f.List:
			m.out = append(m.out, '[')
			for i, v := range vs {
				if err = m.elements(&f, v, i == 0); err != nil {
					break
				}
			}
			m.out = append(m.out, ']')
		default:
			err = m.value(&f, single)
		}
		if err != nil {
			return fieldError(msg, &f, err)
		}
	}
	m.out = append(m.out, '}')
	return nil
}

// lastValue returns the value of a singular field: the last occurrence, or
// for messages all occurrences merged.
func lastValue(f *field, vs []wireValue) wireValue {
	if f.Kind != messageKind || len(vs) == 1 {
		return vs[len(vs)-1]
	}
	merged := wireValue{typ: protowire.BytesType}
	for _, v := range vs {
		merged.buf = append(merged.buf, v.buf...)
	}
	return merged
}

func isDefault(kind Kind, v wireValue) bool {
	switch kind {
	case stringKind, bytesKind, messageKind:
		return len(v.buf) == 0
	default:
		// -0.0 is not the default, it has the sign bit set.
		return v.num == 0
	}
}

// emptyPacked reports an occurrence that contributes no repeated elements.
func emptyPacked(f *field, v wireValue) bool {
	return f.List && packable(f.Kind) && v.typ == protowire.BytesType && len(v.buf) == 0
}

// elements appends the elements of a repeated field in v, which may be
// packed, preceded by a comma unless first.
func (m *marshaler) elements(f *field, v wireValue, first bool) error {
	if v.typ != protowire.BytesType || !packable(f.Kind) {
		if !first {
			m.out = append(m.out, ',')
		}
		return m.value(f, v)
	}
	typ := wireType(f.Kind)
	for packed := v.buf; len(packed) > 0; {
		x := wireValue{typ: typ}
		var n int
		switch typ {
		case protowire.VarintType:
			x.num, n = protowire.ConsumeVarint(packed)
		case protowire.Fixed32Type:
			var y uint32
			y, n = protowire.ConsumeFixed32(packed)
			x.num = uint64(y)
		default:
			x.num, n = protowire.ConsumeFixed64(packed)
		}
		if n < 0 {
			return errSyntax
		}
		packed = packed[n:]
		if !first {
			m.out = append(m.out, ',')
		}
		first = false
		if err := m.value(f, x); err != nil {
			return err
		}
	}
	return nil
}

// mapEntry is the state of rendering a map field.
type mapEntry struct {
	key, value field
	// last is the previous key, valid when count > 0.
	last  wireValue
	count int
}

func (m *marshaler) newMapEntry(f *field) mapEntry {
	entry := m.schema.message(f.Type)
	key, _ := entry.fieldByNumber(1)
	value, _ := entry.fieldByNumber(2)
	return mapEntry{key: key, value: value}
}

// mapEntry checks that the entry in v sorts after the previous one, which
// picobuf encodes them in, so that keys do not repeat; otherwise, it
// reports false and mapValue renders the map. With render, it also appends
// the entry.
func (m *marshaler) mapEntry(e *mapEntry, v wireValue, render bool) (bool, error) {
	key, value, err := decodeEntry(e, v)
	if err != nil {
		return false, err
	}
	if e.count > 0 && compareKeys(e.key.Kind, e.last, key) >= 0 {
		return false, nil
	}
	e.last = key
	e.count++
	if !render {
		return true, nil
	}
	if e.count > 1 {
		m.out = append(m.out, ',')
	}
	if m.out, err = appendKey(m.out, e.key.Kind, key); err != nil {
		return true, err
	}
	m.out = append(m.out, ':')
	return true, m.value(&e.value, value)
}

// compareKeys compares map keys the way picobuf sorts them, after the same
// conversions as appendKey, so that keys that render the same compare equal.
func compareKeys(kind Kind, a, b wireValue) int {
	switch kind {
	case stringKind:
		return bytes.Compare(a.buf, b.buf)
	case boolKind:
		return cmp.Compare(boolInt(a.num != 0), boolInt(b.num != 0))
	case int32Kind, sfixed32Kind:
		return cmp.Compare(int32(a.num), int32(b.num))
	case sint32Kind:
		return cmp.Compare(protowire.DecodeZigZag(a.num&math.MaxUint32), protowire.DecodeZigZag(b.num&math.MaxUint32))
	case int64Kind, sfixed64Kind:
		return cmp.Compare(int64(a.num), int64(b.num))
	case sint64Kind:
		return cmp.Compare(protowire.DecodeZigZag(a.num), protowire.DecodeZigZag(b.num))
	case uint32Kind, fixed32Kind:
		return cmp.Compare(uint32(a.num), uint32(b.num))
	default:
		return cmp.Compare(a.num, b.num)
	}
}

func boolInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// decodeEntry decodes the key and value of a map entry.
func decodeEntry(e *mapEntry, v wireValue) (key, value wireValue, err error) {
	if v.typ != protowire.BytesType {
		return key, value, errSyntax
	}
	key, value = wireValue{typ: wireType(e.key.Kind)}, wireValue{typ: wireType(e.value.Kind)}
	hasValue := false
	for rest := v.buf; len(rest) > 0; {
		num, x, n := consumeValue(rest)
		if n < 0 {
			return key, value, errSyntax
		}
		rest = rest[n:]
		switch {
		case num == 1 && matchesWire(&e.key, x):
			key = x
		case num != 2 || !matchesWire(&e.value, x):
		case hasValue && e.value.Kind == messageKind:
			// Repeated message values merge, see lastValue.
			value.buf = append(slices.Clip(value.buf), x.buf...)
		default:
			value, hasValue = x, true
		}
	}
	return key, value, nil
}

// mapValue renders a map field whose entries may repeat keys, where a later
// entry replaces the value of an earlier one.
func (m *marshaler) mapValue(f *field, vs []wireValue) error {
	e := m.newMapEntry(f)
	type pair struct {
		key   string
		value wireValue
	}
	var pairs []pair
	index := map[string]int{}
	for _, v := range vs {
		key, value, err := decodeEntry(&e, v)
		if err != nil {
			return err
		}
		b, err := appendKey(nil, e.key.Kind, key)
		if err != nil {
			return err
		}
		text := string(b)
		if i, ok := index[text]; ok {
			pairs[i].value = value
			continue
		}
		index[text] = len(pairs)
		pairs = append(pairs, pair{text, value})
	}

	m.out = append(m.out, '{')
	for i, p := range pairs {
		if i > 0 {
			m.out = append(m.out, ',')
		}
		m.out = append(m.out, p.key...)
		m.out = append(m.out, ':')
		if err := m.value(&e.value, p.value); err != nil {
			return err
		}
	}
	m.out = append(m.out, '}')
	return nil
}

// appendKey appends a map key as a JSON string.
func appendKey(out []byte, kind Kind, v wireValue) ([]byte, error) {
	switch kind {
	case stringKind:
		return jsontext.AppendQuote(out, v.buf)
	case int64Kind, sint64Kind, sfixed64Kind, uint64Kind, fixed64Kind:
		// 64-bit values are already quoted.
		return appendScalar(out, kind, v.num), nil
	default:
		out = append(out, '"')
		return append(appendScalar(out, kind, v.num), '"'), nil
	}
}

func (m *marshaler) value(f *field, v wireValue) error {
	if v.typ != wireType(f.Kind) {
		return errSyntax
	}
	var err error
	switch f.Kind {
	case stringKind:
		m.out, err = jsontext.AppendQuote(m.out, v.buf)
	case bytesKind:
		m.out = append(m.out, '"')
		m.out = base64.StdEncoding.AppendEncode(m.out, v.buf)
		m.out = append(m.out, '"')
	case messageKind:
		err = m.message(f.Type, v.buf)
	case enumKind:
		number := int32(v.num)
		if name, ok := m.schema.enumName(f.Type, number); ok {
			m.out, err = jsontext.AppendQuote(m.out, name)
		} else {
			m.out = strconv.AppendInt(m.out, int64(number), 10)
		}
	default:
		m.out = appendScalar(m.out, f.Kind, v.num)
	}
	return err
}

// appendScalar formats a numeric or bool value. 64-bit integers are quoted,
// as ProtoJSON requires.
func appendScalar(out []byte, kind Kind, x uint64) []byte {
	switch kind {
	case boolKind:
		return strconv.AppendBool(out, x != 0)
	case int32Kind:
		return strconv.AppendInt(out, int64(int32(x)), 10)
	case sint32Kind:
		return strconv.AppendInt(out, int64(int32(protowire.DecodeZigZag(x&math.MaxUint32))), 10)
	case sfixed32Kind:
		return strconv.AppendInt(out, int64(int32(uint32(x))), 10)
	case uint32Kind, fixed32Kind:
		return strconv.AppendUint(out, uint64(uint32(x)), 10)
	case int64Kind, sfixed64Kind:
		out = append(out, '"')
		return append(strconv.AppendInt(out, int64(x), 10), '"')
	case sint64Kind:
		out = append(out, '"')
		return append(strconv.AppendInt(out, protowire.DecodeZigZag(x), 10), '"')
	case uint64Kind, fixed64Kind:
		out = append(out, '"')
		return append(strconv.AppendUint(out, x, 10), '"')
	case floatKind:
		return appendFloat(out, float64(math.Float32frombits(uint32(x))), 32)
	case doubleKind:
		return appendFloat(out, math.Float64frombits(x), 64)
	}
	panic("picojson: unhandled kind")
}

// appendFloat formats a float the way protojson does.
func appendFloat(out []byte, n float64, bitSize int) []byte {
	switch {
	case math.IsNaN(n):
		return append(out, `"NaN"`...)
	case math.IsInf(n, +1):
		return append(out, `"Infinity"`...)
	case math.IsInf(n, -1):
		return append(out, `"-Infinity"`...)
	}
	format := byte('f')
	if abs := math.Abs(n); abs != 0 {
		if bitSize == 64 && (abs < 1e-6 || abs >= 1e21) ||
			bitSize == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			format = 'e'
		}
	}
	out = strconv.AppendFloat(out, n, format, -1, bitSize)
	if format == 'e' {
		// Clean up e-09 to e-9.
		if n := len(out); n >= 4 && out[n-4] == 'e' && out[n-3] == '-' && out[n-2] == '0' {
			out[n-2] = out[n-1]
			out = out[:n-1]
		}
	}
	return out
}

// matchesWire treats an occurrence with the wrong wire type as unknown.
// Repeated numeric fields accept both packed and unpacked encodings.
func matchesWire(f *field, v wireValue) bool {
	return v.typ == wireType(f.Kind) || f.List && packable(f.Kind) && v.typ == protowire.BytesType
}

func wireType(kind Kind) protowire.Type {
	switch kind {
	case floatKind, fixed32Kind, sfixed32Kind:
		return protowire.Fixed32Type
	case doubleKind, fixed64Kind, sfixed64Kind:
		return protowire.Fixed64Type
	case stringKind, bytesKind, messageKind:
		return protowire.BytesType
	default:
		return protowire.VarintType
	}
}

func packable(kind Kind) bool {
	return wireType(kind) != protowire.BytesType
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"storj.io/picobuf/internal/protowire"
	"storj.io/picobuf/picojson/internal/jsontext"
)

type unmarshaler struct {
	schema Schema
	dec    *jsontext.Decoder
}

// unmarshalWire parses data as the message at offset in s into wire data.
func unmarshalWire(data []byte, s Schema, offset int32) ([]byte, error) {
	u := unmarshaler{schema: s, dec: jsontext.NewDecoder(bytes.NewReader(data))}
	out, err := u.message(offset, nil)
	if err != nil {
		return nil, err
	}
	if _, err := u.dec.ReadToken(); !errors.Is(err, io.EOF) {
		return nil, errors.New("unexpected data after the message")
	}
	return out, nil
}

func (u *unmarshaler) read(kind jsontext.Kind) (jsontext.Token, error) {
	tok, err := u.dec.ReadToken()
	if err != nil {
		return tok, err
	}
	if tok.Kind() != kind {
		return tok, fmt.Errorf("unexpected %v, want %v", tok.Kind(), kind)
	}
	return tok, nil
}

// message appends the wire data of the message value that the decoder is at.
func (u *unmarshaler) message(offset int32, out []byte) ([]byte, error) {
	msg := u.schema.message(offset)
	if handled, out, err := u.wellKnown(&msg, out); handled {
		return out, err
	}
	if _, err := u.read('{'); err != nil {
		return nil, err
	}
	seen := make([]bool, msg.count)
	var oneofs map[int32]bool
	fields := nameLookup{msg: &msg, it: msg.iter()}
	for u.dec.PeekKind() != '}' {
		tok, err := u.read('"')
		if err != nil {
			return nil, err
		}
		i, f, ok := fields.find(tok.String())
		if !ok {
			return nil, fmt.Errorf("%s: unknown field %q", msg.Name, tok.String())
		}
		if seen[i] {
			return nil, fieldError(&msg, &f, errors.New("field set more than once"))
		}
		seen[i] = true
		// null leaves the field unset.
		if u.dec.PeekKind() == 'n' {
			if _, err := u.dec.ReadToken(); err != nil {
				return nil, err
			}
			continue
		}
		if f.Oneof != 0 {
			if oneofs[f.Oneof] {
				return nil, fieldError(&msg, &f, errors.New("another field of its oneof is already set"))
			}
			if oneofs == nil {
				oneofs = map[int32]bool{}
			}
			oneofs[f.Oneof] = true
		}
		if out, err = u.field(&f, out); err != nil {
			return nil, fieldError(&msg, &f, err)
		}
	}
	if _, err := u.dec.ReadToken(); err != nil {
		return nil, err
	}
	return out, nil
}

func (u *unmarshaler) field(f *field, out []byte) ([]byte, error) {
	var err error
	switch {
	case f.Map:
		entry := u.schema.message(f.Type)
		keyField, _ := entry.fieldByNumber(1)
		valueField, _ := entry.fieldByNumber(2)
		if _, err := u.read('{'); err != nil {
			return nil, err
		}
		var keys map[string]struct{}
		for u.dec.PeekKind() != '}' {
			tok, err := u.read('"')
			if err != nil {
				return nil, err
			}
			out = protowire.AppendTag(out, protowire.Number(f.Number), protowire.BytesType)
			start := len(out)
			if out, err = appendMapKey(&keyField, tok.String(), out); err != nil {
				return nil, err
			}
			// The decoder checks JSON name uniqueness. Numeric keys also
			// need comparison after decoding (for example, "1" and "01").
			if keyField.Kind != stringKind && keyField.Kind != boolKind {
				key := string(out[start:])
				if _, exists := keys[key]; exists {
					return nil, fmt.Errorf("duplicate map key %q", tok.String())
				}
				if keys == nil {
					keys = make(map[string]struct{})
				}
				keys[key] = struct{}{}
			}
			if out, err = u.value(&valueField, out); err != nil {
				return nil, err
			}
			out = insertLength(out, start)
		}
		_, err = u.dec.ReadToken()
	case f.List:
		if _, err := u.read('['); err != nil {
			return nil, err
		}
		for u.dec.PeekKind() != ']' {
			if out, err = u.value(f, out); err != nil {
				return nil, err
			}
		}
		_, err = u.dec.ReadToken()
	default:
		out, err = u.value(f, out)
	}
	return out, err
}

// value appends one value of f, with its tag.
func (u *unmarshaler) value(f *field, out []byte) ([]byte, error) {
	if f.Kind == messageKind {
		out = protowire.AppendTag(out, protowire.Number(f.Number), protowire.BytesType)
		start := len(out)
		out, err := u.message(f.Type, out)
		if err != nil {
			return nil, err
		}
		return insertLength(out, start), nil
	}

	tok, err := u.dec.ReadToken()
	if err != nil {
		return nil, err
	}
	switch f.Kind {
	case boolKind:
		if k := tok.Kind(); k != 't' && k != 'f' {
			return nil, fmt.Errorf("unexpected %v, want a bool", k)
		}
		out = protowire.AppendTag(out, protowire.Number(f.Number), protowire.VarintType)
		return protowire.AppendVarint(out, protowire.EncodeBool(tok.Bool())), nil
	case enumKind:
		switch tok.Kind() {
		case '"':
			number, ok := u.schema.enumNumber(f.Type, tok.String())
			if !ok {
				return nil, fmt.Errorf("unknown enum value %q", tok.String())
			}
			out = protowire.AppendTag(out, protowire.Number(f.Number), protowire.VarintType)
			return protowire.AppendVarint(out, uint64(int64(number))), nil
		case '0':
			return appendText(&field{Number: f.Number, Kind: int32Kind}, tok.String(), out)
		}
		return nil, fmt.Errorf("unexpected %v, want an enum", tok.Kind())
	case stringKind, bytesKind:
		if tok.Kind() != '"' {
			return nil, fmt.Errorf("unexpected %v, want a string", tok.Kind())
		}
	default:
		switch tok.Kind() {
		case '0':
		case '"':
			// Like protojson, a quoted number must be a JSON number.
			text := tok.String()
			special := (f.Kind == floatKind || f.Kind == doubleKind) &&
				(text == "NaN" || text == "Infinity" || text == "-Infinity")
			if !special && !isJSONNumber(text) {
				return nil, fmt.Errorf("invalid number %q", text)
			}
		default:
			return nil, fmt.Errorf("unexpected %v, want a number", tok.Kind())
		}
	}
	return appendText(f, tok.String(), out)
}

// appendText appends a scalar value of f from the text of a JSON string or
// number, with its tag.
func appendText(f *field, text string, out []byte) ([]byte, error) {
	num := protowire.Number(f.Number)
	switch f.Kind {
	case stringKind:
		out = protowire.AppendTag(out, num, protowire.BytesType)
		return protowire.AppendString(out, text), nil
	case bytesKind:
		b, err := decodeBase64(text)
		if err != nil {
			return nil, err
		}
		out = protowire.AppendTag(out, num, protowire.BytesType)
		return protowire.AppendBytes(out, b), nil
	case boolKind:
		b, err := strconv.ParseBool(text)
		if err != nil || (text != "true" && text != "false") {
			return nil, fmt.Errorf("invalid bool %q", text)
		}
		out = protowire.AppendTag(out, num, protowire.VarintType)
		return protowire.AppendVarint(out, protowire.EncodeBool(b)), nil
	case floatKind, doubleKind:
		bits := 64
		if f.Kind == floatKind {
			bits = 32
		}
		x, err := parseFloat(text, bits)
		if err != nil {
			return nil, err
		}
		if f.Kind == floatKind {
			out = protowire.AppendTag(out, num, protowire.Fixed32Type)
			return protowire.AppendFixed32(out, math.Float32bits(float32(x))), nil
		}
		out = protowire.AppendTag(out, num, protowire.Fixed64Type)
		return protowire.AppendFixed64(out, math.Float64bits(x)), nil
	}

	var x uint64
	switch f.Kind {
	case int32Kind, sint32Kind, sfixed32Kind:
		v, err := parseInt(text, 32)
		if err != nil {
			return nil, err
		}
		if f.Kind == sint32Kind {
			x = protowire.EncodeZigZag(v)
		} else {
			x = uint64(v)
		}
	case int64Kind, sint64Kind, sfixed64Kind:
		v, err := parseInt(text, 64)
		if err != nil {
			return nil, err
		}
		if f.Kind == sint64Kind {
			x = protowire.EncodeZigZag(v)
		} else {
			x = uint64(v)
		}
	case uint32Kind, fixed32Kind, uint64Kind, fixed64Kind:
		bits := 64
		if f.Kind == uint32Kind || f.Kind == fixed32Kind {
			bits = 32
		}
		v, err := parseUint(text, bits)
		if err != nil {
			return nil, err
		}
		x = v
	default:
		return nil, fmt.Errorf("unhandled kind %d", f.Kind)
	}

	return appendInteger(f, x, out), nil
}

// appendMapKey uses decimal integer syntax for numeric keys, rather than
// the fractional and exponent forms allowed for scalar values.
func appendMapKey(f *field, text string, out []byte) ([]byte, error) {
	if f.Kind == stringKind || f.Kind == boolKind {
		return appendText(f, text, out)
	}
	bits := 64
	switch f.Kind {
	case int32Kind, sint32Kind, sfixed32Kind, uint32Kind, fixed32Kind:
		bits = 32
	}
	var x uint64
	switch f.Kind {
	case uint32Kind, fixed32Kind, uint64Kind, fixed64Kind:
		v, err := strconv.ParseUint(text, 10, bits)
		if err != nil {
			return nil, fmt.Errorf("invalid map key %q", text)
		}
		x = v
	default:
		v, err := strconv.ParseInt(text, 10, bits)
		if err != nil {
			return nil, fmt.Errorf("invalid map key %q", text)
		}
		x = uint64(v)
		if f.Kind == sint32Kind || f.Kind == sint64Kind {
			x = protowire.EncodeZigZag(v)
		}
	}
	return appendInteger(f, x, out), nil
}

func appendInteger(f *field, x uint64, out []byte) []byte {
	num := protowire.Number(f.Number)
	switch wireType(f.Kind) {
	case protowire.Fixed32Type:
		out = protowire.AppendTag(out, num, protowire.Fixed32Type)
		return protowire.AppendFixed32(out, uint32(x))
	case protowire.Fixed64Type:
		out = protowire.AppendTag(out, num, protowire.Fixed64Type)
		return protowire.AppendFixed64(out, x)
	}
	out = protowire.AppendTag(out, num, protowire.VarintType)
	return protowire.AppendVarint(out, x)
}

// insertLength inserts the length prefix of the value that out holds from
// start, moving the value in place.
func insertLength(out []byte, start int) []byte {
	var prefix [binary.MaxVarintLen64]byte
	return slices.Insert(out, start, protowire.AppendVarint(prefix[:0], uint64(len(out)-start))...)
}

func isJSONNumber(text string) bool {
	v := jsontext.Value(text)
	return text != "" && text == strings.TrimSpace(text) && v.Kind() == '0' && v.IsValid()
}

// parseInt parses an integer. Like protojson, it also accepts integral
// values written with a fraction or an exponent.
func parseInt(text string, bits int) (int64, error) {
	if v, err := strconv.ParseInt(text, 10, bits); err == nil {
		return v, nil
	}
	if integer, ok := integerText(text); ok {
		if v, err := strconv.ParseInt(integer, 10, bits); err == nil {
			return v, nil
		}
	}
	return 0, fmt.Errorf("invalid integer %q", text)
}

func parseUint(text string, bits int) (uint64, error) {
	if v, err := strconv.ParseUint(text, 10, bits); err == nil {
		return v, nil
	}
	if integer, ok := integerText(text); ok {
		if v, err := strconv.ParseUint(integer, 10, bits); err == nil {
			return v, nil
		}
	}
	return 0, fmt.Errorf("invalid unsigned integer %q", text)
}

// integerText expands an integral JSON number into decimal digits without
// rounding. At most 20 digits can fit in a uint64, so even a huge exponent
// must not cause an allocation proportional to its value.
func integerText(text string) (string, bool) {
	negative := strings.HasPrefix(text, "-")
	if negative {
		text = text[1:]
	}
	var exponent string
	if i := strings.IndexAny(text, "eE"); i >= 0 {
		text, exponent = text[:i], text[i+1:]
	}
	whole, fraction, _ := strings.Cut(text, ".")
	whole = strings.TrimLeft(whole, "0")
	fraction = strings.TrimRight(fraction, "0")
	if whole == "" && fraction == "" {
		return "0", true
	}
	var exp int64
	if exponent != "" {
		var err error
		exp, err = strconv.ParseInt(exponent, 10, 32)
		if err != nil {
			return "", false
		}
	}
	var integer string
	if exp >= 0 {
		if int64(len(fraction)) > exp || int64(len(whole))+exp > 20 {
			return "", false
		}
		integer = whole + fraction + strings.Repeat("0", int(exp)-len(fraction))
	} else {
		digits := int64(len(whole)) + exp
		if fraction != "" || digits <= 0 || digits > 20 || strings.TrimRight(whole[digits:], "0") != "" {
			return "", false
		}
		integer = whole[:digits]
	}
	if negative {
		integer = "-" + integer
	}
	return integer, true
}

func parseFloat(text string, bits int) (float64, error) {
	switch text {
	case "NaN":
		return math.NaN(), nil
	case "Infinity":
		return math.Inf(+1), nil
	case "-Infinity":
		return math.Inf(-1), nil
	}
	x, err := strconv.ParseFloat(text, bits)
	if err != nil || math.IsInf(x, 0) || math.IsNaN(x) {
		return 0, fmt.Errorf("invalid number %q", text)
	}
	return x, nil
}

// decodeBase64 accepts standard and URL-safe base64, with or without
// padding, as protojson does.
func decodeBase64(text string) ([]byte, error) {
	enc := base64.StdEncoding
	if strings.ContainsAny(text, "-_") {
		enc = base64.URLEncoding
	}
	if len(text)%4 != 0 {
		enc = enc.WithPadding(base64.NoPadding)
	}
	b, err := enc.DecodeString(text)
	if err != nil {
		return nil, fmt.Errorf("invalid base64: %w", err)
	}
	return b, nil
}

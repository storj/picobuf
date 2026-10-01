// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"bytes"
	"slices"
	"unsafe"

	"storj.io/picobuf/internal/protowire"
)

const (
	fieldDecodingErrored = FieldNumber(-1)
	fieldDecodingDone    = FieldNumber(-2)
)

// Decoder implements decoding of protobuf messages.
type Decoder struct {
	messageDecodeState
	err *parseError

	// Serialized protobuf messages are smaller than 2 GiB, so int32 is
	// sufficient for nesting and repeated-element limits and counters.
	depthBudget         int32
	skipUTF8Budget      int32
	maxRepeatedElements int32
	repeatedElements    int32

	init               bool
	aliasInput         bool
	allowInvalidUTF8   bool
	skipUTF8Validation bool
}

// WithoutUTF8Validation runs fn without validating decoded string fields.
// It is intended for generated code implementing protobuf schema features.
//
// The exemption covers strings decoded directly within the current message,
// including map entry keys and values, but not strings inside nested
// messages, which carry their own UTF-8 validation features. fn cannot
// panic, so the state is restored with a plain call rather than a defer.
func (dec *Decoder) WithoutUTF8Validation(fn func()) {
	prevSkip, prevBudget := dec.skipUTF8Validation, dec.skipUTF8Budget
	dec.skipUTF8Validation, dec.skipUTF8Budget = true, dec.depthBudget
	fn()
	dec.skipUTF8Validation, dec.skipUTF8Budget = prevSkip, prevBudget
}

// skipUTF8 reports whether string validation is currently disabled, per the
// scoping rules of WithoutUTF8Validation.
func (dec *Decoder) skipUTF8() bool {
	return dec.skipUTF8Validation && dec.depthBudget >= dec.skipUTF8Budget-1
}

type messageDecodeState struct {
	pendingField FieldNumber    //nolint: structcheck
	pendingWire  protowire.Type //nolint: structcheck

	buffer     []byte
	groupField FieldNumber // nonzero while reading a group, whose end tag terminates this scope
}

// NewDecoder returns a new Decoder.
func NewDecoder(data []byte) *Decoder {
	return newDecoder(data)
}

func newDecoder(data []byte) *Decoder {
	return &Decoder{
		messageDecodeState: messageDecodeState{buffer: data},
		depthBudget:        protowire.DefaultRecursionLimit,
	}
}

// Unmarshal decodes data into msg using reusable decoder storage. The zero value
// is ready to use. Pass UnmarshalOptions{} for the default decoding options.
// Like UnmarshalOptions.Unmarshal, this merges into msg rather than resetting it.
//
// A Decoder must not be used concurrently or reentered from a decoding callback.
// Its state is cleared before this method returns, including on failure or panic;
// callbacks must not retain it for use after Unmarshal returns. Decoded values
// remain valid according to opts.AliasInput, independently of decoder reuse.
func (dec *Decoder) Unmarshal(data []byte, msg Message, opts UnmarshalOptions) error {
	defer func() { *dec = Decoder{} }()
	if opts.MaxInputSize > 0 && len(data) > opts.MaxInputSize {
		return parseError{message: "input exceeds maximum size"}
	}
	return opts.unmarshal(data, msg, dec)
}

// PendingField returns the next field number in the stream.
func (dec *Decoder) PendingField() FieldNumber { return dec.pendingField }

// Err returns error that occurred during decoding.
func (dec *Decoder) Err() error {
	if dec.err == nil {
		return nil
	}
	return dec.err
}

// enterMessage saves the parent in the caller's stack frame rather than in a
// growing slice. The depth budget counts how many more levels may be entered.
func (dec *Decoder) enterMessage(message []byte) messageDecodeState {
	return dec.enterScope(message, 0)
}

func (dec *Decoder) enterScope(message []byte, group FieldNumber) messageDecodeState {
	parent := dec.messageDecodeState
	tooDeep := dec.depthBudget <= 0
	if tooDeep {
		message = nil
	}
	dec.depthBudget--
	dec.messageDecodeState = messageDecodeState{buffer: message, groupField: group}
	dec.nextField(0)
	if tooDeep {
		dec.fail(0, "exceeded maximum recursion depth")
	}
	return parent
}

func (dec *Decoder) leaveMessage(parent messageDecodeState) {
	dec.messageDecodeState = parent
	dec.depthBudget++
}

// RepeatedMessage decodes a message.
func (dec *Decoder) RepeatedMessage(field FieldNumber, fn func(c *Decoder)) {
	for field == dec.pendingField {
		if !dec.takeRepeated(field) {
			return
		}
		if dec.pendingWire != protowire.BytesType {
			dec.fail(field, "expected wire type Bytes")
			return
		}

		message, n := protowire.ConsumeBytes(dec.buffer)
		parent := dec.enterMessage(message)
		fn(dec)
		dec.leaveMessage(parent)

		dec.nextField(n)
	}
}

// RepeatedMessageCount returns the number of consecutive message values for
// field, capped by the remaining repeated element limit.
func (dec *Decoder) RepeatedMessageCount(field FieldNumber) int {
	return dec.repeatedFieldCount(field, protowire.BytesType)
}

// DecodeRepeatedMessages decodes repeated message values for field into out.
// The elements are allocated in blocks of at most maxElementBlockBytes
// instead of one allocation each, so retaining a single element keeps the
// rest of its block alive.
func DecodeRepeatedMessages[T any](dec *Decoder, field FieldNumber, out *[]*T, decode func(*T, *Decoder)) {
	remaining := dec.RepeatedMessageCount(field)
	*out = slices.Grow(*out, remaining)
	var block []T
	dec.RepeatedMessage(field, func(c *Decoder) {
		var x *T
		if len(block) == 0 && remaining > 0 {
			n := elementBlockLen(remaining, unsafe.Sizeof(*x))
			block, remaining = make([]T, n), remaining-n
		}
		// The count is a lower bound, so fall back on individual
		// allocations once it is exceeded.
		if len(block) > 0 {
			x, block = &block[0], block[1:]
		} else {
			x = new(T)
		}
		c.Loop(func(c *Decoder) { decode(x, c) })
		*out = append(*out, x)
	})
}

// DecodeRepeatedGroups is DecodeRepeatedMessages for delimited messages.
func DecodeRepeatedGroups[T any](dec *Decoder, field FieldNumber, out *[]*T, decode func(*T, *Decoder)) {
	remaining := dec.RepeatedGroupCount(field)
	*out = slices.Grow(*out, remaining)
	var block []T
	dec.RepeatedGroup(field, func(c *Decoder) {
		var x *T
		if len(block) == 0 && remaining > 0 {
			n := elementBlockLen(remaining, unsafe.Sizeof(*x))
			block, remaining = make([]T, n), remaining-n
		}
		// The count is a lower bound, so fall back on individual
		// allocations once it is exceeded.
		if len(block) > 0 {
			x, block = &block[0], block[1:]
		} else {
			x = new(T)
		}
		c.Loop(func(c *Decoder) { decode(x, c) })
		*out = append(*out, x)
	})
}

// maxElementBlockBytes bounds the memory that one retained element of a
// repeated message field can keep alive.
const maxElementBlockBytes = 4 << 10

// elementBlockLen returns how many of the remaining elements of the given
// size to allocate in the next block.
func elementBlockLen(remaining int, size uintptr) int {
	return min(remaining, max(1, maxElementBlockBytes/max(1, int(size))))
}

// RepeatedEnum decodes a repeated enumeration.
func (dec *Decoder) RepeatedEnum(field FieldNumber, add func(x int32)) {
	for field == dec.pendingField {
		switch dec.pendingWire {
		case protowire.BytesType:
			packed, n := protowire.ConsumeBytes(dec.buffer)
			for len(packed) > 0 {
				if !dec.takeRepeated(field) {
					return
				}
				x, xn := protowire.ConsumeVarint(packed)
				if xn < 0 {
					dec.fail(field, "unable to parse Varint")
					return
				}
				add(int32(x))
				packed = packed[xn:]
			}
			dec.nextField(n)
		case protowire.VarintType:
			if !dec.takeRepeated(field) {
				return
			}
			x, n := protowire.ConsumeVarint(dec.buffer)
			if n < 0 {
				dec.fail(field, "unable to parse Varint")
				return
			}
			add(int32(x))
			dec.nextField(n)
		default:
			dec.fail(field, "expected wire type Varint")
			return
		}
	}
}

// TakeRepeated reserves one repeated element when field is pending. It returns
// false on another field or an exhausted limit. Generated custom decoders call
// it before allocating each element; their scalar decoding methods do not count
// repeated values themselves.
func (dec *Decoder) TakeRepeated(field FieldNumber) bool {
	return dec.pendingField == field && dec.takeRepeated(field)
}

func (dec *Decoder) takeRepeated(field FieldNumber) bool {
	if dec.maxRepeatedElements <= 0 {
		return true
	}
	if dec.repeatedElements >= dec.maxRepeatedElements {
		dec.fail(field, "exceeded maximum repeated elements")
		return false
	}
	dec.repeatedElements++
	return true
}

func (dec *Decoder) repeatedReservation(count int) int {
	if dec.maxRepeatedElements <= 0 {
		return count
	}
	remaining := int(dec.maxRepeatedElements - dec.repeatedElements)
	if remaining < count {
		return max(remaining, 0)
	}
	return count
}

// Message decodes a message.
func (dec *Decoder) Message(field FieldNumber, fn func(*Decoder)) {
	if field != dec.pendingField {
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}

	message, n := protowire.ConsumeBytes(dec.buffer)
	parent := dec.enterMessage(message)
	dec.Loop(fn)
	dec.leaveMessage(parent)

	dec.nextField(n)
}

// PresentMessage decodes an always present message.
func (dec *Decoder) PresentMessage(field FieldNumber, fn func(*Decoder)) {
	if field != dec.pendingField {
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}

	message, n := protowire.ConsumeBytes(dec.buffer)
	parent := dec.enterMessage(message)
	dec.Loop(fn)
	dec.leaveMessage(parent)

	dec.nextField(n)
}

// Group decodes a delimited message.
func (dec *Decoder) Group(field FieldNumber, fn func(*Decoder)) {
	if field != dec.pendingField {
		return
	}
	if dec.pendingWire != protowire.StartGroupType {
		dec.fail(field, "expected wire type StartGroup")
		return
	}

	dec.group(field, func(c *Decoder) { c.Loop(fn) })
}

// PresentGroup decodes an always-present delimited message.
func (dec *Decoder) PresentGroup(field FieldNumber, fn func(*Decoder)) {
	dec.Group(field, fn)
}

// RepeatedGroup decodes repeated delimited messages.
func (dec *Decoder) RepeatedGroup(field FieldNumber, fn func(*Decoder)) {
	for field == dec.pendingField {
		if !dec.takeRepeated(field) {
			return
		}
		if dec.pendingWire != protowire.StartGroupType {
			dec.fail(field, "expected wire type StartGroup")
			return
		}

		dec.group(field, fn)
	}
}

// group decodes directly from the parent buffer. nextField consumes the matching
// end tag and stops this scope, leaving the parent's remaining bytes untouched.
func (dec *Decoder) group(field FieldNumber, fn func(*Decoder)) {
	parent := dec.enterScope(dec.buffer, field)
	fn(dec)
	// A callback may leave fields unread, just as with RepeatedMessage.
	if dec.pendingField.IsValid() {
		dec.Loop(func(*Decoder) {})
	}
	consumed := len(parent.buffer) - len(dec.buffer)
	dec.leaveMessage(parent)
	dec.nextField(consumed)
}

// RepeatedGroupCount returns a lower bound on consecutive group values for field,
// capped by the remaining repeated element limit. Lookahead is bounded to avoid
// rescanning entire nested groups before decoding them.
func (dec *Decoder) RepeatedGroupCount(field FieldNumber) int {
	return dec.repeatedFieldCount(field, protowire.StartGroupType)
}

func (dec *Decoder) repeatedFieldCount(field FieldNumber, wire protowire.Type) int {
	if dec.pendingField != field || dec.pendingWire != wire {
		return 0
	}

	buffer := dec.buffer
	if wire == protowire.StartGroupType {
		buffer = buffer[:min(len(buffer), 256)]
	}
	count := 0
	for {
		n := protowire.ConsumeFieldValueDepth(protowire.Number(field), wire, buffer, int(dec.depthBudget))
		if n < 0 {
			break
		}
		count++
		buffer = buffer[n:]
		if len(buffer) == 0 {
			break
		}
		nextField, nextWire, n := protowire.ConsumeTag(buffer)
		if n < 0 || FieldNumber(nextField) != field || nextWire != wire {
			break
		}
		buffer = buffer[n:]
	}
	return dec.repeatedReservation(count)
}

// UnrecognizedFields decodes fields that are not in the provided set.
//
// Fields below 64 are excluded via the exclude bitmask, higher field numbers
// are listed in excludeHigh.
func (dec *Decoder) UnrecognizedFields(exclude uint64, out *[]byte, excludeHigh ...FieldNumber) {
	for dec.pendingField >= 0 {
		if field := dec.pendingField; field < 64 {
			if exclude&(1<<uint64(field)) != 0 {
				return
			}
		} else if containsField(excludeHigh, field) {
			return
		}

		n := protowire.ConsumeFieldValueDepth(protowire.Number(dec.pendingField), dec.pendingWire, dec.buffer, int(dec.depthBudget))
		if n < 0 {
			dec.fail(dec.pendingField, "unable to parse unrecognized field")
			return
		}
		*out = protowire.AppendTag(*out, protowire.Number(dec.pendingField), dec.pendingWire)
		*out = append(*out, dec.buffer[:n]...)
		dec.nextField(n)
	}
}

// RejectUnrecognizedFields fails decoding when a pending field is not in the
// provided set.
//
// Fields below 64 are excluded via the exclude bitmask, higher field numbers
// are listed in excludeHigh.
func (dec *Decoder) RejectUnrecognizedFields(exclude uint64, excludeHigh ...FieldNumber) {
	if field := dec.pendingField; field >= 0 {
		if field < 64 {
			if exclude&(1<<uint64(field)) != 0 {
				return
			}
		} else if containsField(excludeHigh, field) {
			return
		}
		dec.fail(field, "unrecognized field")
	}
}

// containsField reports whether fields contains field.
//
// It does a linear scan, because a message that captures unrecognized fields has
// few fields numbered 64 or above. Sort and binary search if that changes.
func containsField(fields []FieldNumber, field FieldNumber) bool {
	for _, candidate := range fields {
		if candidate == field {
			return true
		}
	}
	return false
}

// Loop loops fields until all messages have been processed.
func (dec *Decoder) Loop(fn func(*Decoder)) {
	if !dec.init {
		dec.nextField(0)
		dec.init = true
	}

	for {
		startingLength := len(dec.buffer)
		fn(dec)
		if !dec.pendingField.IsValid() {
			break
		}
		if len(dec.buffer) == startingLength {
			// we didn't process any of the fields
			n := protowire.ConsumeFieldValueDepth(protowire.Number(dec.pendingField), dec.pendingWire, dec.buffer, int(dec.depthBudget))
			dec.nextField(n)
		}
	}
}

// BytesView decodes a bytes field without copying its contents, regardless of
// UnmarshalOptions.AliasInput. The view aliases the input and must be treated as
// read-only. Copy it if it must remain unchanged when the input is modified or
// reused.
// This is intended for custom decoders that immediately copy into their own
// storage, avoiding the intermediate allocation made by Bytes.
// Like Bytes, it leaves v unchanged when field is not pending or cannot be read.
func (dec *Decoder) BytesView(field FieldNumber, v *[]byte) {
	if field != dec.pendingField {
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}
	x, n := protowire.ConsumeBytes(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Bytes")
		return
	}
	// Limit the capacity, so that append on the view cannot overwrite the
	// rest of the input.
	*v = x[:len(x):len(x)]
	dec.nextField(n)
}

// copyBytes returns x, copied out of the input unless the decoder was asked to
// alias it. See UnmarshalOptions.AliasInput.
func (dec *Decoder) copyBytes(x []byte) []byte {
	if dec.aliasInput {
		return x
	}
	return bytes.Clone(x)
}

// Fail fails the decoding process.
func (dec *Decoder) Fail(field FieldNumber, msg string) {
	dec.fail(field, msg)
}

//go:noinline
func (dec *Decoder) fail(field FieldNumber, msg string) {
	// TODO: use static error types
	dec.pendingField = fieldDecodingErrored
	dec.err = &parseError{field: field, message: msg}
}

type parseError struct {
	field   FieldNumber
	message string
}

func (e parseError) Error() string {
	return "failed while parsing " + e.field.String() + ": " + e.message
}

func (dec *Decoder) nextField(advance int) {
	if dec.err != nil {
		dec.pendingField = fieldDecodingErrored
		return
	}
	if advance < 0 || advance > len(dec.buffer) {
		dec.fail(0, "advance outside buffer")
		return
	}
	dec.buffer = dec.buffer[advance:]
	if len(dec.buffer) == 0 {
		if dec.groupField != 0 {
			dec.fail(dec.groupField, "unterminated group")
			return
		}
		dec.pendingField = fieldDecodingDone
		return
	}

	// Fields 1-15 have a one-byte tag. Avoid the general varint parser for
	// this common case; zero field numbers and longer tags use its checks.
	if tag := dec.buffer[0]; tag >= 8 && tag < 0x80 {
		dec.pendingField = FieldNumber(tag >> 3)
		dec.pendingWire = protowire.Type(tag & 7)
		dec.buffer = dec.buffer[1:]
		dec.checkEndGroup()
		return
	}

	field, wire, n := protowire.ConsumeTag(dec.buffer)
	if n < 0 {
		dec.fail(0, "failed to parse") // TODO: better error message
		return
	}
	// ConsumeTag only rejects field numbers below the minimum, so numbers
	// between MaxValidNumber and MaxInt32 arrive here. Loop treats those as
	// invalid and stops, which would silently discard the rest of the message.
	if !FieldNumber(field).IsValid() {
		dec.fail(FieldNumber(field), "invalid field number")
		return
	}
	dec.buffer = dec.buffer[n:]
	dec.pendingField, dec.pendingWire = FieldNumber(field), wire
	dec.checkEndGroup()
}

// checkEndGroup prevents callbacks from seeing end tags as ordinary fields.
func (dec *Decoder) checkEndGroup() {
	if dec.pendingWire != protowire.EndGroupType {
		return
	}
	if dec.groupField == 0 || dec.pendingField != dec.groupField {
		dec.fail(dec.pendingField, "unexpected end group")
		return
	}
	dec.pendingField = fieldDecodingDone
}

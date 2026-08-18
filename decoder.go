// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"bytes"

	"storj.io/picobuf/internal/protowire"
)

const (
	fieldDecodingErrored = FieldNumber(-1)
	fieldDecodingDone    = FieldNumber(-2)
)

// Decoder implements decoding of protobuf messages.
type Decoder struct {
	messageDecodeState
	stack              []messageDecodeState
	init               bool
	aliasInput         bool
	allowInvalidUTF8   bool
	skipUTF8Validation bool
	skipUTF8Depth      int

	maxRecursionDepth   int
	maxRepeatedElements int
	repeatedElements    int
	err                 error
}

// WithoutUTF8Validation runs fn without validating decoded string fields.
// It is intended for generated code implementing protobuf schema features.
//
// The exemption covers strings decoded directly within the current message,
// including map entry keys and values, but not strings inside nested
// messages, which carry their own UTF-8 validation features. fn cannot
// panic, so the state is restored with a plain call rather than a defer.
func (dec *Decoder) WithoutUTF8Validation(fn func()) {
	prevSkip, prevDepth := dec.skipUTF8Validation, dec.skipUTF8Depth
	dec.skipUTF8Validation, dec.skipUTF8Depth = true, len(dec.stack)
	fn()
	dec.skipUTF8Validation, dec.skipUTF8Depth = prevSkip, prevDepth
}

// skipUTF8 reports whether string validation is currently disabled, per the
// scoping rules of WithoutUTF8Validation.
func (dec *Decoder) skipUTF8() bool {
	return dec.skipUTF8Validation && len(dec.stack) <= dec.skipUTF8Depth+1
}

type messageDecodeState struct {
	pendingField FieldNumber    //nolint: structcheck
	pendingWire  protowire.Type //nolint: structcheck

	buffer []byte
}

// NewDecoder returns a new Decoder.
func NewDecoder(data []byte) *Decoder {
	return newDecoder(data)
}

func newDecoder(data []byte) *Decoder {
	return &Decoder{
		messageDecodeState: messageDecodeState{buffer: data},
		maxRecursionDepth:  protowire.DefaultRecursionLimit,
	}
}

// PendingField returns the next field number in the stream.
func (dec *Decoder) PendingField() FieldNumber { return dec.pendingField }

// Err returns error that occurred during decoding.
func (dec *Decoder) Err() error {
	return dec.err
}

func (dec *Decoder) pushState(message []byte) {
	// Nesting is bounded by the input for a self-referential message, so
	// refuse to descend further rather than exhausting the stack. Still push,
	// so that the caller's matching popState stays balanced.
	tooDeep := len(dec.stack) >= dec.maxRecursionDepth
	if tooDeep {
		message = nil
	}

	dec.stack = append(dec.stack, dec.messageDecodeState)
	dec.messageDecodeState = messageDecodeState{
		buffer: message,
	}
	dec.nextField(0)

	if tooDeep {
		dec.fail(0, "exceeded maximum recursion depth")
	}
}

func (dec *Decoder) popState() {
	if len(dec.stack) == 0 {
		dec.fail(0, "stack mangled")
		return
	}
	dec.messageDecodeState = dec.stack[len(dec.stack)-1]
	dec.stack = dec.stack[:len(dec.stack)-1]
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
		dec.pushState(message)
		fn(dec)
		dec.popState()

		dec.nextField(n)
	}
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

func (dec *Decoder) takeRepeated(field FieldNumber) bool {
	if dec.maxRepeatedElements <= 0 {
		return true
	}
	dec.repeatedElements++
	if dec.repeatedElements > dec.maxRepeatedElements {
		dec.fail(field, "exceeded maximum repeated elements")
		return false
	}
	return true
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
	dec.pushState(message)
	dec.Loop(fn)
	dec.popState()

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
	dec.pushState(message)
	dec.Loop(fn)
	dec.popState()

	dec.nextField(n)
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

		n := protowire.ConsumeFieldValue(protowire.Number(dec.pendingField), dec.pendingWire, dec.buffer)
		if n < 0 {
			dec.fail(dec.pendingField, "unable to parse unrecognized field")
			return
		}
		*out = protowire.AppendTag(*out, protowire.Number(dec.pendingField), dec.pendingWire)
		*out = append(*out, dec.buffer[:n]...)
		dec.nextField(n)
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
			n := protowire.ConsumeFieldValue(protowire.Number(dec.pendingField), dec.pendingWire, dec.buffer)
			dec.nextField(n)
		}
	}
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
	dec.err = parseError{field: field, message: msg}
}

type parseError struct {
	field   FieldNumber
	message string
}

func (e parseError) Error() string {
	return "failed while parsing " + e.field.String() + ": " + e.message
}

func (dec *Decoder) nextField(advance int) {
	if advance < 0 || advance > len(dec.buffer) {
		dec.fail(0, "advance outside buffer")
		return
	}
	dec.buffer = dec.buffer[advance:]
	if len(dec.buffer) == 0 {
		dec.pendingField = fieldDecodingDone
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
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"errors"

	"storj.io/picobuf/internal/protowire"
)

// Encoder implements encoding of protobuf format.
type Encoder struct {
	buffer             []byte
	depth              int
	err                error
	skipUTF8Validation bool
	skipUTF8Depth      int
}

var errEncodeRecursionDepth = errors.New("picobuf: exceeded maximum nesting depth while encoding")

// WithoutUTF8Validation runs fn without validating encoded string fields.
// It is intended for generated code implementing protobuf schema features.
//
// The exemption covers strings encoded directly within the current message,
// including map entry keys and values, but not strings inside nested
// messages, which carry their own UTF-8 validation features. fn cannot
// panic, so the state is restored with a plain call rather than a defer.
func (enc *Encoder) WithoutUTF8Validation(fn func()) {
	prevSkip, prevDepth := enc.skipUTF8Validation, enc.skipUTF8Depth
	enc.skipUTF8Validation, enc.skipUTF8Depth = true, enc.depth
	fn()
	enc.skipUTF8Validation, enc.skipUTF8Depth = prevSkip, prevDepth
}

// skipUTF8 reports whether string validation is currently disabled, per the
// scoping rules of WithoutUTF8Validation.
func (enc *Encoder) skipUTF8() bool {
	return enc.skipUTF8Validation && enc.depth <= enc.skipUTF8Depth+1
}

// Err returns the error that occurred during encoding.
func (enc *Encoder) Err() error { return enc.err }

// Fail fails the encoding process.
func (enc *Encoder) Fail(field FieldNumber, msg string) {
	enc.fail(field, msg)
}

//go:noinline
func (enc *Encoder) fail(field FieldNumber, msg string) {
	if enc.err == nil {
		enc.err = &encodeError{field: field, message: msg}
	}
}

type encodeError struct {
	field   FieldNumber
	message string
}

func (e encodeError) Error() string {
	return "failed while encoding " + e.field.String() + ": " + e.message
}

// enterMessage reports whether another level of nesting may be encoded. A
// self-referential message can nest without bound, which would exhaust the
// stack, so refuse to descend past the limit the decoder also uses.
func (enc *Encoder) enterMessage() bool {
	if enc.depth >= protowire.DefaultRecursionLimit {
		if enc.err == nil {
			enc.err = errEncodeRecursionDepth
		}
		return false
	}
	enc.depth++
	return true
}

// NewEncoder creates a new Encoder.
func NewEncoder() *Encoder {
	return NewEncoderBuffer(make([]byte, 0, 64))
}

// NewEncoderBuffer creates a new encoder using a preallocated buffer.
func NewEncoderBuffer(buffer []byte) *Encoder {
	return &Encoder{buffer: buffer[:0]}
}

// Buffer returns the encoded internal buffer.
func (enc *Encoder) Buffer() []byte { return enc.buffer }

// Message decodes a message.
//
//go:noinline
func (enc *Encoder) Message(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.anyBytes(field, func() bool { return fn(enc) })
}

// AlwaysMessage encodes an message always.
//
//go:noinline
func (enc *Encoder) AlwaysMessage(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.alwaysAnyBytes(field, func() { fn(enc) })
}

// PresentMessage encodes an always present message.
//
//go:noinline
func (enc *Encoder) PresentMessage(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.anyBytes(field, func() bool {
		lengthStart := len(enc.buffer)
		fn(enc)
		return len(enc.buffer) > lengthStart
	})
}

// Group encodes a delimited message when it is present.
//
//go:noinline
func (enc *Encoder) Group(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.anyGroup(field, fn)
}

// AlwaysGroup always encodes a delimited message.
//
//go:noinline
func (enc *Encoder) AlwaysGroup(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.alwaysGroup(field, func() { fn(enc) })
}

// PresentGroup encodes a non-empty always-present delimited message.
//
//go:noinline
func (enc *Encoder) PresentGroup(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.anyGroup(field, func(enc *Encoder) bool {
		start := len(enc.buffer)
		fn(enc)
		return len(enc.buffer) > start
	})
}

func (enc *Encoder) anyGroup(field FieldNumber, fn func(enc *Encoder) bool) bool {
	start := len(enc.buffer)
	enc.buffer = appendTag(enc.buffer, field, protowire.StartGroupType)
	if !enc.enterMessage() {
		enc.buffer = enc.buffer[:start]
		return false
	}
	ok := fn(enc)
	enc.depth--
	if !ok {
		enc.buffer = enc.buffer[:start]
		return false
	}
	enc.buffer = appendTag(enc.buffer, field, protowire.EndGroupType)
	return true
}

func (enc *Encoder) alwaysGroup(field FieldNumber, fn func()) {
	enc.buffer = appendTag(enc.buffer, field, protowire.StartGroupType)
	if enc.enterMessage() {
		fn()
		enc.depth--
	}
	enc.buffer = appendTag(enc.buffer, field, protowire.EndGroupType)
}

// RepeatedEnum encodes a repeated enumeration.
//
//go:noinline
func (enc *Encoder) RepeatedEnum(field FieldNumber, n int, fn func(index uint) int32) {
	if n == 0 {
		return
	}
	enc.alwaysAnyBytes(field, func() {
		for i := range n {
			enc.buffer = protowire.AppendVarint(enc.buffer, uint64(fn(uint(i))))
		}
	})
}

// anyBytes encodes field as Bytes and handles encoding the length.
func (enc *Encoder) anyBytes(field FieldNumber, fn func() bool) bool {
	tagStart := len(enc.buffer)
	enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)
	lengthStart := len(enc.buffer)
	// We'll guess that we need 2 bytes for length.
	// If we need less, then the copy is fast, and needing more is unlikely.
	var lengthBufferPrediction [2]byte
	enc.buffer = append(enc.buffer, lengthBufferPrediction[:]...)
	messageStart := len(enc.buffer)
	// encode the submessage
	ok := enc.enterMessage()
	if ok {
		ok = fn()
		enc.depth--
	}
	if !ok {
		// The message was nil or too deeply nested, we can remove the tag.
		enc.buffer = enc.buffer[:tagStart]
		return false
	}
	messageLength := len(enc.buffer) - messageStart
	bytesForSize := protowire.SizeVarint(uint64(messageLength))
	if bytesForSize == len(lengthBufferPrediction) {
		protowire.PutUvarint(enc.buffer[lengthStart:messageStart], uint64(messageLength))
		return true
	}
	if bytesForSize > len(lengthBufferPrediction) {
		enc.buffer = append(enc.buffer, make([]byte, bytesForSize-len(lengthBufferPrediction))...)
	}

	copy(enc.buffer[lengthStart+bytesForSize:], enc.buffer[messageStart:])
	protowire.PutUvarint(enc.buffer[lengthStart:lengthStart+bytesForSize], uint64(messageLength))
	enc.buffer = enc.buffer[:lengthStart+bytesForSize+messageLength]
	return true
}

// AlwaysAnyBytes encodes field as Bytes and handles encoding the length.
func (enc *Encoder) AlwaysAnyBytes(field FieldNumber, fn func()) bool {
	return enc.alwaysAnyBytes(field, fn)
}

// alwaysAnyBytes encodes field as Bytes and handles encoding the length.
func (enc *Encoder) alwaysAnyBytes(field FieldNumber, fn func()) bool {
	enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)
	lengthStart := len(enc.buffer)
	// We'll guess that we need 2 bytes for length.
	// If we need less, then the copy is fast, and needing more is unlikely.
	var lengthBufferPrediction [2]byte
	enc.buffer = append(enc.buffer, lengthBufferPrediction[:]...)
	messageStart := len(enc.buffer)
	// encode the submessage
	if enc.enterMessage() {
		fn()
		enc.depth--
	}
	messageLength := len(enc.buffer) - messageStart
	bytesForSize := protowire.SizeVarint(uint64(messageLength))
	if bytesForSize == len(lengthBufferPrediction) {
		protowire.PutUvarint(enc.buffer[lengthStart:messageStart], uint64(messageLength))
		return true
	}
	if bytesForSize > len(lengthBufferPrediction) {
		enc.buffer = append(enc.buffer, make([]byte, bytesForSize-len(lengthBufferPrediction))...)
	}

	copy(enc.buffer[lengthStart+bytesForSize:], enc.buffer[messageStart:])
	protowire.PutUvarint(enc.buffer[lengthStart:lengthStart+bytesForSize], uint64(messageLength))
	enc.buffer = enc.buffer[:lengthStart+bytesForSize+messageLength]
	return true
}

// UnrecognizedFields encodes fields that are not in the provided set.
func (enc *Encoder) UnrecognizedFields(out []byte) {
	enc.buffer = append(enc.buffer, out...)
}

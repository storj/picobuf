// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"storj.io/picobuf/internal/protowire"
)

// Encoder encodes the protobuf wire format.
type Encoder struct {
	buffer []byte
	size   int
	err    *encodeError

	// Marshal records prefix widths in preorder during sizing, then consumes
	// them during encoding. They are hints: map traversal can change order.
	inlineHints [2]uint64
	lengthHints *lengthHints

	depth         int32
	skipUTF8Depth int32

	// hintCount is the number of inline hints plus one; zero disables hints.
	hintCount uint8
	hintRead  uint8

	sizing             bool
	skipUTF8Validation bool
}

var errEncodeRecursionDepth = &encodeError{message: "picobuf: exceeded maximum nesting depth while encoding"}

// WithoutUTF8Validation runs fn without validating the string fields that fn
// encodes. Generated code uses it to implement protobuf schema features.
//
// The exemption covers the strings of the current message, including map entry
// keys and values. It does not cover strings inside nested messages, which
// have their own UTF-8 validation features. fn cannot panic, so the encoder
// restores its state with a plain call instead of a defer.
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
func (enc *Encoder) Err() error {
	if enc.err == nil {
		return nil
	}
	return enc.err
}

// IsSizing reports whether the encoder is counting bytes instead of writing
// them. A custom encoder can use it to skip work that only the written bytes
// need.
func (enc *Encoder) IsSizing() bool { return enc.sizing }

func (enc *Encoder) offset() int {
	if enc.sizing {
		return enc.size
	}
	return len(enc.buffer)
}

func (enc *Encoder) truncate(offset int) {
	if enc.sizing {
		enc.size = offset
	} else {
		enc.buffer = enc.buffer[:offset]
	}
}

func (enc *Encoder) appendTag(field FieldNumber, typ protowire.Type) {
	if enc.sizing {
		enc.size += protowire.SizeTag(protowire.Number(field))
	} else {
		enc.buffer = appendTag(enc.buffer, field, typ)
	}
}

func (enc *Encoder) appendVarint(v uint64) {
	if enc.sizing {
		enc.size += protowire.SizeVarint(v)
	} else {
		enc.buffer = protowire.AppendVarint(enc.buffer, v)
	}
}

// Fail fails the encoding with an error for field.
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

func (e *encodeError) Error() string {
	if e == errEncodeRecursionDepth {
		return e.message
	}
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

// NewEncoder creates an Encoder that encodes fields without a message type.
// Call the field methods, check Err, and take the result from Buffer. Unlike
// Marshal, it does not validate the required fields of a message encoded with
// msg.Encode.
func NewEncoder() *Encoder {
	return NewEncoderBuffer(make([]byte, 0, 64))
}

// NewEncoderBuffer is NewEncoder that writes into buffer, starting at
// buffer[:0]. It does not allocate the output while the output fits.
func NewEncoderBuffer(buffer []byte) *Encoder {
	return &Encoder{buffer: buffer[:0]}
}

// MarshalAppend appends the encoding of msg to buffer and reuses the storage
// of the encoder. The zero value is ready to use. Like the MarshalAppend
// function, it sizes the message first and then encodes it, so msg.Encode runs
// twice and must do the same thing both times.
//
// Do not use an Encoder concurrently, and do not reenter it from an encoding
// callback. A callback must not keep the encoder for use after MarshalAppend
// returns. MarshalAppend clears the state of the message, also on error or
// panic, and keeps its scratch space for length prefixes. It does not keep the
// output buffer. The caller owns the returned bytes and decides when to reuse
// them.
func (enc *Encoder) MarshalAppend(buffer []byte, msg Message) ([]byte, error) {
	defer enc.reset()
	return enc.marshalAppend(buffer, msg)
}

// reset clears the state of the encoder, but keeps its scratch space for
// length prefixes.
func (enc *Encoder) reset() {
	hints := enc.lengthHints
	if hints != nil {
		hints.widths = hints.widths[:0]
		hints.read = 0
	}
	*enc = Encoder{lengthHints: hints}
}

// Buffer returns the bytes encoded so far.
func (enc *Encoder) Buffer() []byte { return enc.buffer }

// Message encodes a message field when fn reports that the message is present.
//
//go:noinline
func (enc *Encoder) Message(field FieldNumber, fn func(enc *Encoder) bool) {
	if enc.sizing {
		hint := enc.reserveLengthHint()
		start := enc.size
		if !enc.enterMessage() {
			return
		}
		ok := fn(enc)
		enc.depth--
		if !ok {
			enc.size = start
			return
		}
		messageLength := enc.size - start
		width := protowire.SizeVarint(uint64(messageLength))
		if hint >= 0 {
			enc.setLengthHint(hint, width)
		}
		enc.size += protowire.SizeTag(protowire.Number(field)) + width
		return
	}
	tagStart := len(enc.buffer)
	width := enc.nextLengthHint()
	enc.startBytes(field, width)
	messageStart := len(enc.buffer)
	lengthStart := messageStart - width
	// encode the submessage
	ok := enc.enterMessage()
	if ok {
		ok = fn(enc)
		enc.depth--
	}
	if !ok {
		// The message was nil or too deeply nested, we can remove the tag.
		enc.buffer = enc.buffer[:tagStart]
		return
	}
	if !enc.putLength(lengthStart, messageStart) {
		enc.finishBytes(field, tagStart, lengthStart, messageStart)
	}
}

// AlwaysMessage encodes a message field, also when fn reports that the message
// is absent.
//
//go:noinline
func (enc *Encoder) AlwaysMessage(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.alwaysAnyBytes(field, func() { fn(enc) })
}

// PresentMessage encodes a message field that is always present, unless its
// encoding is empty.
//
//go:noinline
func (enc *Encoder) PresentMessage(field FieldNumber, fn func(enc *Encoder) bool) {
	enc.Message(field, func(enc *Encoder) bool {
		lengthStart := enc.offset()
		fn(enc)
		return enc.offset() > lengthStart
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
		start := enc.offset()
		fn(enc)
		return enc.offset() > start
	})
}

func (enc *Encoder) anyGroup(field FieldNumber, fn func(enc *Encoder) bool) bool {
	start := enc.offset()
	enc.appendTag(field, protowire.StartGroupType)
	if !enc.enterMessage() {
		enc.truncate(start)
		return false
	}
	ok := fn(enc)
	enc.depth--
	if !ok {
		enc.truncate(start)
		return false
	}
	enc.appendTag(field, protowire.EndGroupType)
	return true
}

func (enc *Encoder) alwaysGroup(field FieldNumber, fn func()) {
	enc.appendTag(field, protowire.StartGroupType)
	if enc.enterMessage() {
		fn()
		enc.depth--
	}
	enc.appendTag(field, protowire.EndGroupType)
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
			enc.appendVarint(uint64(fn(uint(i))))
		}
	})
}

// AlwaysAnyBytes encodes field with the bytes wire type. fn writes the
// contents, and AlwaysAnyBytes writes their length.
func (enc *Encoder) AlwaysAnyBytes(field FieldNumber, fn func()) {
	enc.alwaysAnyBytes(field, fn)
}

// alwaysAnyBytes encodes field as Bytes and handles encoding the length.
func (enc *Encoder) alwaysAnyBytes(field FieldNumber, fn func()) {
	if enc.sizing {
		hint := enc.reserveLengthHint()
		start := enc.size
		if enc.enterMessage() {
			fn()
			enc.depth--
		}
		messageLength := enc.size - start
		width := protowire.SizeVarint(uint64(messageLength))
		if hint >= 0 {
			enc.setLengthHint(hint, width)
		}
		enc.size += protowire.SizeTag(protowire.Number(field)) + width
		return
	}
	tagStart := len(enc.buffer)
	width := enc.nextLengthHint()
	enc.startBytes(field, width)
	messageStart := len(enc.buffer)
	lengthStart := messageStart - width
	// encode the submessage
	if enc.enterMessage() {
		fn()
		enc.depth--
	}
	if !enc.putLength(lengthStart, messageStart) {
		enc.finishBytes(field, tagStart, lengthStart, messageStart)
	}
}

// inlineLengthHints is the number of four-bit widths in the two inline words.
// Four bits cover every possible uint64 varint width.
const inlineLengthHints = 32

// reserveLengthHint reserves a preorder slot before visiting children. Zero
// means omitted; keep its slot and any child slots even if the callback returns
// false, since encoding will still visit those callbacks.
func (enc *Encoder) reserveLengthHint() int {
	if enc.hintCount == 0 {
		return -1
	}
	if enc.hintCount <= inlineLengthHints {
		index := int(enc.hintCount) - 1
		enc.hintCount++
		return index
	}
	if enc.lengthHints == nil {
		enc.lengthHints = new(lengthHints)
		enc.lengthHints.widths = enc.lengthHints.inline[:0]
	}
	index := len(enc.lengthHints.widths)
	enc.lengthHints.widths = append(enc.lengthHints.widths, 0)
	return inlineLengthHints + index
}

func (enc *Encoder) setLengthHint(index, width int) {
	if index < inlineLengthHints {
		enc.inlineHints[index/16] |= uint64(width) << (4 * (index % 16))
	} else {
		enc.lengthHints.widths[index-inlineLengthHints] = byte(width)
	}
}

func (enc *Encoder) nextLengthHint() int {
	if index := int(enc.hintRead); index+1 < int(enc.hintCount) {
		width := int(enc.inlineHints[index/16] >> (4 * (index % 16)) & 15)
		enc.hintRead++
		return width
	}
	if hints := enc.lengthHints; hints != nil && hints.read < len(hints.widths) {
		width := int(hints.widths[hints.read])
		hints.read++
		return width
	}
	// Without hints, predict a two-byte length, which fits payloads up to
	// 16 KiB.
	return 2
}

func (enc *Encoder) startBytes(field FieldNumber, width int) {
	// Avoid growing a full buffer for a message expected to be omitted.
	if width != 0 {
		enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)
		enc.buffer = append(enc.buffer, make([]byte, width)...)
	}
}

// putLength handles a matching prediction without entering the correction path.
func (enc *Encoder) putLength(lengthStart, messageStart int) bool {
	length := len(enc.buffer) - messageStart
	if protowire.SizeVarint(uint64(length)) != messageStart-lengthStart {
		return false
	}
	protowire.PutUvarint(enc.buffer[lengthStart:messageStart], uint64(length))
	return true
}

func (enc *Encoder) finishBytes(field FieldNumber, tagStart, lengthStart, messageStart int) {
	messageLength := len(enc.buffer) - messageStart
	width := protowire.SizeVarint(uint64(messageLength))
	missingTag := lengthStart == tagStart
	if missingTag {
		lengthStart += protowire.SizeTag(protowire.Number(field))
	}
	shift := lengthStart + width - messageStart
	if shift != 0 {
		if shift > 0 {
			enc.buffer = append(enc.buffer, make([]byte, shift)...)
		}
		copy(enc.buffer[lengthStart+width:], enc.buffer[messageStart:messageStart+messageLength])
		enc.buffer = enc.buffer[:lengthStart+width+messageLength]
	}
	if missingTag {
		_ = appendTag(enc.buffer[:tagStart], field, protowire.BytesType)
	}
	protowire.PutUvarint(enc.buffer[lengthStart:lengthStart+width], uint64(messageLength))
}

// UnrecognizedFields writes out, which holds the unrecognized fields that
// decoding captured.
func (enc *Encoder) UnrecognizedFields(out []byte) {
	if enc.sizing {
		enc.size += len(out)
	} else {
		enc.buffer = append(enc.buffer, out...)
	}
}

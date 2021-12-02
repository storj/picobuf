// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protowire"
)

const (
	fieldDecodingErrored = FieldNumber(-1)
	fieldDecodingDone    = FieldNumber(-2)
)

// Decoder implements decoding of protobuf messages.
type Decoder struct {
	messageDecodeState
	stack []messageDecodeState

	codec *Codec
	err   error
}

type messageDecodeState struct {
	pendingField FieldNumber
	pendingWire  protowire.Type

	buffer []byte
}

// NewDecoder returns a new Decoder.
func NewDecoder(data []byte) *Decoder {
	codec := &Codec{
		encoding: false,
	}
	codec.decode.codec = codec
	codec.decode.buffer = data
	codec.decode.nextField(0)
	return &codec.decode
}

// Codec returns the associated codec.
func (dec *Decoder) Codec() *Codec { return dec.codec }

// Err returns error that occurred during decoding.
func (dec *Decoder) Err() error {
	return dec.err
}

func (dec *Decoder) pushState(message []byte) {
	dec.stack = append(dec.stack, dec.messageDecodeState)
	dec.messageDecodeState = messageDecodeState{
		buffer: message,
	}
	dec.nextField(0)
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
func (dec *Decoder) RepeatedMessage(field FieldNumber, fn func(c *Codec, index int) bool) {
	for field == dec.pendingField {
		if dec.pendingWire != protowire.BytesType {
			dec.fail(field, "expected wire type Bytes")
			return
		}

		message, n := protowire.ConsumeBytes(dec.buffer)
		dec.pushState(message)
		mode := -1
		dec.Loop(func(c *Codec) bool {
			fn(dec.codec, mode)
			mode = -2
			return true
		})
		dec.popState()

		dec.nextField(n)
	}
}

// Message decodes a message.
func (dec *Decoder) Message(field FieldNumber, fn func(*Codec) bool) {
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
func (dec *Decoder) PresentMessage(field FieldNumber, fn func(*Codec) bool) {
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

// Loop loops fields until all messages have been processed.
func (dec *Decoder) Loop(fn func(*Codec) bool) {
	for {
		startingLength := len(dec.buffer)
		fn(dec.codec)
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

//go:noinline
func (dec *Decoder) fail(field FieldNumber, msg string) {
	// TODO: use static error types
	dec.pendingField = fieldDecodingErrored
	dec.err = fmt.Errorf("failed while parsing %v: %s", field, msg)
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
	dec.buffer = dec.buffer[n:]
	dec.pendingField, dec.pendingWire = FieldNumber(field), wire
}

func init() {
	// silence structcheck linting bug about unused field
	var z messageDecodeState
	z.pendingField = 0
	z.pendingWire = 0
}

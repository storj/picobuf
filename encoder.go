// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"encoding/binary"

	"google.golang.org/protobuf/encoding/protowire"
)

// Encoder implements encoding of protobuf format.
type Encoder struct {
	buffer []byte
	codec  *Codec
}

// NewEncoder creates a new Encoder.
func NewEncoder() *Encoder {
	codec := &Codec{
		encoding: true,
	}
	codec.encode.codec = codec
	return &codec.encode
}

// Codec returns the associated codec.
func (enc *Encoder) Codec() *Codec { return enc.codec }

// Buffer returns the encoded internal buffer.
func (enc *Encoder) Buffer() []byte { return enc.buffer }

// Message encodes an embedded message.
func (enc *Encoder) Message(field FieldNumber, fn func(*Codec)) {
	tagStart := len(enc.buffer)
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	messageStart := len(enc.buffer)
	// add small buffer for encoding the length to avoid overlapping append
	var lengthBuffer [binary.MaxVarintLen64]byte
	enc.buffer = append(enc.buffer, lengthBuffer[:]...)
	paddingStart := len(enc.buffer)
	// encode the submessage
	fn(enc.codec)
	if paddingStart == len(enc.buffer) {
		enc.buffer = enc.buffer[:tagStart]
		return
	}
	// add the message as bytes
	enc.buffer = protowire.AppendBytes(enc.buffer[:messageStart], enc.buffer[paddingStart:])
}

func encodeZigZag32(v int32) uint32 { return (uint32(v) << 1) ^ (uint32(v) >> 31) }

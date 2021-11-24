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
	lengthStart := len(enc.buffer)
	// We'll guess that we need 2 bytes for length.
	// If we need less, then the copy is fast, and needing more is unlikely.
	var lengthBufferPrediction [2]byte
	enc.buffer = append(enc.buffer, lengthBufferPrediction[:]...)
	messageStart := len(enc.buffer)
	// encode the submessage
	fn(enc.codec)
	messageLength := len(enc.buffer) - messageStart
	if messageLength == 0 {
		// The message was empty, we can remove the tag.
		enc.buffer = enc.buffer[:tagStart]
		return
	}
	bytesForSize := protowire.SizeVarint(uint64(messageLength))
	if bytesForSize == len(lengthBufferPrediction) {
		binary.PutUvarint(enc.buffer[lengthStart:messageStart], uint64(messageLength))
		return
	}
	if bytesForSize > len(lengthBufferPrediction) {
		enc.buffer = append(enc.buffer, make([]byte, bytesForSize-len(lengthBufferPrediction))...)
	}
	copy(enc.buffer[lengthStart+bytesForSize:], enc.buffer[messageStart:])
	binary.PutUvarint(enc.buffer[lengthStart:lengthStart+bytesForSize], uint64(messageLength))
	enc.buffer = enc.buffer[:lengthStart+bytesForSize+messageLength]
}

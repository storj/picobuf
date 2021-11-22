// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"encoding/binary"
	"math"

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

// RequiredByte encodes a byte field.
func (enc *Encoder) RequiredByte(field FieldNumber, v *byte) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// RequiredBool encodes a bool field.
func (enc *Encoder) RequiredBool(field FieldNumber, v *bool) {
	if !*v {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(1))
}

// RequiredInt32 encodes a int32 field.
func (enc *Encoder) RequiredInt32(field FieldNumber, v *int32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// RequiredInt64 encodes a int64 field.
func (enc *Encoder) RequiredInt64(field FieldNumber, v *int64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// RequiredUint32 encodes a uint32 field.
func (enc *Encoder) RequiredUint32(field FieldNumber, v *uint32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// RequiredUint64 encodes a uint64 field.
func (enc *Encoder) RequiredUint64(field FieldNumber, v *uint64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, *v)
}

// RequiredSint32 encodes a zig-zag int32 field.
func (enc *Encoder) RequiredSint32(field FieldNumber, v *int32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(encodeZigZag32(*v)))
}

// RequiredSint64 encodes a zig-zag int64 field.
func (enc *Encoder) RequiredSint64(field FieldNumber, v *int64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, protowire.EncodeZigZag(*v))
}

// RequiredFixed32 encodes a fixed uint32 field.
func (enc *Encoder) RequiredFixed32(field FieldNumber, v *uint32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed32Type)
	enc.buffer = protowire.AppendFixed32(enc.buffer, *v)
}

// RequiredFixed64 encodes a fixed uint64 field.
func (enc *Encoder) RequiredFixed64(field FieldNumber, v *uint64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed64Type)
	enc.buffer = protowire.AppendFixed64(enc.buffer, *v)
}

// RequiredSfixed32 encodes a zig-zag fixed int32 field.
func (enc *Encoder) RequiredSfixed32(field FieldNumber, v *int32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed32Type)
	enc.buffer = protowire.AppendFixed32(enc.buffer, encodeZigZag32(*v))
}

// RequiredSfixed64 encodes a zig-zag fixed int64 field.
func (enc *Encoder) RequiredSfixed64(field FieldNumber, v *int64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed64Type)
	enc.buffer = protowire.AppendFixed64(enc.buffer, protowire.EncodeZigZag(*v))
}

// RequiredFloat encodes a float32 field.
func (enc *Encoder) RequiredFloat(field FieldNumber, v *float32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed32Type)
	enc.buffer = protowire.AppendFixed32(enc.buffer, math.Float32bits(*v))
}

// RequiredDouble encodes a float64 field.
func (enc *Encoder) RequiredDouble(field FieldNumber, v *float64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed64Type)
	enc.buffer = protowire.AppendFixed64(enc.buffer, math.Float64bits(*v))
}

// RequiredString encodes a string field.
func (enc *Encoder) RequiredString(field FieldNumber, v *string) {
	if len(*v) == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	enc.buffer = protowire.AppendString(enc.buffer, *v)
}

// RequiredRawString encodes a string field.
func (enc *Encoder) RequiredRawString(field FieldNumber, v *string) {
	if len(*v) == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	enc.buffer = protowire.AppendString(enc.buffer, *v)
}

// RequiredBytes encodes a []byte field.
func (enc *Encoder) RequiredBytes(field FieldNumber, v *[]byte) {
	if len(*v) == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	enc.buffer = protowire.AppendBytes(enc.buffer, *v)
}

// RequiredMessage encodes an embedded message.
func (enc *Encoder) RequiredMessage(field FieldNumber, fn func(*Codec)) {
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

// Bytes returns the encoded internal buffer.
func (enc *Encoder) Bytes() []byte { return enc.buffer }

func encodeZigZag32(v int32) uint32 { return (uint32(v) << 1) ^ (uint32(v) >> 31) }

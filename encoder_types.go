// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
)

// Byte encodes non-default byte protobuf type.
func (enc *Encoder) Byte(field FieldNumber, v *uint8) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// Bool encodes non-default bool protobuf type.
func (enc *Encoder) Bool(field FieldNumber, v *bool) {
	if !*v {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(1))
}

// Int32 encodes non-default int32 protobuf type.
func (enc *Encoder) Int32(field FieldNumber, v *int32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// Int64 encodes non-default int64 protobuf type.
func (enc *Encoder) Int64(field FieldNumber, v *int64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// Uint32 encodes non-default uint32 protobuf type.
func (enc *Encoder) Uint32(field FieldNumber, v *uint32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(*v))
}

// Uint64 encodes non-default uint64 protobuf type.
func (enc *Encoder) Uint64(field FieldNumber, v *uint64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, *v)
}

// Sint32 encodes non-default sint32 protobuf type.
func (enc *Encoder) Sint32(field FieldNumber, v *int32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, uint64(encodeZigZag32(*v)))
}

// Sint64 encodes non-default sint64 protobuf type.
func (enc *Encoder) Sint64(field FieldNumber, v *int64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.VarintType)
	enc.buffer = protowire.AppendVarint(enc.buffer, protowire.EncodeZigZag(*v))
}

// Fixed32 encodes non-default fixed32 protobuf type.
func (enc *Encoder) Fixed32(field FieldNumber, v *uint32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed32Type)
	enc.buffer = protowire.AppendFixed32(enc.buffer, *v)
}

// Fixed64 encodes non-default fixed64 protobuf type.
func (enc *Encoder) Fixed64(field FieldNumber, v *uint64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed64Type)
	enc.buffer = protowire.AppendFixed64(enc.buffer, *v)
}

// Sfixed32 encodes non-default sfixed32 protobuf type.
func (enc *Encoder) Sfixed32(field FieldNumber, v *int32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed32Type)
	enc.buffer = protowire.AppendFixed32(enc.buffer, encodeZigZag32(*v))
}

// Sfixed64 encodes non-default sfixed64 protobuf type.
func (enc *Encoder) Sfixed64(field FieldNumber, v *int64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed64Type)
	enc.buffer = protowire.AppendFixed64(enc.buffer, protowire.EncodeZigZag(*v))
}

// Float encodes non-default float protobuf type.
func (enc *Encoder) Float(field FieldNumber, v *float32) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed32Type)
	enc.buffer = protowire.AppendFixed32(enc.buffer, math.Float32bits(*v))
}

// Double encodes non-default double protobuf type.
func (enc *Encoder) Double(field FieldNumber, v *float64) {
	if *v == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.Fixed64Type)
	enc.buffer = protowire.AppendFixed64(enc.buffer, math.Float64bits(*v))
}

// String encodes non-default string protobuf type.
func (enc *Encoder) String(field FieldNumber, v *string) {
	if len(*v) == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	enc.buffer = protowire.AppendString(enc.buffer, *v)
}

// RawString encodes non-default rawstring protobuf type.
func (enc *Encoder) RawString(field FieldNumber, v *string) {
	if len(*v) == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	enc.buffer = protowire.AppendString(enc.buffer, *v)
}

// Bytes encodes non-default bytes protobuf type.
func (enc *Encoder) Bytes(field FieldNumber, v *[]byte) {
	if len(*v) == 0 {
		return
	}
	enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), protowire.BytesType)
	enc.buffer = protowire.AppendBytes(enc.buffer, *v)
}

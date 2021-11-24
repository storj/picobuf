// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
)

// Byte decodes byte protobuf type.
func (dec *Decoder) Byte(field FieldNumber, v *uint8) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0x0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = byte(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Bool decodes bool protobuf type.
func (dec *Decoder) Bool(field FieldNumber, v *bool) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = false
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = x == 1
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Int32 decodes int32 protobuf type.
func (dec *Decoder) Int32(field FieldNumber, v *int32) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = int32(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Int64 decodes int64 protobuf type.
func (dec *Decoder) Int64(field FieldNumber, v *int64) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = int64(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Uint32 decodes uint32 protobuf type.
func (dec *Decoder) Uint32(field FieldNumber, v *uint32) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0x0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = uint32(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Uint64 decodes uint64 protobuf type.
func (dec *Decoder) Uint64(field FieldNumber, v *uint64) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0x0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = x
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Sint32 decodes sint32 protobuf type.
func (dec *Decoder) Sint32(field FieldNumber, v *int32) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = decodeZigZag32(uint32(x))
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Sint64 decodes sint64 protobuf type.
func (dec *Decoder) Sint64(field FieldNumber, v *int64) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.VarintType {
		dec.fail(field, "expected wire type Varint")
		return
	}
	x, n := protowire.ConsumeVarint(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Varint")
		return
	}
	*v = protowire.DecodeZigZag(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Fixed32 decodes fixed32 protobuf type.
func (dec *Decoder) Fixed32(field FieldNumber, v *uint32) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0x0
		}
		return
	}
	if dec.pendingWire != protowire.Fixed32Type {
		dec.fail(field, "expected wire type Fixed32")
		return
	}
	x, n := protowire.ConsumeFixed32(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Fixed32")
		return
	}
	*v = x
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Fixed64 decodes fixed64 protobuf type.
func (dec *Decoder) Fixed64(field FieldNumber, v *uint64) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0x0
		}
		return
	}
	if dec.pendingWire != protowire.Fixed64Type {
		dec.fail(field, "expected wire type Fixed64")
		return
	}
	x, n := protowire.ConsumeFixed64(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Fixed64")
		return
	}
	*v = x
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Sfixed32 decodes sfixed32 protobuf type.
func (dec *Decoder) Sfixed32(field FieldNumber, v *int32) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.Fixed32Type {
		dec.fail(field, "expected wire type Fixed32")
		return
	}
	x, n := protowire.ConsumeFixed32(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Fixed32")
		return
	}
	*v = decodeZigZag32(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Sfixed64 decodes sfixed64 protobuf type.
func (dec *Decoder) Sfixed64(field FieldNumber, v *int64) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.Fixed64Type {
		dec.fail(field, "expected wire type Fixed64")
		return
	}
	x, n := protowire.ConsumeFixed64(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Fixed64")
		return
	}
	*v = protowire.DecodeZigZag(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Float decodes float protobuf type.
func (dec *Decoder) Float(field FieldNumber, v *float32) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.Fixed32Type {
		dec.fail(field, "expected wire type Fixed32")
		return
	}
	x, n := protowire.ConsumeFixed32(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Fixed32")
		return
	}
	*v = math.Float32frombits(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Double decodes double protobuf type.
func (dec *Decoder) Double(field FieldNumber, v *float64) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = 0
		}
		return
	}
	if dec.pendingWire != protowire.Fixed64Type {
		dec.fail(field, "expected wire type Fixed64")
		return
	}
	x, n := protowire.ConsumeFixed64(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse Fixed64")
		return
	}
	*v = math.Float64frombits(x)
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// String decodes string protobuf type.
func (dec *Decoder) String(field FieldNumber, v *string) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = ""
		}
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}
	x, n := protowire.ConsumeString(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse String")
		return
	}
	*v = x
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// RawString decodes rawstring protobuf type.
func (dec *Decoder) RawString(field FieldNumber, v *string) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = ""
		}
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}
	x, n := protowire.ConsumeString(dec.buffer)
	if n < 0 {
		dec.fail(field, "unable to parse String")
		return
	}
	*v = x
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// Bytes decodes bytes protobuf type.
func (dec *Decoder) Bytes(field FieldNumber, v *[]byte) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			*v = nil
		}
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
	*v = x
	dec.nextField(n)
	dec.filled.Set(int32(field))
}

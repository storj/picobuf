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
}

// RepeatedByte decodes repeated byte protobuf type.
func (dec *Decoder) RepeatedByte(field FieldNumber, v *[]uint8) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Bool decodes bool protobuf type.
func (dec *Decoder) Bool(field FieldNumber, v *bool) {
	if field != dec.pendingField {
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
}

// RepeatedBool decodes repeated bool protobuf type.
func (dec *Decoder) RepeatedBool(field FieldNumber, v *[]bool) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Int32 decodes int32 protobuf type.
func (dec *Decoder) Int32(field FieldNumber, v *int32) {
	if field != dec.pendingField {
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
}

// RepeatedInt32 decodes repeated int32 protobuf type.
func (dec *Decoder) RepeatedInt32(field FieldNumber, v *[]int32) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Int64 decodes int64 protobuf type.
func (dec *Decoder) Int64(field FieldNumber, v *int64) {
	if field != dec.pendingField {
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
}

// RepeatedInt64 decodes repeated int64 protobuf type.
func (dec *Decoder) RepeatedInt64(field FieldNumber, v *[]int64) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Uint32 decodes uint32 protobuf type.
func (dec *Decoder) Uint32(field FieldNumber, v *uint32) {
	if field != dec.pendingField {
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
}

// RepeatedUint32 decodes repeated uint32 protobuf type.
func (dec *Decoder) RepeatedUint32(field FieldNumber, v *[]uint32) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Uint64 decodes uint64 protobuf type.
func (dec *Decoder) Uint64(field FieldNumber, v *uint64) {
	if field != dec.pendingField {
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
}

// RepeatedUint64 decodes repeated uint64 protobuf type.
func (dec *Decoder) RepeatedUint64(field FieldNumber, v *[]uint64) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Sint32 decodes sint32 protobuf type.
func (dec *Decoder) Sint32(field FieldNumber, v *int32) {
	if field != dec.pendingField {
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
}

// RepeatedSint32 decodes repeated sint32 protobuf type.
func (dec *Decoder) RepeatedSint32(field FieldNumber, v *[]int32) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Sint64 decodes sint64 protobuf type.
func (dec *Decoder) Sint64(field FieldNumber, v *int64) {
	if field != dec.pendingField {
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
}

// RepeatedSint64 decodes repeated sint64 protobuf type.
func (dec *Decoder) RepeatedSint64(field FieldNumber, v *[]int64) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Fixed32 decodes fixed32 protobuf type.
func (dec *Decoder) Fixed32(field FieldNumber, v *uint32) {
	if field != dec.pendingField {
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
}

// RepeatedFixed32 decodes repeated fixed32 protobuf type.
func (dec *Decoder) RepeatedFixed32(field FieldNumber, v *[]uint32) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Fixed64 decodes fixed64 protobuf type.
func (dec *Decoder) Fixed64(field FieldNumber, v *uint64) {
	if field != dec.pendingField {
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
}

// RepeatedFixed64 decodes repeated fixed64 protobuf type.
func (dec *Decoder) RepeatedFixed64(field FieldNumber, v *[]uint64) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Sfixed32 decodes sfixed32 protobuf type.
func (dec *Decoder) Sfixed32(field FieldNumber, v *int32) {
	if field != dec.pendingField {
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
}

// RepeatedSfixed32 decodes repeated sfixed32 protobuf type.
func (dec *Decoder) RepeatedSfixed32(field FieldNumber, v *[]int32) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Sfixed64 decodes sfixed64 protobuf type.
func (dec *Decoder) Sfixed64(field FieldNumber, v *int64) {
	if field != dec.pendingField {
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
}

// RepeatedSfixed64 decodes repeated sfixed64 protobuf type.
func (dec *Decoder) RepeatedSfixed64(field FieldNumber, v *[]int64) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Float decodes float protobuf type.
func (dec *Decoder) Float(field FieldNumber, v *float32) {
	if field != dec.pendingField {
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
}

// RepeatedFloat decodes repeated float protobuf type.
func (dec *Decoder) RepeatedFloat(field FieldNumber, v *[]float32) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Double decodes double protobuf type.
func (dec *Decoder) Double(field FieldNumber, v *float64) {
	if field != dec.pendingField {
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
}

// RepeatedDouble decodes repeated double protobuf type.
func (dec *Decoder) RepeatedDouble(field FieldNumber, v *[]float64) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// String decodes string protobuf type.
func (dec *Decoder) String(field FieldNumber, v *string) {
	if field != dec.pendingField {
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
}

// RepeatedString decodes repeated string protobuf type.
func (dec *Decoder) RepeatedString(field FieldNumber, v *[]string) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// RawString decodes rawstring protobuf type.
func (dec *Decoder) RawString(field FieldNumber, v *string) {
	if field != dec.pendingField {
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
}

// RepeatedRawString decodes repeated rawstring protobuf type.
func (dec *Decoder) RepeatedRawString(field FieldNumber, v *[]string) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

// Bytes decodes bytes protobuf type.
func (dec *Decoder) Bytes(field FieldNumber, v *[]byte) {
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
	*v = x
	dec.nextField(n)
}

// RepeatedBytes decodes repeated bytes protobuf type.
func (dec *Decoder) RepeatedBytes(field FieldNumber, v *[][]byte) {
	if field != dec.pendingField {
		return
	}
	panic("not implemented")
}

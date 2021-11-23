// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

// Byte decodes byte protobuf type.
func (dec *Decoder) Byte(field FieldNumber, v *uint8) {}

// Bool decodes bool protobuf type.
func (dec *Decoder) Bool(field FieldNumber, v *bool) {}

// Int32 decodes int32 protobuf type.
func (dec *Decoder) Int32(field FieldNumber, v *int32) {}

// Int64 decodes int64 protobuf type.
func (dec *Decoder) Int64(field FieldNumber, v *int64) {}

// Uint32 decodes uint32 protobuf type.
func (dec *Decoder) Uint32(field FieldNumber, v *uint32) {}

// Uint64 decodes uint64 protobuf type.
func (dec *Decoder) Uint64(field FieldNumber, v *uint64) {}

// Sint32 decodes sint32 protobuf type.
func (dec *Decoder) Sint32(field FieldNumber, v *int32) {}

// Sint64 decodes sint64 protobuf type.
func (dec *Decoder) Sint64(field FieldNumber, v *int64) {}

// Fixed32 decodes fixed32 protobuf type.
func (dec *Decoder) Fixed32(field FieldNumber, v *uint32) {}

// Fixed64 decodes fixed64 protobuf type.
func (dec *Decoder) Fixed64(field FieldNumber, v *uint64) {}

// Sfixed32 decodes sfixed32 protobuf type.
func (dec *Decoder) Sfixed32(field FieldNumber, v *int32) {}

// Sfixed64 decodes sfixed64 protobuf type.
func (dec *Decoder) Sfixed64(field FieldNumber, v *int64) {}

// Float decodes float protobuf type.
func (dec *Decoder) Float(field FieldNumber, v *float32) {}

// Double decodes double protobuf type.
func (dec *Decoder) Double(field FieldNumber, v *float64) {}

// String decodes string protobuf type.
func (dec *Decoder) String(field FieldNumber, v *string) {}

// RawString decodes rawstring protobuf type.
func (dec *Decoder) RawString(field FieldNumber, v *string) {}

// Bytes decodes bytes protobuf type.
func (dec *Decoder) Bytes(field FieldNumber, v *[]byte) {}

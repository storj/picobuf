// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"storj.io/picobuf"
)

// This file converts the google.protobuf wrapper types to plain Go types. A
// wrapper field declared as
//
//	google.protobuf.BoolValue is_valid = 1 [
//	    (pico.field).custom_type = "bool",
//	    (pico.field).custom_serialize = "storj.io/picobuf/picoconv.BoolValue"
//	];
//
// generates a *bool field, where nil means that the wrapper is absent.

// BoolValue converts between google.protobuf.BoolValue and bool.
type BoolValue bool

// PicoEncode implements picobuf.CustomType.
func (v *BoolValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := bool(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Bool(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *BoolValue) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := bool(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Bool(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = BoolValue(value)
}

// Int32Value converts between google.protobuf.Int32Value and int32.
type Int32Value int32

// PicoEncode implements picobuf.CustomType.
func (v *Int32Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := int32(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int32(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *Int32Value) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := int32(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Int32(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = Int32Value(value)
}

// Int64Value converts between google.protobuf.Int64Value and int64.
type Int64Value int64

// PicoEncode implements picobuf.CustomType.
func (v *Int64Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := int64(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *Int64Value) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := int64(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Int64(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = Int64Value(value)
}

// UInt32Value converts between google.protobuf.UInt32Value and uint32.
type UInt32Value uint32

// PicoEncode implements picobuf.CustomType.
func (v *UInt32Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := uint32(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Uint32(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *UInt32Value) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := uint32(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Uint32(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = UInt32Value(value)
}

// UInt64Value converts between google.protobuf.UInt64Value and uint64.
type UInt64Value uint64

// PicoEncode implements picobuf.CustomType.
func (v *UInt64Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := uint64(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Uint64(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *UInt64Value) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := uint64(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Uint64(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = UInt64Value(value)
}

// FloatValue converts between google.protobuf.FloatValue and float32.
type FloatValue float32

// PicoEncode implements picobuf.CustomType.
func (v *FloatValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := float32(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Float(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *FloatValue) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := float32(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Float(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = FloatValue(value)
}

// DoubleValue converts between google.protobuf.DoubleValue and float64.
type DoubleValue float64

// PicoEncode implements picobuf.CustomType.
func (v *DoubleValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := float64(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Double(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *DoubleValue) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := float64(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Double(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = DoubleValue(value)
}

// StringValue converts between google.protobuf.StringValue and string.
type StringValue string

// PicoEncode implements picobuf.CustomType.
func (v *StringValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := string(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.String(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *StringValue) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := string(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.String(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = StringValue(value)
}

// BytesValue converts between google.protobuf.BytesValue and []byte.
type BytesValue []byte

// PicoEncode implements picobuf.CustomType.
func (v *BytesValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if v == nil {
		return
	}
	value := []byte(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Bytes(1, &value)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (v *BytesValue) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}
	value := []byte(*v) // merge with an earlier occurrence
	c.Message(field, func(c *picobuf.Decoder) {
		c.Bytes(1, &value)
	})
	if c.Err() != nil {
		return
	}
	*v = BytesValue(value)
}

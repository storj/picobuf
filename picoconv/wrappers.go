// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"storj.io/picobuf"
)

// This file implements the google.protobuf wrapper types as conversions to
// plain Go types, so that a wrapper field can be declared as e.g.
//
//	google.protobuf.BoolValue is_valid = 1 [
//	    (pico.field).custom_type = "bool",
//	    (pico.field).custom_serialize = "storj.io/picobuf/picoconv.BoolValue"
//	];
//
// which generates a *bool field, where nil means the wrapper is absent.

// BoolValue implements google.protobuf.BoolValue conversion to bool.
type BoolValue bool

// PicoEncode implements custom encoding function.
func (v *BoolValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := bool(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Bool(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// Int32Value implements google.protobuf.Int32Value conversion to int32.
type Int32Value int32

// PicoEncode implements custom encoding function.
func (v *Int32Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := int32(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int32(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// Int64Value implements google.protobuf.Int64Value conversion to int64.
type Int64Value int64

// PicoEncode implements custom encoding function.
func (v *Int64Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := int64(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// UInt32Value implements google.protobuf.UInt32Value conversion to uint32.
type UInt32Value uint32

// PicoEncode implements custom encoding function.
func (v *UInt32Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := uint32(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Uint32(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// UInt64Value implements google.protobuf.UInt64Value conversion to uint64.
type UInt64Value uint64

// PicoEncode implements custom encoding function.
func (v *UInt64Value) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := uint64(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Uint64(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// FloatValue implements google.protobuf.FloatValue conversion to float32.
type FloatValue float32

// PicoEncode implements custom encoding function.
func (v *FloatValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := float32(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Float(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// DoubleValue implements google.protobuf.DoubleValue conversion to float64.
type DoubleValue float64

// PicoEncode implements custom encoding function.
func (v *DoubleValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := float64(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Double(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// StringValue implements google.protobuf.StringValue conversion to string.
type StringValue string

// PicoEncode implements custom encoding function.
func (v *StringValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := string(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.String(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

// BytesValue implements google.protobuf.BytesValue conversion to []byte.
type BytesValue []byte

// PicoEncode implements custom encoding function.
func (v *BytesValue) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if v == nil {
		return false
	}
	value := []byte(*v)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Bytes(1, &value)
		return true
	})
	return true
}

// PicoDecode implements custom decoding function.
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

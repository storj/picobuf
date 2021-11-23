// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

// Byte codes byte protobuf type.
//go:noinline
func (codec *Codec) Byte(field FieldNumber, v *uint8) {
	if codec.encoding {
		codec.encode.Byte(field, v)
	} else {
		codec.decode.Byte(field, v)
	}
}

// Bool codes bool protobuf type.
//go:noinline
func (codec *Codec) Bool(field FieldNumber, v *bool) {
	if codec.encoding {
		codec.encode.Bool(field, v)
	} else {
		codec.decode.Bool(field, v)
	}
}

// Int32 codes int32 protobuf type.
//go:noinline
func (codec *Codec) Int32(field FieldNumber, v *int32) {
	if codec.encoding {
		codec.encode.Int32(field, v)
	} else {
		codec.decode.Int32(field, v)
	}
}

// Int64 codes int64 protobuf type.
//go:noinline
func (codec *Codec) Int64(field FieldNumber, v *int64) {
	if codec.encoding {
		codec.encode.Int64(field, v)
	} else {
		codec.decode.Int64(field, v)
	}
}

// Uint32 codes uint32 protobuf type.
//go:noinline
func (codec *Codec) Uint32(field FieldNumber, v *uint32) {
	if codec.encoding {
		codec.encode.Uint32(field, v)
	} else {
		codec.decode.Uint32(field, v)
	}
}

// Uint64 codes uint64 protobuf type.
//go:noinline
func (codec *Codec) Uint64(field FieldNumber, v *uint64) {
	if codec.encoding {
		codec.encode.Uint64(field, v)
	} else {
		codec.decode.Uint64(field, v)
	}
}

// Sint32 codes sint32 protobuf type.
//go:noinline
func (codec *Codec) Sint32(field FieldNumber, v *int32) {
	if codec.encoding {
		codec.encode.Sint32(field, v)
	} else {
		codec.decode.Sint32(field, v)
	}
}

// Sint64 codes sint64 protobuf type.
//go:noinline
func (codec *Codec) Sint64(field FieldNumber, v *int64) {
	if codec.encoding {
		codec.encode.Sint64(field, v)
	} else {
		codec.decode.Sint64(field, v)
	}
}

// Fixed32 codes fixed32 protobuf type.
//go:noinline
func (codec *Codec) Fixed32(field FieldNumber, v *uint32) {
	if codec.encoding {
		codec.encode.Fixed32(field, v)
	} else {
		codec.decode.Fixed32(field, v)
	}
}

// Fixed64 codes fixed64 protobuf type.
//go:noinline
func (codec *Codec) Fixed64(field FieldNumber, v *uint64) {
	if codec.encoding {
		codec.encode.Fixed64(field, v)
	} else {
		codec.decode.Fixed64(field, v)
	}
}

// Sfixed32 codes sfixed32 protobuf type.
//go:noinline
func (codec *Codec) Sfixed32(field FieldNumber, v *int32) {
	if codec.encoding {
		codec.encode.Sfixed32(field, v)
	} else {
		codec.decode.Sfixed32(field, v)
	}
}

// Sfixed64 codes sfixed64 protobuf type.
//go:noinline
func (codec *Codec) Sfixed64(field FieldNumber, v *int64) {
	if codec.encoding {
		codec.encode.Sfixed64(field, v)
	} else {
		codec.decode.Sfixed64(field, v)
	}
}

// Float codes float protobuf type.
//go:noinline
func (codec *Codec) Float(field FieldNumber, v *float32) {
	if codec.encoding {
		codec.encode.Float(field, v)
	} else {
		codec.decode.Float(field, v)
	}
}

// Double codes double protobuf type.
//go:noinline
func (codec *Codec) Double(field FieldNumber, v *float64) {
	if codec.encoding {
		codec.encode.Double(field, v)
	} else {
		codec.decode.Double(field, v)
	}
}

// String codes string protobuf type.
//go:noinline
func (codec *Codec) String(field FieldNumber, v *string) {
	if codec.encoding {
		codec.encode.String(field, v)
	} else {
		codec.decode.String(field, v)
	}
}

// RawString codes rawstring protobuf type.
//go:noinline
func (codec *Codec) RawString(field FieldNumber, v *string) {
	if codec.encoding {
		codec.encode.RawString(field, v)
	} else {
		codec.decode.RawString(field, v)
	}
}

// Bytes codes bytes protobuf type.
//go:noinline
func (codec *Codec) Bytes(field FieldNumber, v *[]byte) {
	if codec.encoding {
		codec.encode.Bytes(field, v)
	} else {
		codec.decode.Bytes(field, v)
	}
}

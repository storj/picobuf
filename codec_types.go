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

// RepeatedByte codes repeated byte protobuf type.
//go:noinline
func (codec *Codec) RepeatedByte(field FieldNumber, v *[]uint8) {
	if codec.encoding {
		codec.encode.RepeatedByte(field, v)
	} else {
		codec.decode.RepeatedByte(field, v)
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

// RepeatedBool codes repeated bool protobuf type.
//go:noinline
func (codec *Codec) RepeatedBool(field FieldNumber, v *[]bool) {
	if codec.encoding {
		codec.encode.RepeatedBool(field, v)
	} else {
		codec.decode.RepeatedBool(field, v)
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

// RepeatedInt32 codes repeated int32 protobuf type.
//go:noinline
func (codec *Codec) RepeatedInt32(field FieldNumber, v *[]int32) {
	if codec.encoding {
		codec.encode.RepeatedInt32(field, v)
	} else {
		codec.decode.RepeatedInt32(field, v)
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

// RepeatedInt64 codes repeated int64 protobuf type.
//go:noinline
func (codec *Codec) RepeatedInt64(field FieldNumber, v *[]int64) {
	if codec.encoding {
		codec.encode.RepeatedInt64(field, v)
	} else {
		codec.decode.RepeatedInt64(field, v)
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

// RepeatedUint32 codes repeated uint32 protobuf type.
//go:noinline
func (codec *Codec) RepeatedUint32(field FieldNumber, v *[]uint32) {
	if codec.encoding {
		codec.encode.RepeatedUint32(field, v)
	} else {
		codec.decode.RepeatedUint32(field, v)
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

// RepeatedUint64 codes repeated uint64 protobuf type.
//go:noinline
func (codec *Codec) RepeatedUint64(field FieldNumber, v *[]uint64) {
	if codec.encoding {
		codec.encode.RepeatedUint64(field, v)
	} else {
		codec.decode.RepeatedUint64(field, v)
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

// RepeatedSint32 codes repeated sint32 protobuf type.
//go:noinline
func (codec *Codec) RepeatedSint32(field FieldNumber, v *[]int32) {
	if codec.encoding {
		codec.encode.RepeatedSint32(field, v)
	} else {
		codec.decode.RepeatedSint32(field, v)
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

// RepeatedSint64 codes repeated sint64 protobuf type.
//go:noinline
func (codec *Codec) RepeatedSint64(field FieldNumber, v *[]int64) {
	if codec.encoding {
		codec.encode.RepeatedSint64(field, v)
	} else {
		codec.decode.RepeatedSint64(field, v)
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

// RepeatedFixed32 codes repeated fixed32 protobuf type.
//go:noinline
func (codec *Codec) RepeatedFixed32(field FieldNumber, v *[]uint32) {
	if codec.encoding {
		codec.encode.RepeatedFixed32(field, v)
	} else {
		codec.decode.RepeatedFixed32(field, v)
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

// RepeatedFixed64 codes repeated fixed64 protobuf type.
//go:noinline
func (codec *Codec) RepeatedFixed64(field FieldNumber, v *[]uint64) {
	if codec.encoding {
		codec.encode.RepeatedFixed64(field, v)
	} else {
		codec.decode.RepeatedFixed64(field, v)
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

// RepeatedSfixed32 codes repeated sfixed32 protobuf type.
//go:noinline
func (codec *Codec) RepeatedSfixed32(field FieldNumber, v *[]int32) {
	if codec.encoding {
		codec.encode.RepeatedSfixed32(field, v)
	} else {
		codec.decode.RepeatedSfixed32(field, v)
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

// RepeatedSfixed64 codes repeated sfixed64 protobuf type.
//go:noinline
func (codec *Codec) RepeatedSfixed64(field FieldNumber, v *[]int64) {
	if codec.encoding {
		codec.encode.RepeatedSfixed64(field, v)
	} else {
		codec.decode.RepeatedSfixed64(field, v)
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

// RepeatedFloat codes repeated float protobuf type.
//go:noinline
func (codec *Codec) RepeatedFloat(field FieldNumber, v *[]float32) {
	if codec.encoding {
		codec.encode.RepeatedFloat(field, v)
	} else {
		codec.decode.RepeatedFloat(field, v)
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

// RepeatedDouble codes repeated double protobuf type.
//go:noinline
func (codec *Codec) RepeatedDouble(field FieldNumber, v *[]float64) {
	if codec.encoding {
		codec.encode.RepeatedDouble(field, v)
	} else {
		codec.decode.RepeatedDouble(field, v)
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

// RepeatedString codes repeated string protobuf type.
//go:noinline
func (codec *Codec) RepeatedString(field FieldNumber, v *[]string) {
	if codec.encoding {
		codec.encode.RepeatedString(field, v)
	} else {
		codec.decode.RepeatedString(field, v)
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

// RepeatedRawString codes repeated rawstring protobuf type.
//go:noinline
func (codec *Codec) RepeatedRawString(field FieldNumber, v *[]string) {
	if codec.encoding {
		codec.encode.RepeatedRawString(field, v)
	} else {
		codec.decode.RepeatedRawString(field, v)
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

// RepeatedBytes codes repeated bytes protobuf type.
//go:noinline
func (codec *Codec) RepeatedBytes(field FieldNumber, v *[][]byte) {
	if codec.encoding {
		codec.encode.RepeatedBytes(field, v)
	} else {
		codec.decode.RepeatedBytes(field, v)
	}
}

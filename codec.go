// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import "google.golang.org/protobuf/encoding/protowire"

// Codec defines a unified implementation for encoder and decoder to
// minimize generated code usage for messages.
type Codec struct {
	encoding bool
	encode   Encoder
	decode   Decoder
}

// FieldNumber corresponds to a protobuf field number.
type FieldNumber protowire.Number

// IsValid returns whether the field number is in correct range.
func (field FieldNumber) IsValid() bool { return protowire.Number(field).IsValid() }

// RequiredByte executes the appropriate method based on the mode.
//go:noinline
func (codec *Codec) RequiredByte(field FieldNumber, v *byte) {
	if codec.encoding {
		codec.encode.RequiredByte(field, v)
	} else {
		codec.decode.RequiredByte(field, v)
	}
}

// RequiredRawString executes the appropriate method based on the mode.
//go:noinline
func (codec *Codec) RequiredRawString(field FieldNumber, v *string) {
	if codec.encoding {
		codec.encode.RequiredRawString(field, v)
	} else {
		codec.decode.RequiredRawString(field, v)
	}
}

// RequiredBytes executes the appropriate method based on the mode.
//go:noinline
func (codec *Codec) RequiredBytes(field FieldNumber, v *[]byte) {
	if codec.encoding {
		codec.encode.RequiredBytes(field, v)
	} else {
		codec.decode.RequiredBytes(field, v)
	}
}

// RequiredMessage executes the appropriate method based on the mode.
//go:noinline
func (codec *Codec) RequiredMessage(field FieldNumber, fn func(*Codec)) {
	if codec.encoding {
		codec.encode.RequiredMessage(field, fn)
	} else {
		codec.decode.RequiredMessage(field, fn)
	}
}

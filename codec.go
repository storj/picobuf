// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"google.golang.org/protobuf/encoding/protowire"
)

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

// Message codes a message.
//go:noinline
func (codec *Codec) Message(field FieldNumber, isNil func() bool, allocate func(), fn func(*Codec)) {
	if codec.encoding {
		codec.encode.Message(field, isNil, allocate, fn)
	} else {
		codec.decode.Message(field, isNil, allocate, fn)
	}
}

// PresentMessage codes an always present message.
//go:noinline
func (codec *Codec) PresentMessage(field FieldNumber, fn func(*Codec)) {
	if codec.encoding {
		codec.encode.PresentMessage(field, fn)
	} else {
		codec.decode.PresentMessage(field, fn)
	}
}

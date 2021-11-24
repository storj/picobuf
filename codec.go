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

// Message codes an embedded message.
//go:noinline
func (codec *Codec) Message(field FieldNumber, fn func(*Codec)) {
	if codec.encoding {
		codec.encode.Message(field, fn)
	} else {
		codec.decode.Message(field, fn)
	}
}

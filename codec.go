// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"google.golang.org/protobuf/encoding/protowire"
)

// Codec defines a unified implementation for encoder and decoder to
// minimize generated code usage for messages.
type Codec struct {
	decode Decoder
	encode Encoder
}

// FieldNumber corresponds to a protobuf field number.
type FieldNumber protowire.Number

// IsValid returns whether the field number is in correct range.
func (field FieldNumber) IsValid() bool { return protowire.Number(field).IsValid() }

// IsDecoding returns whether codec is in decoding mode.
func (codec *Codec) IsDecoding() bool { return codec.decode.codec != nil }

// Message codes a message.
//go:noinline
func (codec *Codec) Message(field FieldNumber, fn func(*Codec) bool) {
	if codec.decode.codec == nil {
		codec.encode.Message(field, fn)
	} else {
		codec.decode.Message(field, fn)
	}
}

// RepeatedMessage codes a repeated message.
//go:noinline
func (codec *Codec) RepeatedMessage(field FieldNumber, fn func(c *Codec, index int) bool) {
	if codec.decode.codec == nil {
		codec.encode.RepeatedMessage(field, fn)
	} else {
		codec.decode.RepeatedMessage(field, fn)
	}
}

// PresentMessage codes an always present message.
//go:noinline
func (codec *Codec) PresentMessage(field FieldNumber, fn func(*Codec) bool) {
	if codec.decode.codec == nil {
		codec.encode.PresentMessage(field, fn)
	} else {
		codec.decode.PresentMessage(field, fn)
	}
}

// RepeatedEnum codes a repeated enumeration.
//go:noinline
func (codec *Codec) RepeatedEnum(field FieldNumber, fn func(index int) (*int32, bool)) {
	if codec.decode.codec == nil {
		codec.encode.RepeatedEnum(field, fn)
	} else {
		codec.decode.RepeatedEnum(field, fn)
	}
}

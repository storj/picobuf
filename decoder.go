// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

// Decoder implements decoding of protobuf messages.
type Decoder struct{}

// RequiredByte decodes a v when it's in the stream.
func (dec *Decoder) RequiredByte(field FieldNumber, v *byte) {}

// RequiredRawString decodes a raw string when it's in the stream.
func (dec *Decoder) RequiredRawString(field FieldNumber, v *string) {}

// RequiredBytes decodes []byte when it's in the stream.
func (dec *Decoder) RequiredBytes(field FieldNumber, v *[]byte) {}

// RequiredMessage decodes []byte when it's in the stream.
func (dec *Decoder) RequiredMessage(field FieldNumber, fn func(*Codec)) {}

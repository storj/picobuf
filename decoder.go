// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

// Decoder implements decoding of protobuf messages.
type Decoder struct{}

// Message decodes an embedded message.
func (dec *Decoder) Message(field FieldNumber, fn func(*Codec)) {}

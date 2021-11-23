// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import "storj.io/picobuf"

// TypesPico corresponds to Types in types.proto.
type TypesPico struct {
	Int32    int32
	Int64    int64
	Uint32   uint32
	Uint64   uint64
	Sint32   int32
	Sint64   int64
	Fixed32  uint32
	Fixed64  uint64
	Sfixed32 int32
	Sfixed64 int64
	Float    float32
	Double   float64
	Bool     bool
	String   string
	Bytes    []byte
	Message  MessagePico
}

// MessagePico corresponds to Message in types.proto.
type MessagePico struct {
	Int32 int32
}

// Picobuf implements codec visitor.
func (m *TypesPico) Picobuf(codec *picobuf.Codec) {
	codec.Int32(1, &m.Int32)
	codec.Int64(2, &m.Int64)
	codec.Uint32(3, &m.Uint32)
	codec.Uint64(4, &m.Uint64)
	codec.Sint32(5, &m.Sint32)
	codec.Sint64(6, &m.Sint64)
	codec.Fixed32(7, &m.Fixed32)
	codec.Fixed64(8, &m.Fixed64)
	codec.Sfixed32(9, &m.Sfixed32)
	codec.Sfixed64(10, &m.Sfixed64)
	codec.Float(11, &m.Float)
	codec.Double(12, &m.Double)
	codec.Bool(13, &m.Bool)
	codec.String(14, &m.String)
	codec.Bytes(15, &m.Bytes)
	codec.Message(16, m.Message.Picobuf)
}

// Picobuf implements codec visitor.
func (m *MessagePico) Picobuf(codec *picobuf.Codec) {
	codec.Int32(1, &m.Int32)
}

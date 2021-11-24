// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: types.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v3.17.3

package protocompat

import (
	picobuf "storj.io/picobuf"
)

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
	String_  string
	Bytes    []byte
	Message  *MessagePico
}

func (m *TypesPico) Picobuf(c *picobuf.Codec) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
	c.Int64(2, &m.Int64)
	c.Uint32(3, &m.Uint32)
	c.Uint64(4, &m.Uint64)
	c.Sint32(5, &m.Sint32)
	c.Sint64(6, &m.Sint64)
	c.Fixed32(7, &m.Fixed32)
	c.Fixed64(8, &m.Fixed64)
	c.Sfixed32(9, &m.Sfixed32)
	c.Sfixed64(10, &m.Sfixed64)
	c.Float(11, &m.Float)
	c.Double(12, &m.Double)
	c.Bool(13, &m.Bool)
	c.String(14, &m.String_)
	c.Bytes(15, &m.Bytes)
	c.Message(16,
		func() bool { return m.Message == nil },
		func() { m.Message = new(MessagePico) },
		func(c *picobuf.Codec) {
			m.Message.Picobuf(c)
		},
	)
}

type MessagePico struct {
	Int32 int32
}

func (m *MessagePico) Picobuf(c *picobuf.Codec) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
}

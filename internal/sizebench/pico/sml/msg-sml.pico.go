// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: msg-sml.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v5.27.3

package sml

import (
	picobuf "storj.io/picobuf"
	picowire "storj.io/picobuf/picowire"
	strconv "strconv"
)

type Language int32

const (
	Language_UNKNOWN Language = 0
	Language_ENGLISH Language = 1
	Language_SPANISH Language = 3
	Language_FRENCH  Language = 4
	Language_GERMAN  Language = 5
)

func (m Language) String() string {
	switch m {
	case Language_UNKNOWN:
		return "UNKNOWN"
	case Language_ENGLISH:
		return "ENGLISH"
	case Language_SPANISH:
		return "SPANISH"
	case Language_FRENCH:
		return "FRENCH"
	case Language_GERMAN:
		return "GERMAN"
	default:
		return "Language(" + strconv.Itoa(int(m)) + ")"
	}
}

type Nop struct {
}

func (m *Nop) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	return true
}

func (m *Nop) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
}

type Types struct {
	Int32 int32 `json:"int32,omitempty"`
}

func (m *Types) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	return true
}

func (m *Types) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
}

type TypesUnused struct {
	Int32     int32             `json:"int32,omitempty"`
	Int64     int64             `json:"int64,omitempty"`
	Uint32    uint32            `json:"uint32,omitempty"`
	Uint64    uint64            `json:"uint64,omitempty"`
	Sint32    int32             `json:"sint32,omitempty"`
	Sint64    int64             `json:"sint64,omitempty"`
	Fixed32   uint32            `json:"fixed32,omitempty"`
	Fixed64   uint64            `json:"fixed64,omitempty"`
	Sfixed32  int32             `json:"sfixed32,omitempty"`
	Sfixed64  int64             `json:"sfixed64,omitempty"`
	Float     float32           `json:"float,omitempty"`
	Double    float64           `json:"double,omitempty"`
	Bool      bool              `json:"bool,omitempty"`
	String_   string            `json:"string,omitempty"`
	Bytes     []byte            `json:"bytes,omitempty"`
	Message   *Message          `json:"message,omitempty"`
	Language  Language          `json:"language,omitempty"`
	Int32S    []int32           `json:"int32s,omitempty"`
	Int64S    []int64           `json:"int64s,omitempty"`
	Uint32S   []uint32          `json:"uint32s,omitempty"`
	Uint64S   []uint64          `json:"uint64s,omitempty"`
	Sint32S   []int32           `json:"sint32s,omitempty"`
	Sint64S   []int64           `json:"sint64s,omitempty"`
	Fixed32S  []uint32          `json:"fixed32s,omitempty"`
	Fixed64S  []uint64          `json:"fixed64s,omitempty"`
	Sfixed32S []int32           `json:"sfixed32s,omitempty"`
	Sfixed64S []int64           `json:"sfixed64s,omitempty"`
	Floats    []float32         `json:"floats,omitempty"`
	Doubles   []float64         `json:"doubles,omitempty"`
	Bools     []bool            `json:"bools,omitempty"`
	Strings   []string          `json:"strings,omitempty"`
	Bytess    [][]byte          `json:"bytess,omitempty"`
	Messages  []*Message        `json:"messages,omitempty"`
	Languages []Language        `json:"languages,omitempty"`
	Values    map[string]string `json:"values,omitempty"`
}

func (m *TypesUnused) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
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
	c.Message(16, m.Message.Encode)
	c.Int32(17, (*int32)(&m.Language))
	c.RepeatedInt32(18, &m.Int32S)
	c.RepeatedInt64(19, &m.Int64S)
	c.RepeatedUint32(20, &m.Uint32S)
	c.RepeatedUint64(21, &m.Uint64S)
	c.RepeatedSint32(22, &m.Sint32S)
	c.RepeatedSint64(23, &m.Sint64S)
	c.RepeatedFixed32(24, &m.Fixed32S)
	c.RepeatedFixed64(25, &m.Fixed64S)
	c.RepeatedSfixed32(26, &m.Sfixed32S)
	c.RepeatedSfixed64(27, &m.Sfixed64S)
	c.RepeatedFloat(28, &m.Floats)
	c.RepeatedDouble(29, &m.Doubles)
	c.RepeatedBool(30, &m.Bools)
	c.RepeatedString(31, &m.Strings)
	c.RepeatedBytes(32, &m.Bytess)
	for _, x := range m.Messages {
		c.AlwaysMessage(33, x.Encode)
	}
	c.RepeatedEnum(34, len(m.Languages), func(index uint) int32 {
		if index < uint(len(m.Languages)) {
			return (int32)(m.Languages[index])
		}
		return 0
	})
	(*picowire.MapStringString)(&m.Values).PicoEncode(c, 35)
	return true
}

func (m *TypesUnused) Decode(c *picobuf.Decoder) {
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
	c.Message(16, func(c *picobuf.Decoder) {
		if m.Message == nil {
			m.Message = new(Message)
		}
		m.Message.Decode(c)
	})
	c.Int32(17, (*int32)(&m.Language))
	c.RepeatedInt32(18, &m.Int32S)
	c.RepeatedInt64(19, &m.Int64S)
	c.RepeatedUint32(20, &m.Uint32S)
	c.RepeatedUint64(21, &m.Uint64S)
	c.RepeatedSint32(22, &m.Sint32S)
	c.RepeatedSint64(23, &m.Sint64S)
	c.RepeatedFixed32(24, &m.Fixed32S)
	c.RepeatedFixed64(25, &m.Fixed64S)
	c.RepeatedSfixed32(26, &m.Sfixed32S)
	c.RepeatedSfixed64(27, &m.Sfixed64S)
	c.RepeatedFloat(28, &m.Floats)
	c.RepeatedDouble(29, &m.Doubles)
	c.RepeatedBool(30, &m.Bools)
	c.RepeatedString(31, &m.Strings)
	c.RepeatedBytes(32, &m.Bytess)
	c.RepeatedMessage(33, func(c *picobuf.Decoder) {
		x := new(Message)
		c.Loop(x.Decode)
		m.Messages = append(m.Messages, x)
	})
	c.RepeatedEnum(34, func(x int32) {
		m.Languages = append(m.Languages, (Language)(x))
	})
	(*picowire.MapStringString)(&m.Values).PicoDecode(c, 35)
}

type Message struct {
	Int32 int32 `json:"int32,omitempty"`
}

func (m *Message) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	return true
}

func (m *Message) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
}

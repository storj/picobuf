// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: types.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v3.20.0

package pico

import (
	picobuf "storj.io/picobuf"
	picowire "storj.io/picobuf/picowire"
)

type Language int32

const (
	Language_UNKNOWN Language = 0
	Language_ENGLISH Language = 1
	Language_SPANISH Language = 3
	Language_FRENCH  Language = 4
	Language_GERMAN  Language = 5
)

type Types struct {
	Int32           int32
	Int64           int64
	Uint32          uint32
	Uint64          uint64
	Sint32          int32
	Sint64          int64
	Fixed32         uint32
	Fixed64         uint64
	Sfixed32        int32
	Sfixed64        int64
	Float           float32
	Double          float64
	Bool            bool
	String_         string
	Bytes           []byte
	Message         Message
	OptionalMessage *OptionalMessage
	Language        Language
}

func (m *Types) Encode(c *picobuf.Encoder) bool {
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
	c.PresentMessage(16, m.Message.Encode)
	c.Message(17, m.OptionalMessage.Encode)
	c.Int32(18, (*int32)(&m.Language))
	return true
}

func (m *Types) Decode(c *picobuf.Decoder) {
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
	c.PresentMessage(16, m.Message.Decode)
	c.Message(17, func(c *picobuf.Decoder) {
		if m.OptionalMessage == nil {
			m.OptionalMessage = new(OptionalMessage)
		}
		m.OptionalMessage.Decode(c)
	})
	c.Int32(18, (*int32)(&m.Language))
}

type RepeatedTypes struct {
	Int32    []int32
	Int64    []int64
	Uint32   []uint32
	Uint64   []uint64
	Sint32   []int32
	Sint64   []int64
	Fixed32  []uint32
	Fixed64  []uint64
	Sfixed32 []int32
	Sfixed64 []int64
	Float    []float32
	Double   []float64
	Bool     []bool
	String_  []string
	Bytes    [][]byte
	Message  []*Message
	Language []Language
}

func (m *RepeatedTypes) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.RepeatedInt32(1, &m.Int32)
	c.RepeatedInt64(2, &m.Int64)
	c.RepeatedUint32(3, &m.Uint32)
	c.RepeatedUint64(4, &m.Uint64)
	c.RepeatedSint32(5, &m.Sint32)
	c.RepeatedSint64(6, &m.Sint64)
	c.RepeatedFixed32(7, &m.Fixed32)
	c.RepeatedFixed64(8, &m.Fixed64)
	c.RepeatedSfixed32(9, &m.Sfixed32)
	c.RepeatedSfixed64(10, &m.Sfixed64)
	c.RepeatedFloat(11, &m.Float)
	c.RepeatedDouble(12, &m.Double)
	c.RepeatedBool(13, &m.Bool)
	c.RepeatedString(14, &m.String_)
	c.RepeatedBytes(15, &m.Bytes)
	for _, x := range m.Message {
		c.AlwaysMessage(16, x.Encode)
	}
	c.RepeatedEnum(17, len(m.Language), func(index uint) int32 {
		if index < uint(len(m.Language)) {
			return (int32)(m.Language[index])
		}
		return 0
	})
	return true
}

func (m *RepeatedTypes) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.RepeatedInt32(1, &m.Int32)
	c.RepeatedInt64(2, &m.Int64)
	c.RepeatedUint32(3, &m.Uint32)
	c.RepeatedUint64(4, &m.Uint64)
	c.RepeatedSint32(5, &m.Sint32)
	c.RepeatedSint64(6, &m.Sint64)
	c.RepeatedFixed32(7, &m.Fixed32)
	c.RepeatedFixed64(8, &m.Fixed64)
	c.RepeatedSfixed32(9, &m.Sfixed32)
	c.RepeatedSfixed64(10, &m.Sfixed64)
	c.RepeatedFloat(11, &m.Float)
	c.RepeatedDouble(12, &m.Double)
	c.RepeatedBool(13, &m.Bool)
	c.RepeatedString(14, &m.String_)
	c.RepeatedBytes(15, &m.Bytes)
	c.RepeatedMessage(16, func(c *picobuf.Decoder) {
		mm := new(Message)
		c.Loop(mm.Decode)
		m.Message = append(m.Message, mm)
	})
	c.RepeatedEnum(17, func(x int32) {
		m.Language = append(m.Language, (Language)(x))
	})
}

type RepeatedMixed struct {
	Int32   int32
	Message []*Message
}

func (m *RepeatedMixed) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	for _, x := range m.Message {
		c.AlwaysMessage(16, x.Encode)
	}
	return true
}

func (m *RepeatedMixed) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
	c.RepeatedMessage(16, func(c *picobuf.Decoder) {
		mm := new(Message)
		c.Loop(mm.Decode)
		m.Message = append(m.Message, mm)
	})
}

type Message struct {
	Int32 int32
	Int64 int64
}

func (m *Message) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	c.Int64(2, &m.Int64)
	return true
}

func (m *Message) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
	c.Int64(2, &m.Int64)
}

type Person struct {
	Name     string
	Birthday int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64
	Primary  Language
	Spoken   []Language
}

func (m *Person) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Name)
	c.Int64(2, &m.Birthday)
	c.String(3, &m.Phone)
	c.Int32(4, &m.Siblings)
	c.Bool(5, &m.Spouse)
	c.Double(6, &m.Money)
	c.Int32(7, (*int32)(&m.Primary))
	c.RepeatedEnum(8, len(m.Spoken), func(index uint) int32 {
		if index < uint(len(m.Spoken)) {
			return (int32)(m.Spoken[index])
		}
		return 0
	})
	return true
}

func (m *Person) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Name)
	c.Int64(2, &m.Birthday)
	c.String(3, &m.Phone)
	c.Int32(4, &m.Siblings)
	c.Bool(5, &m.Spouse)
	c.Double(6, &m.Money)
	c.Int32(7, (*int32)(&m.Primary))
	c.RepeatedEnum(8, func(x int32) {
		m.Spoken = append(m.Spoken, (Language)(x))
	})
}

type Map struct {
	StringString   map[string]string
	StringInt32    map[string]int32
	StringInt64    map[string]int64
	StringUint32   map[string]uint32
	StringUint64   map[string]uint64
	StringSint32   map[string]int32
	StringSint64   map[string]int64
	StringFixed32  map[string]uint32
	StringFixed64  map[string]uint64
	StringSfixed32 map[string]int32
	StringSfixed64 map[string]int64
	StringFloat    map[string]float32
	StringDouble   map[string]float64
	StringBool     map[string]bool
	StringBytes    map[string][]byte
	Int32String    map[int32]string
	Int64String    map[int64]string
	Uint32String   map[uint32]string
	Uint64String   map[uint64]string
	Sint32String   map[int32]string
	Sint64String   map[int64]string
	Fixed32String  map[uint32]string
	Fixed64String  map[uint64]string
	Sfixed32String map[int32]string
	Sfixed64String map[int64]string
	BoolString     map[bool]string
}

func (m *Map) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	(*picowire.MapStringString)(&m.StringString).PicoEncode(c, 1)
	(*picowire.MapStringInt32)(&m.StringInt32).PicoEncode(c, 2)
	(*picowire.MapStringInt64)(&m.StringInt64).PicoEncode(c, 3)
	(*picowire.MapStringUint32)(&m.StringUint32).PicoEncode(c, 4)
	(*picowire.MapStringUint64)(&m.StringUint64).PicoEncode(c, 5)
	(*picowire.MapStringSint32)(&m.StringSint32).PicoEncode(c, 6)
	(*picowire.MapStringSint64)(&m.StringSint64).PicoEncode(c, 7)
	(*picowire.MapStringFixed32)(&m.StringFixed32).PicoEncode(c, 8)
	(*picowire.MapStringFixed64)(&m.StringFixed64).PicoEncode(c, 9)
	(*picowire.MapStringSfixed32)(&m.StringSfixed32).PicoEncode(c, 10)
	(*picowire.MapStringSfixed64)(&m.StringSfixed64).PicoEncode(c, 11)
	(*picowire.MapStringFloat)(&m.StringFloat).PicoEncode(c, 12)
	(*picowire.MapStringDouble)(&m.StringDouble).PicoEncode(c, 13)
	(*picowire.MapStringBool)(&m.StringBool).PicoEncode(c, 14)
	(*picowire.MapStringBytes)(&m.StringBytes).PicoEncode(c, 16)
	(*picowire.MapInt32String)(&m.Int32String).PicoEncode(c, 17)
	(*picowire.MapInt64String)(&m.Int64String).PicoEncode(c, 18)
	(*picowire.MapUint32String)(&m.Uint32String).PicoEncode(c, 19)
	(*picowire.MapUint64String)(&m.Uint64String).PicoEncode(c, 20)
	(*picowire.MapSint32String)(&m.Sint32String).PicoEncode(c, 21)
	(*picowire.MapSint64String)(&m.Sint64String).PicoEncode(c, 22)
	(*picowire.MapFixed32String)(&m.Fixed32String).PicoEncode(c, 23)
	(*picowire.MapFixed64String)(&m.Fixed64String).PicoEncode(c, 24)
	(*picowire.MapSfixed32String)(&m.Sfixed32String).PicoEncode(c, 25)
	(*picowire.MapSfixed64String)(&m.Sfixed64String).PicoEncode(c, 26)
	(*picowire.MapBoolString)(&m.BoolString).PicoEncode(c, 27)
	return true
}

func (m *Map) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	(*picowire.MapStringString)(&m.StringString).PicoDecode(c, 1)
	(*picowire.MapStringInt32)(&m.StringInt32).PicoDecode(c, 2)
	(*picowire.MapStringInt64)(&m.StringInt64).PicoDecode(c, 3)
	(*picowire.MapStringUint32)(&m.StringUint32).PicoDecode(c, 4)
	(*picowire.MapStringUint64)(&m.StringUint64).PicoDecode(c, 5)
	(*picowire.MapStringSint32)(&m.StringSint32).PicoDecode(c, 6)
	(*picowire.MapStringSint64)(&m.StringSint64).PicoDecode(c, 7)
	(*picowire.MapStringFixed32)(&m.StringFixed32).PicoDecode(c, 8)
	(*picowire.MapStringFixed64)(&m.StringFixed64).PicoDecode(c, 9)
	(*picowire.MapStringSfixed32)(&m.StringSfixed32).PicoDecode(c, 10)
	(*picowire.MapStringSfixed64)(&m.StringSfixed64).PicoDecode(c, 11)
	(*picowire.MapStringFloat)(&m.StringFloat).PicoDecode(c, 12)
	(*picowire.MapStringDouble)(&m.StringDouble).PicoDecode(c, 13)
	(*picowire.MapStringBool)(&m.StringBool).PicoDecode(c, 14)
	(*picowire.MapStringBytes)(&m.StringBytes).PicoDecode(c, 16)
	(*picowire.MapInt32String)(&m.Int32String).PicoDecode(c, 17)
	(*picowire.MapInt64String)(&m.Int64String).PicoDecode(c, 18)
	(*picowire.MapUint32String)(&m.Uint32String).PicoDecode(c, 19)
	(*picowire.MapUint64String)(&m.Uint64String).PicoDecode(c, 20)
	(*picowire.MapSint32String)(&m.Sint32String).PicoDecode(c, 21)
	(*picowire.MapSint64String)(&m.Sint64String).PicoDecode(c, 22)
	(*picowire.MapFixed32String)(&m.Fixed32String).PicoDecode(c, 23)
	(*picowire.MapFixed64String)(&m.Fixed64String).PicoDecode(c, 24)
	(*picowire.MapSfixed32String)(&m.Sfixed32String).PicoDecode(c, 25)
	(*picowire.MapSfixed64String)(&m.Sfixed64String).PicoDecode(c, 26)
	(*picowire.MapBoolString)(&m.BoolString).PicoDecode(c, 27)
}

type OptionalMessage struct {
	Int32 int32
}

func (m *OptionalMessage) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	return true
}

func (m *OptionalMessage) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
}

type CommandMessage struct {
	Class   string
	Command isCommandMessage_Command
}

func (m *CommandMessage) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Class)
	if m, ok := m.Command.(*CommandMessage_Name); ok {
		c.String(2, &m.Name)
	}
	if m, ok := m.Command.(*CommandMessage_Message); ok {
		c.Message(3, m.Message.Encode)
	}
	return true
}

func (m *CommandMessage) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Class)
	if c.PendingField() == 2 {
		var x *CommandMessage_Name
		if z, ok := m.Command.(*CommandMessage_Name); ok {
			x = z
		} else {
			x = new(CommandMessage_Name)
			m.Command = x
		}
		m := x
		c.String(2, &m.Name)
	}
	if c.PendingField() == 3 {
		var x *CommandMessage_Message
		if z, ok := m.Command.(*CommandMessage_Message); ok {
			x = z
		} else {
			x = new(CommandMessage_Message)
			m.Command = x
		}
		m := x
		c.Message(3, func(c *picobuf.Decoder) {
			if m.Message == nil {
				m.Message = new(Message)
			}
			m.Message.Decode(c)
		})
	}
}

type isCommandMessage_Command interface{ isCommandMessage_Command() }

type CommandMessage_Name struct {
	Name string
}

type CommandMessage_Message struct {
	Message *Message
}

func (*CommandMessage_Name) isCommandMessage_Command()    {}
func (*CommandMessage_Message) isCommandMessage_Command() {}

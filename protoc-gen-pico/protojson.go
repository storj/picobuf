// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"maps"
	"strconv"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const picojsonPackage = protogen.GoImportPath("storj.io/picobuf/picojson")

// jsonSchema is the picojson.Schema of a generated file: every message and
// enum that its messages refer to, described from their descriptors, so
// that other packages do not need to be generated with protojson.
type jsonSchema struct {
	ident    string
	messages []*protogen.Message
	enums    []*protogen.Enum
	// records are the encoded messages and then enums, and offset is where
	// each of them starts in the schema.
	records [][]byte
	offset  map[protoreflect.FullName]int
	// strings is where each string is first in the schema.
	strings map[string]int
}

// unsupportedJSON lists the well-known types whose JSON form picojson does
// not implement.
var unsupportedJSON = map[protoreflect.FullName]bool{
	"google.protobuf.Any":       true,
	"google.protobuf.Struct":    true,
	"google.protobuf.Value":     true,
	"google.protobuf.ListValue": true,
	"google.protobuf.NullValue": true,
}

func newJSONSchema(file *protogen.File) *jsonSchema {
	s := &jsonSchema{
		ident:  "picojson_" + file.GoDescriptorIdent.GoName,
		offset: map[protoreflect.FullName]int{},
	}
	seen := map[protoreflect.FullName]bool{}
	var addMessage func(m *protogen.Message)
	addEnum := func(e *protogen.Enum) {
		name := e.Desc.FullName()
		if unsupportedJSON[name] {
			panic("protojson: unsupported type " + string(name))
		}
		if !seen[name] {
			seen[name] = true
			s.enums = append(s.enums, e)
		}
	}
	addMessage = func(m *protogen.Message) {
		name := m.Desc.FullName()
		if unsupportedJSON[name] {
			panic("protojson: unsupported type " + string(name))
		}
		if seen[name] {
			return
		}
		seen[name] = true
		s.messages = append(s.messages, m)
		for _, field := range m.Fields {
			if field.Desc.Kind() == protoreflect.GroupKind {
				panic("protojson: unsupported group field " + string(field.Desc.FullName()))
			}
			if field.Message != nil {
				addMessage(field.Message)
			}
			if field.Enum != nil {
				addEnum(field.Enum)
			}
		}
	}
	var walk func(msgs []*protogen.Message)
	walk = func(msgs []*protogen.Message) {
		for _, m := range msgs {
			addMessage(m)
			walk(m.Messages)
		}
	}
	walk(file.Messages)
	s.encode()
	return s
}

func genJSONMethods(gf *generator, m *protogen.Message) {
	offset := gf.json.offset[m.Desc.FullName()]
	// A value receiver, so that encoding/json uses it for values that are
	// not addressable, such as a message held by value in a map.
	gf.P("func (m ", m.GoIdent, ") MarshalJSON() ([]byte, error) {")
	gf.P("return ", picojsonPackage.Ident("Marshal"), "(&m, ", gf.json.ident, ", ", offset, ")")
	gf.P("}")
	gf.P()
	gf.P("func (m *", m.GoIdent, ") UnmarshalJSON(data []byte) error {")
	// By the encoding/json convention, null is a no-op.
	gf.P("if string(data) == \"null\" {")
	gf.P("return nil")
	gf.P("}")
	gf.P("*m = ", m.GoIdent, "{}")
	gf.P("return ", picojsonPackage.Ident("Unmarshal"), "(data, m, ", gf.json.ident, ", ", offset, ")")
	gf.P("}")
	gf.P()
}

// Field flags, see picojson.Schema.
const (
	jsonFlagList = 1 << iota
	jsonFlagMap
	jsonFlagPresence
	jsonFlagOneof
	jsonFlagJSONName
)

// encode encodes the records. A record refers to others by offset, and the
// size of an offset depends on the offsets before it, so it encodes until
// the offsets no longer change.
func (s *jsonSchema) encode() {
	for {
		var b []byte
		s.records = s.records[:0]
		s.strings = map[string]int{}
		offset := map[protoreflect.FullName]int{}
		add := func(name protoreflect.FullName, start int) {
			offset[name] = start
			s.records = append(s.records, b[start:])
		}
		for _, m := range s.messages {
			start := len(b)
			b = s.appendMessage(b, m)
			add(m.Desc.FullName(), start)
		}
		for _, e := range s.enums {
			start := len(b)
			b = s.appendEnum(b, e)
			add(e.Desc.FullName(), start)
		}
		if maps.Equal(offset, s.offset) {
			return
		}
		s.offset = offset
	}
}

func (s *jsonSchema) appendMessage(b []byte, m *protogen.Message) []byte {
	b = s.appendString(b, string(m.Desc.FullName()))
	b = protowire.AppendVarint(b, uint64(len(m.Fields)))
	for _, field := range m.Fields {
		desc := field.Desc
		var flags byte
		switch {
		case desc.IsMap():
			flags |= jsonFlagMap
		case desc.IsList():
			flags |= jsonFlagList
		case desc.HasPresence():
			flags |= jsonFlagPresence
		}
		oneof := field.Oneof
		if oneof != nil && !oneof.Desc.IsSynthetic() {
			flags |= jsonFlagOneof
		}
		if desc.JSONName() != string(desc.Name()) {
			flags |= jsonFlagJSONName
		}
		b = protowire.AppendVarint(b, uint64(desc.Number()))
		b = append(b, byte(desc.Kind()), flags)
		if flags&jsonFlagOneof != 0 {
			b = protowire.AppendVarint(b, uint64(oneof.Desc.Index()+1))
		}
		switch {
		case field.Message != nil:
			b = protowire.AppendVarint(b, uint64(s.offset[field.Message.Desc.FullName()]))
		case field.Enum != nil:
			b = protowire.AppendVarint(b, uint64(s.offset[field.Enum.Desc.FullName()]))
		}
		b = s.appendString(b, string(desc.Name()))
		if flags&jsonFlagJSONName != 0 {
			b = s.appendString(b, desc.JSONName())
		}
	}
	return b
}

func (s *jsonSchema) appendEnum(b []byte, e *protogen.Enum) []byte {
	b = s.appendString(b, string(e.Desc.FullName()))
	b = protowire.AppendVarint(b, uint64(len(e.Values)))
	for _, v := range e.Values {
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v.Desc.Number())))
		b = s.appendString(b, string(v.Desc.Name()))
	}
	return b
}

// appendString appends str, or a reference to where it is earlier in b
// when that is shorter, see picojson.Schema.
func (s *jsonSchema) appendString(b []byte, str string) []byte {
	if at, ok := s.strings[str]; ok {
		ref := uint64(at)<<1 | 1
		if protowire.SizeVarint(ref) < protowire.SizeVarint(uint64(len(str))<<1)+len(str) {
			return protowire.AppendVarint(b, ref)
		}
	} else {
		s.strings[str] = len(b)
	}
	b = protowire.AppendVarint(b, uint64(len(str))<<1)
	return append(b, str...)
}

// genJSONSchema writes the schema as one record per line.
func genJSONSchema(gf *generator) {
	s := gf.json
	if len(s.records) == 0 {
		return
	}
	gf.P("const ", s.ident, " ", picojsonPackage.Ident("Schema"), " = \"\" +")
	names := make([]protoreflect.FullName, 0, len(s.records))
	for _, m := range s.messages {
		names = append(names, m.Desc.FullName())
	}
	for _, e := range s.enums {
		names = append(names, e.Desc.FullName())
	}
	for i, record := range s.records {
		end := " +"
		if i == len(s.records)-1 {
			end = ""
		}
		gf.P(strconv.Quote(string(record)), end, " // ", s.offset[names[i]], ": ", names[i])
	}
	gf.P()
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"flag"
	"fmt"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:generate protoc -I.. --go_out=paths=source_relative:. pico.proto

const (
	picobufPackage  = protogen.GoImportPath("storj.io/picobuf")
	picowirePackage = protogen.GoImportPath("storj.io/picobuf/picowire")
	strconvPackage  = protogen.GoImportPath("strconv")
)

type config struct{}

func main() {
	var flags flag.FlagSet
	var conf config

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if f.Generate {
				genFile(plugin, f, conf)
			}
		}
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}

type generator struct {
	*protogen.GeneratedFile
	*protogen.File
}

func genFile(plugin *protogen.Plugin, file *protogen.File, conf config) {
	gf := &generator{
		GeneratedFile: plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".pico.go", file.GoImportPath),
		File:          file,
	}

	gf.P("// Code generated by protoc-gen-pico. DO NOT EDIT.")
	gf.P("// source: ", file.Desc.Path())
	gf.P("//")
	gf.P("// versions:")
	{
		picoVersion := "(unknown)"
		if bi, ok := debug.ReadBuildInfo(); ok {
			picoVersion = bi.Main.Version
		}
		gf.P("//     protoc-gen-pico: ", picoVersion)
	}
	{
		protocVersion := "(unknown)"
		if v := plugin.Request.GetCompilerVersion(); v != nil {
			protocVersion = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
			if s := v.GetSuffix(); s != "" {
				protocVersion += "-" + s
			}
		}
		gf.P("//     protoc:          ", protocVersion)
	}
	gf.P()
	gf.P("package ", file.GoPackageName)
	gf.P()

	genWalk(gf, file, file.Enums, file.Messages)
}

func genWalk(gf *generator, file *protogen.File, enums []*protogen.Enum, msgs []*protogen.Message) {
	for _, e := range enums {
		genEnum(gf, e)
	}
	for _, m := range msgs {
		genMessage(gf, m)
	}
	for _, m := range msgs {
		genWalk(gf, file, m.Enums, m.Messages)
	}
}

func genEnum(gf *generator, e *protogen.Enum) {
	gf.P("type ", e.GoIdent, " int32")

	gf.P("const (")
	for _, v := range e.Values {
		gf.P(v.GoIdent, " ", e.GoIdent, " = ", v.Desc.Number())
	}
	gf.P(")")
	gf.P()

	gf.P("func (m ", e.GoIdent, ") String() string {")
	gf.P("switch m {")
	for _, v := range e.Values {
		gf.P("case ", v.GoIdent, ": return ", strconv.Quote(string(v.Desc.Name())))
	}
	itoaName := gf.QualifiedGoIdent(strconvPackage.Ident("Itoa"))
	gf.P("default: return ", strconv.Quote(e.GoIdent.GoName+"("), " + ", itoaName, "(int(m)) + \")\" ")
	gf.P("}")
	gf.P("}")

	gf.P()
}

func genMessage(gf *generator, m *protogen.Message) {
	if m.Desc.IsMapEntry() {
		return
	}

	gf.P("type ", m.GoIdent, " struct {")
	for _, field := range m.Fields {
		genMessageField(gf, m, field)
	}
	if getMessageOpts(m).CaptureUnrecognizedFields {
		gf.P("XXX_unrecognized", " ", "[]byte")
	}
	gf.P("}")

	gf.P()

	genMessageMethods(gf, m)
	genMessageOneofWrapperTypes(gf, m)
}

func genMessageField(gf *generator, m *protogen.Message, field *protogen.Field) {
	if field.Desc.IsWeak() {
		panic("unhandled: weak field")
	}
	if field.Desc.HasDefault() {
		panic("unsupported: default values")
	}

	if oneof := field.Oneof; oneof != nil && !oneof.Desc.IsSynthetic() {
		if oneof.Fields[0] != field {
			return // only generate for first appearance
		}
		gf.P(oneof.GoName, " ", oneofInterfaceName(gf, oneof))
		return
	}

	gf.P(field.GoName, " ", fieldGoType(gf, field), " ", fieldJSONTag(gf, field))
}

func oneofInterfaceName(gf *generator, oneof *protogen.Oneof) string {
	return "is" + oneof.GoIdent.GoName
}

func fieldJSONTag(gf *generator, field *protogen.Field) string {
	// It seems the `field.JSONName()` and the json names std protobuf
	// generates are different. The std protobuf-go generates using
	// the original field name in protobuf.
	return "`json:\"" + string(field.Desc.Name()) + ",omitempty\"`"
}

func genMessageMethods(gf *generator, m *protogen.Message) {
	genMessageEncode(gf, m)
	genMessageDecode(gf, m)
}

func genMessageEncode(gf *generator, m *protogen.Message) {
	gf.P("func (m *", m.GoIdent, ") Encode(c *", picobufPackage.Ident("Encoder"), ") bool {")
	gf.P("if m == nil { return false }")
	defer func() {
		gf.P("return true")
		gf.P("}")
		gf.P()
	}()

	fields := append([]*protogen.Field(nil), m.Fields...)
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Desc.Number() < fields[j].Desc.Number()
	})

	for _, field := range fields {
		genFieldEncode(gf, field)
	}
	if getMessageOpts(m).CaptureUnrecognizedFields {
		gf.P("c.UnrecognizedFields(m.XXX_unrecognized)")
	}
}

func genFieldEncode(gf *generator, field *protogen.Field) {
	info := fieldInfo(gf, field, field.Desc)
	always := false

	if info.oneof {
		gf.P("if m, ok := m.", field.Oneof.GoName, ".(*", oneofWrapperTypeName(gf, field), "); ok {")
		defer gf.P("}")
		always = true
	}

	switch {
	case info.kind == kindCast:
		switch {
		case !info.repeated && info.pointer:
			gf.P("(*", info.castType, ")(m.", field.GoName, ").PicoEncode(c, ", field.Desc.Number(), ")")
		case info.repeated && info.pointer:
			gf.P("for _, x := range m.", field.GoName, " {")
			gf.P("  (*", info.castType, ")(x).PicoEncode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		case !info.repeated && !info.pointer:
			gf.P("(*", info.castType, ")(&m.", field.GoName, ").PicoEncode(c, ", field.Desc.Number(), ")")
		case info.repeated && !info.pointer:
			gf.P("for i := range m.", field.GoName, " {")
			gf.P("  x := &m.", field.GoName, "[i]")
			gf.P("  (*", info.castType, ")(x).PicoEncode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		}

	case info.kind == kindCustom:
		switch {
		case !info.repeated:
			gf.P("m.", field.GoName, ".PicoEncode(c, ", field.Desc.Number(), ")")
		case info.pointer && info.repeated:
			gf.P("for _, x := range m.", field.GoName, " {")
			gf.P("  x.PicoEncode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		case !info.pointer && info.repeated:
			gf.P("for i := range m.", field.GoName, " {")
			gf.P("  x := &m.", field.GoName, "[i]")
			gf.P("  x.PicoEncode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		}

	case info.kind == kindInternal:
		method, ok := codecMethodName[field.Desc.Kind()]
		if !ok {
			panic("unsupported type " + field.Desc.Kind().GoString())
		}
		if info.repeated {
			method = "Repeated" + method
		}
		if always {
			method = "Always" + method
		}

		if info.repeated && info.pointer {
			panic("pointer not supported for repeated internal type " + info.goType)
		}
		if !info.pointer {
			gf.P("c.", method, "(", field.Desc.Number(), ", &m.", field.GoName, ")")
		} else {
			gf.P("if m.", field.GoName, " != nil {")
			gf.P("  c.", method, "(", field.Desc.Number(), ", m.", field.GoName, ")")
			gf.P("}")
		}

	case info.kind == kindMessage:
		switch {
		case info.pointer && !info.repeated:
			gf.P("c.Message(", field.Desc.Number(), ", m.", field.GoName, ".Encode)")
		case info.pointer && info.repeated:
			gf.P("for _, x := range m.", field.GoName, " {")
			gf.P("  c.AlwaysMessage(", field.Desc.Number(), ", x.Encode)")
			gf.P("}")
		case !info.pointer && !info.repeated:
			gf.P("c.PresentMessage(", field.Desc.Number(), ", m.", field.GoName, ".Encode)")
		case !info.pointer && info.repeated:
			gf.P("for i := range m.", field.GoName, " {")
			gf.P("  x := &m.", field.GoName, "[i]")
			gf.P("  c.AlwaysMessage(", field.Desc.Number(), ", x.Encode)")
			gf.P("}")
		}

	case info.kind == kindEnum:
		switch {
		case info.pointer && !info.repeated:
			panic("enum with pointer not supported")
		case info.pointer && info.repeated:
			panic("enum with pointer and repeated not supported")
		case !info.pointer && !info.repeated:
			gf.P("c.Int32(", field.Desc.Number(), ", (*int32)(&m.", field.GoName, "))")
		case !info.pointer && info.repeated:
			gf.P("c.RepeatedEnum(", field.Desc.Number(), ", len(m.", field.GoName, ")", ", func(index uint) int32 {")
			gf.P("  if index < uint(len(m.", field.GoName, ")) {")
			gf.P("     return (int32)(m.", field.GoName, "[index])")
			gf.P("  }")
			gf.P("  return 0")
			gf.P("})")
		}

	case info.kind == kindMap:
		if info.repeated {
			panic("map cannot be repeated")
		}

		panic("unsupported map")
	default:
		panic("unhandled kind")
	}
}

func genMessageDecode(gf *generator, m *protogen.Message) {
	gf.P("func (m *", m.GoIdent, ") Decode(c *", picobufPackage.Ident("Decoder"), ") {")
	gf.P("if m == nil { return }")
	defer func() {
		gf.P("}")
		gf.P()
	}()

	fields := append([]*protogen.Field(nil), m.Fields...)
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Desc.Number() < fields[j].Desc.Number()
	})

	for _, field := range fields {
		genFieldDecode(gf, field)
	}
	if getMessageOpts(m).CaptureUnrecognizedFields {
		gf.P("c.UnrecognizedFields(", fieldsBitSet(m), ", &m.XXX_unrecognized)")
	}
}

func fieldsBitSet(m *protogen.Message) (z uint64) {
	for _, f := range m.Fields {
		num := f.Desc.Number()
		if num >= 64 {
			panic("oh no, too many fields")
		}
		z |= 1 << num
	}
	return z
}

func genFieldDecode(gf *generator, field *protogen.Field) {
	info := fieldInfo(gf, field, field.Desc)

	if info.oneof {
		gf.P("if c.PendingField() == ", field.Desc.Number(), " {")
		gf.P("var x *", oneofWrapperTypeName(gf, field))
		gf.P("if z, ok := m.", field.Oneof.GoName, ".(*", oneofWrapperTypeName(gf, field), "); ok {")
		gf.P("   x = z")
		gf.P("} else {")
		gf.P("   x = new(", oneofWrapperTypeName(gf, field), ")")
		gf.P("   m.", field.Oneof.GoName, " = x")
		gf.P("}")
		gf.P("m := x")
		defer gf.P("}")
	}

	switch {
	case info.kind == kindCast:
		switch {
		case !info.repeated && info.pointer:
			gf.P("if c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  if m.", field.GoName, " == nil {")
			gf.P("    m.", field.GoName, " = new(", info.baseType, ")")
			gf.P("  }")
			gf.P("  (*", info.castType, ")(m.", field.GoName, ").PicoDecode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		case info.repeated && info.pointer:
			gf.P("for c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  x := new(", info.baseType, ")")
			gf.P("  (*", info.castType, ")(x).PicoDecode(c, ", field.Desc.Number(), ")")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", x)")
			gf.P("}")
		case !info.repeated && !info.pointer:
			gf.P("(*", info.castType, ")(&m.", field.GoName, ").PicoDecode(c, ", field.Desc.Number(), ")")
		case info.repeated && !info.pointer:
			gf.P("for c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", ", info.baseType, "{})")
			gf.P("  x := &m.", field.GoName, "[len(m.", field.GoName, ")-1]")
			gf.P("  (*", info.castType, ")(x).PicoDecode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		}

	case info.kind == kindCustom:
		switch {
		case !info.repeated && info.pointer:
			gf.P("if c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  if m.", field.GoName, " == nil {")
			gf.P("    m.", field.GoName, " = new(", info.baseType, ")")
			gf.P("  }")
			gf.P("  m.", field.GoName, ".PicoDecode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		case info.repeated && info.pointer:
			gf.P("for c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  x := new(", info.baseType, ")")
			gf.P("  x.PicoDecode(c, ", field.Desc.Number(), ")")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", x)")
			gf.P("}")
		case !info.repeated && !info.pointer:
			gf.P("m.", field.GoName, ".PicoDecode(c, ", field.Desc.Number(), ")")
		case info.repeated && !info.pointer:
			gf.P("for c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", ", info.baseType, "{})")
			gf.P("  m.", field.GoName, "[len(m.", field.GoName, ")-1].PicoDecode(c, ", field.Desc.Number(), ")")
			gf.P("}")
		}

	case info.kind == kindInternal:
		method, ok := codecMethodName[field.Desc.Kind()]
		if !ok {
			panic("unsupported type " + field.Desc.Kind().GoString())
		}
		if info.repeated {
			method = "Repeated" + method
		}
		if info.repeated && info.pointer {
			panic("pointer not supported for repeated internal type " + info.goType)
		}
		if !info.pointer {
			gf.P("c.", method, "(", field.Desc.Number(), ", &m.", field.GoName, ")")
		} else {
			gf.P("if c.PendingField() == ", field.Desc.Number(), " {")
			gf.P("  m.", field.GoName, " = new(", info.baseType, ")")
			gf.P("  c.", method, "(", field.Desc.Number(), ", m.", field.GoName, ")")
			gf.P("}")
		}

	case info.kind == kindMessage:
		switch {
		case info.pointer && !info.repeated:
			gf.P("c.Message(", field.Desc.Number(), ", func(c *", picobufPackage.Ident("Decoder"), ") {")
			gf.P("  if m.", field.GoName, " == nil {")
			gf.P("    m.", field.GoName, " = new(", info.baseType, ")")
			gf.P("  }")
			gf.P("  m.", field.GoName, ".Decode(c)")
			gf.P("})")
		case info.pointer && info.repeated:
			gf.P("c.RepeatedMessage(", field.Desc.Number(), ", func(c *", picobufPackage.Ident("Decoder"), ") {")
			gf.P("  x := new(", info.baseType, ")")
			gf.P("  c.Loop(x.Decode)")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", x)")
			gf.P("})")
		case !info.pointer && !info.repeated:
			gf.P("c.PresentMessage(", field.Desc.Number(), ", m.", field.GoName, ".Decode)")
		case !info.pointer && info.repeated:
			gf.P("c.RepeatedMessage(", field.Desc.Number(), ", func(c *", picobufPackage.Ident("Decoder"), ") {")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", ", info.baseType, "{})")
			gf.P("  c.Loop(m.", field.GoName, "[len(m.", field.GoName, ")-1].Decode)")
			gf.P("})")
		}

	case info.kind == kindEnum:
		switch {
		case info.pointer && !info.repeated:
			panic("enum with pointer not supported")
		case info.pointer && info.repeated:
			panic("enum with pointer and repeated not supported")
		case !info.pointer && !info.repeated:
			gf.P("c.Int32(", field.Desc.Number(), ", (*int32)(&m.", field.GoName, "))")
		case !info.pointer && info.repeated:
			gf.P("c.RepeatedEnum(", field.Desc.Number(), ", func(x int32) {")
			gf.P("  m.", field.GoName, " = append(m.", field.GoName, ", (", info.baseType, ")(x))")
			gf.P("})")
		}

	case info.kind == kindMap:
		if info.repeated {
			panic("map cannot be repeated")
		}

		panic("unsupported map")
	default:
		panic("unhandled kind")
	}
}

var codecMethodName = map[protoreflect.Kind]string{
	protoreflect.BoolKind:     "Bool",
	protoreflect.Int32Kind:    "Int32",
	protoreflect.Int64Kind:    "Int64",
	protoreflect.Uint32Kind:   "Uint32",
	protoreflect.Uint64Kind:   "Uint64",
	protoreflect.Sint32Kind:   "Sint32",
	protoreflect.Sint64Kind:   "Sint64",
	protoreflect.Fixed32Kind:  "Fixed32",
	protoreflect.Fixed64Kind:  "Fixed64",
	protoreflect.Sfixed32Kind: "Sfixed32",
	protoreflect.Sfixed64Kind: "Sfixed64",
	protoreflect.FloatKind:    "Float",
	protoreflect.DoubleKind:   "Double",
	protoreflect.StringKind:   "String",
	protoreflect.BytesKind:    "Bytes",
}

type fieldKind byte

const (
	kindInternal = fieldKind(iota)
	kindMessage
	kindEnum
	kindMap
	kindCustom
	kindCast
)

type fieldInformation struct {
	baseType string
	goType   string
	castType string
	kind     fieldKind
	pointer  bool
	repeated bool
	oneof    bool

	key   *fieldInformation
	value *fieldInformation
}

func fieldInfo(gf *generator, field *protogen.Field, desc protoreflect.FieldDescriptor) (info fieldInformation) {
	if desc.IsWeak() {
		panic("unhandled: weak field")
	}

	info.kind = kindInternal
	info.pointer = desc.HasPresence() || desc.HasOptionalKeyword()
	info.repeated = desc.IsList()
	info.oneof = desc.ContainingOneof() != nil && !desc.ContainingOneof().IsSynthetic()

	switch desc.Kind() {
	case protoreflect.BoolKind:
		info.baseType = "bool"
	case protoreflect.EnumKind:
		if field == nil {
			panic("unsupported enum")
		}
		info.baseType = gf.QualifiedGoIdent(field.Enum.GoIdent)
		info.kind = kindEnum
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		info.baseType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		info.baseType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		info.baseType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		info.baseType = "uint64"
	case protoreflect.FloatKind:
		info.baseType = "float32"
	case protoreflect.DoubleKind:
		info.baseType = "float64"
	case protoreflect.StringKind:
		info.baseType = "string"
		info.pointer = desc.HasOptionalKeyword()
	case protoreflect.BytesKind:
		info.baseType = "[]byte"
		info.pointer = desc.HasOptionalKeyword()
	case protoreflect.MessageKind:
		if field.Desc.IsMap() {
			info.kind = kindMap

			if keyKind := desc.MapKey().Kind(); !validMapKey[keyKind] {
				panic("invalid map key " + keyKind.String())
			}

			keyInfo := fieldInfo(gf, nil, desc.MapKey())
			info.key = &keyInfo

			valueInfo := fieldInfo(gf, nil, desc.MapValue())
			info.value = &valueInfo

			info.baseType = "map[" + info.key.goType + "]" + info.value.goType

			keyMethod, keyOk := codecMethodName[field.Desc.MapKey().Kind()]
			valueMethod, valueOk := codecMethodName[field.Desc.MapValue().Kind()]
			if keyOk && valueOk {
				info.castType = gf.QualifiedGoIdent(picowirePackage.Ident("Map" + keyMethod + valueMethod))
				info.kind = kindCast
			}
			info.pointer = false
		} else {
			info.kind = kindMessage
			info.pointer = true
			info.baseType = gf.QualifiedGoIdent(field.Message.GoIdent)
		}
	default:
		panic(fmt.Sprintf("unhandled: invalid field kind: %v", field.Desc.Kind()))
	}

	// only messages in oneof should be pointers, by default
	if info.pointer && info.kind != kindMessage && info.oneof {
		info.pointer = false
	}

	if field != nil {
		if getMessageOpts(field.Message).AlwaysPresent {
			info.pointer = false
		}

		opts := getFieldOpts(field)
		if opts.AlwaysPresent {
			info.pointer = false
		}
		if opts.CustomType != "" {
			info.baseType = qualifiedIdent(gf, opts.CustomType)
			info.kind = kindCustom
		}
		if opts.CustomSerialize != "" {
			info.castType = qualifiedIdent(gf, opts.CustomSerialize)
			info.kind = kindCast
		}
	}

	info.goType = info.baseType
	if info.pointer {
		info.goType = "*" + info.goType
	}
	if info.repeated {
		info.goType = "[]" + info.goType
	}

	return info
}

var validMapKey = map[protoreflect.Kind]bool{
	protoreflect.BoolKind:     true,
	protoreflect.EnumKind:     false,
	protoreflect.Int32Kind:    true,
	protoreflect.Sint32Kind:   true,
	protoreflect.Uint32Kind:   true,
	protoreflect.Int64Kind:    true,
	protoreflect.Sint64Kind:   true,
	protoreflect.Uint64Kind:   true,
	protoreflect.Sfixed32Kind: true,
	protoreflect.Fixed32Kind:  true,
	protoreflect.FloatKind:    false,
	protoreflect.Sfixed64Kind: true,
	protoreflect.Fixed64Kind:  true,
	protoreflect.DoubleKind:   false,
	protoreflect.StringKind:   true,
	protoreflect.BytesKind:    false,
	protoreflect.MessageKind:  false,
	protoreflect.GroupKind:    false,
}

func fieldGoType(gf *generator, field *protogen.Field) string {
	info := fieldInfo(gf, field, field.Desc)
	return info.goType
}

func qualifiedIdent(gf *generator, fullyQualifiedIdent string) string {
	p := strings.LastIndexByte(fullyQualifiedIdent, '.')
	if p >= 0 {
		return gf.QualifiedGoIdent(protogen.GoIdent{
			GoName:       fullyQualifiedIdent[p+1:],
			GoImportPath: protogen.GoImportPath(fullyQualifiedIdent[:p]),
		})
	}
	return fullyQualifiedIdent
}

func getMessageOpts(m *protogen.Message) *MessageOptions {
	if m == nil {
		return &MessageOptions{}
	}
	opts, _ := proto.GetExtension(m.Desc.Options(), E_Message).(*MessageOptions)
	if opts == nil {
		return &MessageOptions{}
	}
	return opts
}

func getFieldOpts(f *protogen.Field) *FieldOptions {
	opts, _ := proto.GetExtension(f.Desc.Options(), E_Field).(*FieldOptions)
	if opts == nil {
		return &FieldOptions{}
	}
	return opts
}

func genMessageOneofWrapperTypes(gf *generator, m *protogen.Message) {
	for _, oneof := range m.Oneofs {
		if oneof.Desc.IsSynthetic() {
			continue
		}
		ifName := oneofInterfaceName(gf, oneof)

		gf.P("type ", ifName, " interface {", ifName, "() }")
		gf.P()

		for _, field := range oneof.Fields {
			gf.P("type ", oneofWrapperTypeName(gf, field), " struct {")
			gf.P(field.GoName, " ", fieldGoType(gf, field))
			gf.P("}")
			gf.P()
		}
		gf.P()

		for _, field := range oneof.Fields {
			gf.P("func (*", oneofWrapperTypeName(gf, field), ") ", ifName, "() {}")
		}
		gf.P()
	}
}

func oneofWrapperTypeName(gf *generator, field *protogen.Field) string {
	return field.GoIdent.GoName
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

// PrimitiveType represents a primitive type that the codec supports directly.
type PrimitiveType struct {
	Name      string
	Zero      interface{}
	Wire      protowire.Type
	Append    string
	EncodeFmt string
}

// TypeName returns the corresponding Go type for the primitive.
func (t *PrimitiveType) TypeName() string {
	if _, isBytes := t.Zero.([]byte); isBytes {
		return "[]byte"
	}
	v := reflect.ValueOf(t.Zero)
	return v.Type().String()
}

// WireName returns the tag wire type for the type.
func (t *PrimitiveType) WireName() string {
	switch t.Wire {
	case protowire.VarintType:
		return "protowire.VarintType"
	case protowire.Fixed32Type:
		return "protowire.Fixed32Type"
	case protowire.Fixed64Type:
		return "protowire.Fixed64Type"
	case protowire.BytesType:
		return "protowire.BytesType"
	case protowire.StartGroupType:
		return "protowire.StartGroupType"
	case protowire.EndGroupType:
		return "protowire.EndGroupType"
	default:
		panic("unhandled wire type")
	}
}

var types = []PrimitiveType{
	{"Byte", byte(0), protowire.VarintType, "AppendVarint", "uint64(%s)"},
	{"Bool", bool(false), protowire.VarintType, "AppendVarint", "uint64(1)"},
	{"Int32", int32(0), protowire.VarintType, "AppendVarint", "uint64(%s)"},
	{"Int64", int64(0), protowire.VarintType, "AppendVarint", "uint64(%s)"},
	{"Uint32", uint32(0), protowire.VarintType, "AppendVarint", "uint64(%s)"},
	{"Uint64", uint64(0), protowire.VarintType, "AppendVarint", "%s"},
	{"Sint32", int32(0), protowire.VarintType, "AppendVarint", "uint64(encodeZigZag32(%s))"},
	{"Sint64", int64(0), protowire.VarintType, "AppendVarint", "protowire.EncodeZigZag(%s)"},
	{"Fixed32", uint32(0), protowire.Fixed32Type, "AppendFixed32", "%s"},
	{"Fixed64", uint64(0), protowire.Fixed64Type, "AppendFixed64", "%s"},
	{"Sfixed32", int32(0), protowire.Fixed32Type, "AppendFixed32", "encodeZigZag32(%s)"},
	{"Sfixed64", int64(0), protowire.Fixed64Type, "AppendFixed64", "protowire.EncodeZigZag(%s)"},
	{"Float", float32(0), protowire.Fixed32Type, "AppendFixed32", "math.Float32bits(%s)"},
	{"Double", float64(0), protowire.Fixed64Type, "AppendFixed64", "math.Float64bits(%s)"},
	{"String", string(""), protowire.BytesType, "AppendString", "%s"},
	{"RawString", string(""), protowire.BytesType, "AppendString", "%s"},
	{"Bytes", []byte{}, protowire.BytesType, "AppendBytes", "%s"},
	// {"Message", Message, protowire.BytesType}, // custom implementation
}

func main() {
	var err error

	err = ioutil.WriteFile("encoder_types.go", generateEncoder(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("decoder_types.go", generateDecoder(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("codec_types.go", generateCodec(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func generateEncoder() []byte {
	var b bytes.Buffer
	pf := func(format string, args ...interface{}) {
		fmt.Fprintf(&b, format, args...)
	}
	pf("// Copyright (C) 2021 Storj Labs, Inc.\n")
	pf("// See LICENSE for copying information.\n")
	pf("\n")
	pf("package picobuf\n\n")
	pf(`import ("math";` + "\n\n" + ` "google.golang.org/protobuf/encoding/protowire")`)
	pf("\n")

	for _, t := range types {
		pf("// %s encodes non-default %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
		pf("func (enc *Encoder) %s(field FieldNumber, v *%s) {\n", t.Name, t.TypeName())
		if t.Wire == protowire.BytesType {
			pf("if len(*v) == 0 { return }\n")
		} else if _, isBool := t.Zero.(bool); isBool {
			pf("if !*v { return }\n")
		} else {
			pf("if *v == %v { return }\n", t.Zero)
		}

		pf("enc.buffer = protowire.AppendTag(enc.buffer, protowire.Number(field), %s)\n", t.WireName())
		if strings.Contains(t.EncodeFmt, "%s") {
			pf("enc.buffer = protowire.%s(enc.buffer, "+t.EncodeFmt+")\n", t.Append, "*v")
		} else {
			pf("enc.buffer = protowire.%s(enc.buffer, "+t.EncodeFmt+")\n", t.Append)
		}

		pf("}\n")
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Println(b.String())
		log.Fatal(err)
	}

	return formatted
}

func generateDecoder() []byte {
	var b bytes.Buffer
	pf := func(format string, args ...interface{}) {
		fmt.Fprintf(&b, format, args...)
	}
	pf("// Copyright (C) 2021 Storj Labs, Inc.\n")
	pf("// See LICENSE for copying information.\n")
	pf("\n")
	pf("package picobuf\n\n")

	for _, t := range types {
		pf("// %s decodes %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
		pf("func (dec *Decoder) %s(field FieldNumber, v *%s) {}\n", t.Name, t.TypeName())
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Println(b.String())
		log.Fatal(err)
	}

	return formatted
}

func generateCodec() []byte {
	var b bytes.Buffer
	pf := func(format string, args ...interface{}) {
		fmt.Fprintf(&b, format, args...)
	}
	pf("// Copyright (C) 2021 Storj Labs, Inc.\n")
	pf("// See LICENSE for copying information.\n")
	pf("\n")
	pf("package picobuf\n\n")

	for _, t := range types {
		pf("// %s codes %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
		pf("//go:noinline\n")
		pf("func (codec *Codec) %s(field FieldNumber, v *%s) {\n", t.Name, t.TypeName())
		pf("if codec.encoding {\n")
		pf("  codec.encode.%s(field, v)\n", t.Name)
		pf("} else {\n")
		pf("  codec.decode.%s(field, v)\n", t.Name)
		pf("}\n")
		pf("}\n")
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Println(b.String())
		log.Fatal(err)
	}

	return formatted
}

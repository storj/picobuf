// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

// PrimitiveType represents a primitive type that the codec supports directly.
type PrimitiveType struct {
	Name      string
	Zero      interface{}
	Wire      protowire.Type
	Suffix    string // Suffix corresponds to `AppendSuffix` and `ConsumeSuffix`.
	EncodeFmt string
	DecodeFmt string // TODO: add bounds verification
}

// TypeName returns the corresponding Go type for the primitive.
func (t *PrimitiveType) TypeName() string {
	if _, isBytes := t.Zero.([]byte); isBytes {
		return "[]byte"
	}
	v := reflect.ValueOf(t.Zero)
	return v.Type().String()
}

// ZeroValue returns the zero value.
func (t *PrimitiveType) ZeroValue() string {
	if _, isBytes := t.Zero.([]byte); isBytes {
		return "nil"
	}
	return fmt.Sprintf("%#v", t.Zero)
}

// IsScalar returns whether the type is a scalar value.
func (t *PrimitiveType) IsScalar() bool {
	return t.Wire == protowire.VarintType || t.Wire == protowire.Fixed32Type || t.Wire == protowire.Fixed64Type
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

// ShortWireName returns the wire type for the errors.
func (t *PrimitiveType) ShortWireName() string {
	switch t.Wire {
	case protowire.VarintType:
		return "Varint"
	case protowire.Fixed32Type:
		return "Fixed32"
	case protowire.Fixed64Type:
		return "Fixed64"
	case protowire.BytesType:
		return "Bytes"
	case protowire.StartGroupType:
		return "StartGroup"
	case protowire.EndGroupType:
		return "EndGroup"
	default:
		panic("unhandled wire type")
	}
}

// IsValidMapKey returns whether primitive can be used as a map key.
func (t *PrimitiveType) IsValidMapKey() bool {
	return t.Name != "Float" && t.Name != "Double" && t.Name != "Bytes"
}

var types = []PrimitiveType{
	{"Bool", bool(false), protowire.VarintType, "Varint", "encodeBool64(%s)", "%s == 1"},
	{"Int32", int32(0), protowire.VarintType, "Varint", "uint64(%s)", "int32(%s)"},
	{"Int64", int64(0), protowire.VarintType, "Varint", "uint64(%s)", "int64(%s)"},
	{"Uint32", uint32(0), protowire.VarintType, "Varint", "uint64(%s)", "uint32(%s)"},
	{"Uint64", uint64(0), protowire.VarintType, "Varint", "%s", "%s"},
	{"Sint32", int32(0), protowire.VarintType, "Varint", "uint64(encodeZigZag32(%s))", "decodeZigZag32(uint32(%s))"},
	{"Sint64", int64(0), protowire.VarintType, "Varint", "protowire.EncodeZigZag(%s)", "protowire.DecodeZigZag(%s)"},
	{"Fixed32", uint32(0), protowire.Fixed32Type, "Fixed32", "%s", "%s"},
	{"Fixed64", uint64(0), protowire.Fixed64Type, "Fixed64", "%s", "%s"},
	{"Sfixed32", int32(0), protowire.Fixed32Type, "Fixed32", "encodeZigZag32(%s)", "decodeZigZag32(%s)"},
	{"Sfixed64", int64(0), protowire.Fixed64Type, "Fixed64", "protowire.EncodeZigZag(%s)", "protowire.DecodeZigZag(%s)"},
	{"Float", float32(0), protowire.Fixed32Type, "Fixed32", "math.Float32bits(%s)", "math.Float32frombits(%s)"},
	{"Double", float64(0), protowire.Fixed64Type, "Fixed64", "math.Float64bits(%s)", "math.Float64frombits(%s)"},
	{"String", string(""), protowire.BytesType, "String", "%s", "%s"},
	{"Bytes", []byte{}, protowire.BytesType, "Bytes", "%s", "%s"}, //TODO: reuse the existing bytes buffer to reduce allocs.
}

func main() {
	var err error

	err = os.WriteFile("encoder_types.go", generateEncoder(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("decoder_types.go", generateDecoder(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join("picowire", "map.go"), generateMaps(), 0644)
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
		{ // encoding a single value
			pf("// %s encodes non-default %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
			pf("//go:noinline\n")
			pf("func (enc *Encoder) %s(field FieldNumber, v *%s) {\n", t.Name, t.TypeName())
			if t.Wire == protowire.BytesType {
				pf("if len(*v) == 0 { return }\n")
			} else if _, isBool := t.Zero.(bool); isBool {
				pf("if !*v { return }\n")
			} else {
				pf("if *v == %v { return }\n", t.Zero)
			}

			pf("enc.buffer = appendTag(enc.buffer, field, %s)\n", t.WireName())
			pf("enc.buffer = protowire.Append%s(enc.buffer, "+t.EncodeFmt+")\n", t.Suffix, "*v")
			pf("}\n")
		}

		{ // encoding repeated values
			pf("// Repeated%s encodes non-empty repeated %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
			pf("//go:noinline\n")
			pf("func (enc *Encoder) Repeated%s(field FieldNumber, v *[]%s) {\n", t.Name, t.TypeName())
			pf("if len(*v) == 0 { return }\n")

			if t.IsScalar() {
				// generate packed encoding
				switch t.Wire {
				case protowire.VarintType:
					if t.Name == "Bool" {
						pf("enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)\n")
						pf("enc.buffer = protowire.AppendVarint(enc.buffer, uint64(len(*v)))\n")
						pf("for _, x := range *v {\n")
						pf("    enc.buffer = append(enc.buffer, encodeBool8(x))\n")
						pf("}\n")
					} else {
						pf("enc.alwaysAnyBytes(field, func() {\n")
						pf("    for _, x := range *v {\n")
						pf("         enc.buffer = protowire.Append%s(enc.buffer, "+t.EncodeFmt+")\n", t.Suffix, "x")
						pf("    }\n")
						pf("})\n")
					}
				case protowire.Fixed32Type:
					pf("enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)\n")
					pf("enc.buffer = protowire.AppendVarint(enc.buffer, uint64(len(*v)*4))\n")
					pf("for _, x := range *v {\n")
					pf("    enc.buffer = protowire.Append%s(enc.buffer, "+t.EncodeFmt+")\n", t.Suffix, "x")
					pf("}\n")
				case protowire.Fixed64Type:
					pf("enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)\n")
					pf("enc.buffer = protowire.AppendVarint(enc.buffer, uint64(len(*v)*8))\n")
					pf("for _, x := range *v {\n")
					pf("    enc.buffer = protowire.Append%s(enc.buffer, "+t.EncodeFmt+")\n", t.Suffix, "x")
					pf("}\n")
				default:
					panic("unhandled scalar wire type")
				}

			} else {
				pf("for _, x := range *v {\n")
				pf("    enc.buffer = appendTag(enc.buffer, field, %s)\n", t.WireName())
				pf("    enc.buffer = protowire.Append%s(enc.buffer, "+t.EncodeFmt+")\n", t.Suffix, "x")
				pf("}\n")
			}
			pf("}\n")
		}
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
	pf(`import ("math";` + "\n\n" + ` "google.golang.org/protobuf/encoding/protowire")`)
	pf("\n")

	for _, t := range types {
		{ // decoding single value
			pf("// %s decodes %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
			pf("//go:noinline\n")
			pf("func (dec *Decoder) %s(field FieldNumber, v *%s) {\n", t.Name, t.TypeName())

			pf("if field != dec.pendingField {\n")
			pf("    return\n")
			pf("}\n")

			pf("if dec.pendingWire != %v {\n", t.WireName())
			pf("    dec.fail(field, \"expected wire type %v\")\n", t.ShortWireName())
			pf("    return\n")
			pf("}\n")

			pf("x, n := protowire.Consume%v(dec.buffer)\n", t.Suffix)
			pf("if n < 0 { dec.fail(field, \"unable to parse %v\"); return }\n", t.Suffix)
			pf("*v = "+t.DecodeFmt+"\n", "x")
			pf("dec.nextField(n)\n")
			pf("}\n")
		}

		{ // decoding repeated values
			pf("// Repeated%s decodes repeated %s protobuf type.\n", t.Name, strings.ToLower(t.Name))
			pf("//go:noinline\n")
			pf("func (dec *Decoder) Repeated%s(field FieldNumber, v *[]%s) {\n", t.Name, t.TypeName())
			pf("for field == dec.pendingField {\n")

			pf("switch dec.pendingWire {")
			if t.IsScalar() {
				pf("case protowire.BytesType:\n")
				pf("    packed, n := protowire.ConsumeBytes(dec.buffer)\n")
				pf("    for len(packed) > 0 {\n")
				pf("         x, xn := protowire.Consume%v(packed)\n", t.Suffix)
				pf("         if xn < 0 { dec.fail(field, \"unable to parse %v\"); return }\n", t.Suffix)
				pf("         *v = append(*v, "+t.DecodeFmt+")\n", "x")
				pf("         packed = packed[xn:]\n")
				pf("    }\n")
				pf("    dec.nextField(n)\n")
			}
			pf("case %v:\n", t.WireName())
			pf("    x, n := protowire.Consume%v(dec.buffer)\n", t.Suffix)
			pf("    if n < 0 { dec.fail(field, \"unable to parse %v\"); return }\n", t.Suffix)
			pf("    *v = append(*v, "+t.DecodeFmt+")\n", "x")
			pf("    dec.nextField(n)\n")

			pf("default:\n")
			pf("    dec.fail(field, \"expected wire type %v\")\n", t.ShortWireName())
			pf("    return\n")
			pf("}\n")

			pf("}\n")
			pf("}\n")
		}
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Println(b.String())
		log.Fatal(err)
	}

	return formatted
}

func generateMaps() []byte {
	var b bytes.Buffer
	pf := func(format string, args ...interface{}) {
		fmt.Fprintf(&b, format, args...)
	}
	pf("// Copyright (C) 2021 Storj Labs, Inc.\n")
	pf("// See LICENSE for copying information.\n")
	pf("\n")
	pf("package picowire\n\n")
	pf(`import "storj.io/picobuf"`)
	pf("\n")

	for _, key := range types {
		if !key.IsValidMapKey() {
			continue
		}

		for _, val := range types {
			mapType := "Map" + key.Name + val.Name
			pf("// %s implements map<%s,%s>.\n", mapType, strings.ToLower(key.Name), strings.ToLower(val.Name))
			pf("type %s map[%s]%s\n\n", mapType, key.TypeName(), val.TypeName())

			pf("// PicoEncode encodes as protobuf.\n")
			pf("//go:noinline\n")
			pf("func(m *%s) PicoEncode(enc *picobuf.Encoder, field picobuf.FieldNumber) {\n", mapType)
			pf("   for key, val := range *m {\n")
			pf("       enc.AlwaysAnyBytes(field, func() {\n")
			pf("           enc.%s(1, &key)\n", key.Name)
			pf("           enc.%s(2, &val)\n", val.Name)
			pf("       })\n")
			pf("   }\n")
			pf("}\n\n")

			pf("// PicoDecode decodes as protobuf.\n")
			pf("//go:noinline\n")
			pf("func(m *%s) PicoDecode(dec *picobuf.Decoder, field picobuf.FieldNumber) {\n", mapType)
			pf("   var key %s\n", key.TypeName())
			pf("   var val %s\n", val.TypeName())
			pf("   dec.RepeatedMessage(field, func(c *picobuf.Decoder) {\n")
			pf("       if *m == nil {\n")
			pf("           *m = map[%s]%s{}\n", key.TypeName(), val.TypeName())
			pf("       }\n")
			pf("       c.Loop(func(c *picobuf.Decoder) {\n")
			pf("           c.%s(1, &key)\n", key.Name)
			pf("           c.%s(2, &val)\n", val.Name)
			pf("       })\n")
			pf("       (*m)[key] = val\n")
			pf("   })\n")
			pf("}\n\n")
		}
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		fmt.Println(b.String())
		log.Fatal(err)
	}

	return formatted
}

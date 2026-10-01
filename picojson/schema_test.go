// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"fmt"
	"strconv"
	"testing"

	"storj.io/picobuf/internal/protowire"
)

// wideMessage builds a schema and JSON document for the real JSON-to-wire
// parser. All fields have distinct protobuf and JSON names.
func wideMessage(count int, reversed, aliases bool) (Schema, []byte, []byte) {
	name := func(b []byte, s string) []byte {
		b = protowire.AppendVarint(b, uint64(len(s))<<1)
		return append(b, s...)
	}
	schema := name(nil, "Wide")
	schema = protowire.AppendVarint(schema, uint64(count))
	for i := 1; i <= count; i++ {
		schema = protowire.AppendVarint(schema, uint64(i))
		schema = append(schema, byte(int32Kind), flagJSONName)
		schema = name(schema, fmt.Sprintf("field_%d", i))
		schema = name(schema, fmt.Sprintf("field%d", i))
	}
	data := []byte{'{'}
	var wire []byte
	for i := 1; i <= count; i++ {
		n := i
		if reversed {
			n = count + 1 - i
		}
		if i > 1 {
			data = append(data, ',')
		}
		fieldName := "field" + strconv.Itoa(n)
		if aliases {
			fieldName = "field_" + strconv.Itoa(n)
		}
		data = strconv.AppendQuote(data, fieldName)
		data = append(data, ':')
		data = strconv.AppendInt(data, int64(n), 10)
		wire = protowire.AppendVarint(protowire.AppendTag(wire, protowire.Number(n), protowire.VarintType), uint64(n))
	}
	return Schema(schema), append(data, '}'), wire
}

func TestUnmarshalFieldOrder(t *testing.T) {
	for _, reversed := range []bool{false, true} {
		for _, aliases := range []bool{false, true} {
			t.Run(fmt.Sprintf("reversed=%v/aliases=%v", reversed, aliases), func(t *testing.T) {
				schema, data, want := wideMessage(1000, reversed, aliases)
				got, err := unmarshalWire(data, schema, 0)
				if err != nil {
					t.Fatal(err)
				}
				if string(got) != string(want) {
					t.Fatal("field lookup changed wire values or order")
				}
				// After wrapping, unknown names and duplicate aliases must still fail.
				for _, suffix := range []string{`,"unknown":1}`, `,"field_1":2}`} {
					bad := append(append([]byte(nil), data[:len(data)-1]...), suffix...)
					if _, err := unmarshalWire(bad, schema, 0); err == nil {
						t.Fatalf("accepted invalid suffix %s", suffix)
					}
				}
			})
		}
	}
}

func BenchmarkUnmarshalFieldOrder(b *testing.B) {
	for _, count := range []int{100, 500, 1000} {
		for _, reversed := range []bool{false, true} {
			b.Run(fmt.Sprintf("fields=%d/reversed=%v", count, reversed), func(b *testing.B) {
				schema, data, _ := wideMessage(count, reversed, false)
				b.ReportAllocs()
				b.SetBytes(int64(len(data)))
				b.ResetTimer()
				for b.Loop() {
					if _, err := unmarshalWire(data, schema, 0); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

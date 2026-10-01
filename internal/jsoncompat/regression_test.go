// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package jsoncompat

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"storj.io/picobuf/internal/jsoncompat/pico"
	"storj.io/picobuf/internal/jsoncompat/prot"
	"storj.io/picobuf/internal/protowire"
	"storj.io/picobuf/picowkt"
)

// checkPreservedWire also parses the output with protojson: comparing generic
// JSON values alone would hide duplicate object keys.
func checkPreservedWire(t *testing.T, m *pico.All) {
	t.Helper()
	want, ok := toProt(t, m)
	if !ok {
		t.Fatal("invalid test wire data")
	}
	data, err := m.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	var got prot.All
	if err := protojson.Unmarshal(data, &got); err != nil {
		t.Fatalf("invalid ProtoJSON %s: %v", data, err)
	}
	if !proto.Equal(&got, want) {
		t.Fatalf("JSON %s decoded to %v, want %v", data, &got, want)
	}
	canonical, err := protojson.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !sameJSON(t, data, canonical) {
		t.Fatalf("JSON %s, want %s", data, canonical)
	}
}

func TestJSONEmptyPackedWire(t *testing.T) {
	for _, tt := range []struct {
		name  string
		num   protowire.Number
		value []byte
	}{
		{"varint", 1, []byte{1}},
		{"fixed32", 5, protowire.AppendFixed32(nil, 1)},
		{"fixed64", 6, protowire.AppendFixed64(nil, 1)},
	} {
		empty := protowire.AppendBytes(protowire.AppendTag(nil, tt.num, protowire.BytesType), nil)
		value := protowire.AppendBytes(protowire.AppendTag(nil, tt.num, protowire.BytesType), tt.value)
		for name, parts := range map[string][][]byte{
			"before":  {empty, value},
			"between": {value, empty, value},
			"after":   {value, empty},
			"only":    {empty, empty},
		} {
			var wire []byte
			for _, part := range parts {
				wire = append(wire, part...)
			}
			for _, grouped := range []bool{false, true} {
				path := "stream"
				lists := &pico.Lists{XXX_unrecognized: wire}
				if grouped {
					path = "grouped"
					lists.String_ = []string{"later field"}
				}
				t.Run(tt.name+"/"+name+"/"+path, func(t *testing.T) {
					checkPreservedWire(t, &pico.All{Lists: lists})
				})
			}
		}
	}
}

func TestJSONOneofWire(t *testing.T) {
	name := []byte{10, 1, 'a'}
	count := []byte{24, 3}
	firstMessage := []byte{18, 2, 8, 1}   // scalars.int32 = 1
	secondMessage := []byte{18, 2, 16, 2} // scalars.int64 = 2
	for name, parts := range map[string][][]byte{
		"declaration order":       {name, count},
		"reverse order":           {count, name},
		"same message merges":     {firstMessage, secondMessage},
		"switched message resets": {firstMessage, name, secondMessage},
		"scalar replaces message": {firstMessage, count},
		"message replaces scalar": {name, firstMessage},
		"same scalar replaces":    {count, {24, 4}},
	} {
		t.Run(name, func(t *testing.T) {
			var wire []byte
			for _, part := range parts {
				wire = append(wire, part...)
			}
			checkPreservedWire(t, &pico.All{Oneofs: &pico.Oneofs{XXX_unrecognized: wire}})
		})
	}
}

func TestJSONDuplicateMapWire(t *testing.T) {
	entry := func(key string, value uint64) []byte {
		e := protowire.AppendString(protowire.AppendTag(nil, 1, protowire.BytesType), key)
		e = protowire.AppendVarint(protowire.AppendTag(e, 2, protowire.VarintType), value)
		return protowire.AppendBytes(protowire.AppendTag(nil, 1, protowire.BytesType), e)
	}
	for _, grouped := range []bool{false, true} {
		path := "stream"
		if grouped {
			path = "grouped"
		}
		for name, entries := range map[string][][]byte{
			"adjacent":    {entry("a", 1), entry("a", 3)},
			"nonadjacent": {entry("a", 1), entry("b", 2), entry("a", 3)},
			"escaped":     {entry("a\"", 1), entry("b", 2), entry("a\"", 3)},
			"default key": {entry("", 1), entry("b", 2), entry("", 3)},
		} {
			t.Run(name+"/"+path, func(t *testing.T) {
				maps := &pico.Maps{}
				for _, e := range entries {
					maps.XXX_unrecognized = append(maps.XXX_unrecognized, e...)
				}
				if grouped {
					maps.Int32String = map[int32]string{1: "later field"}
				}
				checkPreservedWire(t, &pico.All{Maps: maps})
			})
		}
	}
	t.Run("generated and preserved", func(t *testing.T) {
		checkPreservedWire(t, &pico.All{Maps: &pico.Maps{
			StringInt64:      map[string]int64{"a": 1},
			XXX_unrecognized: append(entry("b", 2), entry("a", 3)...),
		}})
	})
}

func TestJSONWrongWireType(t *testing.T) {
	for name, m := range map[string]*pico.All{
		"scalar after":      {Scalars: &pico.Scalars{Int32: 3, XXX_unrecognized: []byte{10, 1, 'a'}}},
		"scalar before":     {Scalars: &pico.Scalars{XXX_unrecognized: []byte{10, 1, 'a', 8, 3}}},
		"scalar only":       {Scalars: &pico.Scalars{XXX_unrecognized: []byte{10, 1, 'a'}}},
		"message":           {Scalars: &pico.Scalars{Int32: 3}, XXX_unrecognized: []byte{8, 1}},
		"list":              {Lists: &pico.Lists{XXX_unrecognized: []byte{13, 1, 0, 0, 0, 10, 1, 3, 13, 1, 0, 0, 0, 8, 4}}},
		"list grouped":      {Lists: &pico.Lists{Bool: []bool{true}, XXX_unrecognized: []byte{13, 1, 0, 0, 0, 10, 1, 3}}},
		"map container":     {Maps: &pico.Maps{StringInt64: map[string]int64{"a": 3}, XXX_unrecognized: []byte{8, 1}}},
		"map key and value": {Maps: &pico.Maps{XXX_unrecognized: []byte{10, 10, 10, 1, 'a', 8, 1, 16, 3, 18, 1, 'x'}}},
		"map defaults":      {Maps: &pico.Maps{XXX_unrecognized: []byte{10, 5, 8, 1, 18, 1, 'x'}}},
		"oneof":             {Oneofs: &pico.Oneofs{XXX_unrecognized: []byte{10, 1, 'a', 26, 1, 'x'}}},
		"oneof grouped":     {Oneofs: &pico.Oneofs{XXX_unrecognized: []byte{24, 3, 8, 1}}},
		"wrapper":           {WellKnown: &pico.WellKnown{Int32: &picowkt.Int32Value{Value: 3, XXX_unrecognized: []byte{10, 1, 'a'}}}},
		"timestamp":         {WellKnown: &pico.WellKnown{Timestamp: &picowkt.Timestamp{Seconds: 1, Nanos: 2, XXX_unrecognized: []byte{10, 0, 18, 0}}}},
		"duration":          {WellKnown: &pico.WellKnown{Duration: &picowkt.Duration{Seconds: 1, Nanos: 2, XXX_unrecognized: []byte{10, 0, 18, 0}}}},
		"mask":              {WellKnown: &pico.WellKnown{Mask: &picowkt.FieldMask{Paths: []string{"foo"}, XXX_unrecognized: []byte{8, 1}}}},
	} {
		t.Run(name, func(t *testing.T) { checkPreservedWire(t, m) })
	}
}

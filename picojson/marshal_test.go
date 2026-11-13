// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"slices"
	"testing"

	"storj.io/picobuf/internal/protowire"
)

// testSchema appends message records to a schema, see Schema.
type testSchema []byte

func (s *testSchema) name(text string) {
	*s = protowire.AppendVarint(*s, uint64(len(text))<<1)
	*s = append(*s, text...)
}

// message starts a message with count fields and returns its offset.
func (s *testSchema) message(name string, count int) int32 {
	offset := int32(len(*s))
	s.name(name)
	*s = protowire.AppendVarint(*s, uint64(count))
	return offset
}

// field appends a field, typ is the offset of a message field's type.
func (s *testSchema) field(number int32, kind Kind, flags byte, typ int32, name string) {
	*s = protowire.AppendVarint(*s, uint64(number))
	*s = append(*s, byte(kind), flags)
	if kind == messageKind {
		*s = protowire.AppendVarint(*s, uint64(typ))
	}
	s.name(name)
}

func TestMarshalUnorderedMessage(t *testing.T) {
	// message Leaf { int32 x = 1; }
	// message Unordered { int32 b = 2; int32 a = 1; }
	// message Parent { Unordered u = 1; repeated Leaf leaves = 2; }
	var s testSchema
	leaf := s.message("Leaf", 1)
	s.field(1, int32Kind, 0, 0, "x")
	unordered := s.message("Unordered", 2)
	s.field(2, int32Kind, 0, 0, "b")
	s.field(1, int32Kind, 0, 0, "a")
	parent := s.message("Parent", 2)
	s.field(1, messageKind, 0, unordered, "u")
	s.field(2, messageKind, flagList, leaf, "leaves")

	tests := []struct {
		name      string
		wire      []byte
		want      string
		irregular bool
		unordered []int32
	}{
		// picobuf encodes a before b, which Unordered declares after it.
		{"both set", []byte{0x0a, 4, 0x08, 1, 0x10, 2, 0x12, 2, 0x08, 1, 0x12, 2, 0x08, 2},
			`{"u":{"b":2,"a":1},"leaves":[{"x":1},{"x":2}]}`, false, []int32{unordered}},
		{"repeated", []byte{0x0a, 4, 0x08, 1, 0x10, 2, 0x0a, 4, 0x08, 3, 0x10, 4},
			`{"u":{"b":4,"a":3}}`, true, []int32{unordered}},
		{"one set", []byte{0x0a, 2, 0x10, 2, 0x12, 2, 0x08, 1},
			`{"u":{"b":2},"leaves":[{"x":1}]}`, false, nil},
		// A field set twice is irregular for any message.
		{"irregular", []byte{0x0a, 2, 0x10, 2, 0x12, 4, 0x08, 1, 0x08, 2},
			`{"u":{"b":2},"leaves":[{"x":2}]}`, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := marshaler{schema: Schema(s)}
			if err := m.message(parent, tt.wire); err != nil {
				t.Fatal(err)
			}
			if string(m.out) != tt.want {
				t.Fatalf("got %s, want %s", m.out, tt.want)
			}
			if m.irregular != tt.irregular || !slices.Equal(m.unordered, tt.unordered) {
				t.Fatalf("irregular=%v unordered=%v, want %v %v", m.irregular, m.unordered, tt.irregular, tt.unordered)
			}
		})
	}
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"fmt"
	"strings"
	"testing"

	"storj.io/picobuf/internal/protowire"
)

// nodeSchema is message Node { Node child = 1; string text = 2; }, which
// refers to itself at schema offset zero.
const nodeSchema = Schema("\x08Node\x02" +
	"\x01\x0b\x00\x00\x0achild" +
	"\x02\x09\x00\x08text")

func nestedNode(depth int, text string) []byte {
	return []byte(strings.Repeat(`{"child":`, depth) + `{"text":"` + text + `"}` + strings.Repeat("}", depth))
}

func TestUnmarshalNestedLengths(t *testing.T) {
	// The sizes straddle the one, two, three, and four byte length prefixes.
	for _, size := range []int{0, 1, 127, 128, 16383, 16384, 1 << 21} {
		for _, depth := range []int{1, 3, 200} {
			t.Run(fmt.Sprintf("size=%d/depth=%d", size, depth), func(t *testing.T) {
				text := strings.Repeat("x", size)
				wire, err := unmarshalWire(nestedNode(depth, text), nodeSchema, 0)
				if err != nil {
					t.Fatal(err)
				}
				for level := range depth {
					num, typ, n := protowire.ConsumeTag(wire)
					child, m := protowire.ConsumeBytes(wire[max(n, 0):])
					if num != 1 || typ != protowire.BytesType || m < 0 || n+m != len(wire) {
						t.Fatalf("level %d: child does not span its parent", level)
					}
					wire = child
				}
				num, typ, n := protowire.ConsumeTag(wire)
				got, m := protowire.ConsumeString(wire[max(n, 0):])
				if num != 2 || typ != protowire.BytesType || n+m != len(wire) || got != text {
					t.Fatal("innermost text changed")
				}
			})
		}
	}
}

func BenchmarkUnmarshalNested(b *testing.B) {
	text := strings.Repeat("x", 1<<20)
	for _, depth := range []int{1, 1000, 8000} {
		b.Run(fmt.Sprintf("depth=%d", depth), func(b *testing.B) {
			data := nestedNode(depth, text)
			b.SetBytes(int64(len(data)))
			for b.Loop() {
				if _, err := unmarshalWire(data, nodeSchema, 0); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

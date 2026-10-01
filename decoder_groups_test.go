// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"bytes"
	"testing"

	"storj.io/picobuf"
	"storj.io/picobuf/picowkt"
)

func TestUnknownGroupsRespectDepth(t *testing.T) {
	for _, depth := range []int{1, 2, 100} {
		data := append(bytes.Repeat([]byte{0x0b}, depth), bytes.Repeat([]byte{0x0c}, depth)...)
		err := (picobuf.UnmarshalOptions{MaxRecursionDepth: 1}).Unmarshal(data, new(picowkt.Empty))
		if (err != nil) != (depth > 1) {
			t.Fatalf("depth %d: %v", depth, err)
		}
	}
}

func TestGroupBoundaries(t *testing.T) {
	for _, mode := range []string{"Group", "RepeatedGroup"} {
		for _, data := range [][]byte{{0x0b}, {0x0b, 0x14}, {0x0c}, {0x0b, 0x0b, 0x0c}} {
			if err := picobuf.Unmarshal(data, &stateMessage{mode: mode}); err == nil {
				t.Fatalf("%s accepted %x", mode, data)
			}
		}
		m := &stateMessage{mode: mode}
		if err := picobuf.Unmarshal([]byte{0x0b, 0x10, 7, 0x8c, 0, 0x10, 9}, m); err != nil || m.value != 9 {
			t.Fatalf("%s: value %d, err %v", mode, m.value, err)
		}
	}
}

func BenchmarkNestedGroups(b *testing.B) {
	data := append(bytes.Repeat([]byte{0x0b}, 8000), bytes.Repeat([]byte{0x10, 0}, 10000)...)
	data = append(data, bytes.Repeat([]byte{0x0c}, 8000)...)
	m := &stateMessage{mode: "Group"}
	b.SetBytes(int64(len(data)))
	for b.Loop() {
		if err := picobuf.Unmarshal(data, m); err != nil {
			b.Fatal(err)
		}
	}
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"storj.io/picobuf"
)

func TestUnknownWirePolicies(t *testing.T) {
	// Field 2 is known, but its bytes occurrence must follow the unknown policy.
	data := []byte{0x12, 0, 0x10, 7}
	var ignored IgnoreUnknownMessage
	if err := picobuf.Unmarshal(data, &ignored); err != nil || ignored.Second != 7 {
		t.Fatalf("ignore: %v %+v", err, ignored)
	}
	// Skipped runs of unknown fields must not hide the known fields after them.
	data = []byte{0x08, 1, 0x18, 3, 0x20, 4, 0x28, 5, 0x30, 6, 0x10, 2}
	ignored = IgnoreUnknownMessage{}
	if err := picobuf.Unmarshal(data, &ignored); err != nil || ignored.Second != 2 || ignored.Fourth != 4 {
		t.Fatalf("ignore: %v %+v", err, ignored)
	}
	var rejected RejectUnknownMessage
	if err := picobuf.Unmarshal(data, &rejected); err == nil {
		t.Fatal("reject policy accepted mismatched wire")
	}
}

func BenchmarkDecodeUnrecognizedFields(b *testing.B) {
	// Fields that IgnoreUnknownMessage does not know, as sent by a newer schema.
	var distinct, repeated []byte
	for i := range 64 {
		distinct = protowire.AppendTag(distinct, protowire.Number(1000+i), protowire.VarintType)
		distinct = protowire.AppendVarint(distinct, uint64(i))
		repeated = protowire.AppendTag(repeated, 1000, protowire.BytesType)
		repeated = protowire.AppendBytes(repeated, []byte("hello"))
	}
	for _, bench := range []struct {
		name string
		data []byte
	}{
		{"Distinct", distinct},
		{"Repeated", repeated},
	} {
		b.Run(bench.name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(bench.data)))
			for b.Loop() {
				var m IgnoreUnknownMessage
				if err := picobuf.Unmarshal(bench.data, &m); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"testing"

	"github.com/zeebo/assert"
	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/protocompat/pico"
	"storj.io/picobuf/internal/protocompat/prot"
	"storj.io/picobuf/internal/protowire"
)

// TestMalformed checks that adversarial input is rejected rather than panicking
// or being silently truncated. picobuf may reject more than protobuf-go, since
// it does not implement groups, but must never accept what protobuf-go rejects.
func TestMalformed(t *testing.T) {
	tag := func(num protowire.Number, typ protowire.Type) []byte {
		return protowire.AppendTag(nil, num, typ)
	}
	cat := func(bs ...[]byte) (out []byte) {
		for _, b := range bs {
			out = append(out, b...)
		}
		return out
	}
	varint := func(x uint64) []byte { return protowire.AppendVarint(nil, x) }

	// RepeatedMixed is `int32 int32 = 1; repeated Message message = 16;`.
	tests := []struct {
		name string
		data []byte
	}{
		{"empty", nil},
		{"tag-only", tag(1, protowire.VarintType)},
		{"varint-truncated", cat(tag(1, protowire.VarintType), []byte{0xff, 0xff})},
		{"varint-overlong", cat(tag(1, protowire.VarintType), []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x00})},
		{"fixed32-truncated", cat(tag(1, protowire.Fixed32Type), []byte{1, 2})},
		{"wiretype-6", cat(tag(1, 6), []byte{1})},
		{"wiretype-7", cat(tag(1, 7), []byte{1})},
		{"end-group-only", tag(1, protowire.EndGroupType)},
		{"group-mismatch", cat(tag(1, protowire.StartGroupType), tag(2, protowire.EndGroupType))},
		{"group-unterminated", tag(1, protowire.StartGroupType)},
		{"submsg-len-overflow", cat(tag(16, protowire.BytesType), varint(1<<40))},
		{"submsg-truncated", cat(tag(16, protowire.BytesType), varint(50), []byte{1, 2})},
		{"submsg-len-negative", cat(tag(16, protowire.BytesType), []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f})},
		{"field-above-max", cat(tag(1, protowire.VarintType), varint(7), varint(uint64(1<<29)<<3), varint(1))},
		{"field-at-maxint32", cat(tag(1, protowire.VarintType), varint(7), varint(uint64(1<<31-1)<<3), varint(1))},
		{"unrecognized-truncated", cat(tag(5, protowire.BytesType), []byte{0x7f})},
	}

	for _, test := range tests {
		var reference prot.RepeatedMixed
		rejected := proto.Unmarshal(test.data, &reference) != nil

		var decoded pico.RepeatedMixed
		err := picobuf.Unmarshal(test.data, &decoded)
		if rejected {
			assert.Error(t, err)
		}
		t.Log(test.name, "->", err)
	}
}

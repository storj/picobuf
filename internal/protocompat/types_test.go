// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"testing"

	"github.com/zeebo/assert"
	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
)

func TestTypes(t *testing.T) {
	tests := []TypesPico{
		{},
		{
			Int32:    1,
			Int64:    1,
			Uint32:   1,
			Uint64:   1,
			Sint32:   1,
			Sint64:   1,
			Fixed32:  1,
			Fixed64:  1,
			Sfixed32: 1,
			Sfixed64: 1,
			Float:    1,
			Double:   1,
			Bool:     true,
			String_:  "1",
			Bytes:    []byte{1},
			Message: &MessagePico{
				Int32: 1,
			},
		},
		{
			Message: &MessagePico{},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p Types
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		opts := proto.MarshalOptions{Deterministic: true}
		canonical, err := opts.Marshal(&p)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, canonical, data)

		var got TypesPico
		err = picobuf.Unmarshal(canonical, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestRepeated(t *testing.T) {
	tests := []RepeatedTypesPico{
		{},
		{
			Int32:    []int32{1, 0xFFFF, -1},
			Int64:    []int64{1, 0xFFFFFFFF, -1},
			Uint32:   []uint32{1, 0xFFFF},
			Uint64:   []uint64{1, 0xFFFFFFFF},
			Sint32:   []int32{1, 0xFFFF},
			Sint64:   []int64{1, 0xFFFFFFFF},
			Fixed32:  []uint32{1, 0xFFFF},
			Fixed64:  []uint64{1, 0xFFFFFFFF},
			Sfixed32: []int32{1, 0xFFFF},
			Sfixed64: []int64{1, 0xFFFFFFFF},
			Float:    []float32{1, 1024},
			Double:   []float64{1, 1024},
			Bool:     []bool{true, false, true},
			String_:  []string{"hello", "world"},
			Bytes:    [][]byte{{}, {0, 1}, {0xff}},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p RepeatedTypes
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		opts := proto.MarshalOptions{Deterministic: true}
		canonical, err := opts.Marshal(&p)
		assert.NoError(t, err)
		assert.Equal(t, canonical, data)

		var got RepeatedTypesPico
		err = picobuf.Unmarshal(canonical, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

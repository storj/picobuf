// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"testing"

	"github.com/zeebo/assert"
	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
)

func TestDecodingMixed(t *testing.T) {
	var x1 RepeatedMixedPico
	err := picobuf.Unmarshal([]byte{130, 1, 0, 8, 123, 130, 1, 0}, &x1)
	assert.NoError(t, err)
	assert.DeepEqual(t, x1, RepeatedMixedPico{Int32: 123, Message: []*MessagePico{{}, {}}})

	var x2 RepeatedMixedPico
	err = picobuf.Unmarshal([]byte{130, 1, 4, 16, 56, 8, 109, 8, 123, 130, 1, 4, 16, 56, 8, 109}, &x2)
	assert.NoError(t, err)
	assert.DeepEqual(t, x2, RepeatedMixedPico{Int32: 123, Message: []*MessagePico{{Int32: 109, Int64: 56}, {Int32: 109, Int64: 56}}})
}

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
			Message: MessagePico{
				Int32: 1,
			},
		},
		{
			OptionalMessage: &OptionalMessagePico{},
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
			Message: []*MessagePico{
				{Int32: 1},
				{Int32: 2},
			},
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

func TestMaps(t *testing.T) {
	tests := []MapPico{
		{},
		{
			Values: map[string]string{"a": "b"},
		},
		{
			Values: map[string]string{"": ""},
		},
		{
			Values: map[string]string{"empty": ""},
		},
		{
			Values: map[string]string{"": "v"},
		},
		{
			Values: map[string]string{
				"a": "b",
				"b": "c",
				"c": "d",
			},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p Map
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		opts := proto.MarshalOptions{Deterministic: true}
		canonical, err := opts.Marshal(&p)
		assert.NoError(t, err)

		// encoding of maps is not deterministic
		if len(test.Values) <= 1 {
			_, hasEmptyKey := test.Values[""]
			_, hasEmptyVal := test.Values["empty"]
			if !hasEmptyKey && !hasEmptyVal {
				assert.Equal(t, canonical, data)
			}
		}

		var got MapPico
		err = picobuf.Unmarshal(canonical, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestEnum(t *testing.T) {
	test := PersonPico{
		Primary: Language_ENGLISHPico,
		Spoken: []LanguagePico{
			Language_ENGLISHPico,
			Language_SPANISHPico,
			Language_FRENCHPico,
		},
	}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)

	{
		var got Person
		err = proto.Unmarshal(data, &got)
		assert.NoError(t, err)

		assert.Equal(t, int32(got.Primary), int32(test.Primary))
		assert.Equal(t, len(got.Spoken), len(test.Spoken))

		for i := 0; i < len(got.Spoken); i++ {
			assert.Equal(t, int32(got.Spoken[i]), int32(test.Spoken[i]))
		}
	}

	{
		var got PersonPico
		err = picobuf.Unmarshal(data, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

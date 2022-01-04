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
)

func TestDecodingMixed(t *testing.T) {
	var x1 pico.RepeatedMixed
	err := picobuf.Unmarshal([]byte{130, 1, 0, 8, 123, 130, 1, 0}, &x1)
	assert.NoError(t, err)
	assert.DeepEqual(t, x1, pico.RepeatedMixed{Int32: 123, Message: []*pico.Message{{}, {}}})

	var x2 pico.RepeatedMixed
	err = picobuf.Unmarshal([]byte{130, 1, 4, 16, 56, 8, 109, 8, 123, 130, 1, 4, 16, 56, 8, 109}, &x2)
	assert.NoError(t, err)
	assert.DeepEqual(t, x2, pico.RepeatedMixed{Int32: 123, Message: []*pico.Message{{Int32: 109, Int64: 56}, {Int32: 109, Int64: 56}}})
}

func TestTypes(t *testing.T) {
	tests := []pico.Types{
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
			Message: pico.Message{
				Int32: 1,
			},
			Language: pico.Language_ENGLISH,
		},
		{
			OptionalMessage: &pico.OptionalMessage{},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p prot.Types
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		opts := proto.MarshalOptions{Deterministic: true}
		canonical, err := opts.Marshal(&p)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, canonical, data)

		var got pico.Types
		err = picobuf.Unmarshal(canonical, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestRepeated(t *testing.T) {
	tests := []pico.RepeatedTypes{
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
			Message: []*pico.Message{
				{Int32: 1},
				{Int32: 2},
			},
			Language: []pico.Language{pico.Language_ENGLISH},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p prot.RepeatedTypes
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		opts := proto.MarshalOptions{Deterministic: true}
		canonical, err := opts.Marshal(&p)
		assert.NoError(t, err)
		assert.Equal(t, canonical, data)

		var got pico.RepeatedTypes
		err = picobuf.Unmarshal(canonical, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestMaps(t *testing.T) {
	tests := []pico.Map{
		{},
		{
			StringString: map[string]string{"a": "b"},
		},
		{
			StringString: map[string]string{"": ""},
		},
		{
			StringString: map[string]string{"empty": ""},
		},
		{
			StringString: map[string]string{"": "v"},
		},
		{
			StringString: map[string]string{
				"a": "b",
				"b": "c",
				"c": "d",
			},
		},
		{
			StringString:   map[string]string{"a": "b"},
			StringInt32:    map[string]int32{"a": 1},
			StringInt64:    map[string]int64{"a": 1},
			StringUint32:   map[string]uint32{"a": 1},
			StringUint64:   map[string]uint64{"a": 1},
			StringSint32:   map[string]int32{"a": 1},
			StringSint64:   map[string]int64{"a": 1},
			StringFixed32:  map[string]uint32{"a": 1},
			StringFixed64:  map[string]uint64{"a": 1},
			StringSfixed32: map[string]int32{"a": 1},
			StringSfixed64: map[string]int64{"a": 1},
			StringFloat:    map[string]float32{"a": 1},
			StringDouble:   map[string]float64{"a": 1},
			StringBool:     map[string]bool{"a": true},
			StringBytes:    map[string][]byte{"a": []byte("a")},
			Int32String:    map[int32]string{1: "a"},
			Int64String:    map[int64]string{1: "a"},
			Uint32String:   map[uint32]string{1: "a"},
			Uint64String:   map[uint64]string{1: "a"},
			Sint32String:   map[int32]string{1: "a"},
			Sint64String:   map[int64]string{1: "a"},
			Fixed32String:  map[uint32]string{1: "a"},
			Fixed64String:  map[uint64]string{1: "a"},
			Sfixed32String: map[int32]string{1: "a"},
			Sfixed64String: map[int64]string{1: "a"},
			BoolString:     map[bool]string{true: "a"},
		},
	}

	for _, test := range tests {
		data, err := picobuf.Marshal(&test)
		assert.NoError(t, err)

		var p prot.Map
		err = proto.Unmarshal(data, &p)
		assert.NoError(t, err)

		opts := proto.MarshalOptions{Deterministic: true}
		canonical, err := opts.Marshal(&p)
		assert.NoError(t, err)

		// encoding of maps is not deterministic
		if len(test.StringString) <= 1 {
			_, hasEmptyKey := test.StringString[""]
			_, hasEmptyVal := test.StringString["empty"]
			if !hasEmptyKey && !hasEmptyVal {
				assert.Equal(t, canonical, data)
			}
		}

		var got pico.Map
		err = picobuf.Unmarshal(canonical, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestEnum(t *testing.T) {
	test := pico.Person{
		Primary: pico.Language_ENGLISH,
		Spoken: []pico.Language{
			pico.Language_ENGLISH,
			pico.Language_SPANISH,
			pico.Language_FRENCH,
		},
	}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)

	{
		var got prot.Person
		err = proto.Unmarshal(data, &got)
		assert.NoError(t, err)

		assert.Equal(t, int32(got.Primary), int32(test.Primary))
		assert.Equal(t, len(got.Spoken), len(test.Spoken))

		for i := 0; i < len(got.Spoken); i++ {
			assert.Equal(t, int32(got.Spoken[i]), int32(test.Spoken[i]))
		}
	}

	{
		var got pico.Person
		err = picobuf.Unmarshal(data, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestOneOf(t *testing.T) {
	test := pico.CommandMessage{
		Class: "Hello",
		Command: &pico.CommandMessage_Name{
			Name: "Hello",
		},
	}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)

	{
		var got prot.CommandMessage
		err = proto.Unmarshal(data, &got)
		assert.NoError(t, err)

		assert.Equal(t, got.Class, test.Class)
		name, isName := got.Command.(*prot.CommandMessage_Name)
		assert.True(t, isName)
		assert.Equal(t, name.Name, "Hello")
	}

	{
		var got pico.CommandMessage
		err = picobuf.Unmarshal(data, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

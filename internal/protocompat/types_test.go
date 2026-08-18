// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"math"
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
			Int32:    -2,
			Int64:    -2,
			Sint32:   -2,
			Sint64:   -2,
			Sfixed32: -2,
			Sfixed64: -2,
			Float:    -2,
			Double:   -2,
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

// TestTypesSignedWire checks the signed types against protobuf-go in both
// directions. TestTypes derives its reference message from picobuf's own
// output, so it cannot catch picobuf and the reference agreeing on a wrong
// interpretation of the wire bytes.
func TestTypesSignedWire(t *testing.T) {
	for _, v := range []int64{-2, -1, 0, 1, 2, math.MinInt32, math.MaxInt32, math.MinInt64, math.MaxInt64} {
		p := pico.Types{
			Int32: int32(v), Int64: v,
			Sint32: int32(v), Sint64: v,
			Sfixed32: int32(v), Sfixed64: v,
		}
		r := prot.Types{
			Int32: int32(v), Int64: v,
			Sint32: int32(v), Sint64: v,
			Sfixed32: int32(v), Sfixed64: v,
		}

		// picobuf encodes what protobuf-go reads
		data, err := picobuf.Marshal(&p)
		assert.NoError(t, err)
		var gotProt prot.Types
		assert.NoError(t, proto.Unmarshal(data, &gotProt))
		assert.Equal(t, gotProt.Int32, r.Int32)
		assert.Equal(t, gotProt.Int64, r.Int64)
		assert.Equal(t, gotProt.Sint32, r.Sint32)
		assert.Equal(t, gotProt.Sint64, r.Sint64)
		assert.Equal(t, gotProt.Sfixed32, r.Sfixed32)
		assert.Equal(t, gotProt.Sfixed64, r.Sfixed64)

		// picobuf reads what protobuf-go encodes
		refData, err := proto.Marshal(&r)
		assert.NoError(t, err)
		var gotPico pico.Types
		assert.NoError(t, picobuf.Unmarshal(refData, &gotPico))
		assert.DeepEqual(t, gotPico, p)
	}
}

func TestHighFieldNumbers(t *testing.T) {
	// Field numbers up to 2^29-1 are valid; the tag varint grows past one byte
	// at 16 and the exclude bitmask in UnrecognizedFields stops at 64.
	test := pico.HighFields{Low: 1, Boundary: 2, FirstHigh: 3, Large: 4, Max: 5}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)

	var p prot.HighFields
	assert.NoError(t, proto.Unmarshal(data, &p))
	assert.Equal(t, p.Low, test.Low)
	assert.Equal(t, p.Boundary, test.Boundary)
	assert.Equal(t, p.FirstHigh, test.FirstHigh)
	assert.Equal(t, p.Large, test.Large)
	assert.Equal(t, p.Max, test.Max)

	canonical, err := proto.MarshalOptions{Deterministic: true}.Marshal(&p)
	assert.NoError(t, err)
	assert.Equal(t, canonical, data)

	var got pico.HighFields
	assert.NoError(t, picobuf.Unmarshal(canonical, &got))
	assert.DeepEqual(t, got, test)
}

func TestRepeated(t *testing.T) {
	tests := []pico.RepeatedTypes{
		{},
		{
			Int32:    []int32{1, 0xFFFF, -1},
			Int64:    []int64{1, 0xFFFFFFFF, -1},
			Uint32:   []uint32{1, 0xFFFF},
			Uint64:   []uint64{1, 0xFFFFFFFF},
			Sint32:   []int32{1, 0xFFFF, -1, math.MinInt32},
			Sint64:   []int64{1, 0xFFFFFFFF, -1, math.MinInt64},
			Fixed32:  []uint32{1, 0xFFFF},
			Fixed64:  []uint64{1, 0xFFFFFFFF},
			Sfixed32: []int32{1, 0xFFFF, -1, math.MinInt32},
			Sfixed64: []int64{1, 0xFFFFFFFF, -1, math.MinInt64},
			Float:    []float32{1, 1024, -1},
			Double:   []float64{1, 1024, -1},
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

func TestMapZeroEntries(t *testing.T) {
	// Zero keys and values are omitted on the wire, so each map entry must
	// start from the zero value rather than inheriting the previous entry's.
	test := pico.Map{
		StringInt64: map[string]int64{"a": 1, "b": 0, "": 0},
		BoolString:  map[bool]string{true: "a", false: ""},
	}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)

	var got pico.Map
	err = picobuf.Unmarshal(data, &got)
	assert.NoError(t, err)
	assert.DeepEqual(t, got, test)

	var p prot.Map
	err = proto.Unmarshal(data, &p)
	assert.NoError(t, err)
	assert.DeepEqual(t, p.StringInt64, test.StringInt64)
	assert.DeepEqual(t, p.BoolString, test.BoolString)
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

func TestOneOf_Nil(t *testing.T) {
	test := pico.CommandMessage{
		Class:   "Hello",
		Command: nil,
	}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)

	{
		var got prot.CommandMessage
		err = proto.Unmarshal(data, &got)
		assert.NoError(t, err)

		assert.Equal(t, got.Class, test.Class)
		assert.Nil(t, got.Command)
	}

	{
		var got pico.CommandMessage
		err = picobuf.Unmarshal(data, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

func TestOneOf_Default(t *testing.T) {
	test := pico.CommandMessage{
		Class:   "Hello",
		Command: &pico.CommandMessage_Count{},
	}

	data, err := picobuf.Marshal(&test)
	assert.NoError(t, err)
	assert.Equal(t,
		[]byte{0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x00},
		data,
	)

	{
		var got prot.CommandMessage
		err = proto.Unmarshal(data, &got)
		assert.NoError(t, err)

		assert.Equal(t, got.Class, test.Class)
		cmd, isName := got.Command.(*prot.CommandMessage_Count)
		assert.True(t, isName)
		assert.Equal(t, cmd.Count, 0)
	}

	{
		var got pico.CommandMessage
		err = picobuf.Unmarshal(data, &got)
		assert.NoError(t, err)
		assert.DeepEqual(t, got, test)
	}
}

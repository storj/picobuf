// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"testing"
	"time"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest"
	"storj.io/picobuf/internal/picotest/pic"
)

func TestDecoder_Types(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})

	type example struct {
		varint0      uint32
		varint1      uint32
		varint100    uint32
		varint255    uint32
		stringEmpty  string
		stringHello  string
		bytesZero    []byte
		bytesNumbers []byte
	}

	var fields example
	dec.Loop(func(c *picobuf.Decoder) {
		c.Uint32(1, &fields.varint0)
		c.Uint32(2, &fields.varint1)
		c.Uint32(3, &fields.varint100)
		c.Uint32(4, &fields.varint255)
		c.String(5, &fields.stringEmpty)
		c.String(6, &fields.stringHello)
		c.Bytes(7, &fields.bytesZero)
		c.Bytes(8, &fields.bytesNumbers)
	})

	assert.NoError(t, dec.Err())
	assert.Equal(t, fields, example{
		varint0:      0,
		varint1:      1,
		varint100:    100,
		varint255:    255,
		stringEmpty:  "",
		stringHello:  "hello",
		bytesZero:    []byte{0},
		bytesNumbers: []byte{1, 2, 3, 4},
	})
}

func TestDecoder_OutOfOrder(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})

	type example struct {
		varint0      uint32
		varint1      uint32
		varint100    uint32
		varint255    uint32
		stringEmpty  string
		stringHello  string
		bytesZero    []byte
		bytesNumbers []byte
	}

	var fields example
	dec.Loop(func(c *picobuf.Decoder) {
		c.Bytes(8, &fields.bytesNumbers)
		c.Uint32(1, &fields.varint0)
		c.String(5, &fields.stringEmpty)
		c.Bytes(7, &fields.bytesZero)
		c.Uint32(4, &fields.varint255)
		c.Uint32(2, &fields.varint1)
		c.Uint32(3, &fields.varint100)
		c.String(6, &fields.stringHello)
	})

	assert.NoError(t, dec.Err())
	assert.Equal(t, fields, example{
		varint0:      0,
		varint1:      1,
		varint100:    100,
		varint255:    255,
		stringEmpty:  "",
		stringHello:  "hello",
		bytesZero:    []byte{0},
		bytesNumbers: []byte{1, 2, 3, 4},
	})
}

func TestDecoder_SubMessage(t *testing.T) {
	var got picotest.Person

	err := picobuf.Unmarshal([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x6, 0xa, 0x4, 0x48, 0x6f, 0x6d, 0x65}, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, picotest.Person{
		Name: "Hello",
		Address: &picotest.Address{
			Street: "Home",
		},
	})

	var got2 picotest.Person
	err = picobuf.Unmarshal([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f}, &got2)
	assert.NoError(t, err)
	assert.Equal(t, got2, picotest.Person{
		Name:    "Hello",
		Address: nil,
	})
}

func TestDecoder_Repeated(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
	xs := []int32{}
	dec.Loop(func(c *picobuf.Decoder) { c.RepeatedInt32(4, &xs) })
	assert.Equal(t, xs, []int32{3, 270, 86942})

	dec = picobuf.NewDecoder([]byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05, 0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
	xs = []int32{}
	dec.Loop(func(c *picobuf.Decoder) { c.RepeatedInt32(4, &xs) })
	assert.Equal(t, xs, []int32{3, 270, 86942, 3, 270, 86942})
}

func TestDecoder_Map(t *testing.T) {
	// two map entries
	var m1 picotest.Map
	err := picobuf.Unmarshal([]byte{
		0x0a, 0x04, // field: 1, length: 4
		0x08, 0x01, 0x10, 0x01, // field: 1, val: 1; field: 2, val: 1
		0x0a, 0x04, // field: 1, length: 4
		0x08, 0x02, 0x10, 0x02, // field: 1, val: 2; field: 2, val: 2
	}, &m1)
	assert.NoError(t, err)
	assert.DeepEqual(t, m1.Values, map[int32]int32{1: 1, 2: 2})

	// missing key in map entry
	var m2 picotest.Map
	err = picobuf.Unmarshal([]byte{
		0x0a, 0x02, // field: 1, length: 2
		0x10, 0x01, // field: 2, val: 1
	}, &m2)
	assert.NoError(t, err)
	assert.DeepEqual(t, m2.Values, map[int32]int32{0: 1})

	// missing value in map entry
	var m3 picotest.Map
	err = picobuf.Unmarshal([]byte{
		0x0a, 0x02, // field: 1, length: 2
		0x10, 0x01, // field: 2, val: 1
	}, &m3)
	assert.NoError(t, err)
	assert.DeepEqual(t, m3.Values, map[int32]int32{0: 1})

	// duplicate key and value in map entry
	var m4 picotest.Map
	err = picobuf.Unmarshal([]byte{
		0x0a, 0x08, // field: 1, length: 8
		0x08, 0x01, 0x10, 0x01, // field: 1, val: 1; field: 2, val: 1
		0x08, 0x02, 0x10, 0x02, // field: 1, val: 2; field: 2, val: 2
	}, &m4)
	assert.NoError(t, err)
	assert.DeepEqual(t, m4.Values, map[int32]int32{2: 2})
}

func TestDecoder_CustomMessageTypes(t *testing.T) {
	var decoded picotest.CustomMessageTypes
	err := picobuf.Unmarshal([]byte{
		// Normal
		0x0a, 0x04, 0x08, 0x01, 0x10, 0x02,
		// CustomType
		0x12, 0x04, 0x08, 0x0b, 0x10, 0x0c,
		// PresentCustomType
		0x1a, 0x04, 0x08, 0x15, 0x10, 0x16,
		// CustomTypeCast
		0x22, 0x04, 0x08, 0x1f, 0x10, 0x20,
		// PresentCustomTypeCast
		0x2a, 0x04, 0x08, 0x29, 0x10, 0x2a,
		// RepeatedCustomType
		0x32, 0x04, 0x08, 0x33, 0x10, 0x34,
		0x32, 0x04, 0x08, 0x35, 0x10, 0x36,
		// RepeatedPresentCustomType
		0x3a, 0x04, 0x08, 0x3d, 0x10, 0x3e,
		0x3a, 0x04, 0x08, 0x3f, 0x10, 0x40,
		// RepeatedCustomTypeCast
		0x42, 0x4, 0x8, 0x47, 0x10, 0x48,
		0x42, 0x4, 0x8, 0x49, 0x10, 0x4a,
		// RepeatedPresentCustomTypeCast
		0x4a, 0x4, 0x8, 0x51, 0x10, 0x52,
		0x4a, 0x4, 0x8, 0x53, 0x10, 0x54,
	}, &decoded)
	assert.NoError(t, err)

	assert.DeepEqual(t, &decoded, &picotest.CustomMessageTypes{
		Normal: &picotest.Timestamp{
			Seconds: 1,
			Nanos:   2,
		},
		CustomType: &pic.Timestamp{
			Seconds: 11,
			Nanos:   12,
		},
		PresentCustomType: pic.Timestamp{
			Seconds: 21,
			Nanos:   22,
		},
		CustomTypeCast:        utcTimePtr(31, 32),
		PresentCustomTypeCast: utcTime(41, 42),
		RepeatedCustomType: []*pic.Timestamp{
			{Seconds: 51, Nanos: 52},
			{Seconds: 53, Nanos: 54},
		},
		RepeatedPresentCustomType: []pic.Timestamp{
			{Seconds: 61, Nanos: 62},
			{Seconds: 63, Nanos: 64},
		},
		RepeatedCustomTypeCast: []*time.Time{
			utcTimePtr(71, 72),
			utcTimePtr(73, 74),
		},
		RepeatedPresentCustomTypeCast: []time.Time{
			utcTime(81, 82),
			utcTime(83, 84),
		},
	})
}

func TestDecoder_CustomMessageTypes_Empty(t *testing.T) {
	var decoded picotest.CustomMessageTypes
	err := picobuf.Unmarshal([]byte{}, &decoded)
	assert.NoError(t, err)

	assert.DeepEqual(t, &decoded, &picotest.CustomMessageTypes{
		Normal:                nil,
		CustomType:            nil,
		PresentCustomType:     pic.Timestamp{},
		CustomTypeCast:        nil,
		PresentCustomTypeCast: time.Time{},
	})
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"testing"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest"
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

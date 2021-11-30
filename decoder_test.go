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
		byte0        byte
		byte1        byte
		byte100      byte
		byte255      byte
		stringEmpty  string
		stringHello  string
		bytesZero    []byte
		bytesNumbers []byte
	}

	var fields example
	dec.Loop(func(c *picobuf.Codec) bool {
		c.Byte(1, &fields.byte0)
		c.Byte(2, &fields.byte1)
		c.Byte(3, &fields.byte100)
		c.Byte(4, &fields.byte255)
		c.RawString(5, &fields.stringEmpty)
		c.RawString(6, &fields.stringHello)
		c.Bytes(7, &fields.bytesZero)
		c.Bytes(8, &fields.bytesNumbers)
		return true
	})

	assert.NoError(t, dec.Err())
	assert.Equal(t, fields, example{
		byte0:        0,
		byte1:        1,
		byte100:      100,
		byte255:      255,
		stringEmpty:  "",
		stringHello:  "hello",
		bytesZero:    []byte{0},
		bytesNumbers: []byte{1, 2, 3, 4},
	})
}

func TestDecoder_OutOfOrder(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})

	type example struct {
		byte0        byte
		byte1        byte
		byte100      byte
		byte255      byte
		stringEmpty  string
		stringHello  string
		bytesZero    []byte
		bytesNumbers []byte
	}

	var fields example
	dec.Loop(func(c *picobuf.Codec) bool {
		c.Bytes(8, &fields.bytesNumbers)
		c.Byte(1, &fields.byte0)
		c.RawString(5, &fields.stringEmpty)
		c.Bytes(7, &fields.bytesZero)
		c.Byte(4, &fields.byte255)
		c.Byte(2, &fields.byte1)
		c.Byte(3, &fields.byte100)
		c.RawString(6, &fields.stringHello)
		return true
	})

	assert.NoError(t, dec.Err())
	assert.Equal(t, fields, example{
		byte0:        0,
		byte1:        1,
		byte100:      100,
		byte255:      255,
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
	dec.Loop(func(c *picobuf.Codec) bool {
		c.RepeatedInt32(4, &xs)
		return true
	})
	assert.Equal(t, xs, []int32{3, 270, 86942})

	dec = picobuf.NewDecoder([]byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05, 0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
	xs = []int32{}
	dec.Loop(func(c *picobuf.Codec) bool {
		c.RepeatedInt32(4, &xs)
		return true
	})
	assert.Equal(t, xs, []int32{3, 270, 86942, 3, 270, 86942})
}

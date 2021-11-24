// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"testing"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
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

	dec.Byte(1, &fields.byte0)
	dec.Byte(2, &fields.byte1)
	dec.Byte(3, &fields.byte100)
	dec.Byte(4, &fields.byte255)
	dec.RawString(5, &fields.stringEmpty)
	dec.RawString(6, &fields.stringHello)
	dec.Bytes(7, &fields.bytesZero)
	dec.Bytes(8, &fields.bytesNumbers)

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
	var got Person

	err := picobuf.Unmarshal([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x6, 0xa, 0x4, 0x48, 0x6f, 0x6d, 0x65}, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, Person{
		Name: "Hello",
		Address: Address{
			Street: "Home",
		},
	})

	err = picobuf.Unmarshal([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f}, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, Person{
		Name:    "Hello",
		Address: Address{},
	})
}

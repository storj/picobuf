// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"testing"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
)

func pbyte(v byte) *byte       { return &v }
func pstring(v string) *string { return &v }
func pbytes(v []byte) *[]byte  { return &v }

func TestEncoder_Types(t *testing.T) {
	enc := picobuf.NewEncoder()

	enc.Byte(1, pbyte(0))
	enc.Byte(2, pbyte(1))
	enc.Byte(3, pbyte(100))
	enc.Byte(4, pbyte(255))

	enc.RawString(5, pstring(""))
	enc.RawString(6, pstring("hello"))

	enc.Bytes(7, pbytes([]byte{0}))
	enc.Bytes(8, pbytes([]byte{1, 2, 3, 4}))

	result := enc.Buffer()
	assert.Equal(t, result, []byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})
}

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func (person *Person) Picobuf(codec *picobuf.Codec) {
	codec.RawString(1, &person.Name)
	codec.Message(2, person.Address.Picobuf)
}

func (address *Address) Picobuf(codec *picobuf.Codec) {
	codec.RawString(1, &address.Street)
}

func TestEncoder_SubMessage(t *testing.T) {
	person := &Person{
		Name: "Hello",
		Address: Address{
			Street: "Home",
		},
	}

	data, err := picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x6, 0xa, 0x4, 0x48, 0x6f, 0x6d, 0x65})

	person = &Person{
		Name:    "Hello",
		Address: Address{},
	}

	data, err = picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f})
}

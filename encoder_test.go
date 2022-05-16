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

func puint32(v uint32) *uint32 { return &v }
func pstring(v string) *string { return &v }
func pbytes(v []byte) *[]byte  { return &v }

func TestEncoder_Types(t *testing.T) {
	enc := picobuf.NewEncoder()

	enc.Uint32(1, puint32(0))
	enc.Uint32(2, puint32(1))
	enc.Uint32(3, puint32(100))
	enc.Uint32(4, puint32(255))

	enc.String(5, pstring(""))
	enc.String(6, pstring("hello"))

	enc.Bytes(7, pbytes([]byte{0}))
	enc.Bytes(8, pbytes([]byte{1, 2, 3, 4}))

	result := enc.Buffer()
	assert.Equal(t, result, []byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})
}

func TestEncoder_Repeated(t *testing.T) {
	enc := picobuf.NewEncoder()

	xs := []int32{3, 270, 86942}
	enc.RepeatedInt32(4, &xs)
	result := enc.Buffer()
	assert.Equal(t, result, []byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
}

func TestEncoder_SubMessage(t *testing.T) {
	person := &picotest.Person{
		Name: "Hello",
		Address: &picotest.Address{
			Street: "Home",
		},
	}

	data, err := picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x6, 0xa, 0x4, 0x48, 0x6f, 0x6d, 0x65})

	person = &picotest.Person{
		Name:    "Hello",
		Address: nil,
	}
	data, err = picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f})

	person = &picotest.Person{
		Name:    "Hello",
		Address: &picotest.Address{},
	}
	data, err = picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0})
}

func TestEncoder_CustomMessageTypes(t *testing.T) {
	base := time.Unix(31, 32).UTC()
	msg := &picotest.CustomMessageTypes{
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
		CustomTypeCast:        &base,
		PresentCustomTypeCast: time.Unix(41, 42).UTC(),
		RepeatedCustomType: []*pic.Timestamp{
			{Seconds: 51, Nanos: 52},
			{Seconds: 53, Nanos: 54},
		},
		RepeatedPresentCustomType: []pic.Timestamp{
			{Seconds: 61, Nanos: 62},
			{Seconds: 63, Nanos: 64},
		},
	}
	data, err := picobuf.Marshal(msg)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{
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
	})
}

func TestEncoder_CustomMessageTypes_Empty(t *testing.T) {
	msg := &picotest.CustomMessageTypes{}
	data, err := picobuf.Marshal(msg)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0x1a, 0x0})
}

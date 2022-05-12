// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package pic

import (
	"storj.io/picobuf"
)

// ID is for testing custom type.
type ID [12]byte

// IsZero returns whether the id is zero.
func (id ID) IsZero() bool {
	return id == ID{}
}

// PicoEncode implements picobuf field encoding.
func (id *ID) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if id.IsZero() {
		return false
	}

	p := id[:]
	c.Bytes(field, &p)
	return true
}

// PicoDecode implements picobuf field decoding.
func (id *ID) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}

	var x []byte
	c.Bytes(field, &x)
	if len(x) != 12 {
		c.Fail(field, "ID, expected length 12")
		return
	}
	copy((*id)[:], x)
}

// RawString implements custom encoding for strings.
type RawString string

// PicoEncode implements custom encoding function.
func (id *RawString) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if *id == "" {
		return false
	}

	xs := []byte(*id)
	c.Bytes(field, &xs)

	return true
}

// PicoDecode implements custom decoding function.
func (id *RawString) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}

	var xs []byte
	c.Bytes(field, &xs)
	*id = (RawString)(xs)
}

// Timestamp implements a custom pico encoder/decoder.
type Timestamp struct {
	Seconds int64
	Nanos   int32
}

// PicoEncode implements picobuf field encoding.
func (t *Timestamp) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if t == nil {
		return false
	}

	c.AlwaysMessage(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &t.Seconds)
		c.Int32(2, &t.Nanos)
		return true
	})
	return true
}

// PicoDecode implements picobuf field decoding.
func (t *Timestamp) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	c.Message(field, func(c *picobuf.Decoder) {
		c.Int64(1, &t.Seconds)
		c.Int32(2, &t.Nanos)
	})
}

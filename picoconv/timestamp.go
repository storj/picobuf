// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"time"

	"storj.io/picobuf"
)

// Timestamp implements protobuf timestamp conversion to standard time.Time.
type Timestamp time.Time

// Protobuf timestamps must be between 0001-01-01T00:00:00Z and
// 9999-12-31T23:59:59.999999999Z, inclusive.
const (
	minTimestampSeconds = -62135596800
	maxTimestampSeconds = 253402300799
)

// PicoEncode implements custom encoding function.
func (t *Timestamp) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if t == nil {
		return false
	}
	z := time.Time(*t)
	if z.IsZero() {
		return false
	}

	seconds := z.Unix()
	nanos := int32(z.Nanosecond())
	if seconds < minTimestampSeconds || seconds > maxTimestampSeconds {
		c.Fail(field, "timestamp seconds out of range")
		return false
	}
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &seconds)
		c.Int32(2, &nanos)
		return true
	})

	return true
}

// PicoDecode implements custom decoding function.
func (t *Timestamp) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if c.PendingField() != field {
		return
	}

	var seconds int64
	var nanos int32
	c.Message(field, func(c *picobuf.Decoder) {
		c.Int64(1, &seconds)
		c.Int32(2, &nanos)
	})
	if c.Err() != nil {
		return
	}
	if seconds < minTimestampSeconds || seconds > maxTimestampSeconds {
		c.Fail(field, "timestamp seconds out of range")
		return
	}
	if nanos < 0 || nanos > 999999999 {
		c.Fail(field, "timestamp nanos out of range")
		return
	}

	*t = Timestamp(time.Unix(seconds, int64(nanos)).UTC())
}

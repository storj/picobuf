// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"math"
	"time"

	"storj.io/picobuf"
)

// Duration implements protobuf duration conversion to standard time.Duration.
type Duration time.Duration

// PicoEncode implements custom encoding function.
func (d *Duration) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	if d == nil {
		return false
	}
	z := time.Duration(*d)

	// No range check is needed on encode: time.Duration caps out around
	// ±292 years, well within the ±10000 years protobuf durations allow.
	n := z.Nanoseconds()
	seconds := n / 1e9
	nanos := int32(n - seconds*1e9)
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &seconds)
		c.Int32(2, &nanos)
		return true
	})

	return true
}

// PicoDecode implements custom decoding function.
func (d *Duration) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
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
	if seconds < -315576000000 || seconds > 315576000000 {
		c.Fail(field, "duration seconds out of range")
		return
	}
	if nanos < -999999999 || nanos > 999999999 {
		c.Fail(field, "duration nanos out of range")
		return
	}
	if (seconds < 0 && nanos > 0) || (seconds > 0 && nanos < 0) {
		c.Fail(field, "duration seconds and nanos have different signs")
		return
	}

	z := time.Duration(seconds) * time.Second
	overflow := z/time.Second != time.Duration(seconds)
	z += time.Duration(nanos) * time.Nanosecond
	overflow = overflow || (seconds < 0 && nanos < 0 && z > 0)
	overflow = overflow || (seconds > 0 && nanos > 0 && z < 0)
	if overflow {
		switch {
		case seconds < 0:
			*d = Duration(time.Duration(math.MinInt64))
			return
		case seconds > 0:
			*d = Duration(time.Duration(math.MaxInt64))
			return
		}
	}

	*d = Duration(z)
}

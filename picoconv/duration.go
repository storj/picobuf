// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"math"
	"time"

	"storj.io/picobuf"
)

// Duration converts between a protobuf duration and time.Duration.
type Duration time.Duration

// Protobuf durations must be within about ±10000 years.
const (
	maxDurationSeconds = 315576000000
	maxDurationNanos   = 999999999
)

// PicoEncode implements picobuf.CustomType.
func (d *Duration) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if d == nil {
		return
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
}

// PicoDecode implements picobuf.CustomType.
func (d *Duration) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if !c.MatchField(field, 2) {
		return
	}

	kept := picobuf.KeptMessageValue[secondsNanos](c, d)
	var v secondsNanos
	if kept != nil {
		v = *kept
	} else {
		v = durationParts(time.Duration(*d))
	}
	c.Message(field, v.decode)
	if c.Err() != nil {
		return
	}
	seconds, nanos := v.seconds, v.nanos
	if seconds < -maxDurationSeconds || seconds > maxDurationSeconds {
		c.Fail(field, "duration seconds out of range")
		return
	}
	if nanos < -maxDurationNanos || nanos > maxDurationNanos {
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
		// seconds is nonzero: nanos alone cannot overflow.
		z = math.MaxInt64
		if seconds < 0 {
			z = math.MinInt64
		}
	}
	*d = Duration(z)

	// A later occurrence starts from *d, unless *d saturated and cannot
	// reproduce v.
	if kept != nil {
		*kept = v
	} else if durationParts(z) != v {
		picobuf.KeepMessageValue(c, d, v)
	}
}

func durationParts(z time.Duration) secondsNanos {
	n := int64(z)
	return secondsNanos{n / 1e9, int32(n % 1e9)}
}

var _ picobuf.CustomType = (*Duration)(nil)

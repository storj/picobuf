// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import "storj.io/picobuf"

// Both time conversions retain their wire components until decoding finishes.
// Reconstructing them from time.Time loses the distinction between an unset
// value and year 1; reconstructing from time.Duration loses overflowing seconds.
type secondsNanos struct {
	seconds int64
	nanos   int32
}

func (v *secondsNanos) decode(c *picobuf.Decoder) {
	c.Int64(1, &v.seconds)
	c.Int32(2, &v.nanos)
}

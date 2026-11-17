// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import "storj.io/picobuf"

// secondsNanos holds the wire components of a timestamp or a duration. A field
// that occurs several times merges into the components, which the converted
// value cannot always reproduce: time.Time does not distinguish an unset value
// from year 1, and time.Duration saturates.
type secondsNanos struct {
	seconds int64
	nanos   int32
}

func (v *secondsNanos) decode(c *picobuf.Decoder) {
	c.Int64(1, &v.seconds)
	c.Int32(2, &v.nanos)
}

// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"time"

	"storj.io/picobuf"
)

// Timestamp converts between a protobuf timestamp and time.Time.
type Timestamp time.Time

// Protobuf timestamps must be between 0001-01-01T00:00:00Z and
// 9999-12-31T23:59:59.999999999Z, inclusive.
const (
	minTimestampSeconds = -62135596800
	maxTimestampSeconds = 253402300799
)

// PicoEncode implements picobuf.CustomType.
func (t *Timestamp) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) {
	if t == nil {
		return
	}
	z := time.Time(*t)

	// A zero time is encoded rather than skipped: gogo/protobuf and the
	// official implementation both encode it as seconds=-62135596800, and a
	// message that decodes and re-encodes a peer's bytes has to reproduce
	// them, which signature verification depends on.
	seconds := z.Unix()
	nanos := int32(z.Nanosecond())
	if seconds < minTimestampSeconds || seconds > maxTimestampSeconds {
		c.Fail(field, "timestamp seconds out of range")
		return
	}
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &seconds)
		c.Int32(2, &nanos)
		return true
	})
}

// PicoDecode implements picobuf.CustomType.
func (t *Timestamp) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	if !c.MatchField(field, 2) {
		return
	}

	kept := picobuf.KeptMessageValue[secondsNanos](c, t)
	var v secondsNanos
	if kept != nil {
		v = *kept
	} else {
		v = timestampParts(time.Time(*t))
	}
	c.Message(field, v.decode)
	if c.Err() != nil {
		return
	}
	seconds, nanos := v.seconds, v.nanos
	if seconds < minTimestampSeconds || seconds > maxTimestampSeconds {
		c.Fail(field, "timestamp seconds out of range")
		return
	}
	if nanos < 0 || nanos > 999999999 {
		c.Fail(field, "timestamp nanos out of range")
		return
	}

	*t = Timestamp(time.Unix(seconds, int64(nanos)).UTC())

	// A later occurrence starts from *t, unless *t cannot reproduce v.
	if kept != nil {
		*kept = v
	} else if timestampParts(time.Time(*t)) != v {
		picobuf.KeepMessageValue(c, t, v)
	}
}

// timestampParts returns the components of z. A zero time is the unset value,
// whose components are zero.
func timestampParts(z time.Time) secondsNanos {
	if z.IsZero() {
		return secondsNanos{}
	}
	return secondsNanos{z.Unix(), int32(z.Nanosecond())}
}

var _ picobuf.CustomType = (*Timestamp)(nil)

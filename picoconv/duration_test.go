// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"math"
	"testing"
	"time"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
)

type DurationMessage struct {
	Value Duration
}

func (m *DurationMessage) Encode(c *picobuf.Encoder) bool {
	m.Value.PicoEncode(c, 1)
	return true
}

func (m *DurationMessage) Decode(c *picobuf.Decoder) {
	m.Value.PicoDecode(c, 1)
}

func TestDuration(t *testing.T) {
	for _, tt := range []time.Duration{
		-290 * 365 * 24 * time.Hour,
		-time.Hour,
		-time.Second - time.Nanosecond,
		-time.Second,
		-time.Second + time.Nanosecond,
		-time.Nanosecond,
		0,
		-time.Nanosecond,
		time.Second - time.Nanosecond,
		time.Second,
		time.Second + time.Nanosecond,
		time.Hour,
		290 * 365 * 24 * time.Hour,
	} {
		data, err := picobuf.Marshal(&DurationMessage{Value: Duration(tt)})
		assert.NoError(t, err)

		var r DurationMessage
		err = picobuf.Unmarshal(data, &r)
		assert.NoError(t, err)
		assert.Equal(t, tt, r.Value)
	}
}

type RawDuration struct {
	Seconds int64
	Nanos   int32
}

func (d *RawDuration) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &d.Seconds)
		c.Int32(2, &d.Nanos)
		return true
	})
	return true
}

func (d *RawDuration) PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber) {
	c.Int64(1, &d.Seconds)
	c.Int32(2, &d.Nanos)
}

type RawDurationMessage struct {
	Value RawDuration
}

func (m *RawDurationMessage) Encode(c *picobuf.Encoder) bool {
	m.Value.PicoEncode(c, 1)
	return true
}

func (m *RawDurationMessage) Decode(c *picobuf.Decoder) {
	m.Value.PicoDecode(c, 1)
}

func TestDuration_Overflow(t *testing.T) {
	for _, tt := range []struct {
		seconds  int64
		nanos    int32
		expected time.Duration
	}{
		{seconds: 0, nanos: 0, expected: 0},
		{seconds: 1, nanos: 1, expected: time.Second + time.Nanosecond},
		{seconds: -300 * 365 * 24 * 60 * 60, nanos: 0, expected: math.MinInt64},
		{seconds: 300 * 365 * 24 * 60 * 60, nanos: 0, expected: math.MaxInt64},
		{seconds: math.MinInt64, nanos: math.MinInt32, expected: math.MinInt64},
		{seconds: math.MaxInt64, nanos: math.MaxInt32, expected: math.MaxInt64},
	} {
		data, err := picobuf.Marshal(&RawDurationMessage{RawDuration{Seconds: tt.seconds, Nanos: tt.nanos}})
		assert.NoError(t, err)

		var r DurationMessage
		err = picobuf.Unmarshal(data, &r)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, r.Value)
	}
}

func TestDuration_Encode(t *testing.T) {
	data, err := picobuf.Marshal(&DurationMessage{Value: Duration(time.Second + 2*time.Nanosecond)})
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x0a, 0x04, 0x08, 0x01, 0x10, 0x02}, data)
}

func TestDuration_Decode(t *testing.T) {
	var r DurationMessage
	err := picobuf.Unmarshal([]byte{0x0a, 0x04, 0x08, 0x01, 0x10, 0x02}, &r)
	assert.NoError(t, err)
	assert.Equal(t, Duration(time.Second+2*time.Nanosecond), r.Value)
}

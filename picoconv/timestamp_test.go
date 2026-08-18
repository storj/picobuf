// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"testing"
	"time"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
)

type rawTimestamp struct {
	seconds int64
	nanos   int32
}

func (t *rawTimestamp) PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool {
	c.Message(field, func(c *picobuf.Encoder) bool {
		c.Int64(1, &t.seconds)
		c.Int32(2, &t.nanos)
		return true
	})
	return true
}

type rawTimestampMessage struct {
	value rawTimestamp
}

func (m *rawTimestampMessage) Encode(c *picobuf.Encoder) bool {
	return m.value.PicoEncode(c, 1)
}

func (m *rawTimestampMessage) Decode(c *picobuf.Decoder) {}

type timestampMessage struct {
	value Timestamp
}

func (m *timestampMessage) Encode(c *picobuf.Encoder) bool {
	return m.value.PicoEncode(c, 1)
}

func (m *timestampMessage) Decode(c *picobuf.Decoder) {
	m.value.PicoDecode(c, 1)
}

func TestTimestampEncodeInvalid(t *testing.T) {
	for _, value := range []time.Time{
		time.Date(0, 12, 31, 23, 59, 59, 0, time.UTC),
		time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC),
	} {
		data, err := picobuf.Marshal(&timestampMessage{value: Timestamp(value)})
		assert.Error(t, err)
		assert.Nil(t, data)
	}
}

func TestTimestampInvalid(t *testing.T) {
	for _, value := range []rawTimestamp{
		{seconds: -62135596801},
		{seconds: 253402300800},
		{nanos: -1},
		{nanos: 1000000000},
	} {
		data, err := picobuf.Marshal(&rawTimestampMessage{value: value})
		assert.NoError(t, err)

		var message timestampMessage
		assert.Error(t, picobuf.Unmarshal(data, &message))
	}
}

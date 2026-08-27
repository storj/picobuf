// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv_test

import (
	"math"
	"testing"

	"github.com/zeebo/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"storj.io/picobuf"
	"storj.io/picobuf/picoconv"
)

type fieldMessage struct {
	encode func(c *picobuf.Encoder)
	decode func(c *picobuf.Decoder)
}

func (m *fieldMessage) Encode(c *picobuf.Encoder) bool { m.encode(c); return true }
func (m *fieldMessage) Decode(c *picobuf.Decoder)      { m.decode(c) }

// wrapperField serializes as a message with the wrapper at field 1.
type wrapperField interface {
	PicoEncode(c *picobuf.Encoder, field picobuf.FieldNumber) bool
	PicoDecode(c *picobuf.Decoder, field picobuf.FieldNumber)
}

func checkWrapper(t *testing.T, value wrapperField, decoded wrapperField, want proto.Message) {
	t.Helper()

	got, err := picobuf.Marshal(&fieldMessage{
		encode: func(c *picobuf.Encoder) { value.PicoEncode(c, 1) },
	})
	assert.NoError(t, err)

	payload, err := proto.Marshal(want)
	assert.NoError(t, err)
	expect := append([]byte{0x0a, byte(len(payload))}, payload...)
	assert.Equal(t, got, expect)

	err = picobuf.Unmarshal(got, &fieldMessage{
		decode: func(c *picobuf.Decoder) { decoded.PicoDecode(c, 1) },
	})
	assert.NoError(t, err)
	assert.Equal(t, decoded, value)
}

func TestWrappersMatchProtobuf(t *testing.T) {
	for _, v := range []bool{false, true} {
		x := picoconv.BoolValue(v)
		checkWrapper(t, &x, new(picoconv.BoolValue), wrapperspb.Bool(v))
	}
	for _, v := range []int32{0, 1, -1, 1 << 30} {
		x := picoconv.Int32Value(v)
		checkWrapper(t, &x, new(picoconv.Int32Value), wrapperspb.Int32(v))
	}
	for _, v := range []int64{0, 1, -1, 1 << 60} {
		x := picoconv.Int64Value(v)
		checkWrapper(t, &x, new(picoconv.Int64Value), wrapperspb.Int64(v))
	}
	for _, v := range []uint32{0, 1, 1 << 31} {
		x := picoconv.UInt32Value(v)
		checkWrapper(t, &x, new(picoconv.UInt32Value), wrapperspb.UInt32(v))
	}
	for _, v := range []uint64{0, 1, 1 << 63} {
		x := picoconv.UInt64Value(v)
		checkWrapper(t, &x, new(picoconv.UInt64Value), wrapperspb.UInt64(v))
	}
	for _, v := range []float32{0, float32(math.Copysign(0, -1)), 1.5, -1.5} {
		x := picoconv.FloatValue(v)
		checkWrapper(t, &x, new(picoconv.FloatValue), wrapperspb.Float(v))
	}
	for _, v := range []float64{0, math.Copysign(0, -1), 1.5, -1.5} {
		x := picoconv.DoubleValue(v)
		checkWrapper(t, &x, new(picoconv.DoubleValue), wrapperspb.Double(v))
	}
	for _, v := range []string{"", "hello"} {
		x := picoconv.StringValue(v)
		checkWrapper(t, &x, new(picoconv.StringValue), wrapperspb.String(v))
	}
	for _, v := range [][]byte{nil, {1, 2, 3}} {
		x := picoconv.BytesValue(v)
		checkWrapper(t, &x, new(picoconv.BytesValue), wrapperspb.Bytes(v))
	}
}

func TestWrapperNilEncodesNothing(t *testing.T) {
	var x *picoconv.BoolValue
	got, err := picobuf.Marshal(&fieldMessage{
		encode: func(c *picobuf.Encoder) { x.PicoEncode(c, 1) },
	})
	assert.NoError(t, err)
	assert.Equal(t, len(got), 0)
}

func TestWrapperDecodeMergesOccurrences(t *testing.T) {
	// The wrapper occurs twice, and the second occurrence is empty. Protobuf
	// merges the occurrences, so the value stays true.
	var x picoconv.BoolValue
	err := picobuf.Unmarshal([]byte{0x0a, 0x02, 0x08, 0x01, 0x0a, 0x00}, &fieldMessage{
		decode: func(c *picobuf.Decoder) { x.PicoDecode(c, 1) },
	})
	assert.NoError(t, err)
	assert.Equal(t, bool(x), true)
}

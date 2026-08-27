// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest"
)

func TestOptional(t *testing.T) {
	var o picobuf.Optional[int32]
	assert.That(t, !o.Present())
	assert.That(t, o.IsZero())
	assert.Nil(t, o.Ptr())
	_, ok := o.Get()
	assert.That(t, !ok)
	assert.Equal(t, o.GetOr(7), int32(7))

	o.Set(0)
	assert.That(t, o.Present())
	assert.That(t, !o.IsZero())
	v, ok := o.Get()
	assert.That(t, ok)
	assert.Equal(t, v, int32(0))
	assert.Equal(t, o.GetOr(7), int32(0))
	assert.NotNil(t, o.Ptr())

	o.Clear()
	assert.That(t, !o.Present())

	assert.Equal(t, picobuf.NewOptional[int32](3).GetOr(7), int32(3))
}

func TestOptional_EncodePresence(t *testing.T) {
	data, err := picobuf.Marshal(&picotest.OptionalValues{})
	assert.NoError(t, err)
	assert.Equal(t, len(data), 0)

	// present zero values are emitted, matching explicit presence semantics
	msg := picotest.OptionalValues{
		Int32:   picobuf.NewOptional[int32](0),
		String_: picobuf.NewOptional(""),
	}
	data, err = picobuf.Marshal(&msg)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0x08, 0x00, 0x12, 0x00})
}

func TestOptional_RoundTrip(t *testing.T) {
	msg := picotest.OptionalValues{
		Int32:  picobuf.NewOptional[int32](-3),
		Bytes:  picobuf.NewOptional([]byte{1, 2}),
		Double: picobuf.NewOptional(1.5),
		Enum:   picobuf.NewOptional(picotest.MapEnum_MAP_ENUM_SECOND),
		Time:   picobuf.NewOptional(time.Unix(1234567890, 987654321).UTC()),
	}
	data, err := picobuf.Marshal(&msg)
	assert.NoError(t, err)

	var decoded picotest.OptionalValues
	assert.NoError(t, picobuf.Unmarshal(data, &decoded))
	assert.DeepEqual(t, decoded, msg)
	assert.That(t, !decoded.String_.Present())
}

func TestOptional_ZeroTimePresent(t *testing.T) {
	// picoconv.Timestamp writes the zero time, so its presence survives.
	msg := picotest.OptionalValues{Time: picobuf.NewOptional(time.Time{})}
	data, err := picobuf.Marshal(&msg)
	assert.NoError(t, err)

	var decoded picotest.OptionalValues
	assert.NoError(t, picobuf.Unmarshal(data, &decoded))
	value, ok := decoded.Time.Get()
	assert.That(t, ok && value.IsZero())
}

func TestOptional_JSONNotSupported(t *testing.T) {
	_, err := json.Marshal(&picotest.OptionalValues{Int32: picobuf.NewOptional[int32](0)})
	assert.Error(t, err)
	var decoded picotest.OptionalValues
	assert.Error(t, json.Unmarshal([]byte(`{"int32":0}`), &decoded))
}

func TestOptional_NilMessageGetter(t *testing.T) {
	var m *picotest.OptionalValues
	assert.That(t, !m.GetInt32().Present())
}

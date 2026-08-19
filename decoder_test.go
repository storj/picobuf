// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"testing"
	"time"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest"
	"storj.io/picobuf/internal/picotest/pic"
	"storj.io/picobuf/internal/protowire"
)

type packedScalars struct {
	Unsigned []uint64
	Signed   []int64
}

type packedFixed struct {
	Fixed32 []uint32
	Fixed64 []uint64
}

func (m *packedFixed) Encode(c *picobuf.Encoder) bool {
	c.RepeatedFixed32(1, &m.Fixed32)
	c.RepeatedFixed64(2, &m.Fixed64)
	return true
}

func (m *packedFixed) Decode(c *picobuf.Decoder) {
	c.RepeatedFixed32(1, &m.Fixed32)
	c.RepeatedFixed64(2, &m.Fixed64)
}

func (m *packedScalars) Encode(c *picobuf.Encoder) bool {
	c.RepeatedUint64(1, &m.Unsigned)
	c.RepeatedInt64(2, &m.Signed)
	return true
}

func (m *packedScalars) Decode(c *picobuf.Decoder) {
	c.RepeatedUint64(1, &m.Unsigned)
	c.RepeatedInt64(2, &m.Signed)
}

func TestDecoder_Types(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})

	type example struct {
		varint0      uint32
		varint1      uint32
		varint100    uint32
		varint255    uint32
		stringEmpty  string
		stringHello  string
		bytesZero    []byte
		bytesNumbers []byte
	}

	var fields example
	dec.Loop(func(c *picobuf.Decoder) {
		c.Uint32(1, &fields.varint0)
		c.Uint32(2, &fields.varint1)
		c.Uint32(3, &fields.varint100)
		c.Uint32(4, &fields.varint255)
		c.String(5, &fields.stringEmpty)
		c.String(6, &fields.stringHello)
		c.Bytes(7, &fields.bytesZero)
		c.Bytes(8, &fields.bytesNumbers)
	})

	assert.NoError(t, dec.Err())
	assert.Equal(t, fields, example{
		varint0:      0,
		varint1:      1,
		varint100:    100,
		varint255:    255,
		stringEmpty:  "",
		stringHello:  "hello",
		bytesZero:    []byte{0},
		bytesNumbers: []byte{1, 2, 3, 4},
	})
}

func TestDecoder_OutOfOrder(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})

	type example struct {
		varint0      uint32
		varint1      uint32
		varint100    uint32
		varint255    uint32
		stringEmpty  string
		stringHello  string
		bytesZero    []byte
		bytesNumbers []byte
	}

	var fields example
	dec.Loop(func(c *picobuf.Decoder) {
		c.Bytes(8, &fields.bytesNumbers)
		c.Uint32(1, &fields.varint0)
		c.String(5, &fields.stringEmpty)
		c.Bytes(7, &fields.bytesZero)
		c.Uint32(4, &fields.varint255)
		c.Uint32(2, &fields.varint1)
		c.Uint32(3, &fields.varint100)
		c.String(6, &fields.stringHello)
	})

	assert.NoError(t, dec.Err())
	assert.Equal(t, fields, example{
		varint0:      0,
		varint1:      1,
		varint100:    100,
		varint255:    255,
		stringEmpty:  "",
		stringHello:  "hello",
		bytesZero:    []byte{0},
		bytesNumbers: []byte{1, 2, 3, 4},
	})
}

func TestDecoder_SubMessage(t *testing.T) {
	var got picotest.Person

	err := picobuf.Unmarshal([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x6, 0xa, 0x4, 0x48, 0x6f, 0x6d, 0x65}, &got)
	assert.NoError(t, err)
	assert.Equal(t, got, picotest.Person{
		Name: "Hello",
		Address: &picotest.Address{
			Street: "Home",
		},
	})

	var got2 picotest.Person
	err = picobuf.Unmarshal([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f}, &got2)
	assert.NoError(t, err)
	assert.Equal(t, got2, picotest.Person{
		Name:    "Hello",
		Address: nil,
	})
}

func TestDecoder_Repeated(t *testing.T) {
	dec := picobuf.NewDecoder([]byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
	xs := []int32{}
	dec.Loop(func(c *picobuf.Decoder) { c.RepeatedInt32(4, &xs) })
	assert.Equal(t, xs, []int32{3, 270, 86942})

	dec = picobuf.NewDecoder([]byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05, 0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
	xs = []int32{}
	dec.Loop(func(c *picobuf.Decoder) { c.RepeatedInt32(4, &xs) })
	assert.Equal(t, xs, []int32{3, 270, 86942, 3, 270, 86942})
}

func TestDecoder_PackedScalarsPreallocate(t *testing.T) {
	data := []byte{
		0x0a, 0x10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0x12, 0x10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	}

	got := packedScalars{Unsigned: []uint64{100}, Signed: []int64{-100}}
	assert.NoError(t, picobuf.Unmarshal(data, &got))
	assert.DeepEqual(t, got.Unsigned, []uint64{100, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	assert.DeepEqual(t, got.Signed, []int64{-100, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	if raceEnabled {
		t.Skip("allocation counts differ under the race detector")
	}
	allocs := testing.AllocsPerRun(100, func() {
		var decoded packedScalars
		if err := picobuf.Unmarshal(data, &decoded); err != nil {
			panic(err)
		}
	})
	if allocs > 4 {
		t.Fatalf("packed scalar decode allocated %.0f times, want at most 4", allocs)
	}
}

func TestDecoder_PackedFixedPreallocate(t *testing.T) {
	data := []byte{
		0x0a, 0x08, 1, 0, 0, 0, 2, 0, 0, 0,
		0x12, 0x10, 3, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0,
	}
	got := packedFixed{
		Fixed32: make([]uint32, 1, 4),
		Fixed64: make([]uint64, 1, 4),
	}
	got.Fixed32[0] = 100
	got.Fixed64[0] = 200

	assert.NoError(t, picobuf.Unmarshal(data, &got))
	assert.DeepEqual(t, got.Fixed32, []uint32{100, 1, 2})
	assert.DeepEqual(t, got.Fixed64, []uint64{200, 3, 4})
	assert.Equal(t, cap(got.Fixed32), 4)
	assert.Equal(t, cap(got.Fixed64), 4)
}

func TestDecoder_PackedPreallocationRespectsLimit(t *testing.T) {
	data := []byte{0x0a, 0x10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var got packedScalars
	err := (picobuf.UnmarshalOptions{MaxRepeatedElements: 2}).Unmarshal(data, &got)
	assert.Error(t, err)
	assert.DeepEqual(t, got.Unsigned, []uint64{0, 1})
	assert.Equal(t, cap(got.Unsigned), 2)
}

func TestDecoder_PackedFixedRejectsMalformedPayload(t *testing.T) {
	var got packedFixed
	err := picobuf.Unmarshal([]byte{0x0a, 0x05, 1, 0, 0, 0, 2}, &got)
	assert.Error(t, err)
}

func TestDecoder_PackedVarintRejectsMalformedPayload(t *testing.T) {
	var got packedScalars
	err := picobuf.Unmarshal([]byte{0x0a, 0x01, 0x80}, &got)
	assert.Error(t, err)
}

func TestDecoder_Map(t *testing.T) {
	// two map entries
	var m1 picotest.Map
	err := picobuf.Unmarshal([]byte{
		0x0a, 0x04, // field: 1, length: 4
		0x08, 0x01, 0x10, 0x01, // field: 1, val: 1; field: 2, val: 1
		0x0a, 0x04, // field: 1, length: 4
		0x08, 0x02, 0x10, 0x02, // field: 1, val: 2; field: 2, val: 2
	}, &m1)
	assert.NoError(t, err)
	assert.DeepEqual(t, m1.Values, map[int32]int32{1: 1, 2: 2})

	// missing key in map entry
	var m2 picotest.Map
	err = picobuf.Unmarshal([]byte{
		0x0a, 0x02, // field: 1, length: 2
		0x10, 0x01, // field: 2, val: 1
	}, &m2)
	assert.NoError(t, err)
	assert.DeepEqual(t, m2.Values, map[int32]int32{0: 1})

	// missing value in map entry
	var m3 picotest.Map
	err = picobuf.Unmarshal([]byte{
		0x0a, 0x02, // field: 1, length: 2
		0x10, 0x01, // field: 2, val: 1
	}, &m3)
	assert.NoError(t, err)
	assert.DeepEqual(t, m3.Values, map[int32]int32{0: 1})

	// duplicate key and value in map entry
	var m4 picotest.Map
	err = picobuf.Unmarshal([]byte{
		0x0a, 0x08, // field: 1, length: 8
		0x08, 0x01, 0x10, 0x01, // field: 1, val: 1; field: 2, val: 1
		0x08, 0x02, 0x10, 0x02, // field: 1, val: 2; field: 2, val: 2
	}, &m4)
	assert.NoError(t, err)
	assert.DeepEqual(t, m4.Values, map[int32]int32{2: 2})
}

func TestDecoder_CustomMessageTypes(t *testing.T) {
	var decoded picotest.CustomMessageTypes
	err := picobuf.Unmarshal([]byte{
		// Normal
		0x0a, 0x04, 0x08, 0x01, 0x10, 0x02,
		// CustomType
		0x12, 0x04, 0x08, 0x0b, 0x10, 0x0c,
		// PresentCustomType
		0x1a, 0x04, 0x08, 0x15, 0x10, 0x16,
		// CustomTypeCast
		0x22, 0x04, 0x08, 0x1f, 0x10, 0x20,
		// PresentCustomTypeCast
		0x2a, 0x04, 0x08, 0x29, 0x10, 0x2a,
		// RepeatedCustomType
		0x32, 0x04, 0x08, 0x33, 0x10, 0x34,
		0x32, 0x04, 0x08, 0x35, 0x10, 0x36,
		// RepeatedPresentCustomType
		0x3a, 0x04, 0x08, 0x3d, 0x10, 0x3e,
		0x3a, 0x04, 0x08, 0x3f, 0x10, 0x40,
		// RepeatedCustomTypeCast
		0x42, 0x4, 0x8, 0x47, 0x10, 0x48,
		0x42, 0x4, 0x8, 0x49, 0x10, 0x4a,
		// RepeatedPresentCustomTypeCast
		0x4a, 0x4, 0x8, 0x51, 0x10, 0x52,
		0x4a, 0x4, 0x8, 0x53, 0x10, 0x54,
	}, &decoded)
	assert.NoError(t, err)

	assert.DeepEqual(t, &decoded, &picotest.CustomMessageTypes{
		Normal: &picotest.Timestamp{
			Seconds: 1,
			Nanos:   2,
		},
		CustomType: &pic.Timestamp{
			Seconds: 11,
			Nanos:   12,
		},
		PresentCustomType: pic.Timestamp{
			Seconds: 21,
			Nanos:   22,
		},
		CustomTypeCast:        utcTimePtr(31, 32),
		PresentCustomTypeCast: utcTime(41, 42),
		RepeatedCustomType: []*pic.Timestamp{
			{Seconds: 51, Nanos: 52},
			{Seconds: 53, Nanos: 54},
		},
		RepeatedPresentCustomType: []pic.Timestamp{
			{Seconds: 61, Nanos: 62},
			{Seconds: 63, Nanos: 64},
		},
		RepeatedCustomTypeCast: []*time.Time{
			utcTimePtr(71, 72),
			utcTimePtr(73, 74),
		},
		RepeatedPresentCustomTypeCast: []time.Time{
			utcTime(81, 82),
			utcTime(83, 84),
		},
	})
}

func TestDecoder_CustomMessageTypes_Empty(t *testing.T) {
	var decoded picotest.CustomMessageTypes
	err := picobuf.Unmarshal([]byte{}, &decoded)
	assert.NoError(t, err)

	assert.DeepEqual(t, &decoded, &picotest.CustomMessageTypes{
		Normal:                nil,
		CustomType:            nil,
		PresentCustomType:     pic.Timestamp{},
		CustomTypeCast:        nil,
		PresentCustomTypeCast: time.Time{},
	})
}

// nested is a self-referential message, which picotest does not have.
type nested struct{ Inner *nested }

func (m *nested) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Message(1, m.Inner.Encode)
	return true
}

func (m *nested) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Message(1, func(c *picobuf.Decoder) {
		if m.Inner == nil {
			m.Inner = new(nested)
		}
		m.Inner.Decode(c)
	})
}

// nestedMessage builds depth levels of nested, innermost first.
func nestedMessage(depth int) []byte {
	var data []byte
	for range depth {
		if len(data) > 0x7f {
			t := []byte{0xa}
			t = protowire.AppendVarint(t, uint64(len(data)))
			data = append(t, data...)
			continue
		}
		data = append([]byte{0xa, byte(len(data))}, data...)
	}
	return data
}

func TestDecoder_RecursionLimit(t *testing.T) {
	var shallow nested
	assert.NoError(t, picobuf.Unmarshal(nestedMessage(100), &shallow))

	// Without a limit this exhausts the stack rather than returning an error.
	var deep nested
	assert.Error(t, picobuf.Unmarshal(nestedMessage(protowire.DefaultRecursionLimit+10), &deep))
}

func TestUnmarshalOptions_MaxInputSize(t *testing.T) {
	var message picotest.AllTypes
	err := (picobuf.UnmarshalOptions{MaxInputSize: 1}).Unmarshal([]byte{0x08, 0x01}, &message)
	assert.Error(t, err)

	assert.NoError(t, (picobuf.UnmarshalOptions{MaxInputSize: 2}).Unmarshal([]byte{0x08, 0x01}, &message))
}

func TestUnmarshalOptions_MaxRecursionDepth(t *testing.T) {
	var message nested
	err := (picobuf.UnmarshalOptions{MaxRecursionDepth: 2}).Unmarshal(nestedMessage(3), &message)
	assert.Error(t, err)

	assert.NoError(t, (picobuf.UnmarshalOptions{MaxRecursionDepth: 3}).Unmarshal(nestedMessage(3), &message))
}

func TestUnmarshalOptions_MaxRepeatedElements(t *testing.T) {
	data, err := picobuf.Marshal(&picotest.AllTypes{Int32S: []int32{1, 2, 3}})
	assert.NoError(t, err)

	var message picotest.AllTypes
	err = (picobuf.UnmarshalOptions{MaxRepeatedElements: 2}).Unmarshal(data, &message)
	assert.Error(t, err)

	message = picotest.AllTypes{}
	assert.NoError(t, (picobuf.UnmarshalOptions{MaxRepeatedElements: 3}).Unmarshal(data, &message))
	assert.DeepEqual(t, message.Int32S, []int32{1, 2, 3})
}

func TestUnmarshalOptions_InvalidUTF8(t *testing.T) {
	data := []byte{0x72, 0x01, 0xff, 0xf2, 0x01, 0x01, 0xff}

	var message picotest.AllTypes
	assert.Error(t, picobuf.Unmarshal(data, &message))

	assert.NoError(t, (picobuf.UnmarshalOptions{AllowInvalidUTF8: true}).Unmarshal(data, &message))
	assert.Equal(t, message.String_, string([]byte{0xff}))
	assert.DeepEqual(t, message.Strings, []string{string([]byte{0xff})})
}

func TestUnmarshalOptions_AliasInput(t *testing.T) {
	data, err := picobuf.Marshal(&picotest.AllTypes{
		String_: "hello",
		Bytes:   []byte{1, 2, 3},
		Bytess:  [][]byte{{4}, {5}},
	})
	assert.NoError(t, err)

	// By default the decoded message owns its bytes, so overwriting the input
	// leaves it untouched.
	buffer := append([]byte(nil), data...)
	var copied picotest.AllTypes
	assert.NoError(t, picobuf.Unmarshal(buffer, &copied))
	for i := range buffer {
		buffer[i] = 0xff
	}
	assert.Equal(t, copied.String_, "hello")
	assert.DeepEqual(t, copied.Bytes, []byte{1, 2, 3})
	assert.DeepEqual(t, copied.Bytess, [][]byte{{4}, {5}})

	// With AliasInput the bytes fields point into the input.
	buffer = append([]byte(nil), data...)
	var aliased picotest.AllTypes
	assert.NoError(t, picobuf.UnmarshalOptions{AliasInput: true}.Unmarshal(buffer, &aliased))
	assert.DeepEqual(t, aliased.Bytes, []byte{1, 2, 3})
	for i := range buffer {
		buffer[i] = 0xff
	}
	assert.DeepEqual(t, aliased.Bytes, []byte{0xff, 0xff, 0xff})
	assert.DeepEqual(t, aliased.Bytess, [][]byte{{0xff}, {0xff}})
	// Strings are copied either way.
	assert.Equal(t, aliased.String_, "hello")
}

func TestDecoder_InvalidFieldNumber(t *testing.T) {
	// ConsumeTag only rejects field numbers below the minimum, so numbers from
	// MaxValidNumber up to MaxInt32 reach the decoder. Loop considers those
	// invalid and stops, which used to discard the rest of the message and
	// still report success.
	for _, num := range []uint64{1 << 29, 1<<31 - 1} {
		data := appendField(nil, 1, 7)
		data = protowire.AppendVarint(data, num<<3)
		data = protowire.AppendVarint(data, 1)
		data = appendField(data, 2, 99)

		var m picotest.AllTypes
		assert.Error(t, picobuf.Unmarshal(data, &m))
	}

	// The largest valid field number is still accepted.
	var m picotest.AllTypes
	assert.NoError(t, picobuf.Unmarshal(appendField(nil, 1<<29-1, 1), &m))
}

func TestDecoder_Bool_NonOne(t *testing.T) {
	// Any non-zero varint decodes as true, as protobuf requires.
	var decoded picotest.AllTypes
	err := picobuf.Unmarshal([]byte{0x68, 0x02}, &decoded)
	assert.NoError(t, err)
	assert.That(t, decoded.Bool)
}

func TestDecoder_UnrecognizedFields_Empty(t *testing.T) {
	var decoded picotest.UnknownMessage
	err := picobuf.Unmarshal([]byte{}, &decoded)
	assert.NoError(t, err)

	assert.DeepEqual(t, decoded, picotest.UnknownMessage{
		Second:           0,
		Fourth:           0,
		XXX_unrecognized: nil,
	})
}

func TestDecoder_UnrecognizedFields_Decode(t *testing.T) {
	decode := func(data []byte) picotest.UnknownMessage {
		var decoded picotest.UnknownMessage
		err := picobuf.Unmarshal(data, &decoded)
		assert.NoError(t, err)
		return decoded
	}

	assert.DeepEqual(t, picotest.UnknownMessage{
		Second:           0x22,
		Fourth:           0x44,
		XXX_unrecognized: []byte{0x8, 0x13, 0x18, 0x25},
	}, decode([]byte{
		0x10, 0x22,
		0x20, 0x44,
		0x8, 0x13,
		0x18, 0x25,
	}))

	assert.DeepEqual(t, picotest.UnknownMessage{
		Second:           0x22,
		Fourth:           0x44,
		XXX_unrecognized: []byte{0x8, 0x13, 0x18, 0x25},
	}, decode([]byte{
		0x8, 0x13,
		0x18, 0x25,
		0x10, 0x22,
		0x20, 0x44,
	}))

	assert.DeepEqual(t, picotest.UnknownMessage{
		Second:           0x22,
		Fourth:           0x44,
		XXX_unrecognized: []byte{0x8, 0x13, 0x18, 0x25},
	}, decode([]byte{
		0x8, 0x13,
		0x10, 0x22,
		0x18, 0x25,
		0x20, 0x44,
	}))
}

func TestDecoder_UnrecognizedFields_Truncated(t *testing.T) {
	// field 5, wire type Bytes, with a length that runs past the buffer
	var decoded picotest.UnknownMessage
	err := picobuf.Unmarshal([]byte{0x2a, 0x7f}, &decoded)
	assert.Error(t, err)
}

// appendField appends a varint field to data.
func appendField(data []byte, field picobuf.FieldNumber, x uint64) []byte {
	data = protowire.AppendTag(data, protowire.Number(field), protowire.VarintType)
	return protowire.AppendVarint(data, x)
}

func TestDecoder_UnrecognizedFields_HighFieldNumbers(t *testing.T) {
	// Field 64 arrives after field 100000. Decode offers its fields in
	// ascending order, so by the time UnrecognizedFields runs, field 64 is
	// pending and cannot be picked up until the next Loop pass. It must not be
	// mistaken for an unrecognized field.
	var decoded picotest.HighUnknown
	data := appendField(nil, 100000, 3)
	data = appendField(data, 64, 2)
	assert.NoError(t, picobuf.Unmarshal(data, &decoded))
	assert.DeepEqual(t, decoded, picotest.HighUnknown{Low: 0, High: 2, Higher: 3})

	// Unrecognized fields at or above 64 are still captured, and survive a
	// roundtrip.
	var captured picotest.HighUnknown
	data = appendField(nil, 1, 1)
	data = appendField(data, 64, 2)
	unknown := appendField(nil, 65, 7)
	unknown = appendField(unknown, 200000, 8)
	assert.NoError(t, picobuf.Unmarshal(append(data, unknown...), &captured))
	assert.DeepEqual(t, captured, picotest.HighUnknown{
		Low:              1,
		High:             2,
		XXX_unrecognized: unknown,
	})

	reencoded, err := picobuf.Marshal(&captured)
	assert.NoError(t, err)
	var roundtripped picotest.HighUnknown
	assert.NoError(t, picobuf.Unmarshal(reencoded, &roundtripped))
	assert.DeepEqual(t, roundtripped, captured)
}

func TestCodec_UnrecognizedFields(t *testing.T) {
	initialKnown := picotest.KnownMessage{
		First:  0x11,
		Second: 0x22,
		Third:  0x33,
		Fourth: 0x44,
	}

	knownData, err := picobuf.Marshal(&initialKnown)
	assert.NoError(t, err)
	assert.Equal(t, knownData, []byte{0x8, 0x11, 0x10, 0x22, 0x18, 0x33, 0x20, 0x44})

	unknown := picotest.UnknownMessage{}
	err = picobuf.Unmarshal(knownData, &unknown)
	assert.NoError(t, err)

	assert.DeepEqual(t, unknown, picotest.UnknownMessage{
		Second:           0x22,
		Fourth:           0x44,
		XXX_unrecognized: []byte{0x8, 0x11, 0x18, 0x33},
	})

	unknownData, err := picobuf.Marshal(&unknown)
	assert.NoError(t, err)
	assert.Equal(t, unknownData, []byte{0x10, 0x22, 0x20, 0x44, 0x8, 0x11, 0x18, 0x33})

	parsedKnown := picotest.KnownMessage{}
	err = picobuf.Unmarshal(unknownData, &parsedKnown)
	assert.NoError(t, err)
	assert.DeepEqual(t, parsedKnown, initialKnown)
}

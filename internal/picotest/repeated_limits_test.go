// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"bytes"
	"slices"
	"strings"
	"testing"
	"time"
	"unsafe"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/protowire"
)

func TestCustomRepeatedLimits(t *testing.T) {
	for _, field := range []byte{6, 7, 8, 9} {
		data := bytes.Repeat([]byte{field<<3 | 2, 0}, 3)
		var m CustomMessageTypes
		err := (picobuf.UnmarshalOptions{MaxRepeatedElements: 1}).Unmarshal(data, &m)
		if err == nil {
			t.Fatalf("field %d ignored repeated limit", field)
		}
		n := len(m.RepeatedCustomType) + len(m.RepeatedPresentCustomType) + len(m.RepeatedCustomTypeCast) + len(m.RepeatedPresentCustomTypeCast)
		if n != 1 {
			t.Fatalf("field %d allocated %d values", field, n)
		}
	}
	for _, field := range []byte{3, 6} {
		item := append([]byte{field<<3 | 2, 12}, make([]byte, 12)...)
		var m CustomBytes
		if err := (picobuf.UnmarshalOptions{MaxRepeatedElements: 1}).Unmarshal(bytes.Repeat(item, 3), &m); err == nil {
			t.Fatalf("bytes field %d ignored limit", field)
		}
	}
}

func TestCustomDurationValues(t *testing.T) {
	m := CustomDurations{Values: []time.Duration{0, time.Second, -time.Nanosecond}}
	data, err := picobuf.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}
	var got CustomDurations
	if err := picobuf.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(m.Values, got.Values) {
		t.Fatalf("got %v, want %v", got.Values, m.Values)
	}
}

func TestCustomRepeatedWrongWire(t *testing.T) {
	var m CustomMessageTypes
	// A valid empty timestamp, a varint with the same number, and another timestamp.
	data := []byte{0x42, 0, 0x40, 7, 0x42, 0}
	if err := (picobuf.UnmarshalOptions{MaxRepeatedElements: 2}).Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if len(m.RepeatedCustomTypeCast) != 2 || !bytes.Equal(m.XXX_unrecognized, []byte{0x40, 7}) {
		t.Fatalf("unexpected decoding: %+v", m)
	}
}

func TestRepeatedBytesLimit(t *testing.T) {
	const limit = 1 << 10
	opts := picobuf.UnmarshalOptions{MaxRepeatedBytes: limit}
	// Each case repeats a tiny element that takes far more memory decoded.
	type test struct {
		name  string
		data  []byte
		msg   picobuf.Message
		size  uintptr
		count func() int
	}
	var tests []test
	add := func(name string, data []byte, msg picobuf.Message, size uintptr, count func() int) {
		tests = append(tests, test{name, data, msg, size, count})
	}
	var pointers, values VariationsMessage
	add("message pointers", bytes.Repeat([]byte{3<<3 | 2, 0}, 1000), &pointers,
		unsafe.Sizeof(Message{})+unsafe.Sizeof(&Message{}), func() int { return len(pointers.Rep) })
	add("message values", bytes.Repeat([]byte{6<<3 | 2, 0}, 1000), &values,
		unsafe.Sizeof(Message{}), func() int { return len(values.PresentRep) })
	tag := func(num protowire.Number, typ protowire.Type) []byte { return protowire.AppendTag(nil, num, typ) }
	repeat := func(item []byte) []byte { return bytes.Repeat(item, 1000) }
	var packed AllTypes
	add("packed int64", protowire.AppendBytes(tag(18, protowire.BytesType), make([]byte, 1000)), &packed,
		8, func() int { return len(packed.Int64S) })
	var unpacked AllTypes
	add("unpacked int64", repeat(append(tag(18, protowire.VarintType), 0)), &unpacked,
		8, func() int { return len(unpacked.Int64S) })
	var strs AllTypes
	add("strings", repeat(append(tag(30, protowire.BytesType), 0)), &strs,
		unsafe.Sizeof(""), func() int { return len(strs.Strings) })
	var m Map
	add("map entries", repeat(append(tag(1, protowire.BytesType), 0)), &m,
		8, func() int { return len(m.Values) })

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opts.Unmarshal(tt.data, tt.msg)
			if err == nil || !strings.Contains(err.Error(), "maximum repeated bytes") {
				t.Fatalf("got %v", err)
			}
			if got, most := tt.count(), int(limit/tt.size); got > most {
				t.Fatalf("decoded %d elements, the limit allows %d", got, most)
			}
		})
	}
}

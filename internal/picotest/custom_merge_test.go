// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"testing"
	"time"

	"storj.io/picobuf"
)

type customMergeParent struct{ child CustomMessageTypes }

func (m *customMergeParent) Encode(c *picobuf.Encoder) bool {
	c.Message(1, m.child.Encode)
	return true
}
func (m *customMergeParent) Decode(c *picobuf.Decoder) { c.Message(1, m.child.Decode) }

func TestConvertedMessageMergeAcrossParentOccurrences(t *testing.T) {
	// Each parent occurrence contributes a different timestamp component.
	data := []byte{10, 4, 0x22, 2, 8, 123, 10, 5, 0x22, 3, 16, 0xc8, 3}
	var m customMergeParent
	if err := picobuf.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m.child.CustomTypeCast == nil || !m.child.CustomTypeCast.Equal(time.Unix(123, 456)) {
		t.Fatalf("got %v", m.child.CustomTypeCast)
	}
}

func TestConvertedRepeatedValuesStayDistinct(t *testing.T) {
	data := []byte{0x42, 2, 8, 123, 0x42, 3, 16, 0xc8, 3}
	var m CustomMessageTypes
	if err := picobuf.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	values := m.RepeatedCustomTypeCast
	if len(values) != 2 || !values[0].Equal(time.Unix(123, 0)) || !values[1].Equal(time.Unix(0, 456)) {
		t.Fatalf("got %v", values)
	}
}

// TestConvertedMessageDecodeDoesNotAllocate guards against converted
// timestamps allocating per decoded value, which once tripled the allocations
// of a typical storj order limit.
func TestConvertedMessageDecodeDoesNotAllocate(t *testing.T) {
	timestamp := []byte{8, 123, 16, 0xc8, 3} // seconds=123, nanos=456
	var data []byte
	for _, field := range []byte{5, 9, 9} { // present_custom_type_cast and two repeated values
		data = append(data, field<<3|2, byte(len(timestamp)))
		data = append(data, timestamp...)
	}

	var dec picobuf.Decoder
	m := CustomMessageTypes{RepeatedPresentCustomTypeCast: make([]time.Time, 0, 2)}
	allocs := testing.AllocsPerRun(100, func() {
		m.PresentCustomTypeCast = time.Time{}
		m.RepeatedPresentCustomTypeCast = m.RepeatedPresentCustomTypeCast[:0]
		if err := dec.Unmarshal(data, &m, picobuf.UnmarshalOptions{}); err != nil {
			t.Fatal(err)
		}
	})
	if allocs != 0 {
		t.Fatalf("got %v allocations, want 0", allocs)
	}
	want := time.Unix(123, 456)
	if !m.PresentCustomTypeCast.Equal(want) || len(m.RepeatedPresentCustomTypeCast) != 2 || !m.RepeatedPresentCustomTypeCast[1].Equal(want) {
		t.Fatalf("got %v %v", m.PresentCustomTypeCast, m.RepeatedPresentCustomTypeCast)
	}
}

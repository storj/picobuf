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

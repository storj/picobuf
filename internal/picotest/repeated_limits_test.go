// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"bytes"
	"slices"
	"testing"
	"time"

	"storj.io/picobuf"
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

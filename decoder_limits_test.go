// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"math"
	"testing"
)

func TestRepeatedLimitDoesNotOverflow(t *testing.T) {
	// Exercise the boundary without constructing a multi-gigabyte payload.
	dec := &Decoder{
		maxRepeatedElements: math.MaxInt32,
		repeatedElements:    math.MaxInt32 - 1,
	}
	if !dec.takeRepeated(1) {
		t.Fatal("rejected the last permitted element")
	}
	if dec.takeRepeated(1) || dec.Err() == nil {
		t.Fatal("counter overflow allowed an element beyond the limit")
	}
	if dec.repeatedReservation(10) != 0 {
		t.Fatal("exhausted limit permitted a reservation")
	}
}

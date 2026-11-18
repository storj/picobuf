// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"math"
	"strings"
	"testing"
)

func TestRepeatedLimitDoesNotOverflow(t *testing.T) {
	// Exercise the boundary without constructing a multi-gigabyte payload.
	dec := &Decoder{repeatedLeft: 1, repeatedBytesLeft: math.MaxInt64}
	if !dec.takeRepeated(1, 8) {
		t.Fatal("rejected the last permitted element")
	}
	if dec.takeRepeated(1, 0) || dec.Err() == nil {
		t.Fatal("allowed an element beyond the limit")
	}
	if dec.repeatedReservation(10, 1) != 0 {
		t.Fatal("exhausted limit permitted a reservation")
	}
}

func TestRepeatedReservationWithoutLimit(t *testing.T) {
	// Without a limit the byte budget is math.MaxInt64, which must not
	// truncate where int has 32 bits.
	dec := NewDecoder(nil)
	if got := dec.repeatedReservation(10, 8); got != 10 {
		t.Fatalf("got %d, want 10", got)
	}
}

func TestRepeatedBytesLimit(t *testing.T) {
	dec := &Decoder{repeatedLeft: math.MaxInt32, repeatedBytesLeft: 20}
	if got := dec.repeatedReservation(10, 8); got != 2 {
		t.Fatalf("reservation: got %d, want 2", got)
	}
	if !dec.takeRepeated(1, 16) {
		t.Fatal("rejected an element within the limit")
	}
	if dec.takeRepeated(1, 8) {
		t.Fatal("allowed an element beyond the limit")
	}
	if err := dec.Err(); err == nil || !strings.Contains(err.Error(), "maximum repeated bytes") {
		t.Fatalf("got %v", err)
	}
}

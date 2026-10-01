// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"testing"

	"storj.io/picobuf"
)

func TestUnknownWirePolicies(t *testing.T) {
	// Field 2 is known, but its bytes occurrence must follow the unknown policy.
	data := []byte{0x12, 0, 0x10, 7}
	var ignored IgnoreUnknownMessage
	if err := picobuf.Unmarshal(data, &ignored); err != nil || ignored.Second != 7 {
		t.Fatalf("ignore: %v %+v", err, ignored)
	}
	var rejected RejectUnknownMessage
	if err := picobuf.Unmarshal(data, &rejected); err == nil {
		t.Fatal("reject policy accepted mismatched wire")
	}
}

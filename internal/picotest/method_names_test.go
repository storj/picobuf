// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"testing"

	"storj.io/picobuf"
)

func TestMethodNameGetters(t *testing.T) {
	var m MethodNames
	if err := picobuf.Unmarshal([]byte{10, 1, 'a', 18, 1, 'b', 26, 1, 'c'}, &m); err != nil {
		t.Fatal(err)
	}
	if m.GetEncode__() != "a" || m.GetDecode_() != "b" || m.GetGetEncode_() != "c" {
		t.Fatalf("getters lost fields: %+v", m)
	}
}

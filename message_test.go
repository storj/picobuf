// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"fmt"
	"math"
	"testing"
)

func FuzzFieldNumber_String(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(-1))
	f.Add(int32(1))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		got := FieldNumber(v).String()
		exp := fmt.Sprintf("%v", v)
		if exp != got {
			t.Errorf("invalid result %v != %q", v, got)
		}
	})
}

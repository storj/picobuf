// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package jsoncompat

import (
	"math"
	"strings"
	"testing"

	"storj.io/picobuf/internal/jsoncompat/pico"
	"storj.io/picobuf/picojson"
)

func TestUnmarshalJSONOptions(t *testing.T) {
	data := []byte(`{"int32":[1,2,3]}`)
	var m pico.Lists
	if err := m.UnmarshalJSONOptions(data, picojson.UnmarshalOptions{MaxRepeatedElements: 3}); err != nil {
		t.Fatal(err)
	}
	if len(m.Int32) != 3 {
		t.Fatalf("got %v", m.Int32)
	}
	err := m.UnmarshalJSONOptions(data, picojson.UnmarshalOptions{MaxRepeatedElements: 2})
	if err == nil || !strings.Contains(err.Error(), "maximum repeated elements") {
		t.Fatalf("got %v", err)
	}
}

func TestUnmarshalJSONOptionsLargeDepth(t *testing.T) {
	var m pico.Lists
	opts := picojson.UnmarshalOptions{MaxRecursionDepth: math.MaxInt}
	if err := m.UnmarshalJSONOptions([]byte(`{"int32":[1]}`), opts); err != nil {
		t.Fatal(err)
	}
}

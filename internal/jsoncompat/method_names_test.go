// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package jsoncompat

import (
	"testing"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/jsoncompat/pico"
)

func TestMethodNameFields(t *testing.T) {
	input := []byte(`{"encode":"a","decode":"b","marshalJSON":"c","unmarshalJSON":"d","getEncode":"e"}`)
	var m pico.MethodNames
	if err := m.UnmarshalJSON(input); err != nil {
		t.Fatal(err)
	}
	wire, err := picobuf.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}
	var got pico.MethodNames
	if err := picobuf.Unmarshal(wire, &got); err != nil {
		t.Fatal(err)
	}
	out, err := got.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if !sameJSON(t, input, out) {
		t.Fatalf("got %s, want %s", out, input)
	}
	var oneof pico.MethodOneof
	if err := oneof.UnmarshalJSON([]byte(`{"decode":"selected"}`)); err != nil {
		t.Fatal(err)
	}
	out, err = oneof.MarshalJSON()
	if err != nil || string(out) != `{"decode":"selected"}` {
		t.Fatalf("oneof: %s, %v", out, err)
	}
}

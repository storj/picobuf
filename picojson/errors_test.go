// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picojson

import (
	"errors"
	"strings"
	"testing"
)

type countedError struct{ calls int }

func (e *countedError) Error() string { e.calls++; return "invalid value" }

func TestFieldErrorFormatsOnce(t *testing.T) {
	schema, _, _ := wideMessage(1, false, false)
	msg := schema.message(0)
	f := msg.first()
	cause := new(countedError)
	var err error = cause
	for range 1000 {
		err = fieldError(&msg, &f, err)
	}
	if cause.calls != 0 {
		t.Fatalf("wrapping eagerly formatted the cause %d times", cause.calls)
	}
	if got := err.Error(); got != strings.Repeat("Wide.field_1: ", 1000)+"invalid value" {
		t.Fatal("incorrect field path")
	}
	if cause.calls != 1 || !errors.Is(err, cause) {
		t.Fatal("lost error cause or formatted it more than once")
	}
}

func BenchmarkNestedJSONError(b *testing.B) {
	// A message with one field referring to itself at schema offset zero.
	schema := Schema(string([]byte{8, 'N', 'o', 'd', 'e', 1, 1, byte(messageKind), 0, 0, 10, 'c', 'h', 'i', 'l', 'd'}))
	data := []byte(strings.Repeat(`{"child":`, 4000) + `{"bad":0}` + strings.Repeat("}", 4000))
	b.ReportAllocs()
	for b.Loop() {
		if _, err := unmarshalWire(data, schema, 0); err == nil {
			b.Fatal("accepted unknown field")
		}
	}
}

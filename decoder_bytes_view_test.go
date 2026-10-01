// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"bytes"
	"testing"

	"storj.io/picobuf"
)

func TestDecoderBytesView(t *testing.T) {
	input := []byte{0x0a, 2, 1, 2, 0x12, 2, 3, 4}
	var view, owned []byte
	dec := picobuf.NewDecoder(input)
	dec.Loop(func(dec *picobuf.Decoder) {
		dec.BytesView(1, &view)
		dec.Bytes(2, &owned)
	})
	if err := dec.Err(); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(view, []byte{1, 2}) || &view[0] != &input[2] {
		t.Fatal("BytesView must borrow the field payload")
	}
	if cap(view) != len(view) {
		t.Fatalf("BytesView capacity %d exceeds the payload, append would overwrite the input", cap(view))
	}
	clear(input)
	if !bytes.Equal(owned, []byte{3, 4}) {
		t.Fatal("BytesView changed ownership of the following Bytes field")
	}
}

func TestDecoderBytesViewAbsentAndInvalid(t *testing.T) {
	for _, tc := range []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{"absent", []byte{0x12, 0}, false},
		{"wrong wire type", []byte{0x08, 1}, false},
		{"truncated length", []byte{0x0a, 0x80}, true},
		{"truncated payload", []byte{0x0a, 3, 1}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			view := []byte{9}
			dec := picobuf.NewDecoder(tc.input)
			dec.Loop(func(dec *picobuf.Decoder) {
				dec.BytesView(1, &view)
				var other []byte
				dec.Bytes(2, &other)
			})
			if (dec.Err() != nil) != tc.wantErr {
				t.Fatalf("Err = %v; want error: %v", dec.Err(), tc.wantErr)
			}
			if !bytes.Equal(view, []byte{9}) {
				t.Fatal("absent or invalid field changed the destination")
			}
		})
	}
}

func TestDecoderBytesViewEmpty(t *testing.T) {
	view := []byte{9}
	var next int32
	dec := picobuf.NewDecoder([]byte{0x0a, 0, 0x10, 7})
	dec.Loop(func(dec *picobuf.Decoder) {
		dec.BytesView(1, &view)
		dec.Int32(2, &next)
	})
	if dec.Err() != nil || len(view) != 0 || next != 7 {
		t.Fatalf("empty field: view=%x next=%d err=%v", view, next, dec.Err())
	}
}

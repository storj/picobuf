// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"encoding/hex"
	"testing"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest/pic"
)

func TestGeneratedCustomBytes(t *testing.T) {
	id := pic.ID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	idHex := hex.EncodeToString(id[:])
	zeroHex := hex.EncodeToString(make([]byte, len(pic.ID{})))
	// PresentBasic and PresentOpt store values, so they are always written.
	valuesHex := "220c" + zeroHex + "2a0c" + zeroHex

	for _, test := range []struct {
		name string
		in   CustomBytes
		wire string
	}{
		{name: "absent", in: CustomBytes{}, wire: valuesHex},
		{name: "present zero", in: CustomBytes{Value: new(pic.ID)}, wire: "0a0c" + zeroHex + valuesHex},
		{name: "present", in: CustomBytes{Value: &id}, wire: "0a0c" + idHex + valuesHex},
		{name: "optional present zero", in: CustomBytes{Opt: new(pic.ID)}, wire: "120c" + zeroHex + valuesHex},
		{name: "optional present", in: CustomBytes{Opt: &id}, wire: "120c" + idHex + valuesHex},
	} {
		t.Run(test.name, func(t *testing.T) {
			data, err := picobuf.Marshal(&test.in)
			if err != nil {
				t.Fatal(err)
			}
			if got := hex.EncodeToString(data); got != test.wire {
				t.Fatalf("encoded as %s, want %s", got, test.wire)
			}

			var out CustomBytes
			if err := picobuf.Unmarshal(data, &out); err != nil {
				t.Fatal(err)
			}
			if !equalID(out.Value, test.in.Value) || !equalID(out.Opt, test.in.Opt) {
				t.Fatalf("decoded as {Value: %v, Opt: %v}, want {Value: %v, Opt: %v}",
					out.Value, out.Opt, test.in.Value, test.in.Opt)
			}
		})
	}
}

func TestGeneratedCustomSerializePresentZero(t *testing.T) {
	empty := ""
	data, err := picobuf.Marshal(&Piece{Id: new(pic.ID), Alt: &empty})
	if err != nil {
		t.Fatal(err)
	}
	want := "0a0c" + hex.EncodeToString(make([]byte, len(pic.ID{}))) + "1200"
	if got := hex.EncodeToString(data); got != want {
		t.Fatalf("encoded as %s, want %s", got, want)
	}

	var piece Piece
	if err := picobuf.Unmarshal(data, &piece); err != nil {
		t.Fatal(err)
	}
	if piece.Id == nil || !piece.Id.IsZero() || piece.Alt == nil || *piece.Alt != "" {
		t.Fatalf("present zero piece decoded as %+v", piece)
	}
}

func TestGeneratedCustomBytesInvalidLength(t *testing.T) {
	for _, wire := range []string{"0a00", "0a050102030405"} {
		data, _ := hex.DecodeString(wire)
		var out CustomBytes
		if err := picobuf.Unmarshal(data, &out); err == nil {
			t.Fatalf("decoded %s as %v", wire, out.Value)
		}
	}
}

func equalID(a, b *pic.ID) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

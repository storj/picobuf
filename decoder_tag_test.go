// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"bytes"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
)

// Compare the decoder's tag handling with the reference wire parser, including
// every one-byte tag, nonminimal varints, and invalid or oversized field numbers.
func FuzzDecoderTag(f *testing.F) {
	f.Add([]byte{})
	for tag := 0; tag < 256; tag++ {
		f.Add([]byte{byte(tag)})
	}
	for _, data := range [][]byte{{0x88, 0, 1}, {0x80, 1, 0}, {0x80, 0x80, 0x80, 0x80, 0x10}, {255, 255, 255, 255, 15}, {255, 255, 255, 255, 255, 255, 255, 255, 255, 2}} {
		f.Add(data)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		dec := NewDecoder(data)
		dec.nextField(0)
		if len(data) == 0 {
			if dec.Err() != nil || dec.pendingField != fieldDecodingDone {
				t.Fatal("empty input not complete")
			}
			return
		}
		field, wire, n := protowire.ConsumeTag(data)
		if n < 0 || field < protowire.MinValidNumber || field > protowire.MaxValidNumber || wire == protowire.EndGroupType {
			if dec.Err() == nil {
				t.Fatalf("accepted invalid tag: %x", data)
			}
			return
		}
		if dec.Err() != nil || dec.pendingField != FieldNumber(field) || int(dec.pendingWire) != int(wire) || !bytes.Equal(dec.buffer, data[n:]) {
			t.Fatalf("tag %x: field=%d wire=%d err=%v, want field=%d wire=%d", data, dec.pendingField, dec.pendingWire, dec.Err(), field, wire)
		}
	})
}

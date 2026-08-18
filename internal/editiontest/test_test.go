// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package editiontest

import (
	"testing"

	"storj.io/picobuf"
)

func TestUTF8ValidationFeatures(t *testing.T) {
	var message Message
	if err := picobuf.Unmarshal([]byte{0x0a, 0x01, 0xff}, &message); err != nil {
		t.Fatalf("file-level NONE rejected invalid UTF-8: %v", err)
	}
	if message.Unvalidated != string([]byte{0xff}) {
		t.Fatalf("decoded %q", message.Unvalidated)
	}

	if err := picobuf.Unmarshal([]byte{0x12, 0x01, 0xff}, &message); err == nil {
		t.Fatal("field-level VERIFY accepted invalid UTF-8")
	}
	if err := (picobuf.UnmarshalOptions{AllowInvalidUTF8: true}).Unmarshal([]byte{0x12, 0x01, 0xff}, &message); err != nil {
		t.Fatalf("AllowInvalidUTF8 did not override field-level VERIFY: %v", err)
	}

	if err := picobuf.Unmarshal([]byte{0x1a, 0x03, 0x0a, 0x01, 0xff}, &message); err != nil {
		t.Fatalf("file-level NONE rejected invalid UTF-8 map key: %v", err)
	}
	if err := picobuf.Unmarshal([]byte{0x22, 0x03, 0x0a, 0x01, 0xff}, &message); err == nil {
		t.Fatal("field-level VERIFY accepted invalid UTF-8 map key")
	}

	invalid := string([]byte{0xff})
	if _, err := picobuf.Marshal(&Message{Unvalidated: invalid}); err != nil {
		t.Fatalf("file-level NONE rejected encoding invalid UTF-8: %v", err)
	}
	if _, err := picobuf.Marshal(&Message{Validated: invalid}); err == nil {
		t.Fatal("field-level VERIFY encoded invalid UTF-8")
	}
	if _, err := picobuf.Marshal(&Message{UnvalidatedMap: map[string]string{invalid: ""}}); err != nil {
		t.Fatalf("file-level NONE rejected encoding invalid UTF-8 map key: %v", err)
	}
	if _, err := picobuf.Marshal(&Message{ValidatedMap: map[string]string{invalid: ""}}); err == nil {
		t.Fatal("field-level VERIFY encoded invalid UTF-8 map key")
	}
}

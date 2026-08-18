// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"bytes"
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/protocompat/pico"
	"storj.io/picobuf/internal/protocompat/prot"
)

// FuzzCompat checks picobuf against protobuf-go on the same bytes. Unlike
// FuzzDecode, which can only see whether picobuf contradicts itself, this
// notices picobuf quietly disagreeing with the reference implementation.
func FuzzCompat(f *testing.F) {
	seed, err := picobuf.Marshal(&pico.Types{
		Int32: -2, Int64: -2, Uint32: 1, Uint64: 1,
		Sint32: -2, Sint64: -2, Fixed32: 1, Fixed64: 1,
		Sfixed32: -2, Sfixed64: -2, Float: 1, Double: 1,
		Bool: true, String_: "hello", Bytes: []byte{1, 2},
		Message:         pico.Message{Int32: 1},
		OptionalMessage: &pico.OptionalMessage{Int32: 1},
		Language:        pico.Language_ENGLISH,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(seed)
	f.Add([]byte(nil))

	f.Fuzz(func(t *testing.T, data []byte) {
		// DiscardUnknown, because pico.Types does not capture unknown fields
		// and so cannot echo them back the way protobuf-go does.
		var reference prot.Types
		referenceErr := proto.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, &reference)

		var decoded pico.Types
		err := picobuf.Unmarshal(data, &decoded)

		if referenceErr != nil {
			if err == nil {
				t.Fatalf("accepted input that protobuf-go rejected (%v)\ninput: %x", referenceErr, data)
			}
			return
		}
		// picobuf is allowed to be stricter: it does not implement groups.
		if err != nil {
			return
		}
		// picobuf has no way to represent a present submessage that is empty,
		// so its encoding legitimately differs from the reference there.
		if reference.Message != nil && proto.Size(reference.Message) == 0 {
			return
		}

		referenceData, err := proto.MarshalOptions{Deterministic: true}.Marshal(&reference)
		if err != nil {
			t.Fatal(err)
		}
		decodedData, err := picobuf.Marshal(&decoded)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(referenceData, decodedData) {
			t.Fatalf("decoded differently than protobuf-go\ninput:  %x\ngoogle: %x\npico:   %x",
				data, referenceData, decodedData)
		}
	})
}

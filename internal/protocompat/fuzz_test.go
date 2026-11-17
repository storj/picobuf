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
		// Compare known fields; capture of unknown fields is tested separately.
		// Unknown tags can have several equivalent varint encodings.
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
		if err != nil {
			t.Fatalf("rejected input accepted by protobuf-go: %v\ninput: %x", err, data)
		}
		decoded.XXX_unrecognized = nil
		decoded.Message.XXX_unrecognized = nil
		if decoded.OptionalMessage != nil {
			decoded.OptionalMessage.XXX_unrecognized = nil
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

// FuzzCompatRepeated is FuzzCompat for repeated fields, which decode both
// packed and unpacked.
func FuzzCompatRepeated(f *testing.F) {
	seed, err := proto.Marshal(&prot.RepeatedTypes{
		Int32: []int32{-2, 1, 300}, Int64: []int64{-2, 1 << 40}, Uint32: []uint32{1, 1 << 31},
		Uint64: []uint64{1, 1 << 63}, Sint32: []int32{-2, 2}, Sint64: []int64{-2, 2},
		Fixed32: []uint32{1, 2}, Fixed64: []uint64{1, 2}, Sfixed32: []int32{-2}, Sfixed64: []int64{-2},
		Float: []float32{1, 2}, Double: []float64{1, 2}, Bool: []bool{true, false},
		String_: []string{"a", ""}, Bytes: [][]byte{{1}, nil},
		Message:  []*prot.Message{{Int32: 1}, {}},
		Language: []prot.Language{1, 0},
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(seed)
	// unpacked varints and fixed values
	f.Add([]byte{0x08, 0x01, 0x08, 0x96, 0x02, 0x3d, 1, 0, 0, 0, 0x41, 1, 0, 0, 0, 0, 0, 0, 0})
	f.Add([]byte(nil))

	f.Fuzz(func(t *testing.T, data []byte) {
		var reference prot.RepeatedTypes
		referenceErr := proto.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, &reference)

		var decoded pico.RepeatedTypes
		err := picobuf.Unmarshal(data, &decoded)

		if referenceErr != nil {
			if err == nil {
				t.Fatalf("accepted input that protobuf-go rejected (%v)\ninput: %x", referenceErr, data)
			}
			return
		}
		if err != nil {
			t.Fatalf("rejected input accepted by protobuf-go: %v\ninput: %x", err, data)
		}
		decoded.XXX_unrecognized = nil
		for i := range decoded.Message {
			decoded.Message[i].XXX_unrecognized = nil
		}

		referenceData, err := proto.Marshal(&reference)
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

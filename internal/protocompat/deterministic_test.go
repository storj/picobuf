// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/protocompat/pico"
	"storj.io/picobuf/internal/protocompat/prot"
)

func TestMarshalDeterministic(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	msg := pico.Map{
		StringString:  map[string]string{},
		Int64String:   map[int64]string{},
		StringMessage: map[string]*pico.OptionalMessage{},
	}
	for i := range 200 {
		// entries around 128 bytes long, so that the length prefixes of
		// the entries take one or two bytes, and the hints that the sizing
		// pass recorded in a different order do not match.
		msg.StringString["key-"+strconv.Itoa(i)] = randString(r, 100+r.Intn(60))
		msg.Int64String[r.Int63()] = randString(r, 100+r.Intn(60))
		msg.StringMessage["key-"+strconv.Itoa(i)] = &pico.OptionalMessage{Int32: int32(r.Intn(1 << 20))}
	}

	deterministic := picobuf.MarshalOptions{Deterministic: true}
	sorted, err := deterministic.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}
	var want prot.Map
	if err := proto.Unmarshal(sorted, &want); err != nil {
		t.Fatal(err)
	}
	// google.golang.org/protobuf sorts map entries the same way.
	reference, err := proto.MarshalOptions{Deterministic: true}.Marshal(&want)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(sorted, reference) {
		t.Fatal("deterministic encoding differs from google.golang.org/protobuf")
	}

	differs := false
	for range 10 {
		again, err := deterministic.MarshalAppend([]byte{1, 2, 3}, &msg)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(again[3:], sorted) {
			t.Fatal("deterministic encoding changed")
		}

		data, err := picobuf.Marshal(&msg)
		if err != nil {
			t.Fatal(err)
		}
		if len(data) != len(sorted) {
			t.Fatalf("encoded %d bytes, want %d", len(data), len(sorted))
		}
		differs = differs || !bytes.Equal(data, sorted)
		var got prot.Map
		if err := proto.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(&got, &want) {
			t.Fatal("encoding without Deterministic decodes to a different message")
		}
	}
	if !differs {
		t.Fatal("map entries were sorted without Deterministic")
	}
}

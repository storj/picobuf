// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package jsoncompat

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/jsoncompat/pico"
	"storj.io/picobuf/internal/jsoncompat/prot"
)

func TestMismatchedWireTypes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		pico picobuf.Message
		prot proto.Message
	}{
		{"scalar", []byte{0x0a, 0, 8, 7}, new(pico.Scalars), new(prot.Scalars)},
		{"optional", []byte{0x0a, 0}, new(pico.Optionals), new(prot.Optionals)},
		{"oneof", []byte{0x18, 7, 0x0d, 0, 0, 0, 0}, new(pico.Oneofs), new(prot.Oneofs)},
		{"repeated", []byte{8, 7, 0x0d, 0, 0, 0, 0, 8, 9}, new(pico.Lists), new(prot.Lists)},
		{"message", []byte{8, 7}, new(pico.All), new(prot.All)},
		{"map", []byte{8, 7}, new(pico.Maps), new(prot.Maps)},
		{"custom", []byte{8, 7}, new(pico.Custom), new(prot.Custom)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := proto.Unmarshal(tt.data, tt.prot); err != nil {
				t.Fatal(err)
			}
			if err := picobuf.Unmarshal(tt.data, tt.pico); err != nil {
				t.Fatal(err)
			}
			wire, err := picobuf.Marshal(tt.pico)
			if err != nil {
				t.Fatal(err)
			}
			got := tt.prot.ProtoReflect().New().Interface()
			if err := proto.Unmarshal(wire, got); err != nil {
				t.Fatal(err)
			}
			if !proto.Equal(got, tt.prot) {
				t.Fatalf("got %v, want %v", got, tt.prot)
			}
		})
	}
}

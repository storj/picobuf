// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picotest

import (
	"reflect"
	"testing"

	"storj.io/picobuf"
)

func TestGeneratedGenericMap(t *testing.T) {
	message := GenericMap{Values: map[int32]*GenericMapValue{
		7: {Name: "one"},
	}}

	data, err := picobuf.Marshal(&message)
	if err != nil {
		t.Fatal(err)
	}
	wantData := []byte{
		0x0a, 0x09, 0x08, 0x07, 0x12, 0x05, 0x0a, 0x03, 'o', 'n', 'e',
	}
	if !reflect.DeepEqual(data, wantData) {
		t.Fatalf("encoded %x, want %x", data, wantData)
	}

	var decoded GenericMap
	if err := picobuf.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded.Values, message.Values) {
		t.Fatalf("decoded %#v, want %#v", decoded.Values, message.Values)
	}
}

func TestGeneratedGenericMapEntries(t *testing.T) {
	message := GenericMap{Values: map[int32]*GenericMapValue{
		1: {Name: "one"},
		2: {Name: "two"},
		3: nil,
		4: nil,
	}}

	data, err := picobuf.Marshal(&message)
	if err != nil {
		t.Fatal(err)
	}

	var decoded GenericMap
	if err := picobuf.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	want := map[int32]*GenericMapValue{
		1: {Name: "one"},
		2: {Name: "two"},
		3: {},
		4: {},
	}
	if !reflect.DeepEqual(decoded.Values, want) {
		t.Fatalf("decoded %#v, want %#v", decoded.Values, want)
	}
	if decoded.Values[1] == decoded.Values[2] {
		t.Fatal("entries share a message")
	}
	if decoded.Values[3] == decoded.Values[4] {
		t.Fatal("empty entries share a message")
	}
}

func TestGeneratedEnumMap(t *testing.T) {
	message := EnumMap{Values: map[int32]MapEnum{
		1: MapEnum_MAP_ENUM_FIRST,
	}}

	data, err := picobuf.Marshal(&message)
	if err != nil {
		t.Fatal(err)
	}
	wantData := []byte{0x0a, 0x04, 0x08, 0x01, 0x10, 0x01}
	if !reflect.DeepEqual(data, wantData) {
		t.Fatalf("encoded %x, want %x", data, wantData)
	}

	// The zero enum is omitted from the entry, but the key still decodes.
	data = append(data, 0x0a, 0x02, 0x08, 0x02)

	var decoded EnumMap
	if err := picobuf.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	want := map[int32]MapEnum{
		1: MapEnum_MAP_ENUM_FIRST,
		2: MapEnum_MAP_ENUM_UNSPECIFIED,
	}
	if !reflect.DeepEqual(decoded.Values, want) {
		t.Fatalf("decoded %#v, want %#v", decoded.Values, want)
	}
}

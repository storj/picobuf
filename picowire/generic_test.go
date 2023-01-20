// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import (
	"reflect"
	"testing"

	"storj.io/picobuf"
)

type customMapMessage struct {
	Values MapOf[int32, string, Int32Codec, StringCodec]
}

func (message *customMapMessage) Encode(enc *picobuf.Encoder) bool {
	message.Values.PicoEncode(enc, 1)
	return true
}

func (message *customMapMessage) Decode(dec *picobuf.Decoder) {
	message.Values.PicoDecode(dec, 1)
}

func TestMapOfEncodesProtobufMapEntry(t *testing.T) {
	message := customMapMessage{Values: MapOf[int32, string, Int32Codec, StringCodec]{
		7: "one",
	}}

	data, err := picobuf.Marshal(&message)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0x0a, 0x07, 0x08, 0x07, 0x12, 0x03, 'o', 'n', 'e'}
	if !reflect.DeepEqual(data, want) {
		t.Fatalf("encoded %x, want %x", data, want)
	}
}

func TestMapOfDecodesEachEntryWithProtobufDefaults(t *testing.T) {
	data := []byte{
		0x0a, 0x07, 0x08, 0x07, 0x12, 0x03, 'o', 'n', 'e',
		0x0a, 0x05, 0x12, 0x03, 't', 'w', 'o',
		0x0a, 0x02, 0x08, 0x08,
	}

	var message customMapMessage
	if err := picobuf.Unmarshal(data, &message); err != nil {
		t.Fatal(err)
	}
	want := MapOf[int32, string, Int32Codec, StringCodec]{
		7: "one",
		0: "two",
		8: "",
	}
	if !reflect.DeepEqual(message.Values, want) {
		t.Fatalf("decoded %#v, want %#v", message.Values, want)
	}
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picowire

import (
	"bytes"
	"fmt"
	"testing"

	"storj.io/picobuf"
)

type stringMapMessage struct {
	Values MapStringString
}

func (message *stringMapMessage) Encode(enc *picobuf.Encoder) bool {
	message.Values.PicoEncode(enc, 1)
	return true
}

func (message *stringMapMessage) Decode(dec *picobuf.Decoder) {
	message.Values.PicoDecode(dec, 1)
}

func TestMapEncodingIsSortedByKey(t *testing.T) {
	message := stringMapMessage{Values: MapStringString{
		"b": "2", "a": "1", "c": "3",
	}}

	data, err := picobuf.Marshal(&message)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{
		0x0a, 0x06, 0x0a, 0x01, 'a', 0x12, 0x01, '1',
		0x0a, 0x06, 0x0a, 0x01, 'b', 0x12, 0x01, '2',
		0x0a, 0x06, 0x0a, 0x01, 'c', 0x12, 0x01, '3',
	}
	if !bytes.Equal(data, want) {
		t.Fatalf("encoded %x, want %x", data, want)
	}
}

func BenchmarkMapEncode(b *testing.B) {
	for _, size := range []int{1, 8, 64} {
		message := stringMapMessage{Values: MapStringString{}}
		for i := range size {
			message.Values[string(rune('a'+i%26))+string(rune('0'+i/26))] = "value"
		}
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				if _, err := picobuf.Marshal(&message); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func TestMapOfEncodingIsSortedByKey(t *testing.T) {
	message := customMapMessage{Values: MapOf[int32, string, Int32Codec, StringCodec]{
		9: "nine", 3: "three", 7: "seven",
	}}

	data, err := picobuf.Marshal(&message)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{
		0x0a, 0x09, 0x08, 0x03, 0x12, 0x05, 't', 'h', 'r', 'e', 'e',
		0x0a, 0x09, 0x08, 0x07, 0x12, 0x05, 's', 'e', 'v', 'e', 'n',
		0x0a, 0x08, 0x08, 0x09, 0x12, 0x04, 'n', 'i', 'n', 'e',
	}
	if !bytes.Equal(data, want) {
		t.Fatalf("encoded %x, want %x", data, want)
	}
}

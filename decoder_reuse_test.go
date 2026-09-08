// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"bytes"
	"reflect"
	"testing"
)

type reuseMessage struct {
	value  int32
	data   []byte
	text   string
	values []int32
}

func (*reuseMessage) Encode(*Encoder) bool { return true }
func (m *reuseMessage) Decode(dec *Decoder) {
	dec.Int32(1, &m.value)
	dec.Bytes(2, &m.data)
	dec.String(3, &m.text)
	dec.RepeatedInt32(4, &m.values)
}

func TestDecoderUnmarshalReuse(t *testing.T) {
	var dec Decoder
	for _, tc := range []struct {
		name          string
		data          []byte
		opts          UnmarshalOptions
		payloadOffset int
	}{
		{name: "valid", data: []byte{8, 7, 18, 1, 'a', 26, 1, 'b', 32, 1}, payloadOffset: 4},
		{name: "malformed", data: []byte{18, 2, 'a'}},
		{name: "size limit", data: []byte{8, 7}, opts: UnmarshalOptions{MaxInputSize: 1}},
		{name: "repeated limit", data: []byte{32, 1, 32, 2}, opts: UnmarshalOptions{MaxRepeatedElements: 1}},
		{name: "invalid UTF8 allowed", data: []byte{26, 1, 255}, opts: UnmarshalOptions{AllowInvalidUTF8: true}},
		{name: "invalid UTF8", data: []byte{26, 1, 255}},
		{name: "alias", data: []byte{18, 1, 'a'}, opts: UnmarshalOptions{AliasInput: true}, payloadOffset: 2},
		{name: "empty"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var want, got reuseMessage
			wantErr := tc.opts.Unmarshal(tc.data, &want)
			gotErr := dec.Unmarshal(tc.data, &got, tc.opts)
			if !reflect.DeepEqual(got, want) || !reflect.DeepEqual(gotErr, wantErr) {
				t.Fatalf("reused decode: %+v, %v; fresh: %+v, %v", got, gotErr, want, wantErr)
			}
			if !reflect.DeepEqual(dec, Decoder{}) {
				t.Fatal("decoder retained state after returning")
			}
			if len(got.data) > 0 && (&got.data[0] == &tc.data[tc.payloadOffset]) != tc.opts.AliasInput {
				t.Fatal("incorrect input aliasing")
			}
			// Follow each failure or option change with a successful default decode.
			var next reuseMessage
			nextData := []byte{18, 1, 'z', 32, 1, 32, 2}
			if err := dec.Unmarshal(nextData, &next, UnmarshalOptions{}); err != nil || !bytes.Equal(next.data, []byte{'z'}) || len(next.values) != 2 {
				t.Fatalf("state leaked into next call: %+v, %v", next, err)
			}
			nextData[2] = 'x'
			if next.data[0] != 'z' {
				t.Fatal("input aliasing leaked into default decode")
			}
		})
	}
}

func TestDecoderUnmarshalReuseAllocations(t *testing.T) {
	var dec Decoder
	var msg reuseMessage
	data := []byte{8, 7}
	if allocs := testing.AllocsPerRun(100, func() {
		if err := dec.Unmarshal(data, &msg, UnmarshalOptions{}); err != nil {
			t.Fatal(err)
		}
	}); allocs != 0 {
		t.Fatalf("reusing the decoder allocated %.0f times", allocs)
	}
}

func BenchmarkDecoderReuse(b *testing.B) {
	data := []byte{8, 7}
	for _, reuse := range []bool{false, true} {
		name := "Fresh"
		if reuse {
			name = "Reuse"
		}
		b.Run(name, func(b *testing.B) {
			var dec Decoder
			var msg reuseMessage
			b.ReportAllocs()
			for b.Loop() {
				var err error
				if reuse {
					err = dec.Unmarshal(data, &msg, UnmarshalOptions{})
				} else {
					err = Unmarshal(data, &msg)
				}
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

type reuseRequiredMessage struct{ reuseMessage }

func (m *reuseRequiredMessage) PicoValidateRequired() error {
	if m.value == 0 {
		return MissingRequiredField(1)
	}
	return nil
}

func TestDecoderReuseRequiredAndMerge(t *testing.T) {
	var dec Decoder
	var msg reuseRequiredMessage
	if err := dec.Unmarshal(nil, &msg, UnmarshalOptions{}); err == nil {
		t.Fatal("missing required field accepted")
	}
	if err := dec.Unmarshal([]byte{8, 7}, &msg, UnmarshalOptions{}); err != nil {
		t.Fatal(err)
	}
	if err := dec.Unmarshal(nil, &msg, UnmarshalOptions{}); err != nil || msg.value != 7 {
		t.Fatalf("message was not merged: %+v, %v", msg, err)
	}
}

type reusePanicMessage struct{}

func (*reusePanicMessage) Encode(*Encoder) bool { return true }
func (*reusePanicMessage) Decode(*Decoder)      { panic("decode panic") }

func TestDecoderReusePanic(t *testing.T) {
	var dec Decoder
	func() {
		defer func() {
			if r := recover(); r != "decode panic" {
				t.Fatalf("unexpected panic: %v", r)
			}
		}()
		_ = dec.Unmarshal([]byte{8, 7}, &reusePanicMessage{}, UnmarshalOptions{})
	}()
	if !reflect.DeepEqual(dec, Decoder{}) {
		t.Fatal("decoder retained state after panic")
	}
	var msg reuseMessage
	if err := dec.Unmarshal([]byte{8, 7}, &msg, UnmarshalOptions{}); err != nil || msg.value != 7 {
		t.Fatalf("reuse after panic: %+v, %v", msg, err)
	}
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"bytes"
	"fmt"
	"testing"

	"storj.io/picobuf"
)

func TestEncoderMarshalAppendReuse(t *testing.T) {
	var enc picobuf.Encoder
	for _, depth := range []int{0, 1, 16, 32, 33, 96, 97, 300, 1} {
		t.Run(fmt.Sprint(depth), func(t *testing.T) {
			msg := nestedPayload(depth, []byte{8, 1})
			want, err := picobuf.Marshal(msg)
			if err != nil {
				t.Fatal(err)
			}
			prefix := []byte{1, 2, 3}
			buffer := make([]byte, len(prefix), len(prefix)+len(want))
			copy(buffer, prefix)
			got, err := enc.MarshalAppend(buffer, msg)
			if err != nil || !bytes.Equal(got, append(prefix, want...)) {
				t.Fatalf("encoding: %x, %v", got, err)
			}
			if &got[0] != &buffer[0] {
				t.Fatal("grew a buffer with sufficient capacity")
			}
			if enc.Buffer() != nil || enc.Err() != nil {
				t.Fatal("retained buffer or error")
			}
			// Warmed scratch storage must also handle overflow without allocations.
			if allocs := testing.AllocsPerRun(100, func() {
				if _, err := enc.MarshalAppend(buffer, msg); err != nil {
					t.Fatal(err)
				}
			}); allocs != 0 {
				t.Fatalf("steady-state reuse allocated %.0f times", allocs)
			}
		})
	}
}

func TestEncoderReuseErrorAndPanic(t *testing.T) {
	for _, sizing := range []bool{true, false} {
		for _, panics := range []bool{false, true} {
			t.Run(fmt.Sprintf("sizing=%t/panic=%t", sizing, panics), func(t *testing.T) {
				var enc picobuf.Encoder
				msg := encodeFunc(func(enc *picobuf.Encoder) bool {
					value := int32(7)
					enc.Int32(1, &value)
					if enc.IsSizing() == sizing {
						if panics {
							panic("encode panic")
						}
						enc.Fail(1, "encode failure")
					}
					return true
				})
				func() {
					defer func() {
						r := recover()
						if panics && r != "encode panic" || !panics && r != nil {
							t.Fatalf("unexpected panic: %v", r)
						}
					}()
					data, err := enc.MarshalAppend(nil, msg)
					if !panics && (err == nil || data != nil) {
						t.Fatalf("missing encoding error: %x, %v", data, err)
					}
				}()
				if enc.Buffer() != nil || enc.Err() != nil {
					t.Fatal("retained failed state")
				}
				got, err := enc.MarshalAppend(nil, nestedPayload(1, []byte{8, 1}))
				if err != nil || !bytes.Equal(got, []byte{10, 2, 8, 1}) {
					t.Fatalf("reuse after failure: %x, %v", got, err)
				}
			})
		}
	}
}

func TestEncoderReuseOutputOwnership(t *testing.T) {
	var enc picobuf.Encoder
	first, err := enc.MarshalAppend(nil, nestedPayload(1, []byte{8, 1}))
	if err != nil {
		t.Fatal(err)
	}
	_, err = enc.MarshalAppend(nil, nestedPayload(1, []byte{8, 2}))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, []byte{10, 2, 8, 1}) {
		t.Fatal("later call changed prior output")
	}
}

func BenchmarkEncoderReuse(b *testing.B) {
	for _, depth := range []int{0, 1, 32, 33, 300} {
		b.Run(fmt.Sprint(depth), func(b *testing.B) {
			benchmarkEncoderReuse(b, nestedPayload(depth, []byte{8, 1}))
		})
	}
	b.Run("UploadChunk64KiB", func(b *testing.B) {
		benchmarkEncoderReuse(b, &pieceChunkMessage{upload: true, data: make([]byte, 64<<10)})
	})
}

func benchmarkEncoderReuse(b *testing.B, msg picobuf.Message) {
	for _, reuse := range []bool{false, true} {
		b.Run(fmt.Sprintf("reuse=%t", reuse), func(b *testing.B) {
			var enc picobuf.Encoder
			var buf []byte
			var err error
			b.ReportAllocs()
			for b.Loop() {
				if reuse {
					buf, err = enc.MarshalAppend(buf[:0], msg)
				} else {
					buf, err = picobuf.MarshalAppend(buf[:0], msg)
				}
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

type requiredEncodeMessage struct {
	called  bool
	present bool
}

func (m *requiredEncodeMessage) Encode(*picobuf.Encoder) bool { m.called = true; return true }
func (*requiredEncodeMessage) Decode(*picobuf.Decoder)        {}
func (m *requiredEncodeMessage) PicoValidateRequired() error {
	if !m.present {
		return picobuf.MissingRequiredField(1)
	}
	return nil
}

func TestEncoderReuseRequired(t *testing.T) {
	var enc picobuf.Encoder
	var msg requiredEncodeMessage
	if _, err := enc.MarshalAppend(nil, &msg); err == nil || msg.called {
		t.Fatal("did not validate before encoding")
	}
	msg.present = true
	if _, err := enc.MarshalAppend(nil, &msg); err != nil || !msg.called {
		t.Fatalf("encoding after validation failure: %v", err)
	}
}

func TestEncoderReuseRecursionLimit(t *testing.T) {
	var enc picobuf.Encoder
	var msg encodeFunc
	msg = func(enc *picobuf.Encoder) bool { enc.Message(1, msg); return true }
	_, err := enc.MarshalAppend(nil, msg)
	if err == nil || err.Error() != "picobuf: exceeded maximum nesting depth while encoding" {
		t.Fatalf("recursion limit: %v", err)
	}
	if _, err := enc.MarshalAppend(nil, nestedPayload(1, []byte{8, 1})); err != nil {
		t.Fatalf("reusing after recursion failure: %v", err)
	}
}

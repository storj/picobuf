// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"bytes"
	"fmt"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"storj.io/picobuf"
)

type encodeFunc func(*picobuf.Encoder) bool

func (fn encodeFunc) Encode(enc *picobuf.Encoder) bool { return fn(enc) }
func (fn encodeFunc) Decode(*picobuf.Decoder)          {}

func nestedPayload(depth int, payload []byte) encodeFunc {
	if depth == 0 {
		return func(enc *picobuf.Encoder) bool {
			enc.UnrecognizedFields(payload)
			return true
		}
	}
	child := nestedPayload(depth-1, payload)
	return func(enc *picobuf.Encoder) bool {
		enc.Message(1, child)
		return true
	}
}

func TestMarshalNestedExactCapacity(t *testing.T) {
	for _, depth := range []int{1, 8, 40} {
		for _, size := range []int{0, 127, 128, 16383, 16384, 1 << 21} {
			t.Run(fmt.Sprintf("depth=%d/size=%d", depth, size), func(t *testing.T) {
				payload := bytes.Repeat([]byte{0x08, 0x01}, size/2)
				if size%2 != 0 {
					payload = append([]byte{0x08, 0x80, 0x01}, payload[2:]...)
				}
				msg := nestedPayload(depth, payload)
				checkMarshalEncoding(t, msg, true)
			})
		}
	}
}

func checkMarshalEncoding(t *testing.T, msg picobuf.Message, exactCapacity bool) {
	t.Helper()
	want, err := picobuf.MarshalBuffer(msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	got, err := picobuf.Marshal(msg)
	if err != nil || !bytes.Equal(got, want) {
		t.Fatalf("Marshal differs from single-pass encoding: %v", err)
	}
	prefix := []byte{1, 2, 3}
	buffer := make([]byte, len(prefix), len(prefix)+len(want))
	copy(buffer, prefix)
	got, err = picobuf.MarshalAppend(buffer, msg)
	if err != nil || !bytes.Equal(got, append(prefix, want...)) {
		t.Fatalf("MarshalAppend differs from single-pass encoding: %v", err)
	}
	if exactCapacity && &got[0] != &buffer[0] {
		t.Fatal("MarshalAppend reallocated a buffer with enough capacity")
	}
}

func TestMarshalReorderedLengthHints(t *testing.T) {
	// Force a different sizing order to exercise map-like traversal changes
	// deterministically, including changes in presence and subtree size.
	omitted := encodeFunc(func(*picobuf.Encoder) bool { return false })
	messages := []encodeFunc{
		omitted, nestedPayload(0, nil), nestedPayload(1, []byte{8, 1}),
		nestedPayload(2, bytes.Repeat([]byte{8, 1}, 128)),
		nestedPayload(0, bytes.Repeat([]byte{8, 1}, 32<<10)),
	}
	for _, mode := range []string{"Message", "AlwaysMessage", "PresentMessage", "AlwaysAnyBytes"} {
		t.Run(mode, func(t *testing.T) {
			msg := encodeFunc(func(enc *picobuf.Encoder) bool {
				for i := range messages {
					if enc.IsSizing() {
						i = len(messages) - 1 - i
					}
					field, fn := picobuf.FieldNumber(16+i), messages[i]
					switch mode {
					case "Message":
						enc.Message(field, fn)
					case "AlwaysMessage":
						enc.AlwaysMessage(field, fn)
					case "PresentMessage":
						enc.PresentMessage(field, fn)
					case "AlwaysAnyBytes":
						enc.AlwaysAnyBytes(field, func() { fn(enc) })
					}
				}
				return true
			})
			checkMarshalEncoding(t, msg, false)
		})
	}
}

func TestMarshalOmittedSubtree(t *testing.T) {
	msg := encodeFunc(func(enc *picobuf.Encoder) bool {
		enc.Message(1, func(enc *picobuf.Encoder) bool {
			nestedPayload(40, []byte{8, 1})(enc)
			return false
		})
		enc.Message(2, nestedPayload(0, bytes.Repeat([]byte{8, 1}, 128)))
		enc.Message(3, func(*picobuf.Encoder) bool { return false })
		return true
	})
	checkMarshalEncoding(t, msg, true)
}

// pieceChunkMessage follows the bulk-data variants of Storj's
// PieceUploadRequest and PieceDownloadResponse: offset/data inside a chunk,
// with absent message fields before and after it. No Storj dependency is needed.
type pieceChunkMessage struct {
	upload bool
	offset int64
	data   []byte
}

func (m *pieceChunkMessage) Encode(enc *picobuf.Encoder) bool {
	omitted := func(*picobuf.Encoder) bool { return false }
	chunkField := picobuf.FieldNumber(1)
	if m.upload {
		enc.Message(1, omitted)
		enc.Message(2, omitted)
		chunkField = 3
	}
	enc.Message(chunkField, func(enc *picobuf.Encoder) bool {
		enc.Int64(1, &m.offset)
		enc.Bytes(2, &m.data)
		return true
	})
	enc.Message(chunkField+1, omitted)
	if !m.upload {
		enc.Message(3, omitted)
	}
	return true
}

func (*pieceChunkMessage) Decode(*picobuf.Decoder) {}

func TestMarshalPieceChunk(t *testing.T) {
	for _, upload := range []bool{false, true} {
		for _, size := range []int{32, 64 << 10, 4 << 20} {
			t.Run(fmt.Sprintf("upload=%v/size=%d", upload, size), func(t *testing.T) {
				msg := &pieceChunkMessage{upload: upload, offset: 64 << 10, data: bytes.Repeat([]byte{0xab}, size)}
				chunk := protowire.AppendTag(nil, 1, protowire.VarintType)
				chunk = protowire.AppendVarint(chunk, uint64(msg.offset))
				chunk = protowire.AppendTag(chunk, 2, protowire.BytesType)
				chunk = protowire.AppendBytes(chunk, msg.data)
				field := protowire.Number(1)
				if upload {
					field = 3
				}
				want := protowire.AppendBytes(protowire.AppendTag(nil, field, protowire.BytesType), chunk)
				got, err := picobuf.Marshal(msg)
				if err != nil || !bytes.Equal(got, want) {
					t.Fatalf("incorrect chunk encoding: %v", err)
				}
				checkMarshalEncoding(t, msg, true)
			})
		}
	}
}

func BenchmarkMarshalPieceChunk(b *testing.B) {
	for _, upload := range []bool{true, false} {
		for _, size := range []int{32, 64 << 10, 4 << 20} {
			b.Run(fmt.Sprintf("upload=%v/size=%d", upload, size), func(b *testing.B) {
				msg := &pieceChunkMessage{upload: upload, offset: 64 << 10, data: bytes.Repeat([]byte{0xab}, size)}
				benchmarkMarshal(b, msg)
			})
		}
	}
}

func benchmarkMarshal(b *testing.B, msg picobuf.Message) {
	b.Run("Marshal", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if _, err := picobuf.Marshal(msg); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("MarshalAppend", func(b *testing.B) {
		size, err := picobuf.Size(msg)
		if err != nil {
			b.Fatal(err)
		}
		buffer := make([]byte, 0, size)
		// Warm the reused buffer so a baseline encoder's first growth is not
		// amortized into steady-state allocation measurements.
		buffer, err = picobuf.MarshalAppend(buffer, msg)
		if err != nil {
			b.Fatal(err)
		}
		b.ReportAllocs()
		for b.Loop() {
			buffer, err = picobuf.MarshalAppend(buffer[:0], msg)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkMarshalManySubmessages(b *testing.B) {
	// Model the 39 addressed limits in DownloadSegmentResponse, each with
	// two nested messages. This also exercises spilling the inline hints.
	limit, address := bytes.Repeat([]byte{0xab}, 256), bytes.Repeat([]byte{'a'}, 32)
	msg := encodeFunc(func(enc *picobuf.Encoder) bool {
		for range 39 {
			enc.Message(1, func(enc *picobuf.Encoder) bool {
				enc.Message(1, func(enc *picobuf.Encoder) bool {
					enc.Bytes(1, &limit)
					return true
				})
				enc.Message(2, func(enc *picobuf.Encoder) bool {
					enc.Bytes(1, &address)
					return true
				})
				return true
			})
		}
		return true
	})
	benchmarkMarshal(b, msg)
}

func BenchmarkMarshalLengthHints(b *testing.B) {
	for _, tc := range []struct {
		name        string
		depth, size int
	}{
		{"Flat", 0, 32},
		{"Small", 1, 32},
		{"Chunk64KiB", 1, 64 << 10},
		{"Chunk4MiB", 1, 4 << 20},
		{"Nested", 8, 32},
		{"Spill", 40, 32},
	} {
		b.Run(tc.name, func(b *testing.B) {
			msg := nestedPayload(tc.depth, bytes.Repeat([]byte{0x08, 0x01}, tc.size/2))
			b.ReportAllocs()
			for b.Loop() {
				if _, err := picobuf.Marshal(msg); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func TestMarshalPackedExactCapacity(t *testing.T) {
	// Packed repeated varints reserve a length hint in the sizing pass like
	// any other length-delimited field, so the hints after them stay in sync.
	nums := make([]int64, 200)
	for i := range nums {
		nums[i] = 1
	}
	payload := bytes.Repeat([]byte{1}, 10)
	for _, omitted := range []bool{false, true} {
		t.Run(fmt.Sprintf("omitted=%v", omitted), func(t *testing.T) {
			msg := encodeFunc(func(enc *picobuf.Encoder) bool {
				enc.RepeatedInt64(1, &nums)
				enc.Message(2, func(enc *picobuf.Encoder) bool {
					enc.Bytes(1, &payload)
					return !omitted
				})
				enc.AlwaysRepeatedInt32(3, &[]int32{1, 2, 3})
				enc.Message(4, func(enc *picobuf.Encoder) bool {
					enc.Bytes(1, &payload)
					return true
				})
				return true
			})
			checkMarshalEncoding(t, msg, true)
		})
	}
}

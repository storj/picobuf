// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"math"
	"strconv"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"

	"storj.io/picobuf"
)

// stateMessage visits nested messages without allocating application objects,
// so the tests can isolate the decoder's own nesting overhead.
type stateMessage struct {
	mode  string
	value int32
}

func (*stateMessage) Encode(*picobuf.Encoder) bool { return true }

func (m *stateMessage) Decode(dec *picobuf.Decoder) {
	switch m.mode {
	case "Message":
		dec.Message(1, m.Decode)
	case "PresentMessage":
		dec.PresentMessage(1, m.Decode)
	case "RepeatedMessage":
		dec.RepeatedMessage(1, func(dec *picobuf.Decoder) { dec.Loop(m.Decode) })
	case "Group":
		dec.Group(1, m.Decode)
	case "RepeatedGroup":
		dec.RepeatedGroup(1, func(dec *picobuf.Decoder) { dec.Loop(m.Decode) })
	}
	dec.Int32(2, &m.value)
}

func stateMessageData(mode string, depth int) []byte {
	data := []byte{0x10, 7}
	for i := range depth {
		if mode == "Group" || mode == "RepeatedGroup" {
			data = append(append([]byte{0x0b}, data...), 0x0c)
		} else {
			data = protowire.AppendBytes([]byte{0x0a}, data)
		}
		// A field after each child verifies restoration of the parent's cursor.
		data = append(data, 0x10, byte(8+i))
	}
	return data
}

func TestDecoderNestingAllocations(t *testing.T) {
	for _, mode := range []string{"Message", "PresentMessage", "RepeatedMessage", "Group", "RepeatedGroup"} {
		t.Run(mode, func(t *testing.T) {
			msg := &stateMessage{mode: mode}
			measure := func(data []byte, want int32) float64 {
				return testing.AllocsPerRun(100, func() {
					msg.value = 0
					if err := picobuf.Unmarshal(data, msg); err != nil {
						t.Fatal(err)
					}
					if msg.value != want {
						t.Fatalf("decoded value = %d; want %d", msg.value, want)
					}
				})
			}
			flat := measure(stateMessageData(mode, 0), 7)
			for _, depth := range []int{1, 8, 64} {
				if nested := measure(stateMessageData(mode, depth), int32(7+depth)); nested != flat {
					t.Errorf("depth %d: %.0f allocations, flat: %.0f", depth, nested, flat)
				}
			}
		})
	}
}

func TestDecoderNestingLimits(t *testing.T) {
	for _, mode := range []string{"Message", "PresentMessage", "RepeatedMessage", "Group", "RepeatedGroup"} {
		t.Run(mode, func(t *testing.T) {
			opts := picobuf.UnmarshalOptions{MaxRecursionDepth: 2}
			msg := &stateMessage{mode: mode}
			if err := opts.Unmarshal(stateMessageData(mode, 3), msg); err == nil {
				t.Fatal("accepted a message beyond the recursion limit")
			}
			// Siblings each get the full budget back after the previous child.
			data := stateMessageData(mode, 2)
			data = append(data, data...)
			if err := opts.Unmarshal(data, msg); err != nil {
				t.Fatalf("failed to restore recursion budget between siblings: %v", err)
			}
		})
	}
}

func TestDecoderLargeLimits(t *testing.T) {
	if strconv.IntSize < 64 {
		t.Skip("limits larger than int32 require a 64-bit int")
	}
	data := stateMessageData("RepeatedMessage", 2)
	for _, limit := range []int64{math.MaxInt32, math.MaxInt32 + 1, 1 << 32, math.MaxInt64} {
		opts := picobuf.UnmarshalOptions{
			MaxRecursionDepth: int(limit), MaxRepeatedElements: int(limit),
		}
		msg := &stateMessage{mode: "RepeatedMessage"}
		if err := opts.Unmarshal(data, msg); err != nil {
			t.Fatalf("limit %d wrapped or rejected a small message: %v", limit, err)
		}
	}
}

func BenchmarkDecoderNesting(b *testing.B) {
	for _, mode := range []string{"Message", "Group"} {
		b.Run(mode, func(b *testing.B) {
			data := stateMessageData(mode, 8)
			msg := &stateMessage{mode: mode}
			b.ReportAllocs()
			for b.Loop() {
				if err := picobuf.Unmarshal(data, msg); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

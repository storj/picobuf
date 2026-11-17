// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picoconv

import (
	"math"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/protowire"

	"storj.io/picobuf"
)

func timeParts(field protowire.Number, value int64) []byte {
	body := protowire.AppendVarint(protowire.AppendTag(nil, field, protowire.VarintType), uint64(value))
	return protowire.AppendBytes([]byte{10}, body)
}

func TestTimestampMergesOccurrences(t *testing.T) {
	for _, seconds := range []int64{0, 123, minTimestampSeconds} {
		data := append(timeParts(1, seconds), timeParts(2, 456)...)
		data = append(data, 10, 0) // an empty occurrence preserves the value
		var m timestampMessage
		if err := picobuf.Unmarshal(data, &m); err != nil {
			t.Fatal(err)
		}
		if got := time.Time(m.value); !got.Equal(time.Unix(seconds, 456)) {
			t.Fatalf("got %v, want seconds=%d nanos=456", got, seconds)
		}
	}
}

func TestDurationMergesOccurrences(t *testing.T) {
	data := append(timeParts(1, 123), timeParts(2, 456)...)
	data = append(data, 10, 0)
	var m DurationMessage
	if err := picobuf.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m.Value != Duration(123*time.Second+456) {
		t.Fatalf("got %v", m.Value)
	}
	// Retain raw seconds even when the intermediate value saturates time.Duration.
	data = append(timeParts(1, 315576000000), timeParts(2, 1)...)
	if err := picobuf.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m.Value != Duration(math.MaxInt64) {
		t.Fatalf("lost overflowing seconds: %v", m.Value)
	}
}

func TestConverterDecoderReuse(t *testing.T) {
	var dec picobuf.Decoder
	var m timestampMessage
	if err := dec.Unmarshal(timeParts(1, 123), &m, picobuf.UnmarshalOptions{}); err != nil {
		t.Fatal(err)
	}
	m = timestampMessage{}
	if err := dec.Unmarshal(timeParts(2, 456), &m, picobuf.UnmarshalOptions{}); err != nil {
		t.Fatal(err)
	}
	if !time.Time(m.value).Equal(time.Unix(0, 456)) {
		t.Fatal("decoder retained previous message state")
	}
}

func TestTimestampMergesExistingValue(t *testing.T) {
	for _, tc := range []struct {
		name string
		data []byte
		want time.Time
	}{
		{"empty", []byte{10, 0}, time.Unix(123, 456).UTC()},
		{"seconds", timeParts(1, 789), time.Unix(789, 456).UTC()},
		{"nanos", timeParts(2, 789), time.Unix(123, 789).UTC()},
		{"zero seconds", timeParts(1, 0), time.Unix(0, 456).UTC()},
		{"zero nanos", timeParts(2, 0), time.Unix(123, 0).UTC()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := timestampMessage{value: Timestamp(time.Unix(123, 456).UTC())}
			if err := picobuf.Unmarshal(tc.data, &m); err != nil {
				t.Fatal(err)
			}
			if got := time.Time(m.value); !got.Equal(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDurationMergesExistingValue(t *testing.T) {
	for _, sign := range []int64{1, -1} {
		for _, tc := range []struct {
			name string
			data []byte
			want time.Duration
		}{
			{"empty", []byte{10, 0}, time.Duration(sign * (123e9 + 456))},
			{"seconds", timeParts(1, sign*789), time.Duration(sign * (789e9 + 456))},
			{"nanos", timeParts(2, sign*789), time.Duration(sign * (123e9 + 789))},
			{"zero seconds", timeParts(1, 0), time.Duration(sign * 456)},
			{"zero nanos", timeParts(2, 0), time.Duration(sign * 123e9)},
		} {
			t.Run(tc.name, func(t *testing.T) {
				m := DurationMessage{Value: Duration(sign * (123e9 + 456))}
				if err := picobuf.Unmarshal(tc.data, &m); err != nil {
					t.Fatal(err)
				}
				if got := time.Duration(m.Value); got != tc.want {
					t.Fatalf("got %v, want %v", got, tc.want)
				}
			})
		}
	}
}

func TestConverterDecodeDoesNotAllocate(t *testing.T) {
	timestamp := append(timeParts(1, 123), timeParts(2, 456)...)
	duration := timeParts(1, 123)

	var dec picobuf.Decoder
	var tm timestampMessage
	var dm DurationMessage
	allocs := testing.AllocsPerRun(100, func() {
		tm, dm = timestampMessage{}, DurationMessage{}
		if err := dec.Unmarshal(timestamp, &tm, picobuf.UnmarshalOptions{}); err != nil {
			t.Fatal(err)
		}
		if err := dec.Unmarshal(duration, &dm, picobuf.UnmarshalOptions{}); err != nil {
			t.Fatal(err)
		}
	})
	if allocs != 0 {
		t.Fatalf("got %v allocations, want 0", allocs)
	}
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package perftest

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	pico "storj.io/picobuf/internal/perftest/pico"
	present "storj.io/picobuf/internal/perftest/present"
	prot "storj.io/picobuf/internal/perftest/prot"
)

func TestProfileCrossRoundTrip(t *testing.T) {
	original, data := profileFixture(t)

	var decoded pico.Profile
	if err := picobuf.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	reencoded, err := picobuf.Marshal(&decoded)
	if err != nil {
		t.Fatal(err)
	}

	var roundtrip prot.Profile
	if err := proto.Unmarshal(reencoded, &roundtrip); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(original, &roundtrip) {
		t.Fatal("profile changed during protobuf -> picobuf -> protobuf roundtrip")
	}

	var presentDecoded present.Profile
	if err := picobuf.Unmarshal(data, &presentDecoded); err != nil {
		t.Fatal(err)
	}
	presentData, err := picobuf.Marshal(&presentDecoded)
	if err != nil {
		t.Fatal(err)
	}
	if err := proto.Unmarshal(presentData, &roundtrip); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(original, &roundtrip) {
		t.Fatal("profile changed during always-present picobuf roundtrip")
	}
}

func BenchmarkProfileDecode(b *testing.B) {
	_, data := profileFixture(b)

	b.Run("Protobuf", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		for b.Loop() {
			var profile prot.Profile
			if err := proto.Unmarshal(data, &profile); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("Picobuf", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		for b.Loop() {
			var profile pico.Profile
			if err := picobuf.Unmarshal(data, &profile); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("PicobufPresent", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		for b.Loop() {
			var profile present.Profile
			if err := picobuf.Unmarshal(data, &profile); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkProfileEncode(b *testing.B) {
	protobufProfile, data := profileFixture(b)
	var picobufProfile pico.Profile
	if err := picobuf.Unmarshal(data, &picobufProfile); err != nil {
		b.Fatal(err)
	}
	var presentProfile present.Profile
	if err := picobuf.Unmarshal(data, &presentProfile); err != nil {
		b.Fatal(err)
	}

	b.Run("Protobuf", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		for b.Loop() {
			if _, err := proto.Marshal(protobufProfile); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("Picobuf", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		for b.Loop() {
			if _, err := picobuf.Marshal(&picobufProfile); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("PicobufPresent", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(data)))
		for b.Loop() {
			if _, err := picobuf.Marshal(&presentProfile); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func profileFixture(tb testing.TB) (*prot.Profile, []byte) {
	tb.Helper()
	profile := makeProfile()
	data, err := proto.Marshal(profile)
	if err != nil {
		tb.Fatal(err)
	}
	return profile, data
}

func makeProfile() *prot.Profile {
	const (
		functionCount = 200
		locationCount = 500
		sampleCount   = 2000
	)

	profile := &prot.Profile{
		SampleType: []*prot.ValueType{{Type: 1, Unit: 2}, {Type: 3, Unit: 4}},
		Mapping: []*prot.Mapping{{
			Id: 1, MemoryStart: 0x1000000, MemoryLimit: 0x2000000,
			Filename: 5, BuildId: 6, HasFunctions: true, HasFilenames: true,
			HasLineNumbers: true, HasInlineFrames: true,
		}},
		StringTable:       []string{"", "cpu", "nanoseconds", "alloc_objects", "count", "profile.bin", "build-id", "goroutine"},
		TimeNanos:         1_700_000_000_000_000_000,
		DurationNanos:     30_000_000_000,
		PeriodType:        &prot.ValueType{Type: 1, Unit: 2},
		Period:            10_000_000,
		Comment:           []int64{7},
		DefaultSampleType: 1,
	}

	for i := 1; i <= functionCount; i++ {
		profile.StringTable = append(profile.StringTable,
			fmt.Sprintf("function.%d", i), fmt.Sprintf("file_%d.go", i%25))
		profile.Function = append(profile.Function, &prot.Function{
			Id: uint64(i), Name: int64(len(profile.StringTable) - 2),
			SystemName: int64(len(profile.StringTable) - 2),
			Filename:   int64(len(profile.StringTable) - 1), StartLine: int64(10 + i),
		})
	}
	for i := 1; i <= locationCount; i++ {
		functionID := uint64((i-1)%functionCount + 1)
		profile.Location = append(profile.Location, &prot.Location{
			Id: uint64(i), MappingId: 1, Address: uint64(0x1000000 + i*16),
			Line: []*prot.Line{{FunctionId: functionID, Line: int64(10 + i%200)}},
		})
	}
	for i := 0; i < sampleCount; i++ {
		locations := make([]uint64, 16)
		for frame := range locations {
			locations[frame] = uint64((i*7+frame*13)%locationCount + 1)
		}
		profile.Sample = append(profile.Sample, &prot.Sample{
			LocationId: locations,
			Value:      []int64{int64(1 + i%100), int64(1024 + i*64)},
			Label:      []*prot.Label{{Key: 7, Num: int64(i % 128), NumUnit: 4}},
		})
	}
	return profile
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package protocompat

import (
	"math/rand"
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/protocompat/pico"
	"storj.io/picobuf/internal/protocompat/prot"
)

func BenchmarkProtobuf(b *testing.B) {
	data := generateProtobuf(1000)
	encoded := [][]byte{}
	for i := range data {
		bytes, _ := proto.Marshal(&data[i])
		encoded = append(encoded, bytes)
	}
	b.Run("Encode", func(b *testing.B) {
		b.ReportAllocs()
		for k := 0; k < b.N; k++ {
			for i := range data {
				_, _ = proto.Marshal(&data[i])
			}
		}
	})
	b.Run("Decode", func(b *testing.B) {
		b.ReportAllocs()
		for k := 0; k < b.N; k++ {
			for _, bytes := range encoded {
				var x prot.Person
				_ = proto.Unmarshal(bytes, &x)
			}
		}
	})
}

func BenchmarkPicobuf(b *testing.B) {
	data := generatePicobuf(1000)
	encoded := [][]byte{}
	for i := range data {
		bytes, _ := picobuf.Marshal(&data[i])
		encoded = append(encoded, bytes)
	}
	b.Run("Encode", func(b *testing.B) {
		b.ReportAllocs()
		var buf [64]byte
		for k := 0; k < b.N; k++ {
			for i := range data {
				_, _ = picobuf.MarshalBuffer(&data[i], buf[:])
			}
		}

	})
	b.Run("Decode", func(b *testing.B) {
		b.ReportAllocs()
		for k := 0; k < b.N; k++ {
			for _, bytes := range encoded {
				var x pico.Person
				_ = picobuf.Unmarshal(bytes, &x)
			}
		}
	})
}

func generateProtobuf(n int) []prot.Person {
	r := rand.New(rand.NewSource(int64(n)))
	xs := make([]prot.Person, n)
	for i := range xs {
		xs[i] = prot.Person{
			Name:     randString(r, 16),
			Birthday: r.Int63(),
			Phone:    randString(r, 10),
			Siblings: r.Int31n(5),
			Spouse:   r.Intn(2) == 1,
			Money:    r.Float64(),
			Primary:  prot.Language(r.Intn(5)),
			Spoken:   []prot.Language{prot.Language(r.Intn(5))},
		}
	}
	return xs
}

func generatePicobuf(n int) []pico.Person {
	r := rand.New(rand.NewSource(int64(n)))
	xs := make([]pico.Person, n)
	for i := range xs {
		xs[i] = pico.Person{
			Name:     randString(r, 16),
			Birthday: r.Int63(),
			Phone:    randString(r, 10),
			Siblings: r.Int31n(5),
			Spouse:   r.Intn(2) == 1,
			Money:    r.Float64() * 1000,
			Primary:  pico.Language(r.Intn(5)),
			Spoken:   []pico.Language{pico.Language(r.Intn(5))},
		}
	}
	return xs
}

func randString(r *rand.Rand, n int) string {
	const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	xs := make([]byte, n)
	_, _ = r.Read(xs)
	for i, x := range xs {
		xs[i] = alpha[int(x)%len(alpha)]
	}
	return string(xs)
}

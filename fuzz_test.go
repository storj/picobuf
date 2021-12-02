// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

//go:build go1.18
// +build go1.18

package picobuf_test

import (
	"testing"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest"
)

func FuzzDecode(f *testing.F) {
	x, err := picobuf.Marshal(&picotest.AllTypes{
		Int32:     1,
		Int64:     1,
		Uint32:    1,
		Uint64:    1,
		Sint32:    1,
		Sint64:    1,
		Fixed32:   1,
		Fixed64:   1,
		Sfixed32:  1,
		Sfixed64:  1,
		Float:     1,
		Double:    1,
		Bool:      true,
		String_:   "hello",
		Bytes:     []byte{0x1, 0x2},
		Message:   &picotest.Message{Int32: 1},
		Int32S:    []int32{1, 2},
		Int64S:    []int64{1, 2},
		Uint32S:   []uint32{1, 2},
		Uint64S:   []uint64{1, 2},
		Sint32S:   []int32{1, 2},
		Sint64S:   []int64{1, 2},
		Fixed32S:  []uint32{1, 2},
		Fixed64S:  []uint64{1, 2},
		Sfixed32S: []int32{1, 2},
		Sfixed64S: []int64{1, 2},
		Floats:    []float32{1, 2},
		Doubles:   []float64{1, 2},
		Bools:     []bool{true, true},
		Strings:   []string{"hello", "world"},
		Bytess:    [][]byte{{1}, {2}},
		Messages:  []*picotest.Message{{Int32: 1}, {Int32: 2}},
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(x)

	f.Fuzz(func(t *testing.T, data []byte) {
		var z picotest.AllTypes
		err := picobuf.Unmarshal(data, &z)
		if err != nil {
			return
		}

		_, err = picobuf.Marshal(&z)
		if err != nil {
			t.Fatal(err)
		}
	})
}

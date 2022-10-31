// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package pico

import (
	"fmt"

	"storj.io/picobuf"
)

// Test tests serialization of v.
//
//go:noinline
func Test(v picobuf.Message) {
	data, _ := picobuf.Marshal(v)
	_ = picobuf.Unmarshal(data, v)
	fmt.Println(data)
}

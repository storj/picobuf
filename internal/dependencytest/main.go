// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf"
)

func main() {
	var z byte
	enc := picobuf.NewEncoder()
	enc.Byte(1, &z)
	fmt.Println(enc.Buffer())
}

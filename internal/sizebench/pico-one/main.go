// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf/internal/sizebench/pico"
	"storj.io/picobuf/internal/sizebench/pico/one"
)

func main() {
	fmt.Println([]byte{})
	pico.Test(&one.Nop{})
	pico.Test(&one.Types{})
}

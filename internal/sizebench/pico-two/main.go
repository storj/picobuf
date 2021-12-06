// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf/internal/sizebench/pico"
	"storj.io/picobuf/internal/sizebench/pico/two"
)

func main() {
	fmt.Println([]byte{})
	pico.Test(&two.Nop{})
	pico.Test(&two.Types{})
	pico.Test(&two.Types2{})
}

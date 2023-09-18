// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"storj.io/picobuf/internal/sizebench/pico"
	"storj.io/picobuf/internal/sizebench/pico/two"
)

func main() {
	pico.Test(&two.Nop{})
	pico.Test(&two.Types{})
	pico.Test(&two.Types2{})
}

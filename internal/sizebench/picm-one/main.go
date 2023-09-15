// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"storj.io/picobuf/internal/sizebench/pico"
	"storj.io/picobuf/internal/sizebench/pico/sml"
)

func main() {
	pico.Test(&sml.Nop{})
	pico.Test(&sml.Types{})
}

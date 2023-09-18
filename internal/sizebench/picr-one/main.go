// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"storj.io/picobuf/internal/sizebench/pico"
	"storj.io/picobuf/internal/sizebench/pico/one"
)

func main() {
	pico.Test(&one.Nop{})
	pico.Test(&one.Types{})
}

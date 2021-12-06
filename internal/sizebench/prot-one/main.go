// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf/internal/sizebench/prot"
	"storj.io/picobuf/internal/sizebench/prot/one"
)

func main() {
	fmt.Println([]byte{})
	prot.Test(&one.Nop{})
	prot.Test(&one.Types{})
}

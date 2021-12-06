// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf/internal/sizebench/prot"
	"storj.io/picobuf/internal/sizebench/prot/two"
)

func main() {
	fmt.Println([]byte{})
	prot.Test(&two.Nop{})
	prot.Test(&two.Types{})
	prot.Test(&two.Types2{})
}

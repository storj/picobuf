// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf/internal/sizebench/prot/par"
)

func main() {
	fmt.Println([]byte{})

	var y par.Types
	_, _ = proto.Marshal(&y)
	_ = proto.Unmarshal([]byte{}, &y)
}

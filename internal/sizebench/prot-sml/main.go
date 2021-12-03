// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf/internal/sizebench/prot/sml"
)

func main() {
	fmt.Println([]byte{})

	var z sml.Types
	_, _ = proto.Marshal(&z)
	_ = proto.Unmarshal([]byte{}, &z)
}
// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/sizebench/pico/bas"
)

func main() {
	fmt.Println([]byte{})
	var z bas.Types
	_, _ = picobuf.Marshal(&z)
	_ = picobuf.Unmarshal([]byte{}, &z)
}

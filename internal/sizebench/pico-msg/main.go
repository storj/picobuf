// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/sizebench/pico/bas"
	"storj.io/picobuf/internal/sizebench/pico/snd"
)

func main() {
	fmt.Println([]byte{})
	var z bas.Types
	_, _ = picobuf.Marshal(&z)
	_ = picobuf.Unmarshal([]byte{}, &z)

	var y snd.Types
	_, _ = picobuf.Marshal(&y)
	_ = picobuf.Unmarshal([]byte{}, &y)
}

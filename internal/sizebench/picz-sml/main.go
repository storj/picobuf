// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"reflect"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/sizebench/pico/sml"
)

func main() {
	fmt.Println([]byte{})

	var z sml.Types
	_, _ = picobuf.Marshal(&z)
	_ = picobuf.Unmarshal([]byte{}, &z)

	v := reflect.ValueOf(fmt.Println)
	_ = v.MethodByName("Parsed").Call(nil)

	v = reflect.ValueOf(&z)
	_ = v.MethodByName("Picobuf").Call(nil)
}

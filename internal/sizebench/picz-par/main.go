// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"reflect"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/sizebench/pico/par"
)

func main() {
	fmt.Println([]byte{})

	var y par.Types
	_, _ = picobuf.Marshal(&y)
	_ = picobuf.Unmarshal([]byte{}, &y)

	v := reflect.ValueOf(fmt.Println)
	_ = v.MethodByName("Parsed").Call(nil)

	v = reflect.ValueOf(&y)
	_ = v.MethodByName("Picobuf").Call(nil)
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println([]byte{})

	v := reflect.ValueOf(fmt.Println)
	_ = v.MethodByName("Parsed").Call(nil)
}

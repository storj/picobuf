// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package basz

import (
	"reflect"
)

// CallDynamicMethod calls a method dynamically on v.
// This forces conservative mode for everything reachable from v.
//
//go:noinline
func CallDynamicMethod(v interface{}) {
	reflect.ValueOf(v).MethodByName("Parsed").Call(nil)
}

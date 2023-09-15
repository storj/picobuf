// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package prot

import (
	"google.golang.org/protobuf/proto"
)

// Test tests serialization of v.
//
//go:noinline
func Test(v proto.Message) {
	data, _ := proto.Marshal(v)
	_ = proto.Unmarshal(data, v)
	ignore(data)
}

//go:noinline
func ignore(data []byte) {}

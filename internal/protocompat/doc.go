// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// Package protocompat holds compatibility tests with protobuf.
package protocompat

//go:generate protoc -I../.. -I. --go_out=paths=source_relative:./prot --go_opt=Mtypes.proto=storj.io/picobuf/internal/protocompat/prot types.proto
//go:generate protoc -I../.. -I. --pico_out=paths=source_relative:./pico --pico_opt=Mtypes.proto=storj.io/picobuf/internal/protocompat/pico types.proto

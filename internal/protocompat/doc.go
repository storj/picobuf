// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// Package protocompat holds compatibility tests with protobuf.
package protocompat

//go:generate protoc --go_out=paths=source_relative:. --pico_out=paths=source_relative,suffix=Pico:. types.proto

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// Package protocompat holds compatibility tests with protobuf.
package protocompat

//go:generate protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-drpc_out=paths=source_relative:. types.proto

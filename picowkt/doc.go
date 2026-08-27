// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

// Package picowkt implements the common well-known types with picobuf.
//
// Map the well-known imports to this package when generating:
//
//	protoc --pico_opt=Mgoogle/protobuf/wrappers.proto=storj.io/picobuf/picowkt ...
package picowkt

//go:generate protoc --pico_opt=version_override=dev --pico_out=paths=source_relative:. wrappers.proto timestamp.proto duration.proto empty.proto field_mask.proto

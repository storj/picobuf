// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

// Package jsoncompat compares the ProtoJSON of picojson with protojson.
package jsoncompat

//go:generate protoc -I../.. -I. --go_out=paths=source_relative:./prot --go_opt=Mjsoncompat.proto=storj.io/picobuf/internal/jsoncompat/prot jsoncompat.proto
//go:generate protoc -I../.. -I. --pico_opt=version_override=dev,protojson=true --pico_out=paths=source_relative:./pico --pico_opt=Mjsoncompat.proto=storj.io/picobuf/internal/jsoncompat/pico --pico_opt=Mgoogle/protobuf/duration.proto=storj.io/picobuf/picowkt;picowkt --pico_opt=Mgoogle/protobuf/empty.proto=storj.io/picobuf/picowkt;picowkt --pico_opt=Mgoogle/protobuf/field_mask.proto=storj.io/picobuf/picowkt;picowkt --pico_opt=Mgoogle/protobuf/timestamp.proto=storj.io/picobuf/picowkt;picowkt --pico_opt=Mgoogle/protobuf/wrappers.proto=storj.io/picobuf/picowkt;picowkt jsoncompat.proto

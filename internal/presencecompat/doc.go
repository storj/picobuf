// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

//go:generate protoc -I../.. -I. --go_out=paths=source_relative:./prot --go_opt=Mschema.proto=storj.io/picobuf/internal/presencecompat/prot schema.proto
//go:generate protoc -I../.. -I. --pico_out=paths=source_relative:./pico --pico_opt=Mschema.proto=storj.io/picobuf/internal/presencecompat/pico schema.proto
//go:generate protoc -I../.. -I. --go_out=paths=source_relative:./prot2 --go_opt=Mproto2.proto=storj.io/picobuf/internal/presencecompat/prot2 proto2.proto
//go:generate protoc -I../.. -I. --pico_out=paths=source_relative:./pico2 --pico_opt=Mproto2.proto=storj.io/picobuf/internal/presencecompat/pico2 proto2.proto
//go:generate protoc -I../.. -I. --go_out=paths=source_relative:./editionprot --go_opt=Meditions.proto=storj.io/picobuf/internal/presencecompat/editionprot editions.proto
//go:generate protoc -I../.. -I. --pico_out=paths=source_relative:./editionpico --pico_opt=Meditions.proto=storj.io/picobuf/internal/presencecompat/editionpico editions.proto

// Package presencecompat tests protobuf field-presence interoperability.
package presencecompat

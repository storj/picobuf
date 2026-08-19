// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

// Package perftest contains representative performance fixtures.
package perftest

//go:generate protoc -I../.. -I. --go_out=paths=source_relative:./prot --go_opt=Mprofile.proto=storj.io/picobuf/internal/perftest/prot profile.proto
//go:generate protoc -I../.. -I. --pico_opt=version_override=dev --pico_out=paths=source_relative:./pico --pico_opt=Mprofile.proto=storj.io/picobuf/internal/perftest/pico profile.proto
//go:generate protoc -I../.. -I. --pico_opt=version_override=dev --pico_out=paths=source_relative:./present profile_present.proto

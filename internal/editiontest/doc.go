// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

//go:generate protoc -I../.. -I. --pico_opt=version_override=dev --pico_out=paths=source_relative:. test.proto

// Package editiontest exercises Editions-specific generated code.
package editiontest

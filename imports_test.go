// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"go/build"
	"slices"
	"testing"
)

// Every generated file imports picobuf, so it must stay free of heavy
// dependencies such as encoding/json.
func TestPackageDoesNotImportJSON(t *testing.T) {
	pkg, err := build.ImportDir(".", 0)
	if err != nil {
		t.Fatal(err)
	}
	if slices.Contains(pkg.Imports, "encoding/json") {
		t.Fatal("picobuf imports encoding/json")
	}
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main_test

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zeebo/assert"
)

func TestDependencies(t *testing.T) {
	exe := filepath.Join(t.TempDir(), "main.exe")

	out, err := exec.Command("go", "build", "-o", exe, ".").CombinedOutput()
	if err != nil {
		t.Log(string(out))
	}
	assert.NoError(t, err)

	nmout, err := exec.Command("go", "tool", "nm", exe).CombinedOutput()
	if err != nil {
		t.Log(string(nmout))
	}
	assert.NoError(t, err)

	if strings.Contains(string(nmout), "MethodByName") {
		t.Error("binary contains a reference to \"MethodByName\", which forces compiler to conservative mode")
		t.Fatal()
	}
}

// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

package presencecompat

import (
	"os"
	"os/exec"
	"runtime/debug"
	"testing"

	"google.golang.org/protobuf/proto"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/presencecompat/editionpico"
	"storj.io/picobuf/internal/presencecompat/editionprot"
)

func TestRequiredCycle(t *testing.T) {
	if os.Getenv("PICOBUF_TEST_REQUIRED_CYCLE") == "1" {
		debug.SetMaxStack(32 << 20)
		m := &editionpico.RequiredNode{Value: pointerTo(int32(1))}
		m.Child = m
		if _, err := picobuf.Marshal(m); err == nil {
			t.Fatal("accepted a cyclic message")
		}
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRequiredCycle$")
	cmd.Env = append(os.Environ(), "PICOBUF_TEST_REQUIRED_CYCLE=1")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("validation did not return an error safely: %v\n%.1000s", err, out)
	}
}

func TestRequiredNilRepeated(t *testing.T) {
	reference := &editionprot.RequiredNode{Value: pointerTo(int32(1)), Children: []*editionprot.RequiredNode{nil}}
	if _, err := proto.Marshal(reference); err == nil {
		t.Fatal("reference accepted nil required child")
	}
	m := &editionpico.RequiredNode{Value: pointerTo(int32(1)), Children: []*editionpico.RequiredNode{nil}}
	if _, err := picobuf.Marshal(m); err == nil {
		t.Fatal("encoded nil repeated child without required field")
	}
}

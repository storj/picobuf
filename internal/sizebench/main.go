// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"text/tabwriter"

	"golang.org/x/sync/errgroup"
)

// Note, the names of packages has been chosen to be exactly the same to reduce the difference
// in package names that end up in the binary.

//go:generate protoc --pico_out=paths=source_relative:./pico/bas --pico_opt=Mmsg-bas.proto=storj.io/picobuf/internal/sizebench/pico/bas msg-bas.proto
//go:generate protoc --pico_out=paths=source_relative:./pico/snd --pico_opt=Mmsg-bas.proto=storj.io/picobuf/internal/sizebench/pico/snd msg-bas.proto
//go:generate protoc --pico_out=paths=source_relative:./pico/par --pico_opt=Mmsg-par.proto=storj.io/picobuf/internal/sizebench/pico/par msg-par.proto
//go:generate protoc --pico_out=paths=source_relative:./pico/sml --pico_opt=Mmsg-sml.proto=storj.io/picobuf/internal/sizebench/pico/sml msg-sml.proto
//go:generate protoc --go_out=paths=source_relative:./prot/bas --go_opt=Mmsg-bas.proto=storj.io/picobuf/internal/sizebench/prot/bas msg-bas.proto
//go:generate protoc --go_out=paths=source_relative:./prot/snd --go_opt=Mmsg-bas.proto=storj.io/picobuf/internal/sizebench/prot/snd msg-bas.proto
//go:generate protoc --go_out=paths=source_relative:./prot/par --go_opt=Mmsg-par.proto=storj.io/picobuf/internal/sizebench/prot/par msg-par.proto
//go:generate protoc --go_out=paths=source_relative:./prot/sml --go_opt=Mmsg-sml.proto=storj.io/picobuf/internal/sizebench/prot/sml msg-sml.proto

// OsArch creates a easy map of OS and Arch combinations.
type OsArch struct {
	Os   string
	Arch string
}

var osarches = []OsArch{
	{"linux", "amd64"},
	// {"darwin", "amd64"},
	// {"windows", "amd64"},
	{"js", "wasm"},
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	tempdir, err := ioutil.TempDir("", "sizebench")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.RemoveAll(tempdir) }()

	targets := [][2]string{
		{"base-fmt", "base"},
		{"basz-fmt", "base-conserv"},
		{"pico-lib", "pico-lib"},
		{"pico-msg", "pico-msg"},
		{"pico-par", "pico-partial"},
		{"pico-sml", "pico-small"},
		{"picz-lib", "pico-conserv-lib"},
		{"picz-msg", "pico-conserv-msg"},
		{"picz-par", "pico-conserv-partial"},
		{"picz-sml", "pico-conserv-small"},
		{"prot-lib", "proto-lib"},
		{"prot-msg", "proto-msg"},
		{"prot-par", "proto-partial"},
		{"prot-sml", "proto-small"},
	}

	limiter := make(chan struct{}, 4)

	type idx struct {
		osarch OsArch
		target string
	}

	var mu sync.Mutex
	results := map[idx]result{}

	var g errgroup.Group
	for _, osarch := range osarches {
		for _, target := range targets {
			osarch, target := osarch, target
			g.Go(func() error {
				limiter <- struct{}{}
				defer func() { <-limiter }()

				result := analyze(tempdir, osarch, target[0])
				mu.Lock()
				results[idx{osarch, target[1]}] = result
				mu.Unlock()
				return nil
			})
		}
	}
	_ = g.Wait()

	failed := false
	for _, osarch := range osarches {
		for _, target := range targets {
			r := results[idx{osarch, target[1]}]
			if r.fail != nil {
				fmt.Fprintln(os.Stderr, r.arch, target[0], r.fail, string(r.out))
				failed = true
			}
		}
	}
	if failed {
		return fmt.Errorf("failed to compile")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	defer func() { _ = w.Flush() }()

	for _, osarch := range osarches {
		fmt.Fprintf(w, "goos: %v\n", osarch.Os)
		fmt.Fprintf(w, "goarch: %v\n", osarch.Arch)

		size := func(target string) int64 { return results[idx{osarch, target}].size }

		baseSize := size("base")
		picoMsgSize := size("pico-msg") - size("pico-lib")
		picoLibSize := size("pico-lib") - picoMsgSize - baseSize
		picoUnused := size("pico-partial") - picoMsgSize - picoLibSize - baseSize
		picoSmall := size("pico-small") - baseSize

		fmt.Fprintf(w, "BenchmarkPico/Library\t1\t%12v\n", memorySize(picoLibSize))
		fmt.Fprintf(w, "BenchmarkPico/Message\t1\t%12v\n", memorySize(picoMsgSize))
		fmt.Fprintf(w, "BenchmarkPico/Unused\t1\t%12v\n", memorySize(picoUnused))
		fmt.Fprintf(w, "BenchmarkPico/LibrarySmall\t1\t%12v\n", memorySize(picoSmall))

		baseConservSize := size("base-conserv")
		picoConservMsgSize := size("pico-conserv-msg") - size("pico-conserv-lib")
		picoConservLibSize := size("pico-conserv-lib") - picoConservMsgSize - baseConservSize
		picoConservUnused := size("pico-conserv-partial") - picoConservMsgSize - picoConservLibSize - baseConservSize
		picoConservSmall := size("pico-conserv-small") - baseConservSize

		fmt.Fprintf(w, "BenchmarkPicoConserv/Library\t1\t%12v\n", memorySize(picoConservLibSize))
		fmt.Fprintf(w, "BenchmarkPicoConserv/Message\t1\t%12v\n", memorySize(picoConservMsgSize))
		fmt.Fprintf(w, "BenchmarkPicoConserv/Unused\t1\t%12v\n", memorySize(picoConservUnused))
		fmt.Fprintf(w, "BenchmarkPicoConserv/LibrarySmall\t1\t%12v\n", memorySize(picoConservSmall))

		protoMsgSize := size("proto-msg") - size("proto-lib")
		protoLibSize := size("proto-lib") - protoMsgSize - baseConservSize
		protoUnused := size("proto-partial") - protoMsgSize - protoLibSize - baseConservSize
		protoSmall := size("proto-small") - baseConservSize

		fmt.Fprintf(w, "BenchmarkProto/Library\t1\t%12v\n", memorySize(protoLibSize))
		fmt.Fprintf(w, "BenchmarkProto/Message\t1\t%12v\n", memorySize(protoMsgSize))
		fmt.Fprintf(w, "BenchmarkProto/Unused\t1\t%12v\n", memorySize(protoUnused))
		fmt.Fprintf(w, "BenchmarkProto/LibrarySmall\t1\t%12v\n", memorySize(protoSmall))

		fmt.Fprintf(w, "\n")
	}

	return nil
}

type result struct {
	arch OsArch
	size int64

	out  []byte
	fail error
}

func analyze(tempdir string, osarch OsArch, target string) (r result) {
	r.arch = osarch

	exe := filepath.Join(tempdir, osarch.Os+"_"+osarch.Arch+"_"+target+".exe")
	cmd := exec.Command("go", "build", "-o", exe, "-trimpath", "./"+target)
	cmd.Env = append(os.Environ(),
		"GOOS="+osarch.Os,
		"GOARCH="+osarch.Arch,
	)

	r.out, r.fail = cmd.CombinedOutput()
	if r.fail != nil {
		return
	}

	info, err := os.Stat(exe)
	if err != nil {
		r.fail = err
		return
	}

	r.size = info.Size()
	return
}

func memorySize(size int64) string {
	return fmt.Sprintf("%.3f KB", float64(size)/(1e3))
}

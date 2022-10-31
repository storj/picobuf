// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
)

// Note, the names of packages has been chosen to be exactly the same to reduce the difference
// in package names that end up in the binary.

//go:generate protoc --pico_out=paths=source_relative:./pico/one --pico_opt=Mmsg-one.proto=storj.io/picobuf/internal/sizebench/pico/one msg-one.proto
//go:generate protoc --pico_out=paths=source_relative:./pico/two --pico_opt=Mmsg-two.proto=storj.io/picobuf/internal/sizebench/pico/two msg-two.proto
//go:generate protoc --pico_out=paths=source_relative:./pico/sml --pico_opt=Mmsg-sml.proto=storj.io/picobuf/internal/sizebench/pico/sml msg-sml.proto
//go:generate protoc --go_out=paths=source_relative:./prot/one --go_opt=Mmsg-one.proto=storj.io/picobuf/internal/sizebench/prot/one msg-one.proto
//go:generate protoc --go_out=paths=source_relative:./prot/two --go_opt=Mmsg-two.proto=storj.io/picobuf/internal/sizebench/prot/two msg-two.proto
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
	// {"js", "wasm"},
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	tempdir, err := os.MkdirTemp("", "sizebench")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = os.RemoveAll(tempdir) }()

	exeTargets := []string{
		"base-fmt",
		"basz-fmt",
		"pico-one",
		"pico-two",
		"pico-sml",
		"picz-one",
		"picz-two",
		"picz-sml",
		"prot-one",
		"prot-two",
		"prot-sml",
	}

	pkgTargets := []string{
		"pico/one",
		"pico/two",
	}

	limiter := make(chan struct{}, 4)

	type idx struct {
		osarch OsArch
		target string
	}

	var mu sync.Mutex
	exeResults := map[idx]exeResult{}
	pkgResults := map[idx]pkgResult{}
	failures := []error{}

	var g syncGroup
	for _, osarch := range osarches {
		osarch := osarch
		for _, target := range exeTargets {
			target := target
			g.Go(func() {
				limiter <- struct{}{}
				defer func() { <-limiter }()

				result, err := analyzeExe(tempdir, osarch, target)
				mu.Lock()
				exeResults[idx{osarch, target}] = result
				if err != nil {
					failures = append(failures, fmt.Errorf("%s %s: %w", osarch, target, err))
				}
				mu.Unlock()
			})
		}

		for _, target := range pkgTargets {
			target := target
			g.Go(func() {
				limiter <- struct{}{}
				defer func() { <-limiter }()

				result, err := analyzePkg(tempdir, osarch, target)
				mu.Lock()
				pkgResults[idx{osarch, target}] = result
				if err != nil {
					failures = append(failures, fmt.Errorf("%s %s: %w", osarch, target, err))
				}
				mu.Unlock()
			})
		}
	}
	g.Wait()

	if len(failures) > 0 {
		for _, failure := range failures {
			fmt.Fprintln(os.Stderr, failure)
		}
		return fmt.Errorf("compilation failed")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	defer func() { _ = w.Flush() }()

	for _, osarch := range osarches {
		fmt.Fprintf(w, "goos: %v\n", osarch.Os)
		fmt.Fprintf(w, "goarch: %v\n", osarch.Arch)

		exeSize := func(target string) string {
			size := exeResults[idx{osarch, target}].size
			return fmt.Sprintf("%12v", memorySize(size))
		}
		exeDelta := func(target, base string) string {
			asize := exeResults[idx{osarch, target}].size
			bsize := exeResults[idx{osarch, base}].size
			return fmt.Sprintf("%12v\t%12v-Δ", memorySize(asize), memorySize(asize-bsize))
		}

		pkgSize := func(target string) string {
			size := pkgResults[idx{osarch, target}].size
			return fmt.Sprintf("%12v", memorySize(size))
		}
		pkgSizeMatch := func(target, match string) string {
			r := pkgResults[idx{osarch, target}]
			return fmt.Sprintf("%12v", memorySize(r.symbols.sum(match)))
		}
		pkgDelta := func(target, base string) string {
			asize := pkgResults[idx{osarch, target}].size
			bsize := pkgResults[idx{osarch, base}].size
			return fmt.Sprintf("%12v\t%12v-Δ", memorySize(asize), memorySize(asize-bsize))
		}

		fmt.Fprintf(w, "BenchmarkExeBase\t1\t%v\n", exeSize("base-fmt"))
		fmt.Fprintf(w, "BenchmarkExeBaseConserv\t1\t%v\n", exeSize("basz-fmt"))

		fmt.Fprintf(w, "BenchmarkExePico/One\t1\t%v\n", exeDelta("pico-one", "base-fmt"))
		fmt.Fprintf(w, "BenchmarkExePico/Two\t1\t%v\n", exeDelta("pico-two", "base-fmt"))
		fmt.Fprintf(w, "BenchmarkExePico/Sml\t1\t%v\n", exeDelta("pico-sml", "base-fmt"))

		fmt.Fprintf(w, "BenchmarkPkgPico/One\t1\t%v\n", pkgSize("pico/one"))
		fmt.Fprintf(w, "BenchmarkPkgPico/One/Encode\t1\t%v\t%v-f\n", pkgSizeMatch("pico/one", "one.(*Types).Encode"), pkgSizeMatch("pico/one", "one.(*Types).Encode.func"))
		fmt.Fprintf(w, "BenchmarkPkgPico/One/Decode\t1\t%v\t%v-f\n", pkgSizeMatch("pico/one", "one.(*Types).Decode"), pkgSizeMatch("pico/one", "one.(*Types).Decode.func"))
		fmt.Fprintf(w, "BenchmarkPkgPico/Two\t1\t%v\n", pkgDelta("pico/two", "pico/one"))

		fmt.Fprintf(w, "BenchmarkExePicoConserv/One\t1\t%v\n", exeDelta("picz-one", "basz-fmt"))
		fmt.Fprintf(w, "BenchmarkExePicoConserv/Two\t1\t%v\n", exeDelta("picz-two", "basz-fmt"))
		fmt.Fprintf(w, "BenchmarkExePicoConserv/Sml\t1\t%v\n", exeDelta("picz-sml", "basz-fmt"))

		fmt.Fprintf(w, "BenchmarkExeProto/One\t1\t%v\n", exeDelta("prot-one", "basz-fmt"))
		fmt.Fprintf(w, "BenchmarkExeProto/Two\t1\t%v\n", exeDelta("prot-two", "basz-fmt"))
		fmt.Fprintf(w, "BenchmarkExeProto/Sml\t1\t%v\n", exeDelta("prot-sml", "basz-fmt"))

		fmt.Fprintf(w, "\n")
	}

	return nil
}

type exeResult struct {
	arch OsArch
	size int64
}

func analyzeExe(tempdir string, osarch OsArch, target string) (r exeResult, err error) {
	r.arch = osarch

	exe := filepath.Join(tempdir, osarch.Os+"_"+osarch.Arch+"_"+filepath.Base(target)+".exe")
	cmd := exec.Command("go", "build", "-o", exe, "-trimpath", "./"+target)
	cmd.Env = append(os.Environ(),
		"GOOS="+osarch.Os,
		"GOARCH="+osarch.Arch,
		"CGO_ENABLED=0",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return r, fmt.Errorf("compile failed: %w (%s)", err, string(out))
	}

	info, err := os.Stat(exe)
	if err != nil {
		return r, err
	}

	r.size = info.Size()
	return r, nil
}

type pkgResult struct {
	arch    OsArch
	size    int64
	symbols symbols
}

func analyzePkg(tempdir string, osarch OsArch, target string) (r pkgResult, err error) {
	r.arch = osarch

	exe := filepath.Join(tempdir, osarch.Os+"_"+osarch.Arch+"_"+filepath.Base(target)+".o")
	cmd := exec.Command("go", "build", "-o", exe, "-trimpath", "./"+target)
	cmd.Env = append(os.Environ(),
		"GOOS="+osarch.Os,
		"GOARCH="+osarch.Arch,
		"CGO_ENABLED=0",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return r, fmt.Errorf("compile failed: %w (%s)", err, string(out))
	}

	info, err := os.Stat(exe)
	if err != nil {
		return r, err
	}

	r.size = info.Size()

	r.symbols, err = analyzeSymbols(exe)
	if err != nil {
		return r, err
	}

	return r, nil
}

func memorySize(size int64) string {
	return fmt.Sprintf("%.3f KB", float64(size)/(1e3))
}

type symbols []symbol

func (syms symbols) sum(match string) int64 {
	size := int64(0)
	for _, sym := range syms {
		if strings.Contains(sym.name, match) {
			size += sym.size
		}
	}
	return size
}

type symbol struct {
	name string
	size int64
}

func analyzeSymbols(bin string) (symbols, error) {
	cmd := exec.Command("go", "tool", "nm", "-size", bin)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var syms symbols
	for _, line := range strings.Split(string(out), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		sym, err := parseSymbol(line)
		if err != nil {
			return syms, err
		}
		syms = append(syms, sym)
	}
	return syms, nil
}

func parseSymbol(line string) (sym symbol, err error) {
	tokens := strings.Fields(line[8:])
	if len(tokens) < 2 {
		return sym, fmt.Errorf("invalid sym text: %q", line)
	}

	if size := strings.TrimSpace(tokens[0]); size != "" {
		sym.size, err = strconv.ParseInt(size, 10, 64)
		if err != nil {
			return sym, fmt.Errorf("invalid size: %q", size)
		}
	}

	sym.name = strings.Join(tokens[2:], " ")
	return sym, nil
}

type syncGroup struct{ sync.WaitGroup }

func (g *syncGroup) Go(fn func()) {
	g.Add(1)
	go func() {
		defer g.Done()
		fn()
	}()
}

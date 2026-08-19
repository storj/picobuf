// Copyright (C) 2026 Storj Labs, Inc.
// See LICENSE for copying information.

//go:build !race

package picobuf_test

// raceEnabled reports whether the race detector is active. Race
// instrumentation adds allocations, so allocation-count assertions skip
// themselves when it is on.
const raceEnabled = false

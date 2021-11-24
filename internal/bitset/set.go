// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package bitset

// Small implements a bitset that gives preferential treatment to small values.
type Small struct {
	low  uint64
	rest []uint64
	// TODO: prevent issues with large field numbers.
}

// Set sets the specified bit and returns whether it was already set.
func (set *Small) Set(x int32) bool {
	if x < 0 {
		return false
	}

	if x < 64 {
		wasSet := set.low&(1<<byte(x)) != 0
		set.low |= 1 << byte(x)
		return wasSet
	}

	x -= 64
	w := uint(x)
	bucket, bit := w/64, w%64
	if uint(len(set.rest)) < bucket {
		set.rest = append(set.rest, make([]uint64, int(bucket)-len(set.rest))...)
	}
	wasSet := set.rest[bucket]&(1<<byte(bit)) != 0
	set.rest[bucket] |= 1 << byte(bit)
	return wasSet
}

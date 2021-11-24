// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

func encodeZigZag32(v int32) uint32 {
	return (uint32(v) << 1) ^ (uint32(v) >> 31)
}

func decodeZigZag32(v uint32) int32 {
	return int32(v>>1) ^ int32(v)<<31>>31
}

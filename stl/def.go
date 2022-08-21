package stl

import "math"

const BIGINT = 0x3f3f3f3f
const MAXINT = math.MaxInt32

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxForInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

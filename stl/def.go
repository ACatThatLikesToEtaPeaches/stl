package stl

const INF = 0x3f3f3f3f


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
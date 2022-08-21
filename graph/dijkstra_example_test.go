package graph

import "math"

// There are some example problems using the Dijkstra algorithm.

// 1. LeetCode743. 网络延迟时间
func networkDelayTime(time [][]int, n int, start int) int {
	// 预处理：将题目给的节点标记范围为 [1,n] 转化为 [0,n-1]
	for i := range time {
		time[i][0], time[i][1] = time[i][0]-1, time[i][1]-1
	}
	distTo := Dijkstra(time, n, start-1)

	maxDistance := 0
	for _, d := range distTo {
		if d == math.MaxInt32 {
			return -1
		} else if d > maxDistance {
			maxDistance = d
		}
	}
	return maxDistance
}

// 2. ZOJ1298
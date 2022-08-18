package graph

import (
	"container/heap"
	"datastruct_algorithm/stl"
)

// Dijkstra 单源最短路径算法, 作用于有向有权图
//  -- 输入一幅图和一个起点 start，计算 start 到其他节点的最短距离
// 	假设 有 n 个网络节点，标记为 0 到 n-1
func Dijkstra(weight [][]int, n int, start int) []int {
	type edge struct{ to, weight int }
	// 构造图, from -> [{toId1, weight1}, {toId2, weight2}]
	//  邻接表存储图结构，同时存储权重信息
	graph := make([][]edge, n) // // 节点编号是从 0 开始的，需要一个大小为 n 的邻接表
	for _, w := range weight {
		from, to, value := w[0], w[1], w[2]
		graph[from] = append(graph[from], edge{to: to, weight: value})
	}

	distTo := make([]int, n)
	// base case: init start -> start = 0, others = INF,
	for i := range distTo {
		distTo[i] = stl.INF
	}
	distTo[start] = 0
	hp := &heapImpl{}
	hp.Push(item{id:start, dist: 0})

	for hp.Len() > 0 {
		curItem := heap.Pop(hp).(item)
		curId, curDist := curItem.id, curItem.dist
		if curDist > distTo[curId] { continue }

		for _, adjItem := range graph[curId] {
			adjID, adjDist := adjItem.to, adjItem.weight
			if curDist + adjDist < distTo[adjID] {
				distTo[adjID] = curDist + adjDist
				heap.Push(hp, item{id: adjID, dist: curDist + adjDist})
			}
		}
	}

	return distTo
}

type item struct {
	id   int // current point id in graph
	dist int // distance from start point to cur point
}
type heapImpl []item

func (h heapImpl) Len() int              { return len(h) }
func (h heapImpl) Less(i, j int) bool    { return h[i].dist < h[j].dist }
func (h heapImpl) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *heapImpl) Push(v interface{})   { *h = append(*h, v.(item)) }
func (h *heapImpl) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

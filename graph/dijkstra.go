package graph

import (
	"container/heap"
	"math"
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
	// base case: init start -> start = 0, others = MAXINT,
	for i := 0; i < n; i++ {
		distTo[i] = math.MaxInt32
	}
	distTo[start] = 0
	pq := &priorityQueueImpl{}
	pq.Push(item{id: start, dist: 0})

	for !pq.Empty() {
		curItem := pq.Pop().(item)
		curId, curDist := curItem.id, curItem.dist
		//if curId == end { // if end is specified
		//	return curDist
		//}
		if curDist > distTo[curId] {
			continue
		}

		for _, adjItem := range graph[curId] {
			adjID, adjDist := adjItem.to, adjItem.weight
			if curDist+adjDist < distTo[adjID] {
				distTo[adjID] = curDist + adjDist
				pq.Push(item{id: adjID, dist: curDist + adjDist})
			}
		}
	}

	return distTo
}

func BiDijkstra(weight [][]int, n int, start, end int) int {
	// weight: {{from, to, weight}, {from, to, weight}}
	type edge struct{ to, weight int }
	graph := make([][]edge, n) // // 节点编号是从 0 开始的，需要一个大小为 n 的邻接表
	for _, w := range weight {
		from, to, value := w[0], w[1], w[2]
		graph[from] = append(graph[from], edge{to: to, weight: value})
	}
	var reverse = func() [][]edge {
		newGraph := make([][]edge, n)
		for i, adj := range graph {
			for _, e := range adj {
				newGraph[e.to] = append(newGraph[e.to], edge{to: i, weight: e.weight})
			}
		}
		return newGraph
	}
	rgraph := reverse()

	distF := make([]int, n)
	distR := make([]int, n)
	for i := 0; i < n; i++ {
		distF[start], distR[end] = math.MaxInt32, math.MaxInt32
	}
	distF[start], distR[end] = 0, 0

	fq, rq := &priorityQueueImpl{}, &priorityQueueImpl{}
	fvisited, rvisited := make([]bool, n), make([]bool, n)
	fq.Push(item{id: start, dist: 0})
	fvisited[start] = true
	rq.Push(item{id: end, dist: 0})
	rvisited[end] = true

	for !fq.Empty() && !rq.Empty() {
		fitem := fq.Pop().(item)
		fn, fd := fitem.id, fitem.dist
		if rvisited[fn] {
			fqmin, rqmin := 0, 0
			if !fq.Empty() { fqmin = fq.Top().(item).dist }
			if !rq.Empty() { rqmin = rq.Top().(item).dist }
			if fd+distR[fn] <= fqmin+rqmin {
				return fd + distR[fn]
			}
		}
		for _, adjItem := range graph[fn] {
			adjId, adjDist := adjItem.to, adjItem.weight
			if fvisited[adjId] {
				if fd+adjDist < distF[adjId] {
					distF[adjId] = fd + adjDist
					fq.Push(item{id: adjId, dist: fd + adjDist})
				}
			} else {
				distF[adjId] = fd + adjDist
				fq.Push(item{id: adjId, dist: fd + adjDist})
				fvisited[adjId] = true
			}
		}

		ritem := rq.Pop().(item)
		rn, rd := ritem.id, ritem.dist
		if fvisited[rn] {
			fqmin, rqmin := 0, 0
			if !fq.Empty() { fqmin = fq.Top().(item).dist }
			if !rq.Empty() { rqmin = rq.Top().(item).dist }
			if rd+distF[rn] <= fqmin+rqmin {
				return distF[rn] + rd
			}
		}
		for _, adjItem := range rgraph[rn] {
			adjId, adjDist := adjItem.to, adjItem.weight
			if rvisited[adjId] {
				if rd+adjDist < distR[adjId] {
					distR[adjId] = rd + adjDist
					rq.Push(item{id: adjId, dist: rd + adjDist})
				}
			} else {
				distR[adjId] = rd + adjDist
				rq.Push(item{id: adjId, dist: rd + adjDist})
				rvisited[adjId] = true
			}
		}

	}
	return math.MaxInt32
}

// priorityQueueImpl 优先级队列
type priorityQueueImpl struct{ data heapImpl }

func (pq *priorityQueueImpl) Push(item interface{}) { heap.Push(&pq.data, item) }
func (pq *priorityQueueImpl) Pop() interface{}      { return heap.Pop(&pq.data) }
func (pq *priorityQueueImpl) Empty() bool           { return len(pq.data) == 0 }
func (pq *priorityQueueImpl) Size() int             { return len(pq.data) }
func (pq *priorityQueueImpl) Top() interface{}      { return pq.data[0] }
func (pq *priorityQueueImpl) Remove(it item)        { heap.Remove(&pq.data, it.dist) }

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
func (h *heapImpl) push(vdPair item)     { heap.Push(h, vdPair) }
func (h *heapImpl) pop() item            { return h.Pop().(item) }

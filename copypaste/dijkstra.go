package copypaste

import "container/heap"

type nb struct{ to, wt int }

func dijkstra(g [][]nb, st int) []int {
	const inf int = 1e18
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	q := hp{{st, 0}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if dist[v] < p.d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				q.push(vdPair{w, newD})
			}
		}
	}
	return dist
}

type vdPair struct{ v, d int }
type hp []vdPair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v vdPair)        { heap.Push(h, v) }
func (h *hp) pop() vdPair          { return heap.Pop(h).(vdPair) }

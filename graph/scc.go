package graph

import "datastruct_algorithm/stl"

// Strong Connected Component 强连通分量
// 1. Kosaraju 算法求有向图强连通分量 参考：https://www.cnblogs.com/nullzx/p/6437926.html
// 2. DFS 求无向图连通分量

func KosarajuCalcSCC(bias [][]int, n int) int {
	// bias means all edges: []{{from, to},{from, to}}
	graph := make([][]int, n) // 临接表表示图
	for _, b := range bias {
		from, to := b[0], b[1]
		graph[from] = append(graph[from], to)
	}

	sccID := make([]int, n) // 用于标记每个顶点属于哪个连通分量

	g := &Graph{adjTable: graph, sccID: sccID, nodeCount: n}
	// Step1: 对原图取反，从任意一个顶点开始对反向图进行逆后序DFS遍历
	stk := g.reverse().getOneReversePostOrderStack()

	// Step2: 按照逆后续遍历中栈中的顶点出栈顺序，对原图进行DFS遍历，一次DFS遍历中访问的所有顶点都属于同一强连通分量。
	visited := make([]bool, n)
	for !stk.Empty() {
		cur := stk.Top().(int)
		stk.Pop()
		if !visited[cur] {
			g.dfs2(cur, visited)
			g.sccCount++
		}
	}

	return g.sccCount
}

type Graph struct {
	adjTable  [][]int
	sccID     []int
	nodeCount int
	sccCount  int
}

func (g *Graph) reverse() *Graph {
	newAdjTable := make([][]int, g.nodeCount)
	for i, adj := range g.adjTable {
		for _, to := range adj {
			newAdjTable[to] = append(newAdjTable[to], i)
		}
	}
	return &Graph{adjTable: newAdjTable, sccID: g.sccID, nodeCount: g.nodeCount, sccCount: g.sccCount}
}
func (g *Graph) getOneReversePostOrderStack() stl.Stack { // 求逆后序遍历栈
	stk := stl.NewStack()
	visited := make([]bool, g.nodeCount)
	for i := 0; i < g.nodeCount; i++ { // 这里默认从0开始，从任意节点开始DFS求逆后序栈都是合法的
		g.dfs1(i, visited, stk)
	}
	return stk
}

func (g *Graph) dfs1(i int, visited []bool, stk stl.Stack) {
	if !visited[i] {
		visited[i] = true
		for _, nxt := range g.adjTable[i] {
			g.dfs1(nxt, visited, stk)
		}
		stk.Push(i) // post order result
	}
}

func (g *Graph) dfs2(i int, visited []bool) {
	if !visited[i] {
		visited[i] = true
		g.sccID[i] = g.sccCount
		for _, nxt := range g.adjTable[i] {
			g.dfs2(nxt, visited)
		}
	}
}


// DFS 求无向图连通分量
// DFS 应用：求连通分量以及每个点所属的连通分量 (Connected Component, CC)
// 注意: 这里sccIDs 的值从 1 开始，用0来初始化并表示还未访问

func DFSCalcSCC(g [][]int, n int) (comps [][]int, sccIDs []int) {
	sccIDs = make([]int, n)
	idCnt := 0 // 也可以去掉，用 len(comps)+1 代替
	var comp []int
	var f func(int)
	f = func(v int) {
		sccIDs[v] = idCnt
		comp = append(comp, v)
		for _, w := range g[v] {
			if sccIDs[w] == 0 { //初始化为0，表示未被访问
				f(w)
			}
		}
	}
	for i, id := range sccIDs {
		if id == 0 {
			idCnt++
			comp = []int{}
			f(i)
			comps = append(comps, comp)
		}
	}
	return
}

/* Ignore me: 参考的OI Wiki的实现：https://oi-wiki.org/graph/scc/
// step1: 对原图G进行深度优先遍历，记录每个节点的离开时间num[i]
// step2: 选择具有最晚离开时间的顶点，对反图GT进行遍历，删除能够遍历到的顶点，这些顶点构成一个强连通分量
//  (这里利用了一个事实，即转置图（同图中的每边的方向相反）具有和原图完全一样的强连通分量。)
func KosarajuCalcSCC(bias [][]int, n int) int {
	// bias means all edges: []{{from, to},{from, to}}
	graph := make([][]int, n) // 临接表表示图
	for _, b := range bias {
		from, to := b[0], b[1]
		graph[from] = append(graph[from], to)
	}

	sccID := make([]int, n) // 用于标记每个顶点属于哪个连通分量

	g := &Graph{adjTable: graph, sccID: sccID, nodeCount: n}

	// 第一次DFS：按离开时间保存一个后序遍历栈
	stk := stl.NewStack()
	visited := make([]bool, g.nodeCount)
	for i := 0; i < g.nodeCount; i++ {
		g.dfs1(i, visited, stk)
	}

	// 第二次DFS：按后序遍历出栈顺序，对反图进行遍历。每一轮遍历的结果都属于同一个sccID
	rg := g.reverse()
	visited = make([]bool, n)
	for !stk.Empty() {
		cur := stk.Top().(int)
		stk.Pop()
		if !visited[cur] {
			rg.dfs2(cur, visited)
			rg.sccCount++
		}
	}

	return rg.sccCount
}

*/

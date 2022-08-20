package graph

import "datastruct_algorithm/stl"

// Strong Connected Component 强连通分量
//  1. Tarjan 算法求强连通分量 参考自：https://oi-wiki.org/graph/scc/
//	2. Kosaraju 算法求有向图强连通分量 参考自：https://www.cnblogs.com/nullzx/p/6437926.html
//   注：它们的时间复杂度都是O(n+m)常数级的, 不过常数大小 Tarjan < Kosaraju < Garbow（在 AtCoder 上的测试显示，5e5 的数据下 Tarjan 比 Kosaraju 快了约 100ms）

func TarjanCalcSCC(bias [][]int, n int) (comps [][]int, sccid []int) {
	//	dfn[v] ：深度优先搜索遍历时结点 v 被搜索的次序。
	//	low[v]：在 v 的子树中能够回溯到的最早的已经在栈中的结点。设以 v 为根的子树为 SubTree[v]。
	//		low[v] 定义为以下结点的 dfn 的最小值： SubTree[v] 中的结点；从 SubTree[v] 通过一条不在搜索树上的边能到达的结点。
	g := make([][]int, n) // 临接表表示图
	for _, b := range bias {
		from, to := b[0], b[1]
		g[from] = append(g[from], to)
	}

	dfn := make([]int, n) // 值从 1 开始，初始化为0表示未访问到
	low := make([]int, n)
	dfsClock := 0
	stk := []int{}
	inStk := make([]bool, len(g))

	var tarjan func(int)
	tarjan = func(v int) {
		dfsClock++
		dfn[v] = dfsClock
		low[v] = dfsClock
		stk = append(stk, v) // stk.Push
		inStk[v] = true
		for _, u := range g[v] {
			if dfn[u] == 0 { // case1: 邻接点 u 未被访问：继续对 u 进行深度搜索
				tarjan(u)
				low[v] = min(low[v], low[u])
			} else if inStk[u] { // case2: 邻接点 u 被访问过，且还在当前栈中：根据 low 值的定义，用 dfn[u] 来更新 low[v]
				low[v] = min(low[v], dfn[u])
			} // case3: 邻接点 u 被访问过，且已经不在栈中：说明 v 已搜索完毕，其所在连通分量已被处理，所以不用对其做操作。
		}
		if dfn[v] == low[v] { // 当 v 是连通量的subRoot，把栈里面 v 前面的元素全部取出 作为一个 comp
			comp := []int{}
			for stk[len(stk)-1] != v {
				stktop := stk[len(stk)-1]
				stk = stk[:len(stk)-1] // stk.Pop
				inStk[stktop] = false

				comp = append(comp, stktop)
			}
			comps = append(comps, comp)
		}
		return
	}

	for v, timestamp := range dfn {
		if timestamp == 0 {
			tarjan(v)
		}
	}

	// 由于每个强连通分量都是在它的所有后继强连通分量被求出之后求得的
	// 上面得到的 comps 是拓扑序的逆序
	for i, n := 0, len(comps); i < n/2; i++ {
		comps[i], comps[n-1-i] = comps[n-1-i], comps[i]
	}

	sccid = make([]int, n)
	for i, cp := range comps {
		for _, v := range cp {
			sccid[v] = i
		}
	}

	return comps, sccid
}


func KosarajuCalcSCC(bias [][]int, n int) (comps [][]int, sccIDs []int) {
	// bias means all edges: []{{from, to},{from, to}}
	graph := make([][]int, n) // 临接表表示图
	for _, b := range bias {
		from, to := b[0], b[1]
		graph[from] = append(graph[from], to)
	}
	sccID := make([]int, n) // 用于标记每个顶点属于哪个连通分量,从0开始

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

	comps = make([][]int, g.sccCount)
	for i, sid := range g.sccID {
		comps[sid] = append(comps[sid], i)
	}
	// 由于每个强连通分量都是在它的所有后继强连通分量被求出之后求得的
	// 上面得到的 comps 是拓扑序的逆序
	for i, n := 0, len(comps); i < n/2; i++ {
		comps[i], comps[n-1-i] = comps[n-1-i], comps[i]
	}

	return comps, g.sccID
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

func min(a,b int) int {
	if a < b {return a}
	return b
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

package graph

import "datastruct_algorithm/stl"

// Strong Connected Component 强连通分量
// Kosaraju 算法求强连通分量 参考：https://www.cnblogs.com/nullzx/p/6437926.html

func ConnectedComponents(bias [][]int, n int) int {
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
			g.dfs(cur, visited)
			g.sccCount++
		}
	}

	return g.sccCount
}

func IsInSameConnectedComponents(a int, b int) bool {
	// todo: this can be easy implement by using sccID in Graph after scc run
	return false
}
func GetAllConnected(v int) []int {
	// todo: this can be easy implement by using sccID in Graph after scc run
	return nil
}

type Graph struct {
	adjTable  [][]int
	sccID     []int
	nodeCount int
	sccCount  int
}

func (g *Graph) dfs(i int, visited []bool) {
	if !visited[i] {
		visited[i] = true
		g.sccID[i] = g.sccCount
		for _, nxt := range g.adjTable[i] {
			g.dfs(nxt, visited)
		}
	}
}
func (g *Graph) reverse() *Graph { // 反向图
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
		g.reversePostOrderDFSHelper(i, visited, stk)
	}
	return stk
}
func (g *Graph) reversePostOrderDFSHelper(i int, visited []bool, stk stl.Stack) {
	if !visited[i] {
		visited[i] = true
		for _, nxt := range g.adjTable[i] {
			g.reversePostOrderDFSHelper(nxt, visited, stk)
		}
		stk.Push(i) // post order
	}
}

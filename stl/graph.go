package stl


// "图" 结构，临接表表示
// note: in C++, 通常用list存储临接表，like  vector<list<edge>>& graph, in go, 可以直接用[][]Edge

type Graph [][]Edge
type Edge struct {
	from, to, weight int
}


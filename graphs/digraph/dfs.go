package digraph

import "github.com/zh1014/algorithm/queue"

// Digraph does not support self-loop, but assume the source vertex can reach itself
type DFS struct {
	g      Digraph
	marked []bool
}

func NewDFS(g Digraph, src int) *DFS {
	dfs := &DFS{
		g:      g,
		marked: make([]bool, g.NumV()),
	}
	dfs.doDFS(src)
	return dfs
}

func (dfs *DFS) doDFS(v int) {
	dfs.marked[v] = true
	adj := dfs.g.Adjacent(v)
	for _, w := range adj {
		if !dfs.marked[w] {
			dfs.doDFS(w)
		}
	}
}

func (dfs *DFS) CanReach(v int) bool {
	return dfs.marked[v]
}

func (dfs *DFS) ReachableVertices() *queue.IntQ {
	q := queue.NewIntQ()
	for i, b := range dfs.marked {
		if b {
			q.PushBack(i)
		}
	}
	return q
}

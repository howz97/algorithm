package digraph

import "github.com/howz97/algorithm/queue"

type DFS struct {
	g      Digraph
	marked []bool
}

func (g Digraph) DFS(src int) *DFS {
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

func (dfs *DFS) CanReach(dst int) bool {
	return dfs.marked[dst]
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

func (dfs *DFS) Range(fn func(v int) bool) {
	for v, b := range dfs.marked {
		if !b {
			continue
		}
		if !fn(v) {
			break
		}
	}
}

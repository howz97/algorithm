package digraph

import "github.com/howz97/algorithm/queue"

type TransitiveClosure struct {
	g   Digraph
	all []*DFS
}

func (g Digraph) TransitiveClosure() *TransitiveClosure {
	tc := &TransitiveClosure{
		g:   g,
		all: make([]*DFS, g.NumV()),
	}
	for i := range tc.all {
		// todo optimize: reduce DFS
		tc.all[i] = g.DFS(i)
	}
	return tc
}

func (tc *TransitiveClosure) IsReachable(src, dst int) bool {
	if !tc.g.HasV(src) || !tc.g.HasV(dst) {
		panic(ErrVertexNotExist)
	}
	return tc.all[src].CanReach(dst)
}

func (tc *TransitiveClosure) ReachableVertices(src *queue.IntQ) *queue.IntQ {
	vertices := queue.NewIntQ()
	marked := make([]bool, tc.g.NumV())
	for !src.IsEmpty() {
		rv := tc.all[src.Front()].ReachableVertices()
		for !rv.IsEmpty() {
			marked[rv.Front()] = true
		}
	}
	for i, b := range marked {
		if b {
			vertices.PushBack(i)
		}
	}
	return vertices
}

func (tc *TransitiveClosure) Range(v int, fn func(v int) bool) {
	if v < 0 || v >= len(tc.all) {
		return
	}
	tc.all[v].Range(fn)
}

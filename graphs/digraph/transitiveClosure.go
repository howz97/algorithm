package digraph

import "github.com/zh1014/algorithm/queue"

type TransitiveClosure struct {
	g   Digraph
	all []*DFS
}

func NewTransitiveClosure(g Digraph) *TransitiveClosure {
	tc := &TransitiveClosure{
		g:   g,
		all: make([]*DFS, g.NumV()),
	}
	for i := range tc.all {
		tc.all[i] = NewDFS(g, i)
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

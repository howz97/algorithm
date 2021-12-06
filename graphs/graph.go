package graphs

import (
	"github.com/howz97/algorithm/util"
)

type Graph struct {
	Digraph
}

func NewGraph(size int) *Graph {
	return &Graph{
		Digraph: NewDigraph(size),
	}
}

func (g *Graph) NumEdge() uint {
	return g.Digraph.NumEdge() / 2
}

// AddEdge add edge v1-v2
func (g *Graph) AddEdge(src, dst int) error {
	return g.AddWEdge(src, dst, 1)
}

func (g *Graph) AddWEdge(src, dst int, w float64) error {
	if !g.HasVertical(src) || !g.HasVertical(dst) {
		return ErrVerticalNotExist
	}
	if src == dst {
		return ErrSelfLoop
	}
	g.Digraph[src].Put(util.Integer(dst), w)
	g.Digraph[dst].Put(util.Integer(src), w)
	return nil
}

func (g *Graph) HasCycle() bool {
	marked := make([]bool, g.NumVertical())
	for i, m := range marked {
		if m {
			continue
		}
		if g.detectCycleDFS(i, i, marked) {
			return true
		}
	}
	return false
}

func (g *Graph) detectCycleDFS(last, cur int, marked []bool) bool {
	marked[cur] = true
	found := false
	g.RangeAdj(cur, func(adj int) bool {
		if adj == last { // here is different from digraph
			return true
		}
		if marked[adj] {
			found = true
			return false
		}
		if g.detectCycleDFS(cur, adj, marked) {
			found = true
			return false
		}
		return true
	})
	return found
}

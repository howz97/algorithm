package graph

import (
	"errors"
	"github.com/howz97/algorithm/graphs/digraph"
)

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
)

type Graph struct {
	digraph.Digraph
}

func New(size int) *Graph {
	return &Graph{
		Digraph: digraph.New(size),
	}
}

func (g *Graph) NumEdge() int {
	return g.Digraph.NumEdge() / 2
}

// AddEdge add edge v1-v2
func (g *Graph) AddEdge(v1, v2 int) error {
	if !g.HasVertical(v1) || !g.HasVertical(v2) {
		return ErrVerticalNotExist
	}
	if v1 == v2 {
		return ErrSelfLoop
	}
	g.Digraph[v1].Add(v2)
	g.Digraph[v2].Add(v1)
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

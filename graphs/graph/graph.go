package graph

import (
	"errors"
	"github.com/howz97/algorithm/set"
	"strconv"
)

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
)

type Graph []set.IntSet

func New(numV int) Graph {
	g := make(Graph, numV)
	for i := range g {
		g[i] = make(set.IntSet)
	}
	return g
}

func NewByImport(filename string) Graph {
	// todo
	return nil
}

func (g Graph) NumVertical() int {
	return len(g)
}

func (g Graph) HasVertical(v int) bool {
	return v >= 0 && v < len(g)
}

func (g Graph) NumEdge() int {
	num := 0
	for i := range g {
		num += g[i].Len()
	}
	return num / 2
}

// AddEdge add edge v1-v2
func (g Graph) AddEdge(v1, v2 int) error {
	if !g.HasVertical(v1) || !g.HasVertical(v2) {
		return ErrVerticalNotExist
	}
	if v1 == v2 {
		return ErrSelfLoop
	}
	g[v1].Add(v2)
	g[v2].Add(v1)
	return nil
}

func (g Graph) HasEdge(v1, v2 int) bool {
	if !g.HasVertical(v1) || !g.HasVertical(v2) {
		return false
	}
	return g[v1].Contains(v2)
}

// Adjacent is the adjacent verticals of v
func (g Graph) Adjacent(v int) []int {
	if !g.HasVertical(v) {
		return nil
	}
	return g[v].Traverse()
}

func (g Graph) RangeAdj(v int, fn func(v int) bool) {
	if !g.HasVertical(v) {
		return
	}
	g[v].Range(fn)
}

func (g Graph) String() string {
	out := ""
	for i := range g {
		out += strconv.Itoa(i) + " :"
		adj := g[i].Traverse()
		for j := range adj {
			out += " " + strconv.Itoa(j)
		}
		out += "\n"
	}
	out += "\n"
	return out
}

func (g Graph) HasCycle() bool {
	marked := make([]bool, g.NumVertical())
	for i, b := range marked {
		if b {
			continue
		}
		if g.hasCycleDFS(i, i, marked) {
			return true
		}
	}
	return false
}

func (g Graph) hasCycleDFS(last, cur int, marked []bool) bool {
	if marked[cur] {
		return true
	}
	marked[cur] = true
	hasCycle := false
	g.RangeAdj(cur, func(adj int) bool {
		if adj == last {
			return true
		}
		if g.hasCycleDFS(cur, adj, marked) {
			hasCycle = true
			return false
		}
		return true
	})
	return hasCycle
}

func (g Graph) IsBipartiteGraph() bool {
	marked := make([]bool, g.NumVertical())
	colors := make([]bool, g.NumVertical())
	for i, b := range marked {
		if b {
			continue
		}
		if !g.isBipartiteDFS(i, true, colors, marked) {
			return false
		}
	}
	return true
}

func (g Graph) isBipartiteDFS(cur int, color bool, colors []bool, marked []bool) bool {
	if !marked[cur] {
		marked[cur] = true
		colors[cur] = color
		adjs := g.Adjacent(cur)
		for _, adj := range adjs {
			if !g.isBipartiteDFS(adj, !color, colors, marked) {
				return false
			}
		}
		return true
	} else {
		return colors[cur] == color
	}
}

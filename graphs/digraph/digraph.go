package digraph

import (
	"errors"
	"github.com/zh1014/algorithm/set"
	"strconv"
)

var (
	ErrVerticalNotExist   = errors.New("vertical not exist")
	ErrNotSupportSelfLoop = errors.New("not support self loop")
)

type Digraph []set.Set

func NewDigraph(numV int) Digraph {
	g := make(Digraph, numV)
	for i := range g {
		g[i] = make(set.Set)
	}
	return g
}

func (g Digraph) NumV() int {
	return len(g)
}

func (g Digraph) HasV(v int) bool {
	return v >= 0 && v < g.NumV()
}

func (g Digraph) NumEdge() int {
	nume := 0
	for i := range g {
		nume += g[i].Len()
	}
	return nume
}

func (g Digraph) AddEdge(v1, v2 int) {
	if !g.HasV(v1) || !g.HasV(v2) {
		panic(ErrVerticalNotExist)
	}
	if v1 == v2 {
		panic(ErrNotSupportSelfLoop)
	}
	g[v1].Add(v2)
}

func (g Digraph) HasEdge(v1, v2 int) bool {
	if !g.HasV(v1) || !g.HasV(v2) {
		panic(ErrVerticalNotExist)
	}
	return g[v1].Contains(v2)
}

func (g Digraph) Adjacent(v int) []int {
	if !g.HasV(v) {
		panic(ErrVerticalNotExist)
	}
	return g[v].Traverse()
}

func (g Digraph) String() string {
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

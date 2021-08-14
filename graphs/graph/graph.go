package graph

import (
	"errors"
	"fmt"
	"howz97/algorithm/set"
	"strconv"
)

var (
	errVerticalNotExist   = errors.New("vertical not exist")
	errNotSupportSelfLoop = errors.New("not support self loop")
)

type Graph []set.IntSet

func NewGraph(numV int) Graph {
	g := make(Graph, numV)
	for i := range g {
		g[i] = make(set.IntSet)
	}
	return g
}

func NewByImport(filename string) Graph {
	// TODO
	fmt.Println("NewByImport not support")
	return nil
}

// NumV is the number of verticals in graph
func (g Graph) NumV() int {
	return len(g)
}

func (g Graph) HasV(v int) bool {
	return v >= 0 && v < g.NumV()
}

// NumEdge is the number of edge in graph
func (g Graph) NumEdge() int {
	nume := 0
	for i := range g {
		nume += g[i].Len()
	}
	return nume / 2
}

// AddEdge add edge v1-v2
func (g Graph) AddEdge(v1, v2 int) error {
	if !g.HasV(v1) || !g.HasV(v2) {
		return errVerticalNotExist
	}
	if v1 == v2 {
		return errNotSupportSelfLoop
	}
	g[v1].Add(v2)
	g[v2].Add(v1)
	return nil
}

func (g Graph) HasEdge(v1, v2 int) (bool, error) {
	if !g.HasV(v1) || !g.HasV(v2) {
		return false, errVerticalNotExist
	}
	return g[v1].Contains(v2), nil
}

// Adjacent is the adjacent verticals of v
func (g Graph) Adjacent(v int) ([]int, error) {
	if !g.HasV(v) {
		return nil, errVerticalNotExist
	}
	return g[v].Traverse(), nil
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

func HasCycle(g Graph) bool {
	marked := make([]bool, g.NumV())
	for i, b := range marked {
		if b {
			continue
		}
		if hasCycleDFS(g, i, i, marked) {
			return true
		}
	}
	return false
}

func hasCycleDFS(g Graph, last, cur int, marked []bool) bool {
	if marked[cur] {
		return true
	}
	marked[cur] = true
	adjs, _ := g.Adjacent(cur)
	for _, adj := range adjs {
		if adj == last {
			continue
		}
		if hasCycleDFS(g, cur, adj, marked) {
			return true
		}
	}
	return false
}

func IsBipartiteGraph(g Graph) bool {
	marked := make([]bool, g.NumV())
	colors := make([]bool, g.NumV())
	for i, b := range marked {
		if b {
			continue
		}
		if !isBipartiteDFS(g, i, true, colors, marked) {
			return false
		}
	}
	return true
}

func isBipartiteDFS(g Graph, cur int, color bool, colors []bool, marked []bool) bool {
	if !marked[cur] {
		marked[cur] = true
		colors[cur] = color
		adjs, _ := g.Adjacent(cur)
		for _, adj := range adjs {
			if !isBipartiteDFS(g, adj, !color, colors, marked) {
				return false
			}
		}
		return true
	} else {
		return colors[cur] == color
	}
}

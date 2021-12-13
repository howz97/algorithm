package graphs

import (
	"github.com/howz97/algorithm/util"
)

type Graph struct {
	*Digraph
}

func NewGraph(size uint) *Graph {
	return &Graph{
		Digraph: NewDigraph(size),
	}
}

func (g *Graph) NumEdge() uint {
	return g.Digraph.NumEdge() / 2
}

// AddEdge add edge v1-v2
func (g *Graph) AddEdge(src, dst int) error {
	return g.addWeightedEdge(src, dst, 1)
}

func (g *Graph) addWeightedEdge(src, dst int, w float64) error {
	if !g.HasVertical(src) || !g.HasVertical(dst) {
		return ErrVerticalNotExist
	}
	if src == dst {
		return ErrSelfLoop
	}
	g.Digraph.Edges[src].Put(util.Int(dst), w)
	g.Digraph.Edges[dst].Put(util.Int(src), w)
	return nil
}

func (g *Graph) DelEdge(src, dst int) {
	g.Digraph.DelEdge(src, dst)
	g.Digraph.DelEdge(dst, src)
}

func (g *Graph) IterateWEdge(fn func(int, int, float64) bool) {
	visited := make(map[uint64]struct{})
	g.Digraph.IterateWEdge(func(from int, to int, w float64) bool {
		if _, v := visited[uint64(to)<<32+uint64(from)]; v {
			return true
		}
		visited[uint64(from)<<32+uint64(to)] = struct{}{}
		return fn(from, to, w)
	})
}

func (g *Graph) IterateEdge(fn func(int, int) bool) {
	g.IterateWEdge(func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

func (g *Graph) IterateWEdgeFrom(v int, fn func(int, int, float64) bool) {
	visited := make(map[uint64]struct{})
	g.Digraph.IterateWEdgeFrom(v, func(from int, to int, w float64) bool {
		if _, v := visited[uint64(to)<<32+uint64(from)]; v {
			return true
		}
		visited[uint64(from)<<32+uint64(to)] = struct{}{}
		return fn(from, to, w)
	})
}

func (g *Graph) IterateEdgeFrom(v int, fn func(int, int) bool) {
	g.IterateWEdgeFrom(v, func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
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
	g.IterateAdj(cur, func(adj int) bool {
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

func (g *Graph) TotalWeight() float64 {
	return g.Digraph.TotalWeight() / 2
}

type SubGraphs struct {
	locate    []int   // vertical -> subGraphID
	subGraphs [][]int // subGraphID -> all vertices
}

func (g *Graph) SubGraphs() *SubGraphs {
	tc := &SubGraphs{
		locate: make([]int, g.NumVertical()),
	}
	for i := range tc.locate {
		tc.locate[i] = -1
	}

	subGraphID := 0
	for i, c := range tc.locate {
		if c < 0 {
			dfs := g.ReachableSlice(i)
			for _, v := range dfs {
				tc.locate[v] = subGraphID
			}
			tc.subGraphs = append(tc.subGraphs, dfs)
			subGraphID++
		}
	}
	return tc
}

func (tc *SubGraphs) CanReach(src, dst int) bool {
	if !tc.hasVertical(src) || !tc.hasVertical(dst) {
		return false
	}
	return tc.locate[src] == tc.locate[dst]
}

func (tc *SubGraphs) Range(v int, fn func(v int) bool) {
	if !tc.hasVertical(v) {
		return
	}
	for _, v := range tc.subGraphs[tc.locate[v]] {
		if !fn(v) {
			break
		}
	}
}

func (tc *SubGraphs) hasVertical(v int) bool {
	return v >= 0 || v < len(tc.locate)
}

func (tc *SubGraphs) NumSubGraph() int {
	return len(tc.subGraphs)
}

func (tc *SubGraphs) Locate(v int) int {
	if !tc.hasVertical(v) {
		return -1
	}
	return tc.locate[v]
}

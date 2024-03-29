package graphs

import (
	"github.com/howz97/algorithm/util"
)

func NewGraph(size uint) *Graph {
	return &Graph{
		Digraph: NewDigraph(size),
	}
}

// Graph has no direction
type Graph struct {
	*Digraph
}

// NumEdge get the number of no-direction edges
func (g *Graph) NumEdge() uint {
	return g.Digraph.NumEdge() / 2
}

// AddEdge add an edge
func (g *Graph) AddEdge(a, b int) error {
	return g.addWeightedEdge(a, b, 1)
}

func (g *Graph) addWeightedEdge(src, dst int, w float64) error {
	if !g.HasVert(src) || !g.HasVert(dst) {
		return ErrVerticalNotExist
	}
	if src == dst {
		return ErrSelfLoop
	}
	g.Digraph.edges[src].Put(util.Int(dst), w)
	g.Digraph.edges[dst].Put(util.Int(src), w)
	return nil
}

// DelEdge delete an edge
func (g *Graph) DelEdge(a, b int) {
	g.Digraph.DelEdge(a, b)
	g.Digraph.DelEdge(b, a)
}

// TotalWeight sum the weight of all edges
func (g *Graph) TotalWeight() float64 {
	return g.Digraph.TotalWeight() / 2
}

// IterWEdge iterate all no-direction edges and their weight
func (g *Graph) IterWEdge(fn func(int, int, float64) bool) {
	visited := make(map[uint64]struct{})
	g.Digraph.IterWEdge(func(from int, to int, w float64) bool {
		if _, v := visited[uint64(to)<<32+uint64(from)]; v {
			return true
		}
		visited[uint64(from)<<32+uint64(to)] = struct{}{}
		return fn(from, to, w)
	})
}

// IterEdge iterate all no-direction edges
func (g *Graph) IterEdge(fn func(int, int) bool) {
	g.IterWEdge(func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

// IterWEdgeFrom iterate all reachable edges and their weight from vertical src
func (g *Graph) IterWEdgeFrom(src int, fn func(int, int, float64) bool) {
	visited := make(map[uint64]struct{})
	g.Digraph.IterWEdgeFrom(src, func(from int, to int, w float64) bool {
		if _, v := visited[uint64(to)<<32+uint64(from)]; v {
			return true
		}
		visited[uint64(from)<<32+uint64(to)] = struct{}{}
		return fn(from, to, w)
	})
}

// IterEdgeFrom iterate all reachable edges from vertical src
func (g *Graph) IterEdgeFrom(src int, fn func(int, int) bool) {
	g.IterWEdgeFrom(src, func(a int, b int, _ float64) bool {
		return fn(a, b)
	})
}

// HasCycle check whether graph contains cycle
func (g *Graph) HasCycle() bool {
	marked := make([]bool, g.NumVert())
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
	g.IterAdjacent(cur, func(adj int) bool {
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

type SubGraphs struct {
	locate    []int   // vertical -> subGraphID
	subGraphs [][]int // subGraphID -> all vertices
}

// SubGraphs calculate all sub-graphs of g
func (g *Graph) SubGraphs() *SubGraphs {
	tc := &SubGraphs{
		locate: make([]int, g.NumVert()),
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

// IsConn check whether a and b located in the same sub-graph
func (tc *SubGraphs) IsConn(a, b int) bool {
	return tc.locate[a] == tc.locate[b]
}

// Iterate all vertices of sub-graph where v located
func (tc *SubGraphs) Iterate(v int, fn func(int) bool) {
	for _, v := range tc.subGraphs[tc.locate[v]] {
		if !fn(v) {
			break
		}
	}
}

// NumSubGraph get the number of sub-graphs
func (tc *SubGraphs) NumSubGraph() int {
	return len(tc.subGraphs)
}

// Locate get the ID of sub-graph where v located
func (tc *SubGraphs) Locate(v int) int {
	return tc.locate[v]
}

// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphs

func NewGraph[T any](size uint) *Graph[T] {
	return &Graph[T]{
		Digraph: NewDigraph[T](size),
	}
}

// Graph has no direction
type Graph[T any] struct {
	*Digraph[T]
}

// NumEdge get the number of no-direction edges
func (g *Graph[T]) NumEdge() uint {
	return g.Digraph.NumEdge() / 2
}

// AddEdge add an edge
func (g *Graph[T]) AddEdge(a, b Id) error {
	return g.addWeightedEdge(a, b, 1)
}

func (g *Graph[T]) addWeightedEdge(src, dst Id, w Weight) error {
	if !g.HasVert(src) || !g.HasVert(dst) {
		return ErrInvalidVertex
	}
	if src == dst {
		return ErrInvalidEdge
	}
	g.Digraph.edges[src].Put(dst, w)
	g.Digraph.edges[dst].Put(src, w)
	return nil
}

// DelEdge delete an edge
func (g *Graph[T]) DelEdge(a, b Id) {
	g.Digraph.DelEdge(a, b)
	g.Digraph.DelEdge(b, a)
}

// TotalWeight sum the weight of all edges
func (g *Graph[T]) TotalWeight() Weight {
	return g.Digraph.TotalWeight() / 2
}

// IterWEdge iterate all no-direction edges and their weight
func (g *Graph[T]) IterWEdge(fn func(Id, Id, Weight) bool) {
	visited := make(map[uint64]struct{})
	g.Digraph.IterWEdge(func(from Id, to Id, w Weight) bool {
		if _, v := visited[uint64(to)<<32+uint64(from)]; v {
			return true
		}
		visited[uint64(from)<<32+uint64(to)] = struct{}{}
		return fn(from, to, w)
	})
}

// IterEdge iterate all no-direction edges
func (g *Graph[T]) IterEdge(fn func(Id, Id) bool) {
	g.IterWEdge(func(src Id, dst Id, _ Weight) bool {
		return fn(src, dst)
	})
}

// IterWEdgeFrom iterate all reachable edges and their weight from vertical src
func (g *Graph[T]) IterWEdgeFrom(src Id, fn func(Id, Id, Weight) bool) {
	visited := make(map[uint64]struct{})
	g.Digraph.IterWEdgeFrom(src, func(from Id, to Id, w Weight) bool {
		if _, v := visited[uint64(to)<<32+uint64(from)]; v {
			return true
		}
		visited[uint64(from)<<32+uint64(to)] = struct{}{}
		return fn(from, to, w)
	})
}

// IterEdgeFrom iterate all reachable edges from vertical src
func (g *Graph[T]) IterEdgeFrom(src Id, fn func(Id, Id) bool) {
	g.IterWEdgeFrom(src, func(a Id, b Id, _ Weight) bool {
		return fn(a, b)
	})
}

// HasCycle check whether graph contains cycle
func (g *Graph[T]) HasCycle() bool {
	marked := make([]bool, g.NumVert())
	for i, m := range marked {
		if m {
			continue
		}
		if g.detectCycleDFS(Id(i), Id(i), marked) {
			return true
		}
	}
	return false
}

func (g *Graph[T]) detectCycleDFS(last, cur Id, marked []bool) bool {
	marked[cur] = true
	found := false
	g.IterAdjacent(cur, func(adj Id) bool {
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
	locate    []Id   // vertical -> subGraphID
	subGraphs [][]Id // subGraphID -> all vertices
}

// SubGraphs calculate all sub-graphs of g
func (g *Graph[T]) SubGraphs() *SubGraphs {
	tc := &SubGraphs{
		locate: make([]Id, g.NumVert()),
	}
	for i := range tc.locate {
		tc.locate[i] = -1
	}

	var subGraphID Id
	for i, c := range tc.locate {
		if c < 0 {
			dfs := g.ReachableSlice(Id(i))
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
func (tc *SubGraphs) IsConn(a, b Id) bool {
	return tc.locate[a] >= 0 && tc.locate[a] == tc.locate[b]
}

// Iterate all vertices of sub-graph where v located
func (tc *SubGraphs) Iterate(v Id, fn func(Id) bool) {
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
func (tc *SubGraphs) Locate(v Id) Id {
	return tc.locate[v]
}

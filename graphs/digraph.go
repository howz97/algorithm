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

import (
	"fmt"
	"strconv"

	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/search"
	"gopkg.in/yaml.v2"
)

func NewDigraph[T any](size uint) *Digraph[T] {
	edges := make([]*search.HashMap[Id, float64], 0, size)
	return &Digraph[T]{edges: edges}
}

type Digraph[T any] struct {
	edges    []*search.HashMap[Id, float64]
	vertices []T
}

func (dg *Digraph[T]) AddVertex(vtx T) Id {
	id := Id(len(dg.vertices))
	dg.edges = append(dg.edges, search.NewHashMap[Id, float64]())
	dg.vertices = append(dg.vertices, vtx)
	return id
}

// NumVert get the number of vertices
func (dg *Digraph[T]) NumVert() uint {
	return uint(len(dg.edges))
}

// HasVert indicate whether dg contains vertical v
func (dg *Digraph[T]) HasVert(v Id) bool {
	return int(v) < len(dg.edges)
}

func (dg *Digraph[T]) Vertex(v Id) T {
	return dg.vertices[v]
}

// AddEdge add a new edge
func (dg *Digraph[T]) AddEdge(src, dst Id) error {
	return dg.addWeightedEdge(src, dst, 1)
}

// NumEdge get the number of edges
func (dg *Digraph[T]) NumEdge() uint {
	n := uint(0)
	for i := range dg.edges {
		n += dg.edges[i].Size()
	}
	return n
}

// HasEdge indicate whether dg contains the edge specified by params
func (dg *Digraph[T]) HasEdge(from, to Id) bool {
	if !dg.HasVert(from) || !dg.HasVert(to) {
		return false
	}
	_, ok := dg.edges[from].Get(to)
	return ok
}

func (dg *Digraph[T]) addWeightedEdge(src, dst Id, w float64) error {
	if !dg.HasVert(src) || !dg.HasVert(dst) {
		return ErrInvalidVertex
	}
	if src == dst {
		return ErrInvalidEdge
	}
	// invalid source vertex will panic
	dg.edges[src].Put(dst, w)
	return nil
}

// DelEdge delete an edge
func (dg *Digraph[T]) DelEdge(src, dst Id) {
	dg.edges[src].Del(dst)
}

// GetWeight get the weight of edge
// Zero will be returned if edge not exist
func (dg *Digraph[T]) GetWeight(from, to Id) float64 {
	w, _ := dg.edges[from].Get(to)
	return w
}

// TotalWeight sum the weight of all edges
func (dg *Digraph[T]) TotalWeight() float64 {
	var total float64
	dg.IterWEdge(func(_ Id, _ Id, w float64) bool {
		total += w
		return true
	})
	return total
}

// IterAdjacent iterate all adjacent vertices of v
func (dg *Digraph[T]) IterAdjacent(v Id, fn func(Id) bool) {
	dg.IterWAdjacent(v, func(a Id, _ float64) bool {
		return fn(a)
	})
}

// IterWAdjacent iterate all adjacent vertices and weight of v
func (dg *Digraph[T]) IterWAdjacent(v Id, fn func(Id, float64) bool) {
	dg.edges[v].Range(func(key Id, val float64) bool {
		return fn(key, val)
	})
}

// Adjacent return a slice contains all adjacent vertices of v
func (dg *Digraph[T]) Adjacent(v Id) (adj []Id) {
	dg.IterAdjacent(v, func(a Id) bool {
		adj = append(adj, a)
		return true
	})
	return adj
}

func (dg *Digraph[T]) String() string {
	out := ""
	for i := range dg.edges {
		out += fmt.Sprint(i) + " :"
		dg.IterAdjacent(Id(i), func(j Id) bool {
			out += " " + fmt.Sprint(j)
			return true
		})
		out += "\n"
	}
	out += "\n"
	return out
}

// Reverse all edges of dg
func (dg *Digraph[T]) Reverse() *Digraph[T] {
	rg := NewDigraph[T](dg.NumVert())
	dg.IterWEdge(func(from Id, to Id, w float64) bool {
		rg.addWeightedEdge(to, from, w)
		return true
	})
	return rg
}

// FindNegativeEdge find a negative edge.
// If here is no negative edge, (0, 0) will be returned
func (dg *Digraph[T]) FindNegativeEdge() (src, dst Id) {
	dg.IterWEdge(func(v Id, v2 Id, w float64) bool {
		if w < 0 {
			src = v
			dst = v2
			return false
		}
		return true
	})
	return
}

// FindNegativeEdgeFrom find a reachable negative edge from the specified start vertical
// If here is no negative edge, (0, 0) will be returned
func (dg *Digraph[T]) FindNegativeEdgeFrom(start Id) (src Id, dst Id) {
	dg.IterWEdgeFrom(start, func(v0 Id, v1 Id, w float64) bool {
		if w < 0 {
			src = v0
			dst = v1
			return false
		}
		return true
	})
	return
}

// FindCycle find any directed cycle in dg
func (dg *Digraph[T]) FindCycle() []Id {
	marks := make([]bool, dg.NumVert())
	path := NewPath()
	for v, m := range marks {
		if !m {
			if dg.detectCycleDFS(Id(v), marks, path) {
				return path.Cycle()
			}
		}
	}
	return nil
}

// FindCycleFrom find any directed cycle from vertical v in dg
// But not include cycle that can not be accessed from v
func (dg *Digraph[T]) FindCycleFrom(v Id) *Path {
	marks := make([]bool, dg.NumVert())
	path := NewPath()
	if dg.detectCycleDFS(v, marks, path) {
		return path
	}
	return nil
}

func (dg *Digraph[T]) detectCycleDFS(v Id, marked []bool, path *Path) bool {
	found := false
	dg.IterWAdjacent(v, func(a Id, w float64) bool {
		if marked[a] {
			return true
		}
		if path.HasVert(a) {
			path.Push(v, a, w)
			found = true
			return false
		}
		path.Push(v, a, w)
		found = dg.detectCycleDFS(a, marked, path)
		if !found {
			path.Pop()
		}
		return !found
	})
	marked[v] = true
	return found
}

// Topological return a stack that will pop vertices in topological order
func (dg *Digraph[T]) Topological() (order *basic.Stack[Id]) {
	if dg.FindCycle() != nil {
		return
	}
	order = basic.NewStack[Id](int(dg.NumVert()))
	dg.VetBackDfs(func(v Id) bool {
		order.PushBack(v)
		return true
	})
	return
}

// IterWEdge iterate all edges and their weight in dg
func (dg *Digraph[T]) IterWEdge(fn func(Id, Id, float64) bool) {
	for src, hm := range dg.edges {
		goon := true
		hm.Range(func(dst Id, v float64) bool {
			goon = fn(Id(src), dst, v)
			return goon
		})
		if !goon {
			break
		}
	}
}

// IterEdge iterate all edges in dg
func (dg *Digraph[T]) IterEdge(fn func(Id, Id) bool) {
	dg.IterWEdge(func(src Id, dst Id, _ float64) bool {
		return fn(src, dst)
	})
}

// IterWEdgeFrom iterate all reachable edges and their weight from vertical src
func (dg *Digraph[T]) IterWEdgeFrom(src Id, fn func(Id, Id, float64) bool) {
	dg.IterVertDFS(src, func(v Id) bool {
		goon := true
		dg.IterWAdjacent(v, func(a Id, w float64) bool {
			goon = fn(v, a, w)
			return goon
		})
		return goon
	})
}

// IterEdgeFrom iterate all reachable edges from vertical src
func (dg *Digraph[T]) IterEdgeFrom(src Id, fn func(Id, Id) bool) {
	dg.IterWEdgeFrom(src, func(src Id, dst Id, _ float64) bool {
		return fn(src, dst)
	})
}

// IterVertDFS iterate all reachable vertices from vertical src in DFS order
func (dg *Digraph[T]) IterVertDFS(src Id, fn func(Id) bool) {
	dg.iterUnMarkVetFrom(src, nil, fn)
}

func (dg *Digraph[T]) iterUnMarkVetFrom(src Id, marked []bool, fn func(Id) bool) {
	if !dg.HasVert(src) {
		return
	}
	if len(marked) == 0 {
		marked = make([]bool, dg.NumVert())
	}
	dg.iterUnMarkVetDFS(src, marked, fn)
}

func (dg *Digraph[T]) iterUnMarkVetDFS(v Id, marked []bool, fn func(Id) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	dg.IterAdjacent(v, func(adj Id) bool {
		if !marked[adj] {
			if !dg.iterUnMarkVetDFS(adj, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

// ReachableSlice get a slice contains all reachable vertices from src
func (dg *Digraph[T]) ReachableSlice(src Id) []Id {
	if !dg.HasVert(src) {
		return nil
	}
	var arrived []Id
	dg.IterVertDFS(src, func(v Id) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

// ReachableBits get a bit-map contains all reachable vertices from src
func (dg *Digraph[T]) ReachableBits(src Id) []bool {
	marked := make([]bool, dg.NumVert())
	dg.iterUnMarkVetDFS(src, marked, func(_ Id) bool { return true })
	return marked
}

// VetBackDfs iterate all vertices in Back-DFS order
func (dg *Digraph[T]) VetBackDfs(fn func(Id) bool) {
	marked := make([]bool, dg.NumVert())
	for v := range marked {
		if marked[v] {
			continue
		}
		dg.bDfs(Id(v), marked, fn)
	}
}

// VetBackDfsFrom iterate all reachable vertices from vertical src in RDFS order
func (dg *Digraph[T]) VetBackDfsFrom(src Id, fn func(Id) bool) {
	marked := make([]bool, dg.NumVert())
	dg.bDfs(src, marked, fn)
}

func (dg *Digraph[T]) bDfs(v Id, marked []bool, fn func(Id) bool) bool {
	marked[v] = true
	goon := true // continue DFS or abort
	dg.IterAdjacent(v, func(a Id) bool {
		if !marked[a] {
			if !dg.bDfs(a, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	if !goon {
		return false
	}
	return fn(v)
}

// Bipartite put two colors on all nodes while any connected nodes have different color
func (dg *Digraph[T]) Bipartite() (colors []bool) {
	marks := make([]bool, dg.NumVert())
	colors = make([]bool, dg.NumVert())
	for i, m := range marks {
		if m {
			continue
		}
		if !dg.isBipartiteDFS(Id(i), true, colors, marks) {
			return nil
		}
	}
	return
}

func (dg *Digraph[T]) isBipartiteDFS(cur Id, color bool, colors []bool, marked []bool) bool {
	isBip := true
	if !marked[cur] {
		marked[cur] = true
		colors[cur] = color
		dg.IterAdjacent(cur, func(adj Id) bool {
			if !dg.isBipartiteDFS(adj, !color, colors, marked) {
				isBip = false
				return false
			}
			return true
		})
	} else {
		isBip = colors[cur] == color
	}
	return isBip
}

// IsSameWith check whether dg is the same with other
func (dg *Digraph[T]) IsSameWith(other Digraph[T]) bool {
	if dg.NumVert() != other.NumVert() {
		return false
	}
	if dg.NumEdge() != other.NumEdge() {
		return false
	}
	isSame := true
	dg.IterWEdge(func(from Id, to Id, w float64) bool {
		if w != other.GetWeight(from, to) {
			isSame = false
			return false
		}
		return true
	})
	return isSame
}

// Marshal dg into yaml format
func (dg *Digraph[T]) Marshal() ([]byte, error) {
	m := make(map[string]map[string]float64)
	for v := Id(0); uint(v) < dg.NumVert(); v++ {
		edges := make(map[string]float64)
		dg.IterWAdjacent(v, func(a Id, w float64) bool {
			edges[strconv.Itoa(int(a))] = w
			return true
		})
		m[strconv.Itoa(int(v))] = edges
	}
	return yaml.Marshal(m)
}

// SCC is strong connected components of digraph
// vertices in the same component can access each other
type SCC struct {
	locate     []Id   // vertical -> componentID
	components [][]Id // componentID -> vertices
}

// SCC calculate strong connected components of digraph with kosaraju algorithm
func (dg *Digraph[T]) SCC() *SCC {
	scc := &SCC{
		locate:     make([]Id, dg.NumVert()),
		components: make([][]Id, 1),
	}
	marked := make([]bool, dg.NumVert())
	dg.VetBackDfs(func(v Id) bool {
		if !marked[v] {
			cmpnId := Id(len(scc.components))
			cmpn := make([]Id, 0, 8)
			dg.iterUnMarkVetFrom(v, marked, func(w Id) bool {
				scc.locate[w] = cmpnId
				cmpn = append(cmpn, w)
				return true
			})
			scc.components = append(scc.components, cmpn)
		}
		return true
	})
	return scc
}

// IsStronglyConn check whether src is strongly connected with dst
func (scc *SCC) IsStronglyConn(src, dst Id) bool {
	return scc.locate[src] == scc.locate[dst]
}

// Comp get the strongly connected component ID of vertical v
func (scc *SCC) Comp(v Id) Id {
	return scc.locate[v]
}

func (scc *SCC) IterComponent(c Id, fn func(Id) bool) {
	for _, w := range scc.components[c] {
		if !fn(w) {
			break
		}
	}
}

// NumComponents get the number of components
func (scc *SCC) NumComponents() int {
	return len(scc.components) - 1
}

type Reachable [][]bool

// Reachable save all reachable information of dg
func (dg *Digraph[T]) Reachable() Reachable {
	tc := make(Reachable, dg.NumVert())
	for v := range tc {
		tc[v] = dg.ReachableBits(Id(v))
	}
	return tc
}

// CanReach check whether src can reach dst
func (tc Reachable) CanReach(src, dst Id) bool {
	return tc[src][dst]
}

// Iterate all reachable vertices from src
func (tc Reachable) Iterate(src Id, fn func(v Id) bool) {
	for w, marked := range tc[src] {
		if marked {
			if !fn(Id(w)) {
				break
			}
		}
	}
}

type BFS[T any] struct {
	dg     *Digraph[T]
	src    Id
	marked []bool
	edgeTo []Id
}

// BFS save all BFS information from src
func (dg *Digraph[T]) BFS(src Id) *BFS[T] {
	bfs := &BFS[T]{
		dg:     dg,
		src:    src,
		marked: make([]bool, dg.NumVert()),
		edgeTo: make([]Id, dg.NumVert()),
	}
	q := basic.NewList[Id]()
	bfs.marked[src] = true
	q.PushBack(src)
	for q.Size() > 0 {
		vet := q.Front()
		q.PopFront()
		dg.IterAdjacent(vet, func(adj Id) bool {
			if !bfs.marked[adj] {
				bfs.edgeTo[adj] = vet
				bfs.marked[adj] = true
				q.PushBack(adj)
			}
			return true
		})
	}
	return bfs
}

// CanReach check whether src can reach dst
func (bfs *BFS[T]) CanReach(dst Id) bool {
	return bfs.marked[dst]
}

// ShortestPathTo get the shortest path to dst (ignore weight)
func (bfs *BFS[T]) ShortestPathTo(dst Id) *Path {
	if !bfs.CanReach(dst) {
		return nil
	}
	if dst == bfs.src {
		return nil
	}
	path := NewPath()
	for dst != bfs.src {
		path.Push(bfs.edgeTo[dst], dst, 1)
		dst = bfs.edgeTo[dst]
	}
	// path.Reverse()
	return path
}

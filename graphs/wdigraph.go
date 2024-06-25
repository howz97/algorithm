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
	"math"
	"strconv"

	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/pqueue"
)

const (
	Auto = iota
	Dijkstra
	Topological
	BellmanFord
)

func NewWDigraph[T any](size uint) *WDigraph[T] {
	return &WDigraph[T]{
		Digraph: NewDigraph[T](size),
	}
}

// WDigraph is edge weighted digraph
type WDigraph[T any] struct {
	*Digraph[T]
}

// AddEdge add a weighted and directed edge
func (g *WDigraph[T]) AddEdge(src, dst Id, w float64) {
	g.addWeightedEdge(src, dst, w)
}

// ShortestPathTree get the shortest path tree from src by the specified algorithm
func (g *WDigraph[T]) ShortestPathTree(src Id, alg int) (*PathTree[T], error) {
	if !g.HasVert(src) {
		return nil, ErrInvalidVertex
	}
	var err error
	spt := newShortestPathTree(g, src)
	switch alg {
	case Dijkstra:
		if src, dst := g.FindNegativeEdgeFrom(src); src >= 0 {
			err = fmt.Errorf(fmt.Sprintf("negative edge %d->%d", src, dst))
			break
		}
		spt.initDijkstra(g)
	case Topological:
		if p := g.FindCycleFrom(src); p != nil {
			err = p.Cycle()
			break
		}
		spt.initTopological(g)
	case BellmanFord:
		err = spt.initBellmanFord()
	default:
		if g.FindCycleFrom(src) == nil {
			spt.initTopological(g)
		} else if src, _ := g.FindNegativeEdgeFrom(src); src < 0 {
			spt.initDijkstra(g)
		} else {
			err = spt.initBellmanFord()
		}
	}
	return spt, err
}

// PathTree is shortest path tree
type PathTree[T any] struct {
	wdg    *WDigraph[T]
	src    Id
	distTo []float64
	edgeTo []Id
}

func newShortestPathTree[T any](g *WDigraph[T], src Id) *PathTree[T] {
	spt := &PathTree[T]{
		wdg:    g,
		src:    src,
		distTo: make([]float64, g.NumVert()),
		edgeTo: make([]Id, g.NumVert()),
	}
	for i := range spt.distTo {
		spt.distTo[i] = math.Inf(1)
	}
	spt.distTo[src] = 0
	for i := range spt.edgeTo {
		spt.edgeTo[i] = -1
	}
	return spt
}

// CanReach check whether src can reach dst
func (spt *PathTree[T]) CanReach(dst Id) bool {
	return spt.distTo[dst] != math.Inf(1)
}

// DistanceTo get the distance from src to dst
func (spt *PathTree[T]) DistanceTo(dst Id) float64 {
	return spt.distTo[dst]
}

// PathTo get the path from src to dst
func (spt *PathTree[T]) PathTo(dst Id) *Path {
	src := spt.edgeTo[dst]
	if src < 0 {
		return nil
	}
	path := NewPath()
	for {
		path.Push(src, dst, spt.distTo[dst]-spt.distTo[src])
		dst = src
		src = spt.edgeTo[dst]
		if src < 0 {
			break
		}
	}
	// path.Reverse()
	return path
}

func (spt *PathTree[T]) initDijkstra(g *WDigraph[T]) {
	pq := pqueue.NewFixable[float64, Id](g.NumVert())
	spt.dijkstraRelax(spt.src, pq)
	for pq.Size() > 0 {
		m := pq.Pop()
		spt.dijkstraRelax(m, pq)
	}
}

func (spt *PathTree[T]) dijkstraRelax(v Id, pq *pqueue.Fixable[float64, Id]) {
	spt.wdg.IterWAdjacent(v, func(adj Id, w float64) bool {
		if spt.distTo[v]+w < spt.distTo[adj] {
			inPQ := spt.distTo[adj] != math.Inf(1)
			spt.edgeTo[adj] = v
			spt.distTo[adj] = spt.distTo[v] + w
			if inPQ {
				pq.Fix(spt.distTo[adj], adj)
			} else {
				pq.PushPair(spt.distTo[adj], adj)
			}
		}
		return true
	})
}

func (spt *PathTree[T]) initTopological(g *WDigraph[T]) {
	order := basic.NewStack[Id](int(g.NumVert()))
	g.IterBDFSFrom(spt.src, func(v Id) bool {
		order.PushBack(v)
		return true
	})
	for order.Size() > 0 {
		v := order.Back()
		order.PopBack()
		spt.topologicalRelax(v)
	}
}

func (spt *PathTree[T]) topologicalRelax(v Id) {
	spt.wdg.IterWAdjacent(v, func(a Id, w float64) bool {
		if spt.distTo[v]+w < spt.distTo[a] {
			spt.edgeTo[a] = v
			spt.distTo[a] = spt.distTo[v] + w
		}
		return true
	})
}

func (spt *PathTree[T]) initBellmanFord() error {
	q := basic.NewList[Id]()
	onQ := make([]bool, spt.wdg.NumVert())
	q.PushBack(spt.src)
	onQ[spt.src] = true
	cnt := uint(0)
	for q.Size() > 0 {
		v := q.Front()
		q.PopFront()
		onQ[v] = false
		spt.bellmanFordRelax(v, q, onQ)
		cnt++
		if cnt%spt.wdg.NumVert() == 0 {
			c := spt.toWDigraph().FindCycle()
			if c != nil {
				return c
			}
		}
	}
	return nil
}

func (spt *PathTree[T]) toWDigraph() *WDigraph[T] {
	sptg := NewWDigraph[T](spt.wdg.NumVert())
	for to, from := range spt.edgeTo {
		if from < 0 {
			continue
		}
		sptg.addWeightedEdge(from, Id(to), spt.wdg.GetWeight(from, Id(to)))
	}
	return sptg
}

func (spt *PathTree[T]) bellmanFordRelax(v Id, q *basic.List[Id], onQ []bool) {
	spt.wdg.IterWAdjacent(v, func(adj Id, w float64) bool {
		if spt.distTo[v]+w < spt.distTo[adj] {
			spt.edgeTo[adj] = v
			spt.distTo[adj] = spt.distTo[v] + w
			if !onQ[adj] {
				q.PushBack(adj)
				onQ[adj] = true
			}
		}
		return true
	})
}

func (g *WDigraph[T]) SearcherDijkstra() (*Searcher[T], error) {
	if src, dst := g.FindNegativeEdge(); src >= 0 {
		return nil, fmt.Errorf(fmt.Sprintf("negative edge %d->%d", src, dst))
	}
	sps := &Searcher[T]{
		spt: make([]*PathTree[T], g.NumVert()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, Id(v))
		tree.initDijkstra(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph[T]) SearcherTopological() (*Searcher[T], error) {
	if c := g.FindCycle(); c != nil {
		return nil, c
	}
	sps := &Searcher[T]{
		spt: make([]*PathTree[T], g.NumVert()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, Id(v))
		tree.initTopological(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph[T]) SearcherBellmanFord() (*Searcher[T], error) {
	sps := &Searcher[T]{
		spt: make([]*PathTree[T], g.NumVert()),
	}
	var err error
	for v := range sps.spt {
		tree := newShortestPathTree(g, Id(v))
		err = tree.initBellmanFord()
		if err != nil {
			return nil, err
		}
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph[T]) Searcher() (*Searcher[T], error) {
	sps, err := g.SearcherTopological()
	if err == nil {
		return sps, nil
	}
	sps, err = g.SearcherDijkstra()
	if err == nil {
		return sps, nil
	}
	return g.SearcherBellmanFord()
}

type Searcher[T any] struct {
	spt []*PathTree[T]
}

func (s *Searcher[T]) GetDistance(src, dst Id) float64 {
	return s.spt[src].DistanceTo(dst)
}

func (s *Searcher[T]) GetPath(src, dst Id) *Path {
	return s.spt[src].PathTo(dst)
}

func NewPath() *Path {
	return &Path{
		Stack: basic.NewStack[edge](2),
	}
}

type Path struct {
	*basic.Stack[edge]
}

func (p *Path) Push(from, to Id, w float64) {
	p.Stack.PushBack(edge{
		from:   from,
		to:     to,
		weight: w,
	})
}

func (p *Path) Pop() edge {
	e := p.Stack.Back()
	p.Stack.PopBack()
	return e
}

func (p *Path) Distance() float64 {
	d := 0.0
	p.Iterate(false, func(e edge) bool {
		d += e.weight
		return true
	})
	return d
}

func (p *Path) HasVert(v Id) bool {
	if p.Size() <= 0 {
		return false
	}
	found := false
	p.Iterate(false, func(e edge) bool {
		if e.from == v {
			found = true
			return false
		}
		return true
	})
	if !found {
		found = p.Peek().to == v
	}
	return found
}

func (p *Path) Str() string {
	if p == nil || p.Size() <= 0 {
		return "path not exist"
	}
	i2s := strconv.Itoa
	str := fmt.Sprintf("[TotalDistance=%v]", p.Distance())
	p.Iterate(true, func(e edge) bool {
		str += " " + e.string(i2s)
		return true
	})
	return str
}

func (p *Path) Cycle() *Cycle {
	if p.Size() == 0 {
		return nil
	}
	e := p.Peek()
	x := e.to
	i := p.Find(func(v edge) bool {
		return v.from == x
	})
	path := NewPath()
	p.IterRange(i, p.Size()-1, func(e edge) bool {
		path.Push(e.from, e.to, e.weight)
		return true
	})
	return &Cycle{path}
}

type Cycle struct {
	*Path
}

func (c *Cycle) Error() string {
	return c.Str()
}

func (c *Cycle) Cycle() *Cycle {
	return c
}

type edge struct {
	from, to Id
	weight   float64
}

func (e *edge) string(i2s func(int) string) string {
	return fmt.Sprintf("%s->%s(%.2f)", i2s(int(e.from)), i2s(int(e.to)), e.weight)
}

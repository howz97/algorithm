package ewd

import (
	"fmt"
	pqueue "github.com/zh1014/algorithm/pqueue/binaryheap"
	"github.com/zh1014/algorithm/stack"
	"math"
)

const (
	// the algorithms of finding shortest path
	Dijkstra = iota
	Topological
	BellmanFord
)

type ShortestPathSearcher struct {
	g   EdgeWeightedDigraph
	spt []*ShortestPathTree
}

func NewSPS(g EdgeWeightedDigraph, alg int) (*ShortestPathSearcher, error) {
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumV()),
	}
	// TODO 检查负权重边，环，负权重环
	for src := range sps.spt {
		sps.spt[src] = NewSPT(g, src, alg)
	}
	return sps, nil
}

func (s *ShortestPathSearcher) Distance(src, dst int) float64 {
	if !s.g.HasV(src) && !s.g.HasV(dst) {
		panic(ErrVerticalNotExist)
	}
	return s.spt[src].DistanceTo(dst)
}

func (s *ShortestPathSearcher) Path(src, dst int) *stack.Stack {
	if !s.g.HasV(src) && !s.g.HasV(dst) {
		panic(ErrVerticalNotExist)
	}
	return s.spt[src].PathTo(dst)
}

type ShortestPathTree struct {
	g      EdgeWeightedDigraph
	src    int
	distTo []float64
	edgeTo []*Edge
}

func NewSPT(g EdgeWeightedDigraph, src int, alg int) *ShortestPathTree {
	if !g.HasV(src) {
		panic(ErrVerticalNotExist)
	}
	if alg < 0 || alg > 2 {
		panic(fmt.Sprintf("Invalid algorithm %v (should be in 0 ~ 2)", alg))
	}
	spt := &ShortestPathTree{
		g:      g,
		src:    src,
		distTo: make([]float64, g.NumV()),
		edgeTo: make([]*Edge, g.NumV()),
	}
	for i := range spt.distTo {
		spt.distTo[i] = math.Inf(1)
	}
	spt.distTo[src] = 0
	switch alg {
	case Dijkstra:
		spt.dijkstra()
	case Topological:
		spt.topological(src)
	case BellmanFord:
		spt.bellmanFord()
	}
	return spt
}

func (spt *ShortestPathTree) CanReach(dst int) bool {
	if !spt.g.HasV(dst) {
		panic(ErrVerticalNotExist)
	}
	return spt.distTo[dst] != math.Inf(1)
}

func (spt *ShortestPathTree) DistanceTo(dst int) float64 {
	if !spt.g.HasV(dst) {
		panic(ErrVerticalNotExist)
	}
	return spt.distTo[dst]
}

func (spt *ShortestPathTree) PathTo(dst int) *stack.Stack {
	if !spt.g.HasV(dst) {
		panic(ErrVerticalNotExist)
	}
	s := stack.NewStack(spt.g.NumV() - 1)
	for spt.edgeTo[dst] != nil {
		s.Push(spt.edgeTo[dst])
		dst = spt.edgeTo[dst].from
	}
	return s
}

func (spt *ShortestPathTree) dijkstra() {
	pq := pqueue.NewBinHeap(spt.g.NumV())
	dijkstraRelax(spt.g, spt.src, spt.edgeTo, spt.distTo, pq)
	for !pq.IsEmpty() {
		m := pq.DelMin().(int)
		dijkstraRelax(spt.g, m, spt.edgeTo, spt.distTo, pq)
	}
}

func dijkstraRelax(g EdgeWeightedDigraph, v int, edgeTo []*Edge, distTo []float64, pq *pqueue.BinHeap) {
	adj := g.Adjacent(v)
	for _, e := range adj {
		if distTo[v]+e.weight < distTo[e.to] {
			inPQ := true
			if distTo[e.to] == math.Inf(1) {
				inPQ = false
			}
			edgeTo[e.to] = e
			distTo[e.to] = distTo[v] + e.weight
			if inPQ {
				pq.Update(distTo[e.to], e.to)
			} else {
				pq.Insert(distTo[e.to], e.to)
			}
		}
	}
}

func (spt *ShortestPathTree) topological(src int) {
	marked := make([]bool, spt.g.NumV())
	topoSortStack := stack.NewStackInt(spt.g.NumV())
	reversePostDFS(spt.g, spt.src, marked, topoSortStack)
	for !topoSortStack.IsEmpty() {
		topologicalRelax(spt.g, topoSortStack.Pop(), spt.edgeTo, spt.distTo)
	}
}

func topologicalRelax(g EdgeWeightedDigraph, v int, edgeTo []*Edge, distTo []float64) {
	adj := g.Adjacent(v)
	for _, e := range adj {
		if distTo[v]+e.weight < distTo[e.to] {
			edgeTo[e.to] = e
			distTo[e.to] = distTo[v] + e.weight
		}
	}
}

func (spt *ShortestPathTree) bellmanFord() {

}

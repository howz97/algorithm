package ewd

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	pqueue "github.com/zh1014/algorithm/pqueue/binaryheap"
	"github.com/zh1014/algorithm/queue"
	"github.com/zh1014/algorithm/stack"
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
	var err error
	switch alg {
	case Dijkstra:
		if g.HasNegativeEdge() {
			return nil, errors.New("this digraph contains negative edge")
		}
	case Topological:
		if DetectDirCycle(g) {
			return nil, errors.New("this digraph contains directed cycle")
		}
	}
	for src := range sps.spt {
		sps.spt[src], err = NewSPT(g, src, alg)
		if err != nil {
			return nil, err
		}
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

func NewSPT(g EdgeWeightedDigraph, src int, alg int) (*ShortestPathTree, error) {
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
		spt.topological()
	case BellmanFord:
		negativeCycle := spt.bellmanFord()
		if negativeCycle != nil {
			s := ""
			for !negativeCycle.IsEmpty() {
				e := negativeCycle.Pop().(*Edge)
				s += strconv.Itoa(e.from) + "->" + strconv.Itoa(e.to) + "  "
			}
			return nil, errors.New("weight negative cycle: " + s)
		}
	}
	return spt, nil
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
				pq.Update(int(distTo[e.to]), e.to)
			} else {
				pq.Insert(int(distTo[e.to]), e.to)
			}
		}
	}
}

func (spt *ShortestPathTree) topological() {
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

func (spt *ShortestPathTree) bellmanFord() *stack.Stack {
	needRelax := queue.NewIntQ()
	onQ := make([]bool, spt.g.NumV())
	needRelax.PushBack(spt.src)
	onQ[spt.src] = true
	relaxTimes := 0
	for !needRelax.IsEmpty() {
		v := needRelax.Front()
		onQ[spt.src] = false
		bellmanFordRelax(spt.g, v, spt.edgeTo, spt.distTo, needRelax, onQ)
		relaxTimes++
		if relaxTimes%spt.g.NumV() == 0 {
			if c := spt.findNegativeCycle(); c != nil {
				return c
			}
		}
	}
	return nil
}

func (spt *ShortestPathTree) findNegativeCycle() *stack.Stack {
	g := NewEWD(spt.g.NumV())
	for _, e := range spt.edgeTo {
		if e != nil {
			g.AddEdge(e)
		}
	}
	marked := make([]bool, spt.g.NumV())
	s := stack.NewStack(spt.g.NumV())
	onS := make([]bool, spt.g.NumV())
	for i := 0; i < spt.g.NumV(); i++ {
		if !marked[i] {
			if c := findNC(g, i, marked, s, onS, spt.edgeTo); c != nil {
				return c
			}
		}
	}
	return nil
}

func findNC(g EdgeWeightedDigraph, v int, marked []bool, s *stack.Stack, onS []bool, edgeTo []*Edge) *stack.Stack {
	if onS[v] {
		c := stack.NewStack(g.NumV())
		var weight float64
		for !s.IsEmpty() {
			e := s.Pop().(*Edge)
			weight += e.weight
			c.Push(e)
		}
		if weight < 0 {
			return c
		}
		for !c.IsEmpty() {
			s.Push(c.Pop().(*Edge))
		}
		return nil
	}
	if marked[v] {
		return nil
	}
	marked[v] = true
	onS[v] = true
	s.Push(edgeTo[v])
	adj := g.Adjacent(v)
	for _, e := range adj {
		c := findNC(g, e.to, marked, s, onS, edgeTo)
		if c != nil {
			return c
		}
	}
	onS[v] = false
	s.Pop()
	return nil
}

func bellmanFordRelax(g EdgeWeightedDigraph, v int, edgeTo []*Edge, distTo []float64, needRelax *queue.IntQ, onQ []bool) {
	adj := g.Adjacent(v)
	for _, e := range adj {
		if distTo[v]+e.weight < distTo[e.to] {
			edgeTo[e.to] = e
			distTo[e.to] = distTo[v] + e.weight
			if !onQ[e.to] {
				needRelax.PushBack(e.to)
				onQ[e.to] = true
			}
		}
	}
}

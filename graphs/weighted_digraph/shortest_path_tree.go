package weighted_digraph

import (
	pqueue "github.com/howz97/algorithm/pqueue/binaryheap"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/stack"
	"math"
	"strconv"
)

type ShortestPathTree struct {
	g      EdgeWeightedDigraph
	src    int
	distTo []float64
	edgeTo []*Edge
}

func (g EdgeWeightedDigraph) NewShortestPathTree(src int) *ShortestPathTree {
	if !g.HasV(src) {
		panic(ErrVerticalNotExist)
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
	return spt
}

func (spt *ShortestPathTree) CanReach(dst int) bool {
	if !spt.g.HasV(dst) {
		return false
	}
	return spt.distTo[dst] != math.Inf(1)
}

func (spt *ShortestPathTree) DistanceTo(dst int) float64 {
	if !spt.g.HasV(dst) {
		return math.Inf(1)
	}
	return spt.distTo[dst]
}

func (spt *ShortestPathTree) PathTo(dst int) *stack.Stack {
	if !spt.g.HasV(dst) {
		return nil
	}
	s := stack.NewStack(spt.g.NumV() - 1)
	for spt.edgeTo[dst] != nil {
		s.Push(spt.edgeTo[dst])
		dst = spt.edgeTo[dst].from
	}
	return s
}

// ============================ Dijkstra ============================

func (spt *ShortestPathTree) InitDijkstra() {
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
			inPQ := distTo[e.to] != math.Inf(1)
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

// ============================ Topological ============================

func (spt *ShortestPathTree) InitTopological() {
	marked := make([]bool, spt.g.NumV())
	topoSortStack := stack.NewStackInt(spt.g.NumV())
	spt.g.reversePostDFS(spt.src, marked, topoSortStack)
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

// ============================ BellmanFord ============================

type NegativeCycle struct {
	Stack *stack.Stack
}

func (nc NegativeCycle) Error() string {
	s := ""
	for !nc.Stack.IsEmpty() {
		e := nc.Stack.Pop().(*Edge)
		s += strconv.Itoa(e.from) + "->" + strconv.Itoa(e.to) + "  "
	}
	return "weight negative cycle: " + s
}

func (spt *ShortestPathTree) InitBellmanFord() error {
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
				return NegativeCycle{Stack: c}
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

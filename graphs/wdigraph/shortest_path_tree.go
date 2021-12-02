package wdigraph

import (
	pqueue "github.com/howz97/algorithm/pqueue/binaryheap"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/stack"
	"math"
	"strconv"
)

type ShortestPathTree struct {
	g      WDigraph
	src    int
	distTo []float64
	edgeTo []*Edge
}

func (g WDigraph) NewShortestPathTree(src int) *ShortestPathTree {
	if !g.HasVertical(src) {
		panic(ErrVerticalNotExist)
	}
	spt := &ShortestPathTree{
		g:      g,
		src:    src,
		distTo: make([]float64, g.NumVertical()),
		edgeTo: make([]*Edge, g.NumVertical()),
	}
	for i := range spt.distTo {
		spt.distTo[i] = math.Inf(1)
	}
	spt.distTo[src] = 0
	return spt
}

func (spt *ShortestPathTree) CanReach(dst int) bool {
	if !spt.g.HasVertical(dst) {
		return false
	}
	return spt.distTo[dst] != math.Inf(1)
}

func (spt *ShortestPathTree) DistanceTo(dst int) float64 {
	if !spt.g.HasVertical(dst) {
		return math.Inf(1)
	}
	return spt.distTo[dst]
}

func (spt *ShortestPathTree) PathTo(dst int) *stack.Stack {
	if !spt.g.HasVertical(dst) {
		return nil
	}
	s := stack.New(spt.g.NumVertical() - 1)
	for spt.edgeTo[dst] != nil {
		s.Push(spt.edgeTo[dst])
		dst = spt.edgeTo[dst].from
	}
	return s
}

// ============================ Dijkstra ============================

func (spt *ShortestPathTree) InitDijkstra() {
	pq := pqueue.NewBinHeap(spt.g.NumVertical())
	dijkstraRelax(spt.g, spt.src, spt.edgeTo, spt.distTo, pq)
	for !pq.IsEmpty() {
		m := pq.DelMin().(int)
		dijkstraRelax(spt.g, m, spt.edgeTo, spt.distTo, pq)
	}
}

func dijkstraRelax(g WDigraph, v int, edgeTo []*Edge, distTo []float64, pq *pqueue.BinHeap) {
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
	marked := make([]bool, spt.g.NumVertical())
	topoSortStack := stack.NewInt(spt.g.NumVertical())
	spt.g.reversePostDFS(spt.src, marked, topoSortStack)
	for {
		e, ok := topoSortStack.Pop()
		if !ok {
			break
		}
		topologicalRelax(spt.g, e, spt.edgeTo, spt.distTo)
	}
}

func topologicalRelax(g WDigraph, v int, edgeTo []*Edge, distTo []float64) {
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
	for {
		e, ok := nc.Stack.Pop()
		if !ok {
			break
		}
		eg := e.(*Edge)
		s += strconv.Itoa(eg.from) + "->" + strconv.Itoa(eg.to) + "  "
	}
	return "weight negative cycle: " + s
}

func (spt *ShortestPathTree) InitBellmanFord() error {
	needRelax := queue.NewIntQ()
	onQ := make([]bool, spt.g.NumVertical())
	needRelax.PushBack(spt.src)
	onQ[spt.src] = true
	relaxTimes := 0
	for !needRelax.IsEmpty() {
		v := needRelax.Front()
		onQ[spt.src] = false
		bellmanFordRelax(spt.g, v, spt.edgeTo, spt.distTo, needRelax, onQ)
		relaxTimes++
		if relaxTimes%spt.g.NumVertical() == 0 {
			if c := spt.findNegativeCycle(); c != nil {
				return NegativeCycle{Stack: c}
			}
		}
	}
	return nil
}

func (spt *ShortestPathTree) findNegativeCycle() *stack.Stack {
	g := NewWDigraph(spt.g.NumVertical())
	for _, e := range spt.edgeTo {
		if e != nil {
			g.AddEdge(e)
		}
	}
	marked := make([]bool, spt.g.NumVertical())
	s := stack.New(spt.g.NumVertical())
	onS := make([]bool, spt.g.NumVertical())
	for i := 0; i < spt.g.NumVertical(); i++ {
		if !marked[i] {
			if c := findNC(g, i, marked, s, onS, spt.edgeTo); c != nil {
				return c
			}
		}
	}
	return nil
}

func findNC(g WDigraph, v int, marked []bool, s *stack.Stack, onS []bool, edgeTo []*Edge) *stack.Stack {
	if onS[v] {
		c := stack.New(g.NumVertical())
		var weight float64
		for {
			e, ok := s.Pop()
			if !ok {
				break
			}
			eg := e.(*Edge)
			weight += eg.weight
			c.Push(eg)
		}
		if weight < 0 {
			return c
		}
		for {
			e, ok := c.Pop()
			if !ok {
				break
			}
			s.Push(e)
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

func bellmanFordRelax(g WDigraph, v int, edgeTo []*Edge, distTo []float64, needRelax *queue.IntQ, onQ []bool) {
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

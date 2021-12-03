package wdigraph

import (
	"fmt"
	"github.com/howz97/algorithm/graphs"
	pqueue "github.com/howz97/algorithm/pqueue/binaryheap"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/stack"
	"math"
	"strconv"
)

type ShortestPathTree struct {
	g      *WDigraph
	src    int
	distTo []float64
	edgeTo []int
}

func (g *WDigraph) NewShortestPathTree(src int) *ShortestPathTree {
	if !g.HasVertical(src) {
		return nil
	}
	spt := &ShortestPathTree{
		g:      g,
		src:    src,
		distTo: make([]float64, g.NumVertical()),
		edgeTo: make([]int, g.NumVertical()),
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
	src := spt.edgeTo[dst]
	if src < 0 {
		return nil
	}
	path := stack.New(spt.g.NumVertical())
	for {
		path.Push(dst)
		dst = src
		src = spt.edgeTo[dst]
		if src < 0 {
			break
		}
	}
	path.Push(spt.src)
	return path
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

func dijkstraRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, pq *pqueue.BinHeap) {
	g.RangeWAdj(v, func(adj int, w float64) bool {
		if distTo[v]+w < distTo[adj] {
			inPQ := distTo[adj] != math.Inf(1)
			edgeTo[adj] = v
			distTo[adj] = distTo[v] + w
			if inPQ {
				pq.Update(int(distTo[adj]), adj)
			} else {
				pq.Insert(int(distTo[adj]), adj)
			}
		}
		return true
	})
}

// ============================ Topological ============================

func (spt *ShortestPathTree) InitTopological() {
	order := stack.NewInt(spt.g.NumVertical())
	graphs.RevDFS(spt.g, spt.src, func(v int) bool {
		order.Push(v)
		return true
	})
	for {
		e, ok := order.Pop()
		if !ok {
			break
		}
		topologicalRelax(spt.g, e, spt.edgeTo, spt.distTo)
	}
}

func topologicalRelax(g *WDigraph, v int, edgeTo []int, distTo []float64) {
	g.RangeWAdj(v, func(adj int, w float64) bool {
		if distTo[v]+w < distTo[adj] {
			edgeTo[adj] = v
			distTo[adj] = distTo[v] + w
		}
		return true
	})
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
	g := New(spt.g.NumVertical())
	for dst, src := range spt.edgeTo {
		if src > 0 {
			g.AddEdge(src, dst, spt.g.getWeight(src, dst))
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

func findNC(g *WDigraph, v int, marked []bool, s *stack.Stack, onS []bool, edgeTo []int) *stack.Stack {
	if onS[v] {
		fmt.Println("cycle found! stack is:", s.String(), v)

		start := 0
		s.Iterate(func(x stack.T) bool {
			start++
			if x.(int) == v {
				return false
			}
			return true
		})
		s.Push(v)
		src := v
		var weight float64
		s.IterateRange(start, s.Size(), func(x stack.T) bool {
			dst := x.(int)
			weight += g.getWeight(src, dst)
			src = dst
			return true
		})
		if weight < 0 {
			return s
		}
		return nil
	}
	onS[v] = true
	s.Push(edgeTo[v])
	var nc *stack.Stack
	g.RangeAdj(v, func(a int) bool {
		if marked[a] {
			return true
		}
		//fmt.Printf("detect NC by path %d->%d \n", v, a)
		nc = findNC(g, a, marked, s, onS, edgeTo)
		return nc == nil
	})
	if nc != nil {
		return nc
	}
	onS[v] = false
	s.Pop()
	marked[v] = true
	return nil
}

func bellmanFordRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, needRelax *queue.IntQ, onQ []bool) {
	g.RangeWAdj(v, func(adj int, w float64) bool {
		if distTo[v]+w < distTo[adj] {
			edgeTo[adj] = v
			distTo[adj] = distTo[v] + w
			if !onQ[adj] {
				needRelax.PushBack(adj)
				onQ[adj] = true
			}
		}
		return true
	})
}

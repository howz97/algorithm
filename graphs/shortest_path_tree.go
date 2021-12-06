package graphs

import (
	"errors"
	"fmt"
	pqueue "github.com/howz97/algorithm/pqueue/binaryheap"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/stack"
	"math"
)

const (
	Dijkstra = iota
	Topological
	BellmanFord
)

type ShortestPathTree struct {
	src    int
	distTo []float64
	edgeTo []int
}

func newShortestPathTree(g *WDigraph, src int) *ShortestPathTree {
	spt := &ShortestPathTree{
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

func (g *WDigraph) NewShortestPathTree(src int, alg int) (*ShortestPathTree, error) {
	if !g.HasVertical(src) {
		return nil, ErrVerticalNotExist
	}
	var err error
	spt := newShortestPathTree(g, src)
	switch alg {
	case Dijkstra:
		if src, dst := g.FindNegativeEdgeFrom(src); src >= 0 {
			err = errors.New(fmt.Sprintf("negative edge %d->%d", src, dst))
		}
		spt.initDijkstra(g)
	case Topological:
		if cycle := g.FindCycleFrom(src); cycle != nil {
			err = ErrCycle{Stack: cycle}
		}
		spt.initTopological(g)
	case BellmanFord:
		err = spt.InitBellmanFord(g)
	default:
		err = errors.New(fmt.Sprintf("algorithm %v not supported", alg))
	}
	return spt, err
}

func (spt *ShortestPathTree) CanReach(dst int) bool {
	if !spt.HasVertical(dst) {
		return false
	}
	return spt.distTo[dst] != math.Inf(1)
}

func (spt *ShortestPathTree) DistanceTo(dst int) float64 {
	if !spt.HasVertical(dst) {
		return math.Inf(1)
	}
	return spt.distTo[dst]
}

func (spt *ShortestPathTree) PathTo(dst int) *stack.Stack {
	if !spt.HasVertical(dst) {
		return nil
	}
	src := spt.edgeTo[dst]
	if src < 0 {
		return nil
	}
	path := stack.New(spt.NumVertical())
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

func (spt *ShortestPathTree) NumVertical() int {
	return len(spt.distTo)
}

func (spt *ShortestPathTree) HasVertical(v int) bool {
	return v >= 0 && v < len(spt.distTo)
}

// ============================ Dijkstra ============================

func (spt *ShortestPathTree) initDijkstra(g *WDigraph) {
	pq := pqueue.NewBinHeap(g.NumVertical())
	dijkstraRelax(g, spt.src, spt.edgeTo, spt.distTo, pq)
	for !pq.IsEmpty() {
		m := pq.DelMin().(int)
		dijkstraRelax(g, m, spt.edgeTo, spt.distTo, pq)
	}
}

func dijkstraRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, pq *pqueue.BinHeap) {
	g.IterateWAdj(v, func(adj int, w float64) bool {
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

func (spt *ShortestPathTree) initTopological(g *WDigraph) {
	order := stack.NewInt(int(g.NumVertical()))
	g.IterateRDFSFromVet(spt.src, func(v int) bool {
		order.Push(v)
		return true
	})
	for {
		v, ok := order.Pop()
		if !ok {
			break
		}
		topologicalRelax(g, v, spt.edgeTo, spt.distTo)
	}
}

func topologicalRelax(g *WDigraph, v int, edgeTo []int, distTo []float64) {
	g.IterateWAdj(v, func(a int, w float64) bool {
		if distTo[v]+w < distTo[a] {
			edgeTo[a] = v
			distTo[a] = distTo[v] + w
		}
		return true
	})
}

// ============================ BellmanFord ============================

type ErrCycle struct {
	Stack *stack.IntStack
}

func (nc ErrCycle) Error() string {
	return "weight negative cycle: " + nc.Stack.String()
}

func (spt *ShortestPathTree) InitBellmanFord(g *WDigraph) error {
	needRelax := queue.NewIntQ()
	onQ := make([]bool, spt.NumVertical())
	needRelax.PushBack(spt.src)
	onQ[spt.src] = true
	relaxTimes := uint(0)
	for !needRelax.IsEmpty() {
		v := needRelax.Front()
		onQ[spt.src] = false
		bellmanFordRelax(g, v, spt.edgeTo, spt.distTo, needRelax, onQ)
		relaxTimes++
		if relaxTimes%g.NumVertical() == 0 {
			if c := spt.findNegativeCycle(g); c.Size() > 0 {
				return ErrCycle{Stack: c}
			}
		}
	}
	return nil
}

func (spt *ShortestPathTree) findNegativeCycle(g *WDigraph) *stack.IntStack {
	return g.AnyNegativeCycle() // todo: optimize ?
}

func bellmanFordRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, needRelax *queue.IntQ, onQ []bool) {
	g.IterateWAdj(v, func(adj int, w float64) bool {
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

package graphs

import (
	"fmt"
	"github.com/howz97/algorithm/pq/heap"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/stack"
	"github.com/howz97/algorithm/util"
	"math"
)

const (
	Dijkstra = iota
	Topological
	BellmanFord
)

type Path struct {
	stk      *stack.IntStack
	distance float64
}

func (p *Path) String() string {
	if p == nil {
		return "path not exist"
	}
	str := fmt.Sprintf("(distance=%v): ", p.distance)
	str += fmt.Sprint(p.stk.Pop())
	for p.stk.Size() > 0 {
		str += fmt.Sprint("->", p.stk.Pop())
	}
	return str
}

func (p *Path) Symbol(s *Symbol) string {
	if p == nil {
		return "path not exist"
	}
	str := fmt.Sprintf("(distance=%v): ", p.distance)
	str += fmt.Sprint(s.SymbolOf(p.stk.Pop()))
	for p.stk.Size() > 0 {
		str += fmt.Sprint("->", s.SymbolOf(p.stk.Pop()))
	}
	return str
}

type PathTree struct {
	src    int
	distTo []float64
	edgeTo []int
}

func newShortestPathTree(g *WDigraph, src int) *PathTree {
	spt := &PathTree{
		src:    src,
		distTo: make([]float64, g.NumVert()),
		edgeTo: make([]int, g.NumVert()),
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

func (spt *PathTree) CanReach(dst int) bool {
	if !spt.HasVertical(dst) {
		return false
	}
	return spt.distTo[dst] != math.Inf(1)
}

func (spt *PathTree) DistanceTo(dst int) float64 {
	if !spt.HasVertical(dst) {
		return math.Inf(1)
	}
	return spt.distTo[dst]
}

func (spt *PathTree) PathTo(dst int) *Path {
	if !spt.HasVertical(dst) {
		return nil
	}
	src := spt.edgeTo[dst]
	if src < 0 {
		return nil
	}
	path := stack.NewInt(spt.NumVertical())
	for {
		path.Push(dst)
		dst = src
		src = spt.edgeTo[dst]
		if src < 0 {
			break
		}
	}
	path.Push(spt.src)
	return &Path{
		stk:      path,
		distance: spt.distTo[dst],
	}
}

func (spt *PathTree) NumVertical() int {
	return len(spt.distTo)
}

func (spt *PathTree) HasVertical(v int) bool {
	return v >= 0 && v < len(spt.distTo)
}

func (spt *PathTree) initDijkstra(g *WDigraph) {
	pq := heap.New(g.NumVert())
	dijkstraRelax(g, spt.src, spt.edgeTo, spt.distTo, pq)
	for !pq.IsEmpty() {
		m := pq.Pop().(int)
		dijkstraRelax(g, m, spt.edgeTo, spt.distTo, pq)
	}
}

func dijkstraRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, pq *heap.Heap) {
	g.iterateWAdj(v, func(adj int, w float64) bool {
		if distTo[v]+w < distTo[adj] {
			inPQ := distTo[adj] != math.Inf(1)
			edgeTo[adj] = v
			distTo[adj] = distTo[v] + w
			if inPQ {
				pq.Fix(util.Float(distTo[adj]), adj)
			} else {
				pq.Push(util.Float(distTo[adj]), adj)
			}
		}
		return true
	})
}

func (spt *PathTree) initTopological(g *WDigraph) {
	order := stack.NewInt(int(g.NumVert()))
	g.IterateRDFSFrom(spt.src, func(v int) bool {
		order.Push(v)
		return true
	})
	for order.Size() > 0 {
		v := order.Pop()
		topologicalRelax(g, v, spt.edgeTo, spt.distTo)
	}
}

func topologicalRelax(g *WDigraph, v int, edgeTo []int, distTo []float64) {
	g.iterateWAdj(v, func(a int, w float64) bool {
		if distTo[v]+w < distTo[a] {
			edgeTo[a] = v
			distTo[a] = distTo[v] + w
		}
		return true
	})
}

type ErrCycle struct {
	Stack *stack.IntStack
}

func (nc ErrCycle) Error() string {
	return "weight negative cycle: " + nc.Stack.String()
}

func (spt *PathTree) InitBellmanFord(g *WDigraph) error {
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
		if relaxTimes%g.NumVert() == 0 {
			if c := spt.findNegativeCycle(g); c.Size() > 0 {
				return ErrCycle{Stack: c}
			}
		}
	}
	return nil
}

func (spt *PathTree) findNegativeCycle(g *WDigraph) *stack.IntStack {
	return g.AnyNegativeCycle() // todo: optimize ?
}

func bellmanFordRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, needRelax *queue.IntQ, onQ []bool) {
	g.iterateWAdj(v, func(adj int, w float64) bool {
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

type Searcher struct {
	spt []*PathTree
}

func (s *Searcher) GetDistance(src, dst int) float64 {
	if !s.HasVertical(src) {
		return math.Inf(1)
	}
	return s.spt[src].DistanceTo(dst)
}

func (s *Searcher) GetPath(src, dst int) *Path {
	if !s.HasVertical(src) {
		return nil
	}
	return s.spt[src].PathTo(dst)
}

func (s *Searcher) HasVertical(v int) bool {
	return v >= 0 && v < len(s.spt)
}

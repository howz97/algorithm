package graphs

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/basic/stack"
	"github.com/howz97/algorithm/pq/heap"
)

const (
	Auto = iota
	Dijkstra
	Topological
	BellmanFord
)

func NewWDigraph(size uint) *WDigraph {
	return &WDigraph{
		Digraph: NewDigraph(size),
	}
}

// WDigraph is edge weighted digraph
type WDigraph struct {
	*Digraph
}

// AddEdge add a weighted and directed edge
func (g *WDigraph) AddEdge(src, dst int, w float64) {
	g.addWeightedEdge(src, dst, w)
}

// ShortestPathTree get the shortest path tree from src by the specified algorithm
func (g *WDigraph) ShortestPathTree(src int, alg int) (*PathTree, error) {
	if !g.HasVert(src) {
		return nil, ErrVerticalNotExist
	}
	var err error
	spt := newShortestPathTree(g, src)
	switch alg {
	case Dijkstra:
		if src, dst := g.FindNegativeEdgeFrom(src); src >= 0 {
			err = errors.New(fmt.Sprintf("negative edge %d->%d", src, dst))
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
		err = spt.initBellmanFord(g)
	default:
		if g.FindCycleFrom(src) == nil {
			spt.initTopological(g)
		} else if src, _ := g.FindNegativeEdgeFrom(src); src < 0 {
			spt.initDijkstra(g)
		} else {
			err = spt.initBellmanFord(g)
		}
	}
	return spt, err
}

// PathTree is shortest path tree
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

// CanReach check whether src can reach dst
func (spt *PathTree) CanReach(dst int) bool {
	return spt.distTo[dst] != math.Inf(1)
}

// DistanceTo get the distance from src to dst
func (spt *PathTree) DistanceTo(dst int) float64 {
	return spt.distTo[dst]
}

// PathTo get the path from src to dst
func (spt *PathTree) PathTo(dst int) *Path {
	src := spt.edgeTo[dst]
	if src < 0 {
		return nil
	}
	path := NewPath()
	for {
		path.Push(src, dst, 0)
		dst = src
		src = spt.edgeTo[dst]
		if src < 0 {
			break
		}
	}
	path.Reverse()
	return path
}

func (spt *PathTree) initDijkstra(g *WDigraph) {
	pq := heap.New2[float64, int](g.NumVert())
	dijkstraRelax(g, spt.src, spt.edgeTo, spt.distTo, pq)
	for pq.Size() > 0 {
		m := pq.Pop()
		dijkstraRelax(g, m, spt.edgeTo, spt.distTo, pq)
	}
}

func dijkstraRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, pq *heap.Heap2[float64, int]) {
	g.IterWAdjacent(v, func(adj int, w float64) bool {
		if distTo[v]+w < distTo[adj] {
			inPQ := distTo[adj] != math.Inf(1)
			edgeTo[adj] = v
			distTo[adj] = distTo[v] + w
			if inPQ {
				pq.Fix(distTo[adj], adj)
			} else {
				pq.Push(distTo[adj], adj)
			}
		}
		return true
	})
}

func (spt *PathTree) initTopological(g *WDigraph) {
	order := stack.New[int](int(g.NumVert()))
	g.IterBDFSFrom(spt.src, func(v int) bool {
		order.Push(v)
		return true
	})
	for order.Size() > 0 {
		v := order.Pop()
		topologicalRelax(g, v, spt.edgeTo, spt.distTo)
	}
}

func topologicalRelax(g *WDigraph, v int, edgeTo []int, distTo []float64) {
	g.IterWAdjacent(v, func(a int, w float64) bool {
		if distTo[v]+w < distTo[a] {
			edgeTo[a] = v
			distTo[a] = distTo[v] + w
		}
		return true
	})
}

func (spt *PathTree) initBellmanFord(g *WDigraph) error {
	q := queue.NewLinkQ[int]()
	onQ := make([]bool, g.NumVert())
	q.PushBack(spt.src)
	onQ[spt.src] = true
	cnt := uint(0)
	for q.Size() > 0 {
		v := q.Front()
		onQ[v] = false
		bellmanFordRelax(g, v, spt.edgeTo, spt.distTo, q, onQ)
		cnt++
		if cnt%g.NumVert() == 0 {
			c := spt.toWDigraph(g).FindCycle()
			if c != nil {
				return c
			}
		}
	}
	return nil
}

func (spt *PathTree) toWDigraph(g *WDigraph) *WDigraph {
	sptg := NewWDigraph(g.NumVert())
	for to, from := range spt.edgeTo {
		if from < 0 {
			continue
		}
		sptg.AddEdge(from, to, g.GetWeight(from, to))
	}
	return sptg
}

func bellmanFordRelax(g *WDigraph, v int, edgeTo []int, distTo []float64, q *queue.LinkQ[int], onQ []bool) {
	g.IterWAdjacent(v, func(adj int, w float64) bool {
		if distTo[v]+w < distTo[adj] {
			edgeTo[adj] = v
			distTo[adj] = distTo[v] + w
			if !onQ[adj] {
				q.PushBack(adj)
				onQ[adj] = true
			}
		}
		return true
	})
}

func (g *WDigraph) SearcherDijkstra() (*Searcher, error) {
	if src, dst := g.FindNegativeEdge(); src >= 0 {
		return nil, errors.New(fmt.Sprintf("negative edge %d->%d", src, dst))
	}
	sps := &Searcher{
		spt: make([]*PathTree, g.NumVert()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		tree.initDijkstra(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph) SearcherTopological() (*Searcher, error) {
	if c := g.FindCycle(); c != nil {
		return nil, c
	}
	sps := &Searcher{
		spt: make([]*PathTree, g.NumVert()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		tree.initTopological(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph) SearcherBellmanFord() (*Searcher, error) {
	sps := &Searcher{
		spt: make([]*PathTree, g.NumVert()),
	}
	var err error
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		err = tree.initBellmanFord(g)
		if err != nil {
			return nil, err
		}
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph) Searcher() (*Searcher, error) {
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

type Searcher struct {
	spt []*PathTree
}

func (s *Searcher) GetDistance(src, dst int) float64 {
	return s.spt[src].DistanceTo(dst)
}

func (s *Searcher) GetPath(src, dst int) *Path {
	return s.spt[src].PathTo(dst)
}

func NewPath() *Path {
	return &Path{
		Stack: stack.New[edge](2),
	}
}

type Path struct {
	*stack.Stack[edge]
}

func (p *Path) Push(from, to int, w float64) {
	p.Stack.Push(edge{
		from:   from,
		to:     to,
		weight: w,
	})
}

func (p *Path) Pop() edge {
	e := p.Stack.Pop()
	return e
}

func (p *Path) Distance() float64 {
	d := 0.0
	p.Iterate(func(e edge) bool {
		d += e.weight
		return true
	})
	return d
}

func (p *Path) ContainsVert(v int) bool {
	if p.Size() <= 0 {
		return false
	}
	found := false
	p.Iterate(func(e edge) bool {
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

func (p *Path) Str(s *Symbol) string {
	if p == nil || p.Size() <= 0 {
		return "path not exist"
	}
	var i2s func(int) string
	if s == nil {
		i2s = strconv.Itoa
	} else {
		i2s = s.SymbolOf
	}
	str := fmt.Sprintf("(distance=%v): ", p.Distance())
	p.Iterate(func(e edge) bool {
		str += e.string(i2s) + ", "
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
	i := p.Index(func(v edge) bool {
		return v.from == x
	})
	path := NewPath()
	p.IterateRange(i, p.Size(), func(e edge) bool {
		path.Push(e.from, e.to, e.weight)
		return true
	})
	return &Cycle{path}
}

type Cycle struct {
	*Path
}

func (c *Cycle) Error() string {
	return c.Str(nil)
}

func (c *Cycle) Cycle() *Cycle {
	return c
}

type edge struct {
	from, to int
	weight   float64
}

func (e *edge) string(i2s func(int) string) string {
	return i2s(e.from) + "->" + i2s(e.to)
}

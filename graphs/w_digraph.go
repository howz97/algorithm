package graphs

import (
	"errors"
	"fmt"
	"github.com/howz97/algorithm/stack"
)

// WDigraph is edge weighted digraph without self loop
type WDigraph struct {
	Digraph
}

func NewWDigraph(size uint) *WDigraph {
	return &WDigraph{
		Digraph: *NewDigraph(size),
	}
}

func (g *WDigraph) AddEdge(src, dst int, w float64) error {
	return g.addWeightedEdge(src, dst, w)
}

func (g *WDigraph) String() string {
	bytes, err := g.Marshal()
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (g *WDigraph) FindNegativeEdgeFrom(from int) (src int, dst int) {
	g.IterateWEdgeFrom(from, func(v0 int, v1 int, w float64) bool {
		if w < 0 {
			src = v0
			dst = v1
			return false
		}
		return true
	})
	return -1, -1
}

func (g *WDigraph) AnyNegativeCycle() *stack.IntStack {
	marked := make([]bool, g.NumVertical())
	path := stack.NewInt(4)
	g.IterateWEdge(func(src int, dst int, w float64) bool {
		if w < 0 {
			if !marked[src] {
				if g.DetectCycleDFS(src, marked, path) {
					return false
				}
			}
		}
		return true
	})
	return path
}

func (g *WDigraph) NewShortestPathTree(src int, alg int) (*PathTree, error) {
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

func (g *WDigraph) SearcherDijkstra() (*Searcher, error) {
	if src, dst := g.FindNegativeEdge(); src >= 0 {
		return nil, errors.New(fmt.Sprintf("negative edge %d->%d", src, dst))
	}
	sps := &Searcher{
		spt: make([]*PathTree, g.NumVertical()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		tree.initDijkstra(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph) SearcherTopological() (*Searcher, error) {
	if cycle := g.FindCycle(); cycle != nil {
		return nil, ErrCycle{Stack: cycle}
	}
	sps := &Searcher{
		spt: make([]*PathTree, g.NumVertical()),
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
		spt: make([]*PathTree, g.NumVertical()),
	}
	var err error
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		err = tree.InitBellmanFord(g)
		if err != nil {
			return nil, err
		}
		sps.spt[v] = tree
	}
	return sps, nil
}

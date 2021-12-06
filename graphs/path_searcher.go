package graphs

import (
	"errors"
	"fmt"
	"github.com/howz97/algorithm/stack"
	"math"
)

type PathSearcher struct {
	spt []*ShortestPathTree
}

func (g *WDigraph) SearcherDijkstra() (*PathSearcher, error) {
	if src, dst := g.FindNegativeEdge(); src >= 0 {
		return nil, errors.New(fmt.Sprintf("negative edge %d->%d", src, dst))
	}
	sps := &PathSearcher{
		spt: make([]*ShortestPathTree, g.NumVertical()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		tree.initDijkstra(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph) SearcherTopological() (*PathSearcher, error) {
	if cycle := g.FindCycle(); cycle != nil {
		return nil, ErrCycle{Stack: cycle}
	}
	sps := &PathSearcher{
		spt: make([]*ShortestPathTree, g.NumVertical()),
	}
	for v := range sps.spt {
		tree := newShortestPathTree(g, v)
		tree.initTopological(g)
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g *WDigraph) SearcherBellmanFord() (*PathSearcher, error) {
	sps := &PathSearcher{
		spt: make([]*ShortestPathTree, g.NumVertical()),
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

func (s *PathSearcher) Distance(src, dst int) float64 {
	if !s.HasVertical(src) && !s.HasVertical(dst) {
		return math.Inf(1)
	}
	return s.spt[src].DistanceTo(dst)
}

func (s *PathSearcher) Path(src, dst int) *stack.Stack {
	if !s.HasVertical(src) && !s.HasVertical(dst) {
		return nil
	}
	return s.spt[src].PathTo(dst)
}

func (s *PathSearcher) PrintPath(src, dst int) {
	if src == dst {
		return
	}
	p := s.Path(src, dst)
	if p == nil {
		fmt.Printf("%d can not access %d\n", src, dst)
		return
	}
	p.Pop()
	fmt.Print("PATH: ", src)
	for {
		v, ok := p.Pop()
		if !ok {
			break
		}
		fmt.Print("->", v.(int))
	}
	fmt.Printf(" (distance %v)\n", s.Distance(src, dst))
}

func (s *PathSearcher) HasVertical(v int) bool {
	return v >= 0 && v < len(s.spt)
}

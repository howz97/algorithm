package weighted_digraph

import (
	"errors"
	"howz97/algorithm/stack"
)

type ShortestPathSearcher struct {
	g   EdgeWeightedDigraph
	spt []*ShortestPathTree
}

func (g EdgeWeightedDigraph) GenSearcherDijkstra() (*ShortestPathSearcher, error) {
	if g.HasNegativeEdge() {
		return nil, errors.New("this digraph contains negative edge")
	}
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumV()),
	}
	for v := range sps.spt {
		tree := g.NewShortestPathTree(v)
		tree.InitDijkstra()
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g EdgeWeightedDigraph) GenSearcherTopological() (*ShortestPathSearcher, error) {
	if g.DetectDirCycle() {
		return nil, errors.New("this digraph contains directed cycle")
	}
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumV()),
	}
	for v := range sps.spt {
		tree := g.NewShortestPathTree(v)
		tree.InitTopological()
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g EdgeWeightedDigraph) GenSearcherBellmanFord() (*ShortestPathSearcher, error) {
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumV()),
	}
	var err error
	for v := range sps.spt {
		tree := g.NewShortestPathTree(v)
		err = tree.InitBellmanFord()
		if err != nil {
			return nil, err
		}
		sps.spt[v] = tree
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

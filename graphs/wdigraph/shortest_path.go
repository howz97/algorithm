package wdigraph

import (
	"errors"
	"fmt"
	"github.com/howz97/algorithm/stack"
)

type ShortestPathSearcher struct {
	g   WDigraph
	spt []*ShortestPathTree
}

func (g WDigraph) GenSearcherDijkstra() (*ShortestPathSearcher, error) {
	if g.HasNegativeEdge() {
		return nil, errors.New("this digraph contains negative edge")
	}
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumVertical()),
	}
	for v := range sps.spt {
		tree := g.NewShortestPathTree(v)
		tree.InitDijkstra()
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g WDigraph) GenSearcherTopological() (*ShortestPathSearcher, error) {
	if g.DetectDirCycle() {
		return nil, errors.New("this digraph contains directed cycle")
	}
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumVertical()),
	}
	for v := range sps.spt {
		tree := g.NewShortestPathTree(v)
		tree.InitTopological()
		sps.spt[v] = tree
	}
	return sps, nil
}

func (g WDigraph) GenSearcherBellmanFord() (*ShortestPathSearcher, error) {
	sps := &ShortestPathSearcher{
		g:   g,
		spt: make([]*ShortestPathTree, g.NumVertical()),
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
	if !s.g.HasVertical(src) && !s.g.HasVertical(dst) {
		panic(ErrVerticalNotExist)
	}
	return s.spt[src].DistanceTo(dst)
}

func (s *ShortestPathSearcher) Path(src, dst int) *stack.Stack {
	if !s.g.HasVertical(src) && !s.g.HasVertical(dst) {
		panic(ErrVerticalNotExist)
	}
	return s.spt[src].PathTo(dst)
}

func (s *ShortestPathSearcher) PrintPath(src, dst int) {
	p := s.Path(src, dst)
	fmt.Print("PATH: ", src)
	for {
		e, ok := p.Pop()
		if !ok {
			break
		}
		eg := e.(*Edge)
		fmt.Print("->", eg.to)
	}
	fmt.Printf(" (distance %v)\n", s.Distance(src, dst))
}

package wdigraph

import (
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/stack"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// WDigraph is edge weighted digraph without self loop
type WDigraph struct {
	digraph.Digraph
	graphs.Weight
}

func New(size int) *WDigraph {
	return &WDigraph{
		Digraph: digraph.New(size),
		Weight:  graphs.NewWeight(size),
	}
}

func LoadWDigraph(filename string) (*WDigraph, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var m map[int]map[int]float64
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		return nil, err
	}
	g := New(len(m))
	for src, adj := range m {
		for dst, w := range adj {
			err = g.AddEdge(src, dst, w)
			if err != nil {
				return nil, err
			}
		}
	}
	return g, nil
}

func (g *WDigraph) AddEdge(src, dst int, w float64) error {
	err := g.Digraph.AddEdge(src, dst)
	if err != nil {
		return err
	}
	g.SetWeight(src, dst, w)
	return nil
}

// RangeWAdj range adjacent vertices of v
func (g *WDigraph) RangeWAdj(v int, fn func(int, float64) bool) {
	g.RangeAdj(v, func(a int) bool {
		return fn(a, g.GetWeight(v, a))
	})
}

func (g *WDigraph) String() string {
	bytes, err := graphs.MarshalWGraph(g)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (g *WDigraph) IterateEdge(fn func(int, int, float64) bool) {
	g.Digraph.IterateEdge(func(src int, dst int) bool {
		return fn(src, dst, g.GetWeight(src, dst))
	})
}

func (g *WDigraph) IterateEdgeFrom(v int, fn func(int, int, float64) bool) {
	g.Digraph.IterateEdgeFrom(v, func(src int, dst int) bool {
		return fn(src, dst, g.GetWeight(src, dst))
	})
}

func (g *WDigraph) FindNegativeEdgeFrom(from int) (src int, dst int) {
	g.IterateEdgeFrom(from, func(v0 int, v1 int, w float64) bool {
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
	g.IterateEdge(func(src int, dst int, w float64) bool {
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

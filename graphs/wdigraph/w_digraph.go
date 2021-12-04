package wdigraph

import (
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/stack"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// WDigraph is edge weighted digraph without self loop
type WDigraph struct {
	digraph.Digraph
	weight []*hash_map.Chaining
}

func New(size int) *WDigraph {
	weight := make([]*hash_map.Chaining, size)
	for i := range weight {
		weight[i] = hash_map.New()
	}
	return &WDigraph{
		Digraph: digraph.New(size),
		weight:  weight,
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
	g.weight[src].Put(search.Integer(dst), w)
	return nil
}

// RangeWAdj range adjacent vertices of v
func (g *WDigraph) RangeWAdj(v int, fn func(int, float64) bool) {
	g.RangeAdj(v, func(adj int) bool {
		w := g.weight[v].Get(search.Integer(adj)).(float64)
		return fn(adj, w)
	})
}

func (g *WDigraph) FindNegativeEdge() (src, dst int) {
	src, dst = -1, -1
	for v0, hm := range g.weight {
		hm.Range(func(key hash_map.Key, val search.T) bool {
			if val.(float64) < 0 {
				src = v0
				dst = int(key.(search.Integer))
				return false
			}
			return true
		})
	}
	return
}

func (g *WDigraph) GetWeight(src, dst int) float64 {
	return g.weight[src].Get(search.Integer(dst)).(float64)
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

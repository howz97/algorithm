package graphs

import (
	"github.com/howz97/algorithm/stack"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// WDigraph is edge weighted digraph without self loop
type WDigraph struct {
	Digraph
}

func NewWDigraph(size uint) *WDigraph {
	return &WDigraph{
		Digraph: NewDigraph(size),
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
	g := NewWDigraph(uint(len(m)))
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
	return g.addWeightedEdge(src, dst, w)
}

func (g *WDigraph) String() string {
	bytes, err := MarshalWGraph(g)
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

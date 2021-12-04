package wdigraph

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/stack"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func ImportEWD(filename string) (*WDigraph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scan := bufio.NewScanner(file)
	if !scan.Scan() {
		return nil, errors.New("eof")
	}
	headLine := scan.Text()
	headSli := strings.Split(headLine, " ")
	numV, err := strconv.Atoi(headSli[0])
	if err != nil {
		return nil, err
	}
	ewd := New(numV)
	for scan.Scan() {
		lineSli := strings.Split(scan.Text(), " ")
		from, err := strconv.Atoi(lineSli[0])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(lineSli[1])
		if err != nil {
			return nil, err
		}
		weight, err := strconv.ParseFloat(lineSli[2], 64)
		if err != nil {
			fmt.Println("can not parse float", lineSli[2])
			return nil, err
		}
		ewd.AddEdge(from, to, weight)
	}
	return ewd, nil
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

func (g *WDigraph) HasNegativeEdge() bool {
	found := false
	for _, hm := range g.weight {
		hm.Range(func(_ hash_map.Key, val search.T) bool {
			if val.(float64) < 0 {
				found = true
				return false
			}
			return true
		})
	}
	return found
}

func (g *WDigraph) getWeight(src, dst int) float64 {
	if !g.HasEdge(src, dst) {
		panic(fmt.Sprintf("edge %d->%d not exist: %s", src, dst, g.String()))
	}
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
		return fn(src, dst, g.getWeight(src, dst))
	})
}

func (g *WDigraph) IterateEdgeFrom(v int, fn func(int, int, float64) bool) {
	g.Digraph.IterateEdgeFrom(v, func(src int, dst int) bool {
		return fn(src, dst, g.getWeight(src, dst))
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

type Edge struct {
	from, to int
	weight   float64
}

func (e *Edge) From() int {
	return e.from
}

func (e *Edge) To() int {
	return e.to
}

func (e *Edge) GetWeight() float64 {
	return e.weight
}

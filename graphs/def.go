package graphs

import (
	"errors"
	"fmt"
	"github.com/howz97/algorithm/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
	ErrInputFormat      = errors.New("input format error")
)

func readYaml(filename string, isSymbol bool) (*Digraph, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var g *Digraph
	if isSymbol {
		var m map[string]map[string]float64
		yaml.Unmarshal(file, &m)
		if err != nil {
			return nil, err
		}
		g = &Digraph{Symbol: NewSymbolGraph()}
		for v := range m {
			g.Symbol.scanVertical(v)
		}
		for from, adj := range m {
			for to, w := range adj {
				err = g.addWeightedEdge(g.Symbol.str2int[from], g.Symbol.str2int[to], w)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		var m map[int]map[int]float64
		err = yaml.Unmarshal(file, &m)
		if err != nil {
			return nil, err
		}
		g = NewDigraph(uint(len(m)))
		for from, adj := range m {
			for to, w := range adj {
				err = g.addWeightedEdge(from, to, w)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return g, nil
}

func LoadDigraph(filename string, isSymbol bool) (*Digraph, error) {
	dg, err := readYaml(filename, isSymbol)
	if err != nil {
		return nil, err
	}
	dropWeight(dg)
	return dg, nil
}

func LoadGraph(filename string, isSymbol bool) (*Graph, error) {
	dg, err := readYaml(filename, isSymbol)
	if err != nil {
		return nil, err
	}
	err = checkNoDirection(dg)
	if err != nil {
		return nil, err
	}
	dropWeight(dg)
	return &Graph{Digraph: *dg}, nil
}

func dropWeight(dg *Digraph) {
	dg.IterateEdge(func(from int, to int) bool {
		dg.Edges[from].Put(util.Int(to), 1)
		return true
	})
}

func checkNoDirection(dg *Digraph) error {
	var err error
	dg.IterateWEdge(func(from, to int, w float64) bool {
		wr, ok := dg.GetWeight(to, from)
		if !ok {
			err = errors.New(fmt.Sprintf("edge %d->%d has direction", from, to))
			return false
		}
		if wr != w {
			err = errors.New(fmt.Sprintf("edge %d->%d has weight %v, but %d->%d has weight %v",
				from, to, w, to, from, wr))
			return false
		}
		return true
	})
	return err
}

func LoadWDigraph(filename string, isSymbol bool) (*WDigraph, error) {
	dg, err := readYaml(filename, isSymbol)
	if err != nil {
		return nil, err
	}
	return &WDigraph{Digraph: *dg}, nil
}

func LoadWGraph(filename string, isSymbol bool) (*WGraph, error) {
	dg, err := readYaml(filename, isSymbol)
	if err != nil {
		return nil, err
	}
	err = checkNoDirection(dg)
	if err != nil {
		return nil, err
	}
	return &WGraph{Graph{Digraph: *dg}}, nil
}

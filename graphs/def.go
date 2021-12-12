package graphs

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/howz97/algorithm/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
	ErrInvalidYaml      = errors.New("input yaml file is invalid")
)

func readYaml(filename string) (*Digraph, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	isSymbol := false
	if bytes.IndexByte(content, '"') >= 0 {
		isSymbol = true
	}
	var g *Digraph
	if isSymbol {
		var m map[string]map[string]float64
		err = yaml.Unmarshal(content, &m)
		if err != nil {
			return nil, err
		}
		if len(m) == 0 {
			return nil, ErrInvalidYaml
		}
		g = NewDigraph(uint(len(m)))
		g.Symbol = NewSymbolGraph()
		for v := range m {
			g.scanVertical(v)
		}
		for from, adj := range m {
			for to, w := range adj {
				err = g.addWeightedEdge(g.syb2vet[from], g.syb2vet[to], w)
				if err != nil {
					return nil, err
				}
			}
		}
	} else {
		var m map[int]map[int]float64
		err = yaml.Unmarshal(content, &m)
		if err != nil {
			return nil, err
		}
		if len(m) == 0 {
			return nil, ErrInvalidYaml
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

func LoadDigraph(filename string) (*Digraph, error) {
	dg, err := readYaml(filename)
	if err != nil {
		return nil, err
	}
	dropWeight(dg)
	return dg, nil
}

func LoadGraph(filename string) (*Graph, error) {
	dg, err := readYaml(filename)
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

func LoadWDigraph(filename string) (*WDigraph, error) {
	dg, err := readYaml(filename)
	if err != nil {
		return nil, err
	}
	return &WDigraph{Digraph: *dg}, nil
}

func LoadWGraph(filename string) (*WGraph, error) {
	dg, err := readYaml(filename)
	if err != nil {
		return nil, err
	}
	err = checkNoDirection(dg)
	if err != nil {
		return nil, err
	}
	return &WGraph{Graph{Digraph: *dg}}, nil
}

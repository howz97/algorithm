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

type IGraph interface {
	HasVertical(v int) bool
	NumVertical() uint
	IterateAdj(v int, fn func(a int) bool)
	AddEdge(v1, v2 int) error
	HasEdge(v1, v2 int) bool
}

func ReadYaml(filename string) (Digraph, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var m map[int]map[int]float64
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		return nil, err
	}

	g := NewDigraph(uint(len(m)))
	for from, adj := range m {
		for to, w := range adj {
			err = g.addWeightedEdge(from, to, w)
			if err != nil {
				return nil, err
			}
		}
	}
	return g, nil
}

func LoadDigraph(filename string) (Digraph, error) {
	dg, err := ReadYaml(filename)
	if err != nil {
		return nil, err
	}
	dropWeight(dg)
	return dg, nil
}

func LoadGraph(filename string) (*Graph, error) {
	dg, err := ReadYaml(filename)
	if err != nil {
		return nil, err
	}
	err = checkNoDirection(dg)
	if err != nil {
		return nil, err
	}
	dropWeight(dg)
	return &Graph{Digraph: dg}, nil
}

func dropWeight(dg Digraph) {
	dg.IterateEdge(func(from int, to int) bool {
		dg[from].Put(util.Int(to), 1)
		return true
	})
}

func checkNoDirection(dg Digraph) error {
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
	dg, err := ReadYaml(filename)
	if err != nil {
		return nil, err
	}
	return &WDigraph{Digraph: dg}, nil
}

func LoadWGraph(filename string) (*WGraph, error) {
	dg, err := ReadYaml(filename)
	if err != nil {
		return nil, err
	}
	err = checkNoDirection(dg)
	if err != nil {
		return nil, err
	}
	return &WGraph{Graph{Digraph: dg}}, nil
}
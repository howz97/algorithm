package digraph

import (
	"errors"
	"github.com/howz97/algorithm/set"
	"strconv"
)

var (
	ErrVertexNotExist     = errors.New("vertex not exist")
	ErrNotSupportSelfLoop = errors.New("not support self loop")
)

type Digraph []set.IntSet

func New(numV int) Digraph {
	g := make(Digraph, numV)
	for i := range g {
		g[i] = make(set.IntSet)
	}
	return g
}

func NewBy2DSli(sli [][]int) (Digraph, error) {
	dg := New(len(sli))
	var err error
	for src, s := range sli {
		for _, dst := range s {
			err = dg.AddEdge(src, dst)
			if err != nil {
				return nil, err
			}
		}
	}
	return dg, nil
}

func (dg Digraph) NumVertical() int {
	return len(dg)
}

func (dg Digraph) HasVertical(v int) bool {
	return v >= 0 && v < dg.NumVertical()
}

func (dg Digraph) NumEdge() int {
	nume := 0
	for i := range dg {
		nume += dg[i].Len()
	}
	return nume
}

func (dg Digraph) AddEdge(v1, v2 int) error {
	if !dg.HasVertical(v1) || !dg.HasVertical(v2) {
		return ErrVertexNotExist
	}
	if v1 == v2 {
		return ErrNotSupportSelfLoop
	}
	dg[v1].Add(v2)
	return nil
}

func (dg Digraph) HasEdge(v1, v2 int) bool {
	if !dg.HasVertical(v1) || !dg.HasVertical(v2) {
		panic(ErrVertexNotExist)
	}
	return dg[v1].Contains(v2)
}

func (dg Digraph) Adjacent(v int) []int {
	if !dg.HasVertical(v) {
		panic(ErrVertexNotExist)
	}
	return dg[v].Traverse()
}

func (dg Digraph) RangeAdj(v int, fn func(v int) bool) {
	if !dg.HasVertical(v) {
		return
	}
	dg[v].Range(fn)
}

func (dg Digraph) String() string {
	out := ""
	for i := range dg {
		out += strconv.Itoa(i) + " :"
		adj := dg.Adjacent(i)
		for _, j := range adj {
			out += " " + strconv.Itoa(j)
		}
		out += "\n"
	}
	out += "\n"
	return out
}

func (dg Digraph) Reverse() Digraph {
	rg := New(dg.NumVertical())
	for v := 0; v < dg.NumVertical(); v++ {
		adj := dg.Adjacent(v)
		for _, w := range adj {
			rg.AddEdge(w, v)
		}
	}
	return rg
}

func (dg Digraph) HasDir() bool {
	return true
}

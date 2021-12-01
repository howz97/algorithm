package digraph

import (
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/set"
	"github.com/howz97/algorithm/stack"
	"github.com/howz97/algorithm/util"
	"strconv"
)

type Digraph []set.IntSet

func New(size int) Digraph {
	g := make(Digraph, size)
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
	n := 0
	for i := range dg {
		n += dg[i].Len()
	}
	return n
}

func (dg Digraph) AddEdge(v1, v2 int) error {
	if !dg.HasVertical(v1) || !dg.HasVertical(v2) {
		return graphs.ErrVerticalNotExist
	}
	if v1 == v2 {
		return graphs.ErrSelfLoop
	}
	dg[v1].Add(v2)
	return nil
}

func (dg Digraph) HasEdge(v1, v2 int) bool {
	if !dg.HasVertical(v1) || !dg.HasVertical(v2) {
		return false
	}
	return dg[v1].Contains(v2)
}

func (dg Digraph) Adjacent(v int) []int {
	if !dg.HasVertical(v) {
		return nil
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
		dg.RangeAdj(i, func(j int) bool {
			out += " " + strconv.Itoa(j)
			return true
		})
		out += "\n"
	}
	out += "\n"
	return out
}

func (dg Digraph) Reverse() Digraph {
	rg := New(dg.NumVertical())
	for v := 0; v < dg.NumVertical(); v++ {
		dg.RangeAdj(v, func(w int) bool {
			rg.AddEdge(w, v)
			return true
		})
	}
	return rg
}

func (dg Digraph) HasCycle() bool {
	return dg.getCycle() != nil
}

func (dg Digraph) GetCycle() []int {
	pathStack := dg.getCycle()
	if pathStack == nil {
		return nil
	}
	path := make([]int, 0, pathStack.Size())
	w, _ := pathStack.Pop()
	path = append(path, w)
	for {
		v, _ := pathStack.Pop()
		path = append(path, v)
		if v == w {
			break
		}
	}
	util.ReverseInts(path)
	return path
}

func (dg Digraph) getCycle() *stack.IntStack {
	marks := make([]bool, dg.NumVertical())
	path := stack.NewInt(4)
	for v, m := range marks {
		if !m {
			if dg.detectCycleDFS(v, marks, path) {
				return path
			}
		}
	}
	return nil
}

func (dg Digraph) detectCycleDFS(v int, marked []bool, path *stack.IntStack) bool {
	path.Push(v)
	found := false
	dg.RangeAdj(v, func(w int) bool {
		if !marked[w] {
			return true
		}
		if path.Contains(w) {
			path.Push(w)
			found = true
			return false
		}
		found = dg.detectCycleDFS(w, marked, path)
		return !found
	})
	if found {
		return true
	}
	path.Pop()
	marked[v] = true
	return false
}

func (dg Digraph) ReversePostOrder() *stack.IntStack {
	order := stack.NewInt(dg.NumVertical())
	graphs.RevDFSAll(dg, func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func (dg Digraph) Topological() *stack.IntStack {
	if dg.HasCycle() {
		return nil
	}
	return dg.ReversePostOrder()
}

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

func (dg Digraph) AddEdge(src, dst int) error {
	if !dg.HasVertical(src) || !dg.HasVertical(dst) {
		return graphs.ErrVerticalNotExist
	}
	if src == dst {
		return graphs.ErrSelfLoop
	}
	dg[src].Add(dst)
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

func (dg Digraph) RangeAdj(v int, fn func(int) bool) {
	if !dg.HasVertical(v) {
		return
	}
	dg[v].Iterate(fn)
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
	path := ParseCycleInStack(pathStack)
	return path
}

func (dg Digraph) getCycle() *stack.IntStack {
	marks := make([]bool, dg.NumVertical())
	path := stack.NewInt(4)
	for v, m := range marks {
		if !m {
			if dg.DetectCycleDFS(v, marks, path) {
				return path
			}
		}
	}
	return nil
}

// IsOnCycle detect whether vertical locate on a cycle
func (dg Digraph) IsOnCycle(v int) bool {
	return dg.getCycleBy(v) != nil
}

func (dg Digraph) GetCycleBy(v int) []int {
	pathStack := dg.getCycleBy(v)
	if pathStack == nil {
		return nil
	}
	path := ParseCycleInStack(pathStack)
	return path
}

func (dg Digraph) getCycleBy(v int) *stack.IntStack {
	marks := make([]bool, dg.NumVertical())
	path := stack.NewInt(4)
	if dg.DetectCycleDFS(v, marks, path) {
		return path
	}
	return nil
}

func ParseCycleInStack(stk *stack.IntStack) []int {
	path := make([]int, 0, stk.Size())
	w, _ := stk.Pop()
	path = append(path, w)
	for {
		v, _ := stk.Pop()
		path = append(path, v)
		if v == w {
			break
		}
	}
	util.ReverseInts(path)
	return path
}

func (dg Digraph) DetectCycleDFS(v int, marked []bool, path *stack.IntStack) bool {
	path.Push(v)
	found := false
	dg.RangeAdj(v, func(w int) bool {
		if marked[w] {
			return true
		}
		if path.Contains(w) {
			path.Push(w)
			found = true
			return false
		}
		found = dg.DetectCycleDFS(w, marked, path)
		return !found
	})
	marked[v] = true
	if !found {
		path.Pop()
	}
	return found
}

// Topological return a stack that pop vertices in topological order
func (dg Digraph) Topological() *stack.IntStack {
	if dg.HasCycle() {
		return nil
	}
	order := stack.NewInt(dg.NumVertical())
	graphs.RevDFSAll(dg, func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func (dg Digraph) IterateEdge(fn func(int, int) bool) {
	for src, adj := range dg {
		goon := true
		adj.Iterate(func(dst int) bool {
			goon = fn(src, dst)
			return goon
		})
		if !goon {
			break
		}
	}
}

func (dg Digraph) IterateEdgeFrom(v int, fn func(int, int) bool) {
	graphs.RangeDFS(dg, v, func(v int) bool {
		goon := true
		dg.RangeAdj(v, func(a int) bool {
			goon = fn(v, a)
			return goon
		})
		return goon
	})
}

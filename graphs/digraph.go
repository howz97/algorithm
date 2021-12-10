package graphs

import (
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/stack"
	"github.com/howz97/algorithm/util"
	"strconv"
)

type Digraph []*hash_map.Chaining

func NewDigraph(size uint) Digraph {
	dg := make([]*hash_map.Chaining, size)
	for i := range dg {
		dg[i] = hash_map.New()
	}
	return dg
}

func NewDigraphBy2DSli(sli [][]int) (Digraph, error) {
	dg := NewDigraph(uint(len(sli)))
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

func (dg Digraph) NumVertical() uint {
	return uint(len(dg))
}

func (dg Digraph) HasVertical(v int) bool {
	return v >= 0 && v < len(dg)
}

func (dg Digraph) NumEdge() uint {
	n := uint(0)
	for i := range dg {
		n += dg[i].Size()
	}
	return n
}

func (dg Digraph) AddEdge(src, dst int) error {
	return dg.addWeightedEdge(src, dst, 1)
}

func (dg Digraph) addWeightedEdge(src, dst int, w float64) error {
	if !dg.HasVertical(src) || !dg.HasVertical(dst) {
		return ErrVerticalNotExist
	}
	if src == dst {
		return ErrSelfLoop
	}
	dg[src].Put(util.Int(dst), w)
	return nil
}

func (dg Digraph) HasEdge(v1, v2 int) bool {
	if !dg.HasVertical(v1) || !dg.HasVertical(v2) {
		return false
	}
	return dg[v1].Get(util.Int(v2)) != nil
}

func (dg Digraph) IterateAdj(v int, fn func(int) bool) {
	dg.IterateWAdj(v, func(a int, _ float64) bool {
		return fn(a)
	})
}

func (dg Digraph) IterateWAdj(v int, fn func(int, float64) bool) {
	if !dg.HasVertical(v) {
		return
	}
	dg[v].Range(func(key hash_map.Key, val util.T) bool {
		return fn(int(key.(util.Int)), val.(float64))
	})
}

func (dg Digraph) Adjacent(v int) (adj []int) {
	dg.IterateAdj(v, func(a int) bool {
		adj = append(adj, a)
		return true
	})
	return adj
}

func (dg Digraph) String() string {
	out := ""
	for i := range dg {
		out += strconv.Itoa(i) + " :"
		dg.IterateAdj(i, func(j int) bool {
			out += " " + strconv.Itoa(j)
			return true
		})
		out += "\n"
	}
	out += "\n"
	return out
}

func (dg Digraph) Reverse() Digraph {
	rg := NewDigraph(dg.NumVertical())
	for v := 0; v < int(dg.NumVertical()); v++ {
		dg.IterateAdj(v, func(w int) bool {
			rg.AddEdge(w, v)
			return true
		})
	}
	return rg
}

func (dg Digraph) FindNegativeEdge() (src, dst int) {
	src, dst = -1, -1
	dg.IterateWEdge(func(v int, v2 int, w float64) bool {
		if w < 0 {
			src = v
			dst = v2
			return false
		}
		return true
	})
	return
}

func (dg Digraph) iterateAdj(v int, fn func(int, int, float64) bool) {
	dg[v].Range(func(key hash_map.Key, val util.T) bool {
		return fn(v, int(key.(util.Int)), val.(float64))
	})
}

func (dg Digraph) HasCycle() bool {
	return dg.FindCycle() != nil
}

func (dg Digraph) FindCycle() *stack.IntStack {
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

func (dg Digraph) FindCycleFrom(v int) *stack.IntStack {
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
	dg.IterateAdj(v, func(w int) bool {
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
	order := stack.NewInt(int(dg.NumVertical()))
	dg.IterateVetRDFS(func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func (dg Digraph) IterateWEdge(fn func(int, int, float64) bool) {
	for src, hm := range dg {
		goon := true
		hm.Range(func(dst hash_map.Key, v util.T) bool {
			goon = fn(src, int(dst.(util.Int)), v.(float64))
			return goon
		})
		if !goon {
			break
		}
	}
}

func (dg Digraph) IterateEdge(fn func(int, int) bool) {
	dg.IterateWEdge(func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

func (dg Digraph) IterateWEdgeFrom(v int, fn func(int, int, float64) bool) {
	dg.IterateVetDFS(v, func(v int) bool {
		goon := true
		dg.IterateWAdj(v, func(a int, w float64) bool {
			goon = fn(v, a, w)
			return goon
		})
		return goon
	})
}

func (dg Digraph) IterateEdgeFrom(v int, fn func(int, int) bool) {
	dg.IterateWEdgeFrom(v, func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

func (dg Digraph) IterateVetDFS(src int, fn func(int) bool) {
	dg.IterateUnMarkVetDFS(src, nil, fn)
}

func (dg Digraph) IterateUnMarkVetDFS(src int, marked []bool, fn func(int) bool) {
	if !dg.HasVertical(src) {
		return
	}
	if len(marked) == 0 {
		marked = make([]bool, dg.NumVertical())
	}
	dg.iterateUnMarkVetDFS(src, marked, fn)
}

func (dg Digraph) iterateUnMarkVetDFS(v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	dg.IterateAdj(v, func(adj int) bool {
		if !marked[adj] {
			if !dg.iterateUnMarkVetDFS(adj, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

func (dg Digraph) ReachableSlice(src int) []int {
	if !dg.HasVertical(src) {
		return nil
	}
	var arrived []int
	dg.IterateVetDFS(src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

func (dg Digraph) ReachableBits(src int) []bool {
	if !dg.HasVertical(src) {
		return nil
	}
	marked := make([]bool, dg.NumVertical())
	dg.iterateUnMarkVetDFS(src, marked, func(_ int) bool { return true })
	return marked
}

func (dg Digraph) IterateVetRDFS(fn func(int) bool) {
	marked := make([]bool, dg.NumVertical())
	for v := range marked {
		if marked[v] {
			continue
		}
		dg.rDFS(v, marked, fn)
	}
}

func (dg Digraph) IterateRDFSFromVet(src int, fn func(int) bool) {
	if !dg.HasVertical(src) {
		return
	}
	marked := make([]bool, dg.NumVertical())
	dg.rDFS(src, marked, fn)
}

func (dg Digraph) RDFSOrderVertical(src int) *stack.IntStack {
	order := stack.NewInt(int(dg.NumVertical()))
	dg.IterateRDFSFromVet(src, func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func (dg Digraph) rDFS(v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	goon := true // continue DFS or abort
	dg.IterateAdj(v, func(a int) bool {
		if !marked[a] {
			if !dg.rDFS(a, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	if !goon {
		return false
	}
	return fn(v)
}

func (dg Digraph) IsBipartiteGraph() bool {
	marks := make([]bool, dg.NumVertical())
	colors := make([]bool, dg.NumVertical())
	for i, m := range marks {
		if m {
			continue
		}
		if !dg.isBipartiteDFS(i, true, colors, marks) {
			return false
		}
	}
	return true
}

func (dg Digraph) isBipartiteDFS(cur int, color bool, colors []bool, marked []bool) bool {
	isBip := true
	if !marked[cur] {
		marked[cur] = true
		colors[cur] = color
		dg.IterateAdj(cur, func(adj int) bool {
			if !dg.isBipartiteDFS(adj, !color, colors, marked) {
				isBip = false
				return false
			}
			return true
		})
	} else {
		isBip = colors[cur] == color
	}
	return isBip
}

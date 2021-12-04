package graph

type TransitiveClosure struct {
	locate    []int   // vertical -> subGraphID
	subGraphs [][]int // subGraphID -> all vertices
}

func (g *Graph) TransitiveClosure() *TransitiveClosure {
	tc := &TransitiveClosure{
		locate: make([]int, g.NumVertical()),
	}
	for i := range tc.locate {
		tc.locate[i] = -1
	}

	subGraphID := 0
	for i, c := range tc.locate {
		if c < 0 {
			dfs := g.ReachableSlice(i)
			for _, v := range dfs {
				tc.locate[v] = subGraphID
			}
			tc.subGraphs = append(tc.subGraphs, dfs)
			subGraphID++
		}
	}
	return tc
}

func (tc *TransitiveClosure) CanReach(src, dst int) bool {
	if !tc.hasVertical(src) || !tc.hasVertical(dst) {
		return false
	}
	return tc.locate[src] == tc.locate[dst]
}

func (tc *TransitiveClosure) Range(v int, fn func(v int) bool) {
	if !tc.hasVertical(v) {
		return
	}
	for _, v := range tc.subGraphs[tc.locate[v]] {
		if !fn(v) {
			break
		}
	}
}

func (tc *TransitiveClosure) hasVertical(v int) bool {
	return v >= 0 || v < len(tc.locate)
}

func (tc *TransitiveClosure) NumSubGraph() int {
	return len(tc.subGraphs)
}

func (tc *TransitiveClosure) Locate(v int) int {
	if !tc.hasVertical(v) {
		return -1
	}
	return tc.locate[v]
}

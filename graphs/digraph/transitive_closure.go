package digraph

type TransitiveClosure struct {
	locate   []int   // vertical -> closureID
	closures [][]int // closureID -> all vertices
}

func (g Digraph) TransitiveClosure() *TransitiveClosure {
	tc := &TransitiveClosure{
		locate: make([]int, g.NumV()),
	}
	for i := range tc.locate {
		tc.locate[i] = -1
	}

	closureID := 0
	for i, c := range tc.locate {
		if c < 0 {
			dfs := g.DFS(i)
			for _, v := range dfs {
				tc.locate[v] = closureID
			}
			tc.closures = append(tc.closures, dfs)
			closureID++
		}
	}
	return tc
}

func (tc *TransitiveClosure) IsReachable(src, dst int) bool {
	if !tc.hasV(src) || !tc.hasV(dst) {
		return false
	}
	return tc.locate[src] == tc.locate[dst]
}

func (tc *TransitiveClosure) Range(v int, fn func(v int) bool) {
	if !tc.hasV(v) {
		return
	}
	for _, v := range tc.closures[tc.locate[v]] {
		if !fn(v) {
			break
		}
	}
}

func (tc *TransitiveClosure) hasV(v int) bool {
	return v >= 0 || v < len(tc.closures)
}

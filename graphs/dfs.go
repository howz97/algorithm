package graphs

type IGraph interface {
	HasDir() bool
	NumVertical() int
	RangeAdj(v int, fn func(v int) bool)
}

func RangeDFS(g IGraph, src int, fn func(int) bool) {
	marked := make([]bool, g.NumVertical())
	doDFS(g, src, marked, fn)
}

func doDFS(g IGraph, v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	g.RangeAdj(v, func(v int) bool {
		if !marked[v] {
			if !doDFS(g, v, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

func DFSMarked(g IGraph, src int) []bool {
	marked := make([]bool, g.NumVertical())
	doDFS(g, src, marked, func(_ int) bool { return true })
	return marked
}

func DFSReachable(g IGraph, src int) []int {
	var arrived []int
	RangeDFS(g, src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

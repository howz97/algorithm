package graphs

type IGraph interface {
	HasVertical(v int) bool
	NumVertical() int
	AddEdge(v1, v2 int) error
	HasEdge(v1, v2 int) bool
	RangeAdj(v int, fn func(v int) bool)
}

// DFS will infinitely perform recursion on meeting a ring
func DFS(g IGraph, src int, fn func(int) bool) bool {
	if !fn(src) {
		return false
	}
	goon := true
	g.RangeAdj(src, func(adj int) bool {
		if !fn(adj) {
			goon = false
			return false
		}
		goon = DFS(g, adj, fn)
		return goon
	})
	return goon
}

func RangeReachable(g IGraph, src int, fn func(int) bool) {
	if !g.HasVertical(src) {
		return
	}
	marked := make([]bool, g.NumVertical())
	rangeReachable(g, src, marked, fn)
}

func rangeReachable(g IGraph, v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	g.RangeAdj(v, func(adj int) bool {
		if !marked[adj] {
			if !rangeReachable(g, adj, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

func ReachableBits(g IGraph, src int) []bool {
	if !g.HasVertical(src) {
		return nil
	}
	marked := make([]bool, g.NumVertical())
	rangeReachable(g, src, marked, func(_ int) bool { return true })
	return marked
}

func ReachableSlice(g IGraph, src int) []int {
	if !g.HasVertical(src) {
		return nil
	}
	var arrived []int
	RangeReachable(g, src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

package graphs

func RangeDFS(g ITraverse, src int, fn func(int) bool) {
	RangeUnMarkDFS(g, src, nil, fn)
}

func RangeUnMarkDFS(g ITraverse, src int, marked []bool, fn func(int) bool) {
	if !g.HasVertical(src) {
		return
	}
	if len(marked) == 0 {
		marked = make([]bool, g.NumVertical())
	}
	rangeDFS(g, src, marked, fn)
}

func rangeDFS(g ITraverse, v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	g.RangeAdj(v, func(adj int) bool {
		if !marked[adj] {
			if !rangeDFS(g, adj, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

func RevDFSAll(g ITraverse, fn func(int) bool) {
	marked := make([]bool, g.NumVertical())
	for v := range marked {
		if marked[v] {
			continue
		}
		revDFS(g, v, marked, fn)
	}
}

func RevDFS(g ITraverse, src int, fn func(int) bool) {
	if !g.HasVertical(src) {
		return
	}
	marked := make([]bool, g.NumVertical())
	revDFS(g, src, marked, fn)
}

func revDFS(g ITraverse, v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	goon := true // continue DFS or abort
	g.RangeAdj(v, func(adj int) bool {
		if !marked[adj] {
			if !rangeDFS(g, adj, marked, fn) {
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

func ReachableBits(g ITraverse, src int) []bool {
	if !g.HasVertical(src) {
		return nil
	}
	marked := make([]bool, g.NumVertical())
	rangeDFS(g, src, marked, func(_ int) bool { return true })
	return marked
}

func ReachableSlice(g ITraverse, src int) []int {
	if !g.HasVertical(src) {
		return nil
	}
	var arrived []int
	RangeDFS(g, src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

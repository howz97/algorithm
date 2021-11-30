package graphs

func IsBipartiteGraph(g IGraph) bool {
	marks := make([]bool, g.NumVertical())
	colors := make([]bool, g.NumVertical())
	for i, m := range marks {
		if m {
			continue
		}
		if !isBipartiteDFS(g, i, true, colors, marks) {
			return false
		}
	}
	return true
}

func isBipartiteDFS(g IGraph, cur int, color bool, colors []bool, marked []bool) bool {
	isBip := true
	if !marked[cur] {
		marked[cur] = true
		colors[cur] = color
		g.RangeAdj(cur, func(adj int) bool {
			if !isBipartiteDFS(g, adj, !color, colors, marked) {
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

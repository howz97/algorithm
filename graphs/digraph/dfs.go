package digraph

func (g Digraph) DFS(src int) []int {
	marked := make([]bool, g.NumV())
	doDFS(src, g, marked)
	var arrived []int
	for v, m := range marked {
		if m {
			arrived = append(arrived, v)
		}
	}
	return arrived
}

func doDFS(v int, g Digraph, marked []bool) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, w := range adj {
		if !marked[w] {
			doDFS(w, g, marked)
		}
	}
}

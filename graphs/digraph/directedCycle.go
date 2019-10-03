package digraph

import "github.com/zh1014/algorithm/set"

func DetectDirCycle(g Digraph) bool {
	marked := make([]bool, g.NumV())
	s := set.New()
	for i, b := range marked {
		if !b {
			if detectDirCycle(g, i, marked, s) {
				return true
			}
		}
	}
	return false
}

func detectDirCycle(g Digraph, v int, marked []bool, s set.Set) bool {
	if s.Contains(v) {
		return true
	}
	s.Add(v)
	marked[v] = true
	adj := g.Adjacent(v)
	for _, w := range adj {
		if !marked[w] && detectDirCycle(g, w, marked, s) {
			return true
		}
	}
	s.Remove(v)
	return false
}

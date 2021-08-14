package weighted_digraph

import "howz97/algorithm/set"

func DetectDirCycle(g EdgeWeightedDigraph) bool {
	marked := make([]bool, g.NumV())
	s := set.NewIntSet()
	for i, b := range marked {
		if !b {
			if detectDirCycle(g, i, marked, s) {
				return true
			}
		}
	}
	return false
}

func detectDirCycle(g EdgeWeightedDigraph, v int, marked []bool, s set.IntSet) bool {
	if s.Contains(v) {
		return true
	}
	s.Add(v)
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		w := e.to
		if !marked[w] && detectDirCycle(g, w, marked, s) {
			return true
		}
	}
	s.Remove(v)
	return false
}

package wdigraph

import "github.com/howz97/algorithm/set"

func (g WDigraph) DetectDirCycle() bool {
	marked := make([]bool, g.NumVertical())
	s := set.NewIntSet()
	for i, b := range marked {
		if !b {
			if g.detectDirCycle(i, marked, s) {
				return true
			}
		}
	}
	return false
}

func (g WDigraph) detectDirCycle(v int, marked []bool, s set.IntSet) bool {
	if s.Contains(v) {
		return true
	}
	s.Add(v)
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		w := e.to
		if !marked[w] && g.detectDirCycle(w, marked, s) {
			return true
		}
	}
	s.Remove(v)
	return false
}

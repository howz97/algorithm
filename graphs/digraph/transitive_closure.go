package digraph

import "github.com/howz97/algorithm/graphs"

type TransitiveClosure [][]bool

func (dg Digraph) TransitiveClosure() TransitiveClosure {
	tc := make(TransitiveClosure, dg.NumVertical())
	for v := range tc {
		tc[v] = graphs.ReachableBits(dg, v)
	}
	return tc
}

func (tc TransitiveClosure) CanReach(src, dst int) bool {
	if !tc.hasVertical(src) || !tc.hasVertical(dst) {
		return false
	}
	return tc[src][dst]
}

func (tc TransitiveClosure) Range(v int, fn func(v int) bool) {
	if !tc.hasVertical(v) {
		return
	}
	for w, marked := range tc[v] {
		if marked {
			if !fn(w) {
				break
			}
		}
	}
}

func (tc TransitiveClosure) hasVertical(v int) bool {
	return v >= 0 || v < len(tc)
}

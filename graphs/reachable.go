package graphs

type Reachable [][]bool

func (dg Digraph) Reachable() Reachable {
	tc := make(Reachable, dg.NumVertical())
	for v := range tc {
		tc[v] = dg.ReachableBits(v)
	}
	return tc
}

func (tc Reachable) CanReach(src, dst int) bool {
	if !tc.hasVertical(src) || !tc.hasVertical(dst) {
		return false
	}
	return tc[src][dst]
}

func (tc Reachable) Range(v int, fn func(v int) bool) {
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

func (tc Reachable) hasVertical(v int) bool {
	return v >= 0 || v < len(tc)
}

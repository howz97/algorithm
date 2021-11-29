package digraph

type SCC struct {
	marked []bool
	id     []int
	g      Digraph
	count  int
}

func NewSCC(g Digraph) *SCC {
	scc := &SCC{
		marked: make([]bool, g.NumV()),
		id:     make([]int, g.NumV()),
		g:      g,
	}
	topOrderStack := ReversePostOrder(g.Reverse())
	for {
		v, ok := topOrderStack.Pop()
		if !ok {
			break
		}
		if !scc.marked[v] {
			scc.markID(v, scc.count)
			scc.count++
		}
	}
	return scc
}

func (scc *SCC) markID(v, sccid int) {
	scc.marked[v] = true
	scc.id[v] = sccid
	adj := scc.g.Adjacent(v)
	for _, w := range adj {
		if !scc.marked[w] {
			scc.markID(w, sccid)
		}
	}
}

func (scc *SCC) IsStronglyConnected(src, dst int) bool {
	if !scc.has(src) || !scc.has(dst) {
		return false
	}
	return scc.id[src] == scc.id[dst]
}

func (scc *SCC) GetID(v int) int {
	if !scc.has(v) {
		panic("invalid vertical")
	}
	return scc.id[v]
}

func (scc *SCC) NumSCC() int {
	return scc.count
}

func (scc *SCC) has(v int) bool {
	return v >= 0 && v < len(scc.marked)
}

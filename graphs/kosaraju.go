package graphs

// SCC is strong connected components of digraph
// vertices in the same component can access each other
type SCC struct {
	locate     []int   // vertical -> componentID
	components [][]int // componentID -> vertices
}

// SCC calculate strong connected components of digraph with kosaraju algorithm
func (dg Digraph) SCC() *SCC {
	scc := &SCC{
		locate: make([]int, dg.NumVertical()),
	}
	marked := make([]bool, dg.NumVertical())
	dg.IterateVetRDFS(func(v int) bool {
		if !marked[v] {
			c := make([]int, 0, 8)
			dg.IterateUnMarkVetDFS(v, marked, func(w int) bool {
				scc.locate[w] = len(scc.components)
				c = append(c, w)
				return true
			})
			scc.components = append(scc.components, c)
		}
		return true
	})
	return scc
}

func (scc *SCC) IsStronglyConn(src, dst int) bool {
	if !scc.hasVertical(src) || !scc.hasVertical(dst) {
		return false
	}
	return scc.locate[src] == scc.locate[dst]
}

func (scc *SCC) GetCompID(v int) int {
	if !scc.hasVertical(v) {
		return -1
	}
	return scc.locate[v]
}

func (scc *SCC) RangeComponent(v int, fn func(int) bool) {
	if !scc.hasVertical(v) {
		return
	}
	for _, w := range scc.components[scc.locate[v]] {
		if !fn(w) {
			break
		}
	}
}

func (scc *SCC) NumComponents() int {
	return len(scc.components)
}

func (scc *SCC) hasVertical(v int) bool {
	return v >= 0 && v < len(scc.locate)
}

package graph

type Connectivity struct {
	marked []bool
	count  int
	id     []int
}

func (g Graph) NewConnectivity() *Connectivity {
	cd := &Connectivity{
		marked: make([]bool, g.NumVertical()),
		count:  0,
		id:     make([]int, g.NumVertical()),
	}
	for i, b := range cd.marked {
		if !b {
			cd.count++
			cd.dfs(g, i)
		}
	}
	return cd
}

func (cd *Connectivity) dfs(g Graph, src int) {
	cd.id[src] = cd.count
	cd.marked[src] = true
	adjs := g.Adjacent(src)
	for _, adj := range adjs {
		if !cd.marked[adj] {
			cd.dfs(g, adj)
		}
	}
}

func (cd *Connectivity) IsConnected(v1, v2 int) (bool, error) {
	if !cd.hasV(v1) || !cd.hasV(v2) {
		return false, ErrVerticalNotExist
	}
	return cd.id[v1] == cd.id[v2], nil
}

func (cd *Connectivity) NumSubGraph() int {
	return cd.count
}

func (cd *Connectivity) SubGraphIDOf(v int) (int, error) {
	if !cd.hasV(v) {
		return 0, ErrVerticalNotExist
	}
	return cd.id[v], nil
}

func (cd *Connectivity) hasV(v int) bool {
	return v >= 0 && v < len(cd.marked)
}

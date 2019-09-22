package graph

type ConnvtyDetector struct {
	marked []bool
	count  int
	id     []int
}

func NewConnectivity(g Graph) *ConnvtyDetector {
	cd := &ConnvtyDetector{
		marked: make([]bool, g.NumV()),
		count:  0,
		id:     make([]int, g.NumV()),
	}
	for i, b := range cd.marked {
		if !b {
			cd.count++
			cd.dfs(g, i)
		}
	}
	return cd
}

func (cd *ConnvtyDetector) dfs(g Graph, src int) {
	cd.id[src] = cd.count
	adjs, _ := g.Adjacent(src)
	for _, adj := range adjs {
		if !cd.marked[adj] {
			cd.dfs(g, adj)
		}
	}
}

func (cd *ConnvtyDetector) IsConnected(v1, v2 int) (bool, error) {
	if !cd.hasV(v1) || !cd.hasV(v2) {
		return false, errVerticalNotExist
	}
	return cd.id[v1] == cd.id[v2], nil
}

func (cd *ConnvtyDetector) NumSubGraph() int {
	return cd.count
}

func (cd *ConnvtyDetector) SubGraphIDOf(v int) (int, error) {
	if !cd.hasV(v) {
		return 0, errVerticalNotExist
	}
	return cd.id[v], nil
}

func (cd *ConnvtyDetector)hasV(v int) bool {
	return v >=0 && v < len(cd.marked)
}

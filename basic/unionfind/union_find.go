package unionfind

const (
	verticalOverflow = "vertical overflow"
)

type UnionFind struct {
	id    []int
	numcc int
}

func New(numV int) *UnionFind {
	uf := &UnionFind{
		id:    make([]int, numV),
		numcc: numV,
	}
	for i := range uf.id {
		uf.id[i] = i
	}
	return uf
}

func (uf *UnionFind) NumConnectedComponent() int {
	return uf.numcc
}

func (uf *UnionFind) Find(v int) int {
	if v < 0 || v >= len(uf.id) {
		panic(verticalOverflow)
	}
	return uf.id[v]
}

func (uf *UnionFind) IsConnected(v1, v2 int) bool {
	if v1 < 0 || v1 >= len(uf.id) || v2 < 0 || v2 >= len(uf.id) {
		panic(verticalOverflow)
	}
	return uf.id[v1] == uf.id[v2]
}

func (uf *UnionFind) Union(v1, v2 int) {
	if v1 < 0 || v1 >= len(uf.id) || v2 < 0 || v2 >= len(uf.id) {
		panic(verticalOverflow)
	}
	if uf.IsConnected(v1, v2) {
		return
	}
	v2id := uf.id[v2]
	for i := range uf.id {
		if uf.id[i] == v2id {
			uf.id[i] = uf.id[v1]
		}
	}
	uf.numcc--
}

package wgraph

type edgeSet map[*Edge]struct{}

func NewEdgeSet() edgeSet {
	return make(map[*Edge]struct{})
}

func (s edgeSet) add(e *Edge) {
	s[e] = struct{}{}
}

func (s edgeSet) contains(e *Edge) bool {
	_, contain := s[e]
	return contain
}

func (s edgeSet) remove(e *Edge) {
	delete(s, e)
}

func (s edgeSet) clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s edgeSet) len() int {
	return len(s)
}

func (s edgeSet) traverse() []*Edge {
	result := make([]*Edge, 0)
	for e := range s {
		result = append(result, e)
	}
	return result
}

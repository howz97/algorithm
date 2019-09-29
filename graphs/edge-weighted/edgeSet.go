package edge_weighted

type edgeSet map[*edge]struct{}

func New() edgeSet {
	return make(map[*edge]struct{})
}

func (s edgeSet) add(e *edge) {
	s[e] = struct{}{}
}

func (s edgeSet) contains(e *edge) bool {
	_, contain := s[e]
	return contain
}

func (s edgeSet) remove(e *edge) {
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

func (s edgeSet) traverse() []*edge {
	result := make([]*edge, 0)
	for e := range s {
		result = append(result, e)
	}
	return result
}

package graphs

import (
	"errors"
	"gopkg.in/yaml.v2"
)

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
	ErrInputFormat      = errors.New("input format error")
)

type ITraverse interface {
	HasVertical(v int) bool
	NumVertical() uint
	IterateAdj(v int, fn func(a int) bool)
}

type IGraph interface {
	ITraverse
	AddEdge(v1, v2 int) error
	HasEdge(v1, v2 int) bool
}

type IWGraph interface {
	ITraverse
	AddEdge(v1, v2 int, w float64) error
	HasEdge(v1, v2 int) bool
	IterateWAdj(v int, fn func(a int, w float64) bool)
}

func MarshalWGraph(g IWGraph) ([]byte, error) {
	m := make(map[int]map[int]float64)
	for v := 0; v < int(g.NumVertical()); v++ {
		edges := make(map[int]float64)
		g.IterateWAdj(v, func(a int, w float64) bool {
			edges[a] = w
			return true
		})
		m[v] = edges
	}
	return yaml.Marshal(m)
}

package graphs

import "errors"

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
	ErrInputFormat      = errors.New("input format error")
)

type ITraverse interface {
	HasVertical(v int) bool
	NumVertical() int
	RangeAdj(v int, fn func(v int) bool)
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
}

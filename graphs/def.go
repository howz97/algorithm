package graphs

import "errors"

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
	ErrSelfLoop         = errors.New("not support self loop")
)

type IGraph interface {
	HasVertical(v int) bool
	NumVertical() int
	AddEdge(v1, v2 int) error
	HasEdge(v1, v2 int) bool
	RangeAdj(v int, fn func(v int) bool)
}

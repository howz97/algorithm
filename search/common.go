package search

import (
	"fmt"
)

type Searcher[Cmp comparable, T any] interface {
	Put(key Cmp, val T)
	Get(key Cmp) T
	Del(key Cmp)
	Clean()
	Size() uint
}

type ITraversal interface {
	fmt.Stringer
	IsNil() bool
	Left() ITraversal
	Right() ITraversal
}

package search

import (
	"golang.org/x/exp/constraints"
)

type Searcher[Ord constraints.Ordered, T any] interface {
	Put(key Ord, val T)
	Get(key Ord) (T, bool)
	Del(key Ord)
	Clean()
	Size() uint
}

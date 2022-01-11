package search

import "github.com/howz97/algorithm/util"

type Searcher interface {
	Put(key util.Comparable, val util.T)
	Get(key util.Comparable) util.T
	Del(key util.Comparable)
	Clean()
	Size() uint
}

type ITraversal interface {
	IsNil() bool
	Val() util.T
	Left() ITraversal
	Right() ITraversal
}

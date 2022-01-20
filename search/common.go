package search

import (
	"fmt"
	"github.com/howz97/algorithm/util"
)

type Searcher interface {
	Put(key util.Comparable, val util.T)
	Get(key util.Comparable) util.T
	Del(key util.Comparable)
	Size() uint
}

type ITraversal interface {
	fmt.Stringer
	IsNil() bool
	Left() ITraversal
	Right() ITraversal
}

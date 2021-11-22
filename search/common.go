package search

import "fmt"

const StrNilNode = "#"

const (
	Equal Result = iota
	Less
	More
)

type Result int

type T interface{}

type Cmp interface {
	Cmp(other Cmp) Result
}

type Searcher interface {
	Put(key Cmp, value T)
	Get(k Cmp) T
	Del(key Cmp)
	Clean()
	//Size() uint
}

type ITraversal interface {
	IsNil() bool
	Left() ITraversal
	Right() ITraversal
	fmt.Stringer
}

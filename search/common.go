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
	Insert(key Cmp, value T)
	Find(k Cmp) T
	Delete(key Cmp)
	Clean()
	//Size() uint
}

type ITraversal interface {
	IsNil() bool
	Left() ITraversal
	Right() ITraversal
	fmt.Stringer
}

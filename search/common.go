package search

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
	Put(key Cmp, val T)
	Get(key Cmp) T
	Del(key Cmp)
	Clean()
}

type ITraversal interface {
	IsNil() bool
	Key() Cmp
	Val() T
	Left() ITraversal
	Right() ITraversal
}

type KeyVal struct {
	key   Cmp
	value T
}

func (kv *KeyVal) Key() Cmp {
	return kv.key
}

func (kv *KeyVal) Val() T {
	return kv.value
}

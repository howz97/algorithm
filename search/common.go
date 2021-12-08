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
	Key() util.Comparable
	Val() util.T
	Left() ITraversal
	Right() ITraversal
}

type KeyVal struct {
	key   util.Comparable
	value util.T
}

func (kv *KeyVal) Key() util.Comparable {
	return kv.key
}

func (kv *KeyVal) Val() util.T {
	return kv.value
}

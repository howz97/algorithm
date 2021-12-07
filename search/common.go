package search

import "github.com/howz97/algorithm/util"

type Searcher interface {
	Put(key util.Cmp, val util.T)
	Get(key util.Cmp) util.T
	Del(key util.Cmp)
	Clean()
	Size() uint
}

type ITraversal interface {
	IsNil() bool
	Key() util.Cmp
	Val() util.T
	Left() ITraversal
	Right() ITraversal
}

type KeyVal struct {
	key   util.Cmp
	value util.T
}

func (kv *KeyVal) Key() util.Cmp {
	return kv.key
}

func (kv *KeyVal) Val() util.T {
	return kv.value
}

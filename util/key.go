package util

import (
	"github.com/howz97/algorithm/search"
	"unsafe"
)

type Integer int // todo: rename to Int

func (v Integer) Hash() uint {
	return uint(v)
}

func (v Integer) Cmp(other search.Cmp) search.Result {
	o := other.(Integer)
	if v < o {
		return search.Less
	} else if v > o {
		return search.More
	} else {
		return search.Equal
	}
}

type Float float64

func (v Float) Hash() uint {
	return *(*uint)(unsafe.Pointer(&v))
}

func (v Float) Cmp(other search.Cmp) search.Result {
	o := other.(Float)
	if v < o {
		return search.Less
	} else if v > o {
		return search.More
	} else {
		return search.Equal
	}
}

const littlePrime = 31

type Str string

func (v Str) Hash() uint {
	h := uint(0)
	for i := 0; i < len(v); i++ {
		h = h*littlePrime + uint(v[i])
	}
	return h
}

func (v Str) Cmp(other search.Cmp) search.Result {
	o := other.(Str)
	if v < o {
		return search.Less
	} else if v > o {
		return search.More
	} else {
		return search.Equal
	}
}

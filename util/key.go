package util

import (
	"unsafe"
)

type Result int

const (
	Equal Result = iota
	Less
	More
)

type T interface{}

type Comparable interface {
	Cmp(other Comparable) Result
}

type Int int

func (v Int) Hash() uint {
	return uint(v)
}

func (v Int) Cmp(other Comparable) Result {
	o := other.(Int)
	if v < o {
		return Less
	} else if v > o {
		return More
	} else {
		return Equal
	}
}

type Float float64

func (v Float) Hash() uint {
	return *(*uint)(unsafe.Pointer(&v))
}

func (v Float) Cmp(other Comparable) Result {
	o := other.(Float)
	if v < o {
		return Less
	} else if v > o {
		return More
	} else {
		return Equal
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

func (v Str) Cmp(other Comparable) Result {
	o := other.(Str)
	if v < o {
		return Less
	} else if v > o {
		return More
	} else {
		return Equal
	}
}

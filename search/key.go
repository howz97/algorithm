package search

import "unsafe"

type Integer int

func (v Integer) Hash() uint {
	return uint(v)
}

func (v Integer) Cmp(other Cmp) Result {
	o := other.(Integer)
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

func (v Float) Cmp(other Cmp) Result {
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

func (v Str) Cmp(other Cmp) Result {
	o := other.(Str)
	if v < o {
		return Less
	} else if v > o {
		return More
	} else {
		return Equal
	}
}

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

type Int int

func (v Int) Hash() uintptr {
	return uintptr(v)
}

type Float float64

func (v Float) Hash() uintptr {
	return *(*uintptr)(unsafe.Pointer(&v))
}

const littlePrime = 31

type Str string

func (v Str) Hash() uintptr {
	h := uintptr(0)
	for i := 0; i < len(v); i++ {
		h = h*littlePrime + uintptr(v[i])
	}
	return h
}

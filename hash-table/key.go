package hashtable

import "unsafe"

type Key interface {
	HashCode() int
	Equal(h Key) bool
}

type Integer int

func (i Integer) HashCode() int {
	return int(i)
}

func (i Integer) Equal(k Key) bool {
	i2, ok := k.(Integer)
	if !ok {
		return false
	}
	return i == i2
}

type Float float64

func (f Float) HashCode() int {
	return *(*int)(unsafe.Pointer(&f))
}

func (f Float) Equal(k Key) bool {
	f2, ok := k.(Float)
	if !ok {
		return false
	}
	return f == f2
}

const littlePrime = 31

type Str string

func (s Str) HashCode() int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = h*littlePrime + int(s[i])
	}
	return h
}

func (s Str) Equal(k Key) bool {
	s2, ok := k.(Str)
	if !ok {
		return false
	}
	return s == s2
}

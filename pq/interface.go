package pq

import "cmp"

type PriorQueue[P cmp.Ordered] interface {
	Push(P)
	Pop() P
	Top() P
	Size() int
}

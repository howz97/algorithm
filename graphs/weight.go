package graphs

import (
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
)

type Weight []*hash_map.Chaining

func NewWeight(size int) Weight {
	w := make([]*hash_map.Chaining, size)
	for i := range w {
		w[i] = hash_map.New()
	}
	return w
}

func (w Weight) SetWeight(src, dst int, v float64) {
	w[src].Put(search.Integer(dst), v)
}

func (w Weight) FindNegativeEdge() (src, dst int) {
	src, dst = -1, -1
	for v0, hm := range w {
		hm.Range(func(key hash_map.Key, val search.T) bool {
			if val.(float64) < 0 {
				src = v0
				dst = int(key.(search.Integer))
				return false
			}
			return true
		})
	}
	return
}

func (w Weight) GetWeight(src, dst int) float64 {
	return w[src].Get(search.Integer(dst)).(float64)
}
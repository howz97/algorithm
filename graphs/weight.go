package graphs

import (
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/util"
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
	w[src].Put(util.Integer(dst), v)
}

func (w Weight) FindNegativeEdge() (src, dst int) {
	src, dst = -1, -1
	for v0, hm := range w {
		hm.Range(func(key hash_map.Key, val search.T) bool {
			if val.(float64) < 0 {
				src = v0
				dst = int(key.(util.Integer))
				return false
			}
			return true
		})
	}
	return
}

func (w Weight) Iterate(fn func(int,int,float64)bool) {
	for src, hm := range w {
		goon := true
		hm.Range(func(dst hash_map.Key, v search.T) bool {
			goon = fn(src, int(dst.(util.Integer)), v.(float64))
			return goon
		})
		if !goon {
			break
		}
	}
}

func (w Weight) IterateAdj(v int, fn func(int,int,float64)bool) {
	w[v].Range(func(key hash_map.Key, val search.T) bool {
		return fn(v, int(key.(util.Integer)), val.(float64))
	})
}

func (w Weight) HasVet(v int) bool {
	return v >= 0 && v < len(w)
}

func (w Weight) GetWeight(src, dst int) float64 {
	return w[src].Get(util.Integer(dst)).(float64)
}
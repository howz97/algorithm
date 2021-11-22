package integer

import (
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
)

type HashMap struct {
	hash_map.Chaining
}

func NewHashMap() *HashMap {
	return &HashMap{Chaining: *hash_map.New()}
}

func (m *HashMap) Put(key search.Cmp, val search.T) {
	m.Chaining.Put(key.(hash_map.Key), val)
}

func (m *HashMap) Get(key search.Cmp) search.T {
	return m.Chaining.Get(key.(hash_map.Key))
}

func (m *HashMap) Del(key search.Cmp) {
	m.Chaining.Del(key.(hash_map.Key))
}

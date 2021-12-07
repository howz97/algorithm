package integer

import (
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/util"
)

type HashMap struct {
	hash_map.Chaining
}

func NewHashMap() *HashMap {
	return &HashMap{Chaining: *hash_map.New()}
}

func (m *HashMap) Put(key util.Cmp, val util.T) {
	m.Chaining.Put(key.(hash_map.Key), val)
}

func (m *HashMap) Get(key util.Cmp) util.T {
	return m.Chaining.Get(key.(hash_map.Key))
}

func (m *HashMap) Del(key util.Cmp) {
	m.Chaining.Del(key.(hash_map.Key))
}

// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"math/rand"
	"testing"

	"github.com/howz97/algorithm/util"
)

func TestBenchmark_Put_OrderKeys(t *testing.T) {
	// const benchmark = 10000000
	const benchmark = 30000

	stdMap := make(map[util.Int]int)
	elapsed := util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			stdMap[util.Int(i)] = i
		}
	})
	t.Logf("stdMap.Put cost %v", elapsed)

	hm := NewHashMap[util.Int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(util.Int(i), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)

	avl := NewAVL[util.Int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(util.Int(i), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)

	bt := NewBinTree[util.Int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Put(util.Int(i), i)
		}
	})
	t.Logf("binarytree.Put cost %v", elapsed)
}

func TestBenchmark_RandKeys(t *testing.T) {
	const benchmark = 200000

	stdMap := make(map[util.Int]int)
	elapsed := util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			stdMap[util.Int(rand.Intn(benchmark))] = i
		}
	})
	t.Logf("stdMap.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			_ = stdMap[util.Int(rand.Intn(benchmark))]
		}
	})
	t.Logf("stdMap.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			delete(stdMap, util.Int(rand.Intn(benchmark)))
		}
	})
	t.Logf("stdMap.Del cost %v", elapsed)

	hm := NewHashMap[util.Int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(util.Int(rand.Intn(benchmark)), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Get(util.Int(rand.Intn(benchmark)))
		}
	})
	t.Logf("hashmap.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Del(util.Int(rand.Intn(benchmark)))
		}
	})
	t.Logf("hashmap.Del cost %v", elapsed)

	avl := NewAVL[int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(rand.Intn(benchmark), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Get(rand.Intn(benchmark))
		}
	})
	t.Logf("avl.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Del(rand.Intn(benchmark))
		}
	})
	t.Logf("avl.Del cost %v", elapsed)

	bt := NewBinTree[int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Put(rand.Intn(benchmark), i)
		}
	})
	t.Logf("binarytree.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Get(rand.Intn(benchmark))
		}
	})
	t.Logf("binarytree.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Del(rand.Intn(benchmark))
		}
	})
	t.Logf("binarytree.Del cost %v", elapsed)
}

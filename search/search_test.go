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

func TestMap(t *testing.T) {
	const maxKey = 100000
	rnd := rand.New(rand.NewSource(0))

	m := make(map[util.Int]int)
	var maps [4]Searcher[util.Int, int]
	maps[0] = NewBinTree[util.Int, int]()
	maps[1] = NewAVL[util.Int, int]()
	maps[2] = NewRBTree[util.Int, int]()
	maps[3] = NewHashMap[util.Int, int]()

	putFn := func(cnt int) {
		for i := 0; i < cnt; i++ {
			v := rnd.Intn(maxKey)
			m[util.Int(v)] = v
			for i := range maps {
				maps[i].Put(util.Int(v), v)
			}
		}
	}

	delFn := func(cnt int) {
		for k := range m {
			if cnt <= 0 {
				break
			}
			cnt--
			delete(m, k)
			for i := range maps {
				maps[i].Del(k)
			}
		}
	}

	verifyFn := func() {
		for k, v := range m {
			for i := range maps {
				v2, ok := maps[i].Get(k)
				if !ok {
					t.Fatalf("key %v should exist", k)
				}
				if v != v2 {
					t.Fatalf("%v != %v", v, v2)
				}
			}
		}
		for i := range maps {
			sz := maps[i].Size()
			if sz != uint(len(m)) {
				t.Fatalf("size %v != %v", sz, uint(len(m)))
			}
		}
	}

	putFn(10000)
	verifyFn()
	for i := 0; i < 100; i++ {
		putFn(rnd.Intn(1000))
		verifyFn()
		delFn(rnd.Intn(1000))
		verifyFn()
	}
	delFn(len(m))
	verifyFn()
}

func TestTraversal(t *testing.T) {
	data := []string{"26", "-", "5", "*", "3", "+", "2"}
	avl := NewAVL[int, string]()
	for i, v := range data {
		avl.Put(i, v)
	}
	PrintTree(avl.Root(), func(nd *AvlNode[int, string]) string { return nd.Value() })

	str := ""
	RevOrder(avl.Root(), func(nd *AvlNode[int, string]) bool {
		str += nd.Value()
		return true
	})
	if str != "2+3*5-26" {
		t.Fatalf("ReverseOrder: %s", str)
	}

	str = ""
	Inorder(avl.Root(), func(nd *AvlNode[int, string]) bool {
		str += nd.Value()
		return true
	})
	if str != "26-5*3+2" {
		t.Fatalf("InOrder: %s", str)
	}

	str = ""
	PreorderRecur(avl.Root(), func(nd *AvlNode[int, string]) bool {
		str += nd.Value()
		return true
	})
	if str != "*-265+32" {
		t.Fatalf("PreOrder: %s", str)
	}

	str = ""
	Postorder(avl.Root(), func(nd *AvlNode[int, string]) bool {
		str += nd.Value()
		return true
	})
	if str != "265-32+*" {
		t.Fatalf("SufOrder: %s", str)
	}

	str = ""
	LevelOrder(avl.Root(), func(nd *AvlNode[int, string]) bool {
		str += nd.Value()
		return true
	})
	if str != "*-+26532" {
		t.Fatalf("LevelOrder: %s", str)
	}
}

func TestPreOrder(t *testing.T) {
	bt := NewBinTree[int, int]()
	nodes := []int{50, 20, 10, 15, 30, 25, 40, 35, 45, 60, 55, 70}
	for _, v := range nodes {
		bt.Put(v, v)
	}
	i := 0
	PreorderRecur(bt.Root(), func(n *BtNode[int, int]) bool {
		if nodes[i] != n.value {
			t.FailNow()
		}
		i++
		return true
	})
}

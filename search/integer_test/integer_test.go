package integer

import (
	"math/rand"
	"testing"

	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
	"github.com/howz97/algorithm/search/binarytree"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/search/redblack"
	"github.com/howz97/algorithm/util"
)

func TestAVL(t *testing.T) {
	LoopTest(t, avltree.New[util.Int, int]())
}

func TestBinaryTree(t *testing.T) {
	LoopTest(t, binarytree.New[util.Int, int]())
}

func TestRedBlack(t *testing.T) {
	LoopTest(t, redblack.New[util.Int, int]())
}

func TestHashMap(t *testing.T) {
	LoopTest(t, hash_map.New[util.Int, int]())
}

func LoopTest(t *testing.T, s search.Searcher[util.Int, int]) {
	verify := make(map[int]int)
	BulkDelete(verify, s, 100)
	BulkInsert(verify, s, 100)
	BulkDelete(verify, s, 1000)
	BulkInsert(verify, s, 500)
	BulkDelete(verify, s, 100)
	for i := 0; i < 20; i++ {
		BulkInsert(verify, s, rand.Intn(1000))
		VerifyResult(t, verify, s)
		BulkDelete(verify, s, rand.Intn(1000))
		VerifyResult(t, verify, s)
	}
}

func BulkInsert(verify map[int]int, s search.Searcher[util.Int, int], cnt int) {
	for i := 0; i < cnt; i++ {
		k := rand.Int()
		s.Put(util.Int(k), k)
		verify[k] = k
	}
}

func BulkDelete(verify map[int]int, s search.Searcher[util.Int, int], cnt int) {
	for i := 0; i < cnt; i++ {
		k := rand.Int()
		s.Del(util.Int(k))
		delete(verify, k)
	}
}

func VerifyResult(t *testing.T, verify map[int]int, s search.Searcher[util.Int, int]) {
	for k, v := range verify {
		vGot, _ := s.Get(util.Int(k))
		if vGot != v {
			t.Fatalf("key %v has wrong value %v, should be %v", k, vGot, v)
		}
	}
	if uint(len(verify)) != s.Size() {
		t.Fatalf("size not equal %d != %d", len(verify), s.Size())
	}
}

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

	hm := hash_map.New[util.Int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(util.Int(i), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)

	avl := avltree.New[util.Int, int]()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(util.Int(i), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)

	bt := binarytree.New[util.Int, int]()
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

	hm := hash_map.New[util.Int, int]()
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

	avl := avltree.New[int, int]()
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

	bt := binarytree.New[int, int]()
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

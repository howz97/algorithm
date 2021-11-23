package integer

import (
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
	"github.com/howz97/algorithm/search/binarytree"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/search/redblack"
	"github.com/howz97/algorithm/util"
	"math/rand"
	"strconv"
	"testing"
)

const n = 1000

func TestAVL(t *testing.T) {
	DifferentKVType(t, avltree.New())
}

func TestBinaryTree(t *testing.T) {
	DifferentKVType(t, binarytree.New())
}

func TestRedBlack(t *testing.T) {
	DifferentKVType(t, redblack.New())
}

func TestHashMap(t *testing.T) {
	DifferentKVType(t, NewHashMap())
}

func DifferentKVType(t *testing.T, s search.Searcher) {
	t.Logf("start test different types of k-v ...")
	LoopTest(t, s, IntStrKV)
	t.Logf("int-str passed")

	LoopTest(t, s, FloatIntKV)
	t.Logf("float-int passed")

	LoopTest(t, s, StrIntKV)
	t.Logf("str-int passed")
}

func IntStrKV() (search.Cmp, search.T) {
	k := rand.Intn(n)
	v := strconv.Itoa(k)
	return search.Integer(k), v
}

func FloatIntKV() (search.Cmp, search.T) {
	v := rand.Intn(n)
	k := float64(v) / n
	return search.Float(k), v
}

func StrIntKV() (search.Cmp, search.T) {
	k := alphabet.Ascii.RandString(2) // length of string hugely affect cost of BST
	v := rand.Intn(n)
	return search.Str(k), v
}

func LoopTest(t *testing.T, s search.Searcher, kvfn func() (search.Cmp, search.T)) {
	s.Clean()
	verify := make(map[search.Cmp]search.T)
	for i := 0; i < 200; i++ {
		BulkInsert(verify, s, n, kvfn)
		VerifyResult(t, verify, s)
		BulkDelete(verify, s, n, kvfn)
		VerifyResult(t, verify, s)
	}
}

func BulkInsert(verify map[search.Cmp]search.T, s search.Searcher, cnt int, kvfn func() (search.Cmp, search.T)) {
	for i := 0; i < cnt; i++ {
		k, v := kvfn()
		s.Put(k, v)
		verify[k] = v
	}
}

func BulkDelete(verify map[search.Cmp]search.T, s search.Searcher, cnt int, gen func() (search.Cmp, search.T)) {
	for i := 0; i < cnt; i++ {
		k, _ := gen()
		s.Del(k)
		delete(verify, k)
	}
}

func VerifyResult(t *testing.T, verify map[search.Cmp]search.T, s search.Searcher) {
	for k, v := range verify {
		vGot := s.Get(k)
		if vGot != v {
			t.Fatalf("key %v has wrong value %v, should be %v", k, vGot, v)
		}
	}
	if uint(len(verify)) != s.Size() {
		t.Fatalf("size not equal %d != %d", len(verify), s.Size())
	}
}

func TestBenchmark_Put_OrderKeys(t *testing.T) {
	const benchmark = 10000000
	//const benchmark = 30000

	stdMap := make(map[search.Integer]int)
	elapsed := util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			stdMap[search.Integer(i)] = i
		}
	})
	t.Logf("stdMap.Put cost %v", elapsed)

	hm := hash_map.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(search.Integer(i), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)

	avl := avltree.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(search.Integer(i), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)

	//bt := binarytree.New()
	//elapsed = util.ExecCost(func() {
	//	for i := 0; i < benchmark; i++ {
	//		bt.Put(search.Integer(i), i)
	//	}
	//})
	//t.Logf("binarytree.Put cost %v", elapsed)
}

func TestBenchmark_RandKeys(t *testing.T) {
	const benchmark = 20000000

	stdMap := make(map[search.Integer]int)
	elapsed := util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			stdMap[search.Integer(rand.Intn(n))] = i
		}
	})
	t.Logf("stdMap.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			_ = stdMap[search.Integer(rand.Intn(n))]
		}
	})
	t.Logf("stdMap.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			delete(stdMap, search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("stdMap.Del cost %v", elapsed)

	hm := hash_map.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(search.Integer(rand.Intn(n)), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Get(search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("hashmap.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Del(search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("hashmap.Del cost %v", elapsed)

	avl := avltree.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(search.Integer(rand.Intn(n)), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Get(search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("avl.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Del(search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("avl.Del cost %v", elapsed)

	bt := binarytree.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Put(search.Integer(rand.Intn(n)), i)
		}
	})
	t.Logf("binarytree.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Get(search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("binarytree.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Del(search.Integer(rand.Intn(n)))
		}
	})
	t.Logf("binarytree.Del cost %v", elapsed)
}

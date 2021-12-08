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

func IntStrKV() (util.Comparable, util.T) {
	k := rand.Intn(n)
	v := strconv.Itoa(k)
	return util.Int(k), v
}

func FloatIntKV() (util.Comparable, util.T) {
	v := rand.Intn(n)
	k := float64(v) / n
	return util.Float(k), v
}

func StrIntKV() (util.Comparable, util.T) {
	k := alphabet.Ascii.RandString(2) // length of string hugely affect cost of BST
	v := rand.Intn(n)
	return util.Str(k), v
}

func LoopTest(t *testing.T, s search.Searcher, kvfn func() (util.Comparable, util.T)) {
	s.Clean()
	verify := make(map[util.Comparable]util.T)
	for i := 0; i < 200; i++ {
		BulkInsert(verify, s, n, kvfn)
		VerifyResult(t, verify, s)
		BulkDelete(verify, s, n, kvfn)
		VerifyResult(t, verify, s)
	}
}

func BulkInsert(verify map[util.Comparable]util.T, s search.Searcher, cnt int, kvfn func() (util.Comparable, util.T)) {
	for i := 0; i < cnt; i++ {
		k, v := kvfn()
		s.Put(k, v)
		verify[k] = v
	}
}

func BulkDelete(verify map[util.Comparable]util.T, s search.Searcher, cnt int, gen func() (util.Comparable, util.T)) {
	for i := 0; i < cnt; i++ {
		k, _ := gen()
		s.Del(k)
		delete(verify, k)
	}
}

func VerifyResult(t *testing.T, verify map[util.Comparable]util.T, s search.Searcher) {
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

	stdMap := make(map[util.Int]int)
	elapsed := util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			stdMap[util.Int(i)] = i
		}
	})
	t.Logf("stdMap.Put cost %v", elapsed)

	hm := hash_map.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(util.Int(i), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)

	avl := avltree.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(util.Int(i), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)

	//bt := binarytree.New()
	//elapsed = util.ExecCost(func() {
	//	for i := 0; i < benchmark; i++ {
	//		bt.Put(search.Int(i), i)
	//	}
	//})
	//t.Logf("binarytree.Put cost %v", elapsed)
}

func TestBenchmark_RandKeys(t *testing.T) {
	const benchmark = 20000000

	stdMap := make(map[util.Int]int)
	elapsed := util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			stdMap[util.Int(rand.Intn(n))] = i
		}
	})
	t.Logf("stdMap.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			_ = stdMap[util.Int(rand.Intn(n))]
		}
	})
	t.Logf("stdMap.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			delete(stdMap, util.Int(rand.Intn(n)))
		}
	})
	t.Logf("stdMap.Del cost %v", elapsed)

	hm := hash_map.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Put(util.Int(rand.Intn(n)), i)
		}
	})
	t.Logf("hashmap.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Get(util.Int(rand.Intn(n)))
		}
	})
	t.Logf("hashmap.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			hm.Del(util.Int(rand.Intn(n)))
		}
	})
	t.Logf("hashmap.Del cost %v", elapsed)

	avl := avltree.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Put(util.Int(rand.Intn(n)), i)
		}
	})
	t.Logf("avl.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Get(util.Int(rand.Intn(n)))
		}
	})
	t.Logf("avl.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			avl.Del(util.Int(rand.Intn(n)))
		}
	})
	t.Logf("avl.Del cost %v", elapsed)

	bt := binarytree.New()
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Put(util.Int(rand.Intn(n)), i)
		}
	})
	t.Logf("binarytree.Put cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Get(util.Int(rand.Intn(n)))
		}
	})
	t.Logf("binarytree.Get cost %v", elapsed)
	elapsed = util.ExecCost(func() {
		for i := 0; i < benchmark; i++ {
			bt.Del(util.Int(rand.Intn(n)))
		}
	})
	t.Logf("binarytree.Del cost %v", elapsed)
}

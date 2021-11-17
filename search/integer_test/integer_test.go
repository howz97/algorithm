package integer_test

import (
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
	"github.com/howz97/algorithm/search/binarytree"
	"math/rand"
	"strconv"
	"testing"
)

const n = 100

func TestAVL(t *testing.T) {
	DifferentKVType(t, func() search.Searcher {
		return avltree.New()
	})
}

func TestBinaryTree(t *testing.T) {
	DifferentKVType(t, func() search.Searcher {
		return binarytree.New()
	})
}

func DifferentKVType(t *testing.T, newFn func() search.Searcher) {
	s := newFn()
	LoopTest(t, s, IntStrKV)
	s = newFn()
	LoopTest(t, s, FloatIntKV)
	s = newFn()
	LoopTest(t, s, StrIntKV)
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
	k := alphabet.Ascii.RandString(20)
	v := rand.Intn(n)
	return search.Str(k), v
}

func LoopTest(t *testing.T, s search.Searcher, kvfn func() (search.Cmp, search.T)) {
	for i := 0; i < 1000; i++ {
		BulkInsert(t, s, n, kvfn)
		BulkDelete(t, s, n, kvfn)
	}
}

func BulkInsert(t *testing.T, s search.Searcher, cnt int, kvfn func() (search.Cmp, search.T)) {
	inserted := make(map[search.Cmp]search.T)
	for i := 0; i < cnt; i++ {
		k, v := kvfn()
		s.Insert(k, v)
		inserted[k] = v
	}
	for k, v := range inserted {
		vGot := s.Find(k)
		if vGot != v {
			t.Fatalf("get wrong value %v, should be %d", vGot, v)
		}
	}
}

func BulkDelete(t *testing.T, s search.Searcher, cnt int, gen func() (search.Cmp, search.T)) {
	deleted := make(map[search.Cmp]struct{})
	for i := 0; i < cnt; i++ {
		k, _ := gen()
		s.Delete(k)
		deleted[k] = struct{}{}
	}
	for k := range deleted {
		v := s.Find(k)
		if v != nil {
			t.Fatalf("delete failed")
		}
	}
}

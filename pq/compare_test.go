package pq

import (
	"fmt"
	"github.com/howz97/algorithm/pq/binomial"
	"github.com/howz97/algorithm/pq/heap"
	"github.com/howz97/algorithm/pq/leftist"
	. "github.com/howz97/algorithm/util"
	"math/rand"
	"testing"
)

const (
	mergeCnt  = 1000
	mergeSize = 10000
)

func TestCompareMerge(t *testing.T) {
	t.Logf("LeftistMerge cost %v", ExecCost(LeftistMerge))
	t.Logf("BinomialMerge cost %v", ExecCost(BinomialMerge))
}

func LeftistMerge() {
	left := leftist.New()
	for i := 0; i < mergeCnt; i++ {
		l2 := leftist.New()
		n := rand.Intn(mergeSize)
		for j := 0; j < n; j++ {
			l2.Push(Int(rand.Int()))
		}
		left.Merge(l2)
	}
	fmt.Printf("LeftistMerge final size %d \n", left.Size())
}

func BinomialMerge() {
	b := binomial.New()
	for i := 0; i < mergeCnt; i++ {
		l2 := binomial.New()
		n := rand.Intn(mergeSize)
		for j := 0; j < n; j++ {
			l2.Push(Int(rand.Int()))
		}
		b.Merge(l2)
	}
	fmt.Printf("BinomialMerge final size %d \n", b.Size())
}

const pushCnt = 10000000

func TestComparePush(t *testing.T) {
	t.Logf("HeapPush cost %v", ExecCost(HeapPush))
	t.Logf("LeftistPush cost %v", ExecCost(LeftistPush))
	t.Logf("BinomialPush cost %v", ExecCost(BinomialPush))
}

func HeapPush() {
	left := heap.New(pushCnt)
	for i := 0; i < pushCnt; i++ {
		left.Push(Int(i), i)
	}
}

func LeftistPush() {
	left := leftist.New()
	for i := 0; i < pushCnt; i++ {
		left.Push(Int(i))
	}
}

func BinomialPush() {
	left := binomial.New()
	for i := 0; i < pushCnt; i++ {
		left.Push(Int(i))
	}
}

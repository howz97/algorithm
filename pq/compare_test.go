package pq

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/howz97/algorithm/pq/binomial"
	"github.com/howz97/algorithm/pq/heap"
	"github.com/howz97/algorithm/pq/leftist"
	. "github.com/howz97/algorithm/util"
)

const (
	mergeCnt = 3000
)

func TestCompareMerge(t *testing.T) {
	t.Log("Start to compare Merge...")

	rand.Seed(1)
	t.Logf("LeftistMerge cost %v", ExecCost(LeftistMerge))
	rand.Seed(1)
	t.Logf("BinomialMerge cost %v", ExecCost(BinomialMerge))
}

func LeftistMerge() {
	left := leftist.New()
	for i := 0; i < mergeCnt; i++ {
		l2 := leftist.New()
		n := rand.Intn(mergeCnt)
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
		n := rand.Intn(mergeCnt)
		for j := 0; j < n; j++ {
			l2.Push(Int(rand.Int()))
		}
		b.Merge(l2)
	}
	fmt.Printf("BinomialMerge final size %d \n", b.Size())
}

const pushCnt = 1000000

func TestComparePush(t *testing.T) {
	t.Log("Start to compare Push...")

	rand.Seed(1)
	t.Logf("HeapPush cost %v", ExecCost(HeapPush))
	rand.Seed(1)
	t.Logf("LeftistPush cost %v", ExecCost(LeftistPush))
	rand.Seed(1)
	t.Logf("BinomialPush cost %v", ExecCost(BinomialPush))
}

func HeapPush() {
	left := heap.New[int, int](pushCnt)
	for i := 0; i < pushCnt; i++ {
		left.Push(rand.Int(), i)
	}
}

func LeftistPush() {
	left := leftist.New()
	for i := 0; i < pushCnt; i++ {
		left.Push(Int(rand.Int()))
	}
}

func BinomialPush() {
	left := binomial.New()
	for i := 0; i < pushCnt; i++ {
		left.Push(Int(rand.Int()))
	}
}

func TestComparePop(t *testing.T) {
	t.Log("Start to compare Pop...")

	rand.Seed(1)
	h := heap.New[int, int](0)
	for i := 0; i < pushCnt; i++ {
		h.Push(rand.Int(), i)
	}
	t.Logf("HeapPop cost %v", ExecCost(func() {
		for h.Size() > 0 {
			h.Pop()
		}
	}))

	rand.Seed(1)
	left := leftist.New()
	for i := 0; i < pushCnt; i++ {
		left.Push(Int(rand.Int()))
	}
	t.Logf("LeftistPop cost %v", ExecCost(func() {
		for left.Size() > 0 {
			left.Pop()
		}
	}))

	rand.Seed(1)
	b := binomial.New()
	for i := 0; i < pushCnt; i++ {
		b.Push(Int(rand.Int()))
	}
	t.Logf("BinomialPop cost %v", ExecCost(func() {
		for b.Size() > 0 {
			b.Pop()
		}
	}))
}

const (
	integerLoop = 300
	integerPush = 1000
)

func TestCompareInteger(t *testing.T) {
	t.Log("Start to compare Push&Merge&Pop...")

	rand.Seed(1)
	t.Logf("LeftistInteger cost %v", ExecCost(LeftistInteger))
	rand.Seed(1)
	t.Logf("BinomialInteger cost %v", ExecCost(BinomialInteger))
}

func LeftistInteger() {
	left := leftist.New()
	for cnt := 0; cnt < integerLoop; cnt++ {
		for i := 0; i < integerPush; i++ {
			left.Push(Int(rand.Int()))
		}
		n := rand.Intn(left.Size() * 2)
		l2 := leftist.New()
		for i := 0; i < n; i++ {
			l2.Push(Int(rand.Int()))
		}
		left.Merge(l2)
		n = rand.Intn(left.Size())
		for i := 0; i < n; i++ {
			left.Pop()
		}
	}
	fmt.Printf("final size %d \n", left.Size())
}

func BinomialInteger() {
	b := binomial.New()
	for cnt := 0; cnt < integerLoop; cnt++ {
		for i := 0; i < integerPush; i++ {
			b.Push(Int(rand.Int()))
		}
		n := rand.Intn(b.Size() * 2)
		b2 := binomial.New()
		for i := 0; i < n; i++ {
			b2.Push(Int(rand.Int()))
		}
		b.Merge(b2)
		n = rand.Intn(b.Size())
		for i := 0; i < n; i++ {
			b.Pop()
		}
	}
	fmt.Printf("final size %d \n", b.Size())
}

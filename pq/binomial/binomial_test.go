package binomial

import (
	"fmt"
	"math"
	"testing"
)

func Test_BQ(t *testing.T) {
	bq := New[int]()
	for i := 0; i <= 50; i++ {
		bq.Push(i)
	}
	bq1 := New[int]()
	for i := 51; i < 100; i++ {
		bq1.Push(i)
	}
	bq.Merge(bq1)
	for i := 0; i < 100; i++ {
		m := bq.Pop()
		if m != i {
			fmt.Printf("minimal element should be %v instead of %v\n", i, m)
			t.Fatal()
		}
	}
}

func Test_DelMin(t *testing.T) {
	bq := New[int]()
	var err error
	for i := 100; i < 200; i++ {
		bq.Push(i)
	}
	for i := 100; i < 200; i++ {
		m := bq.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if m != i {
			fmt.Printf("minimal element should be %v instead of %v\n", i, m)
			t.Fatal()
		}
	}
}

func Test_Insert(t *testing.T) {
	bq := New[int]()
	for i := 0; i < 198; i++ {
		bq.Push(i)
	}
	if bq.Size() != 198 {
		t.Fatal()
	}
	if !(bq.trees[0] == nil &&
		bq.trees[1] != nil &&
		bq.trees[2] != nil &&
		bq.trees[3] == nil &&
		bq.trees[4] == nil &&
		bq.trees[5] == nil &&
		bq.trees[6] != nil &&
		bq.trees[7] != nil) {
		t.Fatal()
	}
}

func Test_Merge(t *testing.T) {
	bq := newBQSize(90)
	bq1 := newBQSize(108)
	bq.Merge(bq1)
	fmt.Println(bq.Size())
	if bq.Size() != 90+108 {
		t.Fatal()
	}
	if !(bq.trees[0] == nil &&
		bq.trees[1] != nil &&
		bq.trees[2] != nil &&
		bq.trees[3] == nil &&
		bq.trees[4] == nil &&
		bq.trees[5] == nil &&
		bq.trees[6] != nil &&
		bq.trees[7] != nil) {
		t.Fatal()
	}
}

func newBQSize(size int) *Binomial[int] {
	if size < 0 {
		panic("size < 0")
	}
	maxTrees := int(math.Logb(float64(size))) + 1
	bq := New[int]()
	for i := 0; i <= maxTrees; i++ {
		if 1<<uint(i)&size != 0 {
			bq.trees[i] = binomialTree(i)
		}
	}
	bq.size = size
	return bq
}

func binomialTree(height int) *node[int] {
	if height < 0 {
		panic("height < 0")
	}
	if height == 0 {
		return &node[int]{
			p: 1,
		}
	}
	t1 := binomialTree(height - 1)
	t2 := binomialTree(height - 1)
	t2.sibling = t1.son
	t1.son = t2
	return t1
}

func TestExample(t *testing.T) {
	ExampleBinomial()
}

func ExampleBinomial() {
	b := New[int]()
	b.Push(1)
	b.Push(9)
	b.Push(9)
	b.Push(7)
	b2 := New[int]()
	b2.Push(13)
	b2.Push(11)
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	//Output: 1,7,9,9,11,13,
}

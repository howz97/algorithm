package binomial

import (
	"fmt"
	"math"
	"testing"
)

func Test_BQ(t *testing.T) {
	bq := New()
	var err error
	for i := 0; i <= 50; i++ {
		err = bq.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	bq1 := NewWithMaxTrees(33)
	for i := 51; i < 100; i++ {
		err = bq1.Insert(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	err = bq.Merge(bq1)
	if err != nil {
		t.Fatal(err)
	}
	m := 0
	for i := 0; i < 100; i++ {
		m, err = bq.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if m != i {
			fmt.Printf("minimal element should be %v instead of %v\n", i, m)
			t.Fatal()
		}
		fmt.Printf("%v deleted!!!\n", m)
	}
	if !bq.IsEmpty() {
		t.Fatal()
	}
}

func Test_DelMin(t *testing.T) {
	bq := New()
	var err error
	for i := 100; i < 200; i++ {
		err = bq.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	m := 0
	for i := 100; i < 200; i++ {
		m, err = bq.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if m != i {
			fmt.Printf("minimal element should be %v instead of %v\n", i, m)
			t.Fatal()
		}
		fmt.Printf("%v deleted!!! left-size:%v\n", m, bq.Size())
	}
}

func Test_Insert(t *testing.T) {
	bq := New()
	var err error
	for i := 0; i < 198; i++ {
		err = bq.Push(i)
		if err != nil {
			t.Fatal(err)
		}
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
	err := bq.Merge(bq1)
	if err != nil {
		t.Fatal(err)
	}
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

func newBQSize(size int) *Binomial {
	if size < 0 {
		panic("size < 0")
	}
	maxTrees := int(math.Logb(float64(size))) + 1
	bq := New()
	for i := 0; i <= maxTrees; i++ {
		if 1<<uint(i)&size != 0 {
			bq.trees[i] = binomialTree(i)
		}
	}
	bq.size = size
	return bq
}

func binomialTree(height int) *node {
	if height < 0 {
		panic("height < 0")
	}
	if height == 0 {
		return &node{
			k: 1,
		}
	}
	t1 := binomialTree(height - 1)
	t2 := binomialTree(height - 1)
	t2.nextSibling = t1.leftSon
	t1.leftSon = t2
	return t1
}

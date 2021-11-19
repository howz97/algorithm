package binarytree

import (
	"github.com/howz97/algorithm/search"
	"testing"
)

type student struct {
	id   int
	name string
}

func (s *student) ID() int {
	return s.id
}

func Test_searchtree(t *testing.T) {
	searchT := New()
	for i := 1; i <= 7; i++ {
		searchT.Insert(search.Integer(i), i)
	}
	for i := 16; i >= 10; i-- {
		searchT.Insert(search.Integer(i), i)
	}
	searchT.Insert(search.Integer(8), 8)
	searchT.Insert(search.Integer(9), 9)

	if searchT.FindMin().(int) != 1 {
		t.Fatal("FindMin failed")
	}
	if searchT.FindMax().(int) != 16 {
		t.Fatal("FindMax failed")
	}
	for i := 1; i <= 16; i++ {
		if searchT.Find(search.Integer(i)).(int) != i {
			t.Fatal("Find failed")
		}
		searchT.Delete(search.Integer(i))
		if searchT.Find(search.Integer(i)) != nil {
			t.Fatal("Delete failed")
		}
	}
}

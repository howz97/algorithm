package searchtree

import (
	"fmt"
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
	st := NewSearchTree(&student{id: 10, name: "s10"})
	Insert(st, &student{id: 6, name: "s6"})
	Insert(st, &student{id: 7, name: "s7"})
	Insert(st, &student{id: 13, name: "s13"})
	Insert(st, &student{id: 5, name: "s5"})
	Insert(st, &student{id: 9, name: "s9"})
	Insert(st, &student{id: 2, name: "s2"})
	Insert(st, &student{id: 14, name: "s14"})
	Insert(st, &student{id: 15, name: "s15"})
	Insert(st, &student{id: 11, name: "s11"})
	Insert(st, &student{id: 1, name: "s1"})

	fmt.Println(st.Elem.ID())
	fmt.Println(st.leftSon.Elem.ID())
	fmt.Println(st.rightSon.Elem.ID())

	s9node := Find(st, 9)
	s9 := s9node.Elem.(*student)
	if s9.name != "s9" {
		t.Fatal("find s9")
	}

	minnode := FindMin(st)
	min := minnode.Elem.(*student)
	if min.name != "s1" {
		t.Fatal("find min")
	}

	maxnode := FindMax(st)
	max := maxnode.Elem.(*student)
	if max.name != "s15" {
		t.Fatal("find max")
	}

	DeleteMin(st)
	minnode = FindMin(st)
	min = minnode.Elem.(*student)
	if min.name != "s2" {
		t.Fatal("find second small")
	}

	DeleteMax(st)
	maxnode = FindMax(st)
	max = maxnode.Elem.(*student)
	if max.name != "s14" {
		t.Fatal("find second big")
	}

	Delete(st, 9)
	s9node = Find(st, 9)
	if s9node != nil {
		t.Fatal("delete s9")
	}
	Delete(st, 10)
	s10node := Find(st, 10)
	if s10node != nil {
		t.Fatal("delete s10")
	}
}

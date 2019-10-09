package ewd

import (
	"fmt"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	g, err := ImportEWD("./tinyEWDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	s := TopologicalSort(g)
	for !s.IsEmpty() {
		fmt.Print(s.Pop(), " ")
	}
}

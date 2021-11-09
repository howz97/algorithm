package weighted_digraph

import (
	"fmt"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	g, err := ImportEWD("./tinyEWDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	s := g.TopologicalSort()
	for !s.IsEmpty() {
		fmt.Print(s.Pop(), " ")
	}
}
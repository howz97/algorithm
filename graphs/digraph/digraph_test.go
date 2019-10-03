package digraph

import (
	"fmt"
	"testing"
)

func TestSCC_IsStronglyConnected(t *testing.T) {
	g := NewDigraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(5, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 2)
	g.AddEdge(3, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 0)
	g.AddEdge(6, 0)
	g.AddEdge(6, 4)
	g.AddEdge(6, 9)
	g.AddEdge(9, 10)
	g.AddEdge(10, 12)
	g.AddEdge(12, 9)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)
	g.AddEdge(11, 4)
	g.AddEdge(7, 6)
	g.AddEdge(7, 8)
	g.AddEdge(8, 7)
	g.AddEdge(8, 9)
	fmt.Println("number of edge: ", g.NumEdge())
	scc := NewSCC(g)
	fmt.Println("number of SCC:", scc.NumSCC())
	for i := 0; i < g.NumV(); i++ {
		fmt.Printf("SCC ID of vertical(%v): %v\n", i, scc.GetID(i))
	}
	if !scc.IsStronglyConnected(1, 1) {
		t.Fatal()
	}
	if !scc.IsStronglyConnected(0, 4) {
		t.Fatal()
	}
	if !scc.IsStronglyConnected(9, 11) {
		t.Fatal()
	}
	if scc.IsStronglyConnected(1, 0) {
		t.Fatal()
	}
	if scc.IsStronglyConnected(11, 8) {
		t.Fatal()
	}
}

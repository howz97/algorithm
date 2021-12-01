package digraph

import (
	"fmt"
	"testing"
)

func TestSCC_IsStronglyConnected(t *testing.T) {
	g := New(13)
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
	scc := g.SCC()
	fmt.Println("number of SCC:", scc.NumComponents())
	for i := 0; i < g.NumVertical(); i++ {
		fmt.Printf("SCC ID of vertical(%v): %v\n", i, scc.GetCompID(i))
	}
	if !scc.IsStronglyConn(1, 1) {
		t.Fatal()
	}
	if !scc.IsStronglyConn(0, 4) {
		t.Fatal()
	}
	if !scc.IsStronglyConn(9, 11) {
		t.Fatal()
	}
	if scc.IsStronglyConn(1, 0) {
		t.Fatal()
	}
	if scc.IsStronglyConn(11, 8) {
		t.Fatal()
	}
}
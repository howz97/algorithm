package unionfind

import (
	"fmt"
	"testing"
)

func TestUnionFind_Union(t *testing.T) {
	uf := New(8)
	fmt.Printf("NumConnectedComponent: %v\n", uf.NumConnectedComponent())
	uf.Union(1, 7)
	uf.Union(7, 0)
	uf.Union(0, 2)
	uf.Union(2, 3)
	uf.Union(5, 7)
	fmt.Println(uf.IsConnected(1, 3))
	fmt.Printf("NumConnectedComponent: %v\n", uf.NumConnectedComponent())
}

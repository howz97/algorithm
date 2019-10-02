package unionfind

import (
	"fmt"
	"testing"
)

func TestUnionFind_Union(t *testing.T) {
	uf := NewUF(10)
	fmt.Printf("NumConnectedComponent: %v\n", uf.NumConnectedComponent())
	fmt.Println(uf.IsConnected(1,2))
	uf.Union(1,2)
	uf.Union(0,1)
	fmt.Println(uf.IsConnected(0,2))
	uf.Union(3,4)
	uf.Union(3,5)
	fmt.Printf("NumConnectedComponent: %v\n", uf.NumConnectedComponent())
}

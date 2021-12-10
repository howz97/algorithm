package graphs

import (
	"fmt"
	"testing"
)

func TestEdgeWeightedGraph_Kruskal(t *testing.T) {
	g := NewWGraph(8) // 算法4th P399 图4.3.10, 不同的是这里权重使用int
	g.AddEdge(0, 2, 26)
	g.AddEdge(0, 4, 38)
	g.AddEdge(0, 6, 58)
	g.AddEdge(0, 7, 17)
	g.AddEdge(1, 2, 36)
	g.AddEdge(1, 3, 29)
	g.AddEdge(1, 5, 32)
	g.AddEdge(1, 7, 19)
	g.AddEdge(2, 3, 17)
	g.AddEdge(2, 6, 40)
	g.AddEdge(2, 7, 34)
	g.AddEdge(3, 6, 52)
	g.AddEdge(4, 5, 35)
	g.AddEdge(4, 6, 93)
	g.AddEdge(4, 7, 37)
	g.AddEdge(5, 7, 28)
	fmt.Printf("Number of edges: %v\n", g.NumEdge())

	fmt.Println("MST:")
	mst := g.Kruskal()
	for !mst.IsEmpty() {
		v := mst.Front()
		e := v.(*Edge)
		fmt.Printf("%v--%v, %v\n", e.from, e.to, e.weight)
	}
}

func TestEdgeWeightedGraph_LazyPrim(t *testing.T) {
	g := NewWGraph(8) // 算法4th P399 图4.3.10, 不同的是这里权重使用int
	g.AddEdge(0, 2, 26)
	g.AddEdge(0, 4, 38)
	g.AddEdge(0, 6, 58)
	g.AddEdge(0, 7, 17)
	g.AddEdge(1, 2, 36)
	g.AddEdge(1, 3, 29)
	g.AddEdge(1, 5, 32)
	g.AddEdge(1, 7, 19)
	g.AddEdge(2, 3, 17)
	g.AddEdge(2, 6, 40)
	g.AddEdge(2, 7, 34)
	g.AddEdge(3, 6, 52)
	g.AddEdge(4, 5, 35)
	g.AddEdge(4, 6, 93)
	g.AddEdge(4, 7, 37)
	g.AddEdge(5, 7, 28)

	t.Logf("MST:\n %s \n", g.LazyPrim().String())
}

func TestEdgeWeightedGraph_Prim(t *testing.T) {
	g := NewWGraph(8) // 算法4th P399 图4.3.10, 不同的是这里权重使用int
	g.AddEdge(0, 2, 26)
	g.AddEdge(0, 4, 8)
	g.AddEdge(0, 6, 8)
	g.AddEdge(0, 7, 7)
	g.AddEdge(1, 2, 6)
	g.AddEdge(1, 3, 9)
	g.AddEdge(1, 5, 2)
	g.AddEdge(1, 7, 9)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 6, 0)
	g.AddEdge(2, 7, 4)
	g.AddEdge(3, 6, 2)
	g.AddEdge(4, 5, 5)
	g.AddEdge(4, 6, 3)
	g.AddEdge(4, 7, 7)
	g.AddEdge(5, 7, 8)
}

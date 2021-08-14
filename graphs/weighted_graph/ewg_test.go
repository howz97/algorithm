package weighted_graph

import (
	"fmt"
	"testing"
)

func TestEdgeWeightedGraph_Kruskal(t *testing.T) {
	g := NewEWG(8) // 算法4th P399 图4.3.10, 不同的是这里权重使用int
	g.AddEdge(&Edge{v: 0, w: 2, weight: 26})
	g.AddEdge(&Edge{v: 0, w: 4, weight: 38})
	g.AddEdge(&Edge{v: 0, w: 6, weight: 58})
	g.AddEdge(&Edge{v: 0, w: 7, weight: 17})
	g.AddEdge(&Edge{v: 1, w: 2, weight: 36})
	g.AddEdge(&Edge{v: 1, w: 3, weight: 29})
	g.AddEdge(&Edge{v: 1, w: 5, weight: 32})
	g.AddEdge(&Edge{v: 1, w: 7, weight: 19})
	g.AddEdge(&Edge{v: 2, w: 3, weight: 17})
	g.AddEdge(&Edge{v: 2, w: 6, weight: 40})
	g.AddEdge(&Edge{v: 2, w: 7, weight: 34})
	g.AddEdge(&Edge{v: 3, w: 6, weight: 52})
	g.AddEdge(&Edge{v: 4, w: 5, weight: 35})
	g.AddEdge(&Edge{v: 4, w: 6, weight: 93})
	g.AddEdge(&Edge{v: 4, w: 7, weight: 37})
	g.AddEdge(&Edge{v: 5, w: 7, weight: 28})
	fmt.Printf("Number of edges: %v\n", g.NumE())
	all := g.AllEdges()
	if all.Size() != g.NumE() {
		t.Fatal()
	}

	fmt.Println("MST:")
	mst := g.Kruskal()
	for !mst.IsEmpty() {
		v := mst.Front()
		e := v.(*Edge)
		fmt.Printf("%v--%v, %v\n", e.v, e.w, e.weight)
	}
}

func TestEdgeWeightedGraph_LazyPrim(t *testing.T) {
	g := NewEWG(8) // 算法4th P399 图4.3.10, 不同的是这里权重使用int
	g.AddEdge(&Edge{v: 0, w: 2, weight: 26})
	g.AddEdge(&Edge{v: 0, w: 4, weight: 38})
	g.AddEdge(&Edge{v: 0, w: 6, weight: 58})
	g.AddEdge(&Edge{v: 0, w: 7, weight: 17})
	g.AddEdge(&Edge{v: 1, w: 2, weight: 36})
	g.AddEdge(&Edge{v: 1, w: 3, weight: 29})
	g.AddEdge(&Edge{v: 1, w: 5, weight: 32})
	g.AddEdge(&Edge{v: 1, w: 7, weight: 19})
	g.AddEdge(&Edge{v: 2, w: 3, weight: 17})
	g.AddEdge(&Edge{v: 2, w: 6, weight: 40})
	g.AddEdge(&Edge{v: 2, w: 7, weight: 34})
	g.AddEdge(&Edge{v: 3, w: 6, weight: 52})
	g.AddEdge(&Edge{v: 4, w: 5, weight: 35})
	g.AddEdge(&Edge{v: 4, w: 6, weight: 93})
	g.AddEdge(&Edge{v: 4, w: 7, weight: 37})
	g.AddEdge(&Edge{v: 5, w: 7, weight: 28})

	mstf := g.LazyPrim()
	fmt.Println("MST:")
	mst := mstf.MST(0)
	for !mst.IsEmpty() {
		v := mst.Front()
		e := v.(*Edge)
		fmt.Printf("%v--%v, %v\n", e.v, e.w, e.weight)
	}
}

func TestEdgeWeightedGraph_Prim(t *testing.T) {
	g := NewEWG(8) // 算法4th P399 图4.3.10, 不同的是这里权重使用int
	g.AddEdge(&Edge{v: 0, w: 2, weight: 26})
	g.AddEdge(&Edge{v: 0, w: 4, weight: 38})
	g.AddEdge(&Edge{v: 0, w: 6, weight: 58})
	g.AddEdge(&Edge{v: 0, w: 7, weight: 17})
	g.AddEdge(&Edge{v: 1, w: 2, weight: 36})
	g.AddEdge(&Edge{v: 1, w: 3, weight: 29})
	g.AddEdge(&Edge{v: 1, w: 5, weight: 32})
	g.AddEdge(&Edge{v: 1, w: 7, weight: 19})
	g.AddEdge(&Edge{v: 2, w: 3, weight: 17})
	g.AddEdge(&Edge{v: 2, w: 6, weight: 40})
	g.AddEdge(&Edge{v: 2, w: 7, weight: 34})
	g.AddEdge(&Edge{v: 3, w: 6, weight: 52})
	g.AddEdge(&Edge{v: 4, w: 5, weight: 35})
	g.AddEdge(&Edge{v: 4, w: 6, weight: 93})
	g.AddEdge(&Edge{v: 4, w: 7, weight: 37})
	g.AddEdge(&Edge{v: 5, w: 7, weight: 28})

	mstf := g.Prim()
	fmt.Println("MST:")
	mst := mstf.MST(0)
	for !mst.IsEmpty() {
		v := mst.Front()
		e := v.(*Edge)
		fmt.Printf("%v--%v, %v\n", e.v, e.w, e.weight)
	}
}

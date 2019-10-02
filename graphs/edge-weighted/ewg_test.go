package ewg

import (
	"fmt"
	"testing"
)

func TestEdgeWeightedGraph(t *testing.T) {
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

	fmt.Println("MST:")
	mst := g.Kruskal()
	fmt.Println(mst.IsEmpty())
	for !mst.IsEmpty() {
		v, _ := mst.Front()
		e := v.(*Edge)
		fmt.Printf("%v--%v, %v\n", e.v, e.w, e.weight)
	}
}

package edge_weighted

import (
	"errors"
	"fmt"
	"github.com/zh1014/algorithm/queue"
)

var (
	errVerticalNotExist   = errors.New("vertical not exist")
	errNotSupportSelfLoop = errors.New("not support self loop")
)

type EdgeWeightedGraph []edgeSet

func (ewg EdgeWeightedGraph) NumV() int {
	return len(ewg)
}

func (ewg EdgeWeightedGraph) NumE() int {
	nume := 0
	for i := range ewg {
		nume += ewg[i].len()
	}
	return nume / 2
}

func (ewg EdgeWeightedGraph) AddEdge(e *edge) {
	v1 := e.getOne()
	v2 := e.getAnother(v1)
	if !ewg.HasV(v1) || !ewg.HasV(v2) {
		panic(errVerticalNotExist)
	}
	if v1 == v2 {
		panic(errNotSupportSelfLoop)
	}
	ewg[v1].add(e)
	ewg[v2].add(e)
}

func (ewg EdgeWeightedGraph) Adjacent(v int) []*edge {
	if !ewg.HasV(v) {
		panic(errVerticalNotExist)
	}
	return ewg[v].traverse()
}

func (ewg EdgeWeightedGraph) AllEdges() []*edge {
	marked := make([]bool, ewg.NumV())
	edgesQ := queue.NewInterfaceQ()
	for i, b := range marked {
		if !b {
			ewg.dfsAllEdges(i, marked, edgesQ)
		}
	}
	edges := make([]*edge, ewg.NumV())
	for i := 0; !edgesQ.IsEmpty(); i++ {
		e := edgesQ.Front().(*edge)
		edges[i] = e
	}
	return edges
}

func (ewg EdgeWeightedGraph) dfsAllEdges(v int, marked []bool, edges *queue.InterfaceQ) {
	adj := ewg.Adjacent(v)
	for _, e := range adj {
		v2 := e.getAnother(v)
		if !marked[v2] {
			edges.PushBack(e)
			marked[v2] = true
			ewg.dfsAllEdges(v2, marked, edges)
		}
	}
}

func (ewg EdgeWeightedGraph) HasV(v int) bool {
	return v >= 0 && v < ewg.NumV()
}

type edge struct {
	v, w   int
	weight float64
}

func (e *edge) getOne() int {
	return e.v
}

func (e *edge) getAnother(v int) int {
	if v == e.v {
		return e.w
	} else if v == e.w {
		return e.v
	} else {
		panic(fmt.Sprintf("edge %v-%v(%v) does not contains vertical %v", e.v, e.w, e.weight, v))
	}
}

func (e *edge) getWeight() float64 {
	return e.weight
}

func (ewg EdgeWeightedGraph) LazyPrim() []*edge {

}

func (ewg EdgeWeightedGraph) Prim() []*edge {

}

func (ewg EdgeWeightedGraph) Kruskal() []*edge {

}

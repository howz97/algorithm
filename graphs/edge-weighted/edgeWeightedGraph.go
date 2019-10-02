package edge_weighted

import (
	"errors"
	"fmt"
	pqueue "github.com/zh1014/algorithm/pqueue/binaryheap"
	"github.com/zh1014/algorithm/queue"
	"math"
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
	edgesQ := queue.NewQueen(ewg.NumV())
	for i, b := range marked {
		if !b {
			ewg.dfsAllEdges(i, marked, edgesQ)
		}
	}
	edges := make([]*edge, ewg.NumV())
	for i := 0; !edgesQ.IsEmpty(); i++ {
		f,_ := edgesQ.Front()
		edges[i] = f.(*edge)
	}
	return edges
}

func (ewg EdgeWeightedGraph) dfsAllEdges(v int, marked []bool, edges *queue.Queen) {
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
	weight int
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

func (e *edge) getWeight() int {
	return e.weight
}

func (ewg EdgeWeightedGraph) LazyPrim() *MSTForest {
	marked := make([]bool, ewg.NumV())
	f := newMSTForest()
	for i, b := range marked {
		if !b {
			f.addMST(lazyPrim(ewg, i, marked))
		}
	}
	return f
}

func lazyPrim(g EdgeWeightedGraph, v int,marked []bool) []*edge {
	pq := pqueue.NewBinHeap(g.NumV()-1)
	marked[v] = true
	vadj := g.Adjacent(v)
	mst := make([]*edge, 0)
	for i := range vadj {
		pq.Insert(vadj[i].weight, vadj[i])
	}
	for !pq.IsEmpty() {
		_, m := pq.DelMin()
		e := m.(*edge)
		if marked[e.v] && marked[e.w] {
			continue
		}
		mst = append(mst, e)
		if !marked[e.v] {
			lazyPrimVisit(g, e.v, marked, pq)
		}
		if !marked[e.w] {
			lazyPrimVisit(g, e.w, marked, pq)
		}
	}
	return mst
}

func lazyPrimVisit(g EdgeWeightedGraph, v int, marked []bool, pq *pqueue.BinHeap) {
	marked[v] = true
	vadj := g.Adjacent(v)
	for _, e := range vadj {
		if !marked[e.getAnother(v)] {
			pq.Insert(e.weight, e)
		}
	}
}

func (ewg EdgeWeightedGraph) Prim() *MSTForest {
	marked := make([]bool, ewg.NumV())
	distTo := make([]int, ewg.NumV())
	for i := range distTo {
		distTo[i] = math.MaxInt64
	}
	f := newMSTForest()
	for i, b := range marked {
		if !b {
			f.addMST(prim(ewg, i, marked, distTo))
		}
	}
	return f
}

func prim(g EdgeWeightedGraph, v int,marked []bool, distTo []int) []*edge {
	pq := pqueue.NewBinHeap(g.NumV()-1)
	marked[v] = true
	vadj := g.Adjacent(v)
	mst := make([]*edge, 0)
	for i := range vadj {
		pq.Insert(vadj[i].weight, vadj[i])
	}
	for !pq.IsEmpty() {
		_, m := pq.DelMin()
		e := m.(*edge)
		mst = append(mst, e)
		if !marked[e.v] {
			primVisit(g, e.v, marked, pq, distTo)
		}
		if !marked[e.w] {
			primVisit(g, e.w, marked, pq, distTo)
		}
	}
	return mst
}

func primVisit(g EdgeWeightedGraph, v int, marked []bool, pq *pqueue.BinHeap, distTo []int) {
	marked[v] = true
	vadj := g.Adjacent(v)
	for _, e := range vadj {
		w := e.getAnother(v)
		if !marked[w]&& e.weight < distTo[w] {
			distTo[w] = e.weight
			pq.Insert(e.weight, e)
		}
	}
}

func (ewg EdgeWeightedGraph) Kruskal() *MSTForest {

}

type MSTForest [][]*edge

func newMSTForest() *MSTForest {
	f := make(MSTForest, 0, 1)
	return &f
}

func (f *MSTForest) MST(sub int) []*edge {
	if sub < 0 ||sub >= len(*f) {
		panic(fmt.Sprintf("subgraph %v does not exist", sub))
	}
	return (*f)[sub]
}

func (f *MSTForest) addMST(mst []*edge) {
	*f = append(*f, mst)
}

func (f *MSTForest) NumSubgraph() int {
	return len(*f)
}

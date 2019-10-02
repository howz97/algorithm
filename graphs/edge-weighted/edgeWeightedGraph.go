package ewg

import (
	"errors"
	"fmt"
	pqueue "github.com/zh1014/algorithm/pqueue/binaryheap"
	"github.com/zh1014/algorithm/queue"
	unionfind "github.com/zh1014/algorithm/union-find"
	"math"
)

var (
	ErrVerticalNotExist   = errors.New("vertical not exist")
	ErrNotSupportSelfLoop = errors.New("not support self loop")
)

type EdgeWeightedGraph []edgeSet

func NewEWG(numV int) EdgeWeightedGraph {
	g := make(EdgeWeightedGraph, numV)
	for i := range g {
		g[i] = NewEdgeSet()
	}
	return g
}

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

func (ewg EdgeWeightedGraph) AddEdge(e *Edge) {
	v1 := e.EitherV()
	v2 := e.Another(v1)
	if !ewg.HasV(v1) || !ewg.HasV(v2) {
		panic(ErrVerticalNotExist)
	}
	if v1 == v2 {
		panic(ErrNotSupportSelfLoop)
	}
	ewg[v1].add(e)
	ewg[v2].add(e)
}

func (ewg EdgeWeightedGraph) Adjacent(v int) []*Edge {
	if !ewg.HasV(v) {
		panic(ErrVerticalNotExist)
	}
	return ewg[v].traverse()
}

func (ewg EdgeWeightedGraph) AllEdges() []*Edge {
	marked := make([]bool, ewg.NumV())
	edgesQ := queue.NewQueen(ewg.NumV())
	for i, b := range marked {
		if !b {
			ewg.dfsAllEdges(i, marked, edgesQ)
		}
	}
	edges := make([]*Edge, ewg.NumV())
	for i := 0; !edgesQ.IsEmpty(); i++ {
		f, _ := edgesQ.Front()
		edges[i] = f.(*Edge)
	}
	return edges
}

func (ewg EdgeWeightedGraph) dfsAllEdges(v int, marked []bool, edges *queue.Queen) {
	adj := ewg.Adjacent(v)
	for _, e := range adj {
		v2 := e.Another(v)
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

type Edge struct {
	v, w   int
	weight int
}

func (e *Edge) EitherV() int {
	return e.v
}

func (e *Edge) Another(v int) int {
	if v == e.v {
		return e.w
	} else if v == e.w {
		return e.v
	} else {
		panic(fmt.Sprintf("Edge %v-%v(%v) does not contains vertical %v", e.v, e.w, e.weight, v))
	}
}

func (e *Edge) GetWeight() int {
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

func lazyPrim(g EdgeWeightedGraph, v int, marked []bool) *queue.LinkedQueue {
	pq := pqueue.NewBinHeap(g.NumV() - 1)
	marked[v] = true
	vadj := g.Adjacent(v)
	mst := queue.NewLinkedQueue()
	for i := range vadj {
		pq.Insert(vadj[i].weight, vadj[i])
	}
	for !pq.IsEmpty() {
		m := pq.DelMin()
		e := m.(*Edge)
		if marked[e.v] && marked[e.w] {
			continue
		}
		mst.PushBack(e)
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
		if !marked[e.Another(v)] {
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

func prim(g EdgeWeightedGraph, v int, marked []bool, distTo []int) *queue.LinkedQueue {
	pq := pqueue.NewBinHeap(g.NumV() - 1)
	marked[v] = true
	vadj := g.Adjacent(v)
	mst := queue.NewLinkedQueue()
	for i := range vadj {
		pq.Insert(vadj[i].weight, vadj[i])
	}
	for !pq.IsEmpty() {
		m := pq.DelMin()
		e := m.(*Edge)
		mst.PushBack(e)
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
		w := e.Another(v)
		if !marked[w] && e.weight < distTo[w] {
			distTo[w] = e.weight
			pq.Insert(e.weight, e)
		}
	}
}

func (ewg EdgeWeightedGraph) Kruskal() *queue.Queen {
	mst := queue.NewQueen(ewg.NumV())
	uf := unionfind.NewUF(ewg.NumV())
	pq := pqueue.NewBinHeap(ewg.NumE())
	allEdge := ewg.AllEdges()
	for _, e := range allEdge {
		pq.Insert(e.weight, e)
	}
	for !pq.IsEmpty() {
		min := pq.DelMin()
		minE := min.(*Edge)
		if uf.IsConnected(minE.v, minE.w) {
			continue
		}
		uf.Union(minE.v, minE.w)
		mst.PushBack(minE)
	}
	return mst
}

type MSTForest []*queue.LinkedQueue

func newMSTForest() *MSTForest {
	f := make(MSTForest, 0, 1)
	return &f
}

// 根据连通分量的id获取它的最小生成树
func (f *MSTForest) MST(cc int) *queue.LinkedQueue {
	if cc < 0 || cc >= len(*f) {
		panic(fmt.Sprintf("subgraph %v does not exist", cc))
	}
	return (*f)[cc]
}

func (f *MSTForest) addMST(mst *queue.LinkedQueue) {
	*f = append(*f, mst)
}

func (f *MSTForest) NumConnectedComponent() int {
	return len(*f)
}

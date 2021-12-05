package wgraph

import (
	"errors"
	"fmt"
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/graphs/graph"
	pqueue "github.com/howz97/algorithm/pqueue/binaryheap"
	"github.com/howz97/algorithm/queue"
	unionfind "github.com/howz97/algorithm/union-find"
)

var (
	ErrVerticalNotExist   = errors.New("vertical not exist")
	ErrNotSupportSelfLoop = errors.New("not support self loop")
)

type WGraph struct {
	graph.Graph
	graphs.Weight
}

func NewWGraph(size int) *WGraph {
	return &WGraph{
		Graph: *graph.New(size),
		Weight: graphs.NewWeight(size),
	}
}

func (g WGraph) AddEdge(src, dst int, w float64) error {
	if err := g.Graph.AddEdge(src, dst); err != nil {
		return err
	}
	g.SetWeight(src, dst, w)
	g.SetWeight(dst, src, w)
	return nil
}

func (g WGraph) Adjacent(v int) []*Edge {
	if !g.HasVertical(v) {
		panic(ErrVerticalNotExist)
	}
	return g[v].traverse()
}

func (g WGraph) AllEdges() *queue.Queen {
	edges := queue.NewQueen(g.NumEdge())
	for i := range g {
		adj := g.Adjacent(i)
		for _, e := range adj {
			if e.Another(i) > i { // self-loop not supported
				edges.PushBack(e)
			}
		}
	}
	return edges
}

func (g WGraph) LazyPrim() *MSTForest {
	marked := make([]bool, g.NumVertical())
	f := newMSTForest()
	for i, b := range marked {
		if !b {
			f.addMST(lazyPrim(g, i, marked))
		}
	}
	return f
}

func lazyPrim(g WGraph, v int, marked []bool) *queue.LinkedQueue {
	pq := pqueue.NewBinHeap(g.NumEdge())
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

func lazyPrimVisit(g WGraph, v int, marked []bool, pq *pqueue.BinHeap) {
	marked[v] = true
	vadj := g.Adjacent(v)
	for _, e := range vadj {
		if !marked[e.Another(v)] {
			pq.Insert(e.weight, e)
		}
	}
}

func (g WGraph) Prim() *MSTForest {
	marked := make([]bool, g.NumVertical())
	edgeTo := make([]*Edge, g.NumVertical())
	f := newMSTForest()
	for i, b := range marked {
		if !b {
			f.addMST(prim(g, i, marked, edgeTo))
		}
	}
	return f
}

func prim(g WGraph, v int, marked []bool, edgeTo []*Edge) *queue.LinkedQueue {
	pq := pqueue.NewBinHeap(g.NumVertical() - 1)
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		w := e.Another(v)
		pq.Insert(e.weight, w)
		edgeTo[w] = e
	}
	mst := queue.NewLinkedQueue()
	for !pq.IsEmpty() {
		m := pq.DelMin()
		w := m.(int)
		mst.PushBack(edgeTo[w])
		primVisit(g, w, marked, pq, edgeTo)
	}
	return mst
}

func primVisit(g WGraph, v int, marked []bool, pq *pqueue.BinHeap, edgeTo []*Edge) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		w := e.Another(v)
		if marked[w] {
			continue
		}
		if edgeTo[w] == nil {
			pq.Insert(e.weight, w)
			edgeTo[w] = e
		} else if e.weight < edgeTo[w].weight {
			pq.Update(e.weight, w)
			edgeTo[w] = e
		}
	}
}

// Kruskal 该实现仅支持连通图
func (g WGraph) Kruskal() *queue.Queen {
	mst := queue.NewQueen(g.NumVertical() - 1)
	uf := unionfind.NewUF(g.NumVertical())
	pq := pqueue.NewBinHeap(g.NumEdge())

	allEdge := g.AllEdges()
	for !allEdge.IsEmpty() {
		e := allEdge.Front().(*Edge)
		pq.Insert(e.weight, e)
	}
	for !mst.IsFull() {
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

package graphs

import (
	"fmt"
	pqueue "github.com/howz97/algorithm/pq/heap"
	"github.com/howz97/algorithm/queue"
	unionfind "github.com/howz97/algorithm/union-find"
	"github.com/howz97/algorithm/util"
)

type WGraph struct {
	Graph
}

func NewWGraph(size uint) *WGraph {
	return &WGraph{
		Graph: *NewGraph(size),
	}
}

func (g WGraph) AddEdge(src, dst int, w float64) error {
	return g.Graph.addWeightedEdge(src, dst, w)
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
	pq := pqueue.New(g.NumEdge())
	marked[v] = true
	mst := queue.NewLinkedQueue()
	g.iterateAdj(v, func(src int, dst int, w float64) bool {
		pq.Push(util.Float(w), &Edge{ // fixme pq Cmp
			from:   src,
			to:     dst,
			weight: w,
		})
		return true
	})
	for !pq.IsEmpty() {
		m := pq.Pop()
		e := m.(*Edge)
		if marked[e.from] && marked[e.to] {
			continue
		}
		mst.PushBack(e)
		if !marked[e.from] {
			lazyPrimVisit(g, e.from, marked, pq)
		}
		if !marked[e.to] {
			lazyPrimVisit(g, e.to, marked, pq)
		}
	}
	return mst
}

func lazyPrimVisit(g WGraph, v int, marked []bool, pq *pqueue.Heap) {
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, w float64) bool {
		if !marked[a] {
			pq.Push(util.Float(w), &Edge{
				from:   v,
				to:     a,
				weight: w,
			})
		}
		return true
	})
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
	pq := pqueue.New(g.NumVertical())
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, w float64) bool {
		pq.Push(util.Float(w), a)
		edgeTo[a] = &Edge{
			from:   v,
			to:     a,
			weight: w,
		}
		return true
	})
	mst := queue.NewLinkedQueue()
	for !pq.IsEmpty() {
		m := pq.Pop()
		w := m.(int)
		mst.PushBack(edgeTo[w])
		primVisit(g, w, marked, pq, edgeTo)
	}
	return mst
}

func primVisit(g WGraph, v int, marked []bool, pq *pqueue.Heap, edgeTo []*Edge) {
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, wt float64) bool {
		if marked[a] {
			return true
		}
		e := &Edge{
			from:   v,
			to:     a,
			weight: wt,
		}
		if edgeTo[a] == nil {
			pq.Push(util.Float(wt), a)
			edgeTo[a] = e
		} else if e.weight < edgeTo[a].weight {
			pq.Fix(util.Float(e.weight), a)
			edgeTo[a] = e
		}
		return true
	})
}

// Kruskal 该实现仅支持连通图
func (g WGraph) Kruskal() *queue.Queen {
	mst := queue.NewQueen(int(g.NumVertical()))
	uf := unionfind.NewUF(int(g.NumVertical()))
	pq := pqueue.New(g.NumEdge())
	g.IterateWEdge(func(src int, dst int, w float64) bool {
		pq.Push(util.Float(w), &Edge{ // fixme pq Cmp Key
			from:   src,
			to:     dst,
			weight: w,
		})
		return true
	})
	for !mst.IsFull() {
		min := pq.Pop()
		minE := min.(*Edge)
		if uf.IsConnected(minE.from, minE.to) {
			continue
		}
		uf.Union(minE.from, minE.to)
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
	from, to int
	weight   float64
}

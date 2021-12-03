package graphs

import (
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/stack"
	"github.com/howz97/algorithm/util"
)

type BFS struct {
	src    int
	marked []bool
	edgeTo []int
}

func NewBFS(g ITraverse, src int) *BFS {
	if !g.HasVertical(src) {
		return nil
	}
	bfs := &BFS{
		src:    src,
		marked: make([]bool, g.NumVertical()),
		edgeTo: make([]int, g.NumVertical()),
	}
	q := queue.NewIntQ()
	bfs.marked[src] = true
	q.PushBack(src)
	for !q.IsEmpty() {
		edge := q.Front()
		g.RangeAdj(edge, func(adj int) bool {
			if !bfs.marked[adj] {
				bfs.edgeTo[adj] = edge
				bfs.marked[adj] = true
				q.PushBack(adj)
			}
			return true
		})
	}
	return bfs
}

func (bfs *BFS) CanReach(dst int) bool {
	if !bfs.checkVertical(dst) {
		return false
	}
	return bfs.marked[dst]
}

func (bfs *BFS) ShortestPathTo(dst int) []int {
	if !bfs.CanReach(dst) {
		return nil
	}
	if dst == bfs.src {
		return []int{dst}
	}
	path := make([]int, 0, 2)
	path = append(path, dst)
	for bfs.edgeTo[dst] != bfs.src {
		dst = bfs.edgeTo[dst]
		path = append(path, dst)
	}
	path = append(path, bfs.src)
	util.ReverseInts(path)
	return path
}

func (bfs *BFS) checkVertical(v int) bool {
	return v >= 0 && v < len(bfs.marked)
}

func RangeDFS(g ITraverse, src int, fn func(int) bool) {
	RangeUnMarkDFS(g, src, nil, fn)
}

func RangeUnMarkDFS(g ITraverse, src int, marked []bool, fn func(int) bool) {
	if !g.HasVertical(src) {
		return
	}
	if len(marked) == 0 {
		marked = make([]bool, g.NumVertical())
	}
	rangeDFS(g, src, marked, fn)
}

func rangeDFS(g ITraverse, v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	g.RangeAdj(v, func(adj int) bool {
		if !marked[adj] {
			if !rangeDFS(g, adj, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

func RevDFSAll(g ITraverse, fn func(int) bool) {
	marked := make([]bool, g.NumVertical())
	for v := range marked {
		if marked[v] {
			continue
		}
		revDFS(g, v, marked, fn)
	}
}

func RevDFS(g ITraverse, src int, fn func(int) bool) { // fixme
	if !g.HasVertical(src) {
		return
	}
	marked := make([]bool, g.NumVertical())
	revDFS(g, src, marked, fn)
}

func RevDFSOrder(g ITraverse, src int) *stack.IntStack {
	order := stack.NewInt(g.NumVertical())
	RevDFS(g, src, func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func revDFS(g ITraverse, v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	goon := true // continue DFS or abort
	g.RangeAdj(v, func(a int) bool {
		if !marked[a] {
			if !revDFS(g, a, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	if !goon {
		return false
	}
	return fn(v)
}

func ReachableBits(g ITraverse, src int) []bool {
	if !g.HasVertical(src) {
		return nil
	}
	marked := make([]bool, g.NumVertical())
	rangeDFS(g, src, marked, func(_ int) bool { return true })
	return marked
}

func ReachableSlice(g ITraverse, src int) []int {
	if !g.HasVertical(src) {
		return nil
	}
	var arrived []int
	RangeDFS(g, src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

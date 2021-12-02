package graphs

import (
	"github.com/howz97/algorithm/queue"
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

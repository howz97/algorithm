package graph

import (
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/sort"
)

type BFS struct {
	src    int
	marked []bool

	// the adjacent vertical (edgeTo[i]) of the destination vertical (i) on the shortest path
	edgeTo []int
}

func NewBFS(g Graph, src int) (*BFS, error) {
	if !g.HasV(src) {
		return nil, errVerticalNotExist
	}
	bfs := &BFS{
		src:    src,
		marked: make([]bool, g.NumV()),
		edgeTo: make([]int, g.NumV()),
	}
	q := queue.NewIntQ()
	bfs.marked[src] = true
	q.PushBack(src)
	for !q.IsEmpty() {
		edge := q.Front()
		adjs, _ := g.Adjacent(edge)
		for _, adj := range adjs {
			if bfs.marked[adj] {
				continue
			}
			bfs.edgeTo[adj] = edge
			bfs.marked[adj] = true
			q.PushBack(adj)
		}
	}
	return bfs, nil
}

func (bfs *BFS) IsMarked(v int) bool {
	if v < 0 || v >= len(bfs.marked) {
		return false
	}
	return bfs.marked[v]
}

func (bfs *BFS) ShortestPathTo(dst int) []int {
	if dst < 0 || dst >= len(bfs.marked) || !bfs.marked[dst] {
		return nil
	}
	if dst == bfs.src {
		return []int{dst}
	}
	path := append(make([]int, 0, 2), dst)
	for bfs.edgeTo[dst] != bfs.src {
		dst = bfs.edgeTo[dst]
		path = append(path, dst)
	}
	path = append(path, bfs.src)
	sort.Reverse(path)
	return path
}

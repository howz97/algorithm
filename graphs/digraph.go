package graphs

import (
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/stack"
	"github.com/howz97/algorithm/util"
	"gopkg.in/yaml.v2"
	"strconv"
)

type Digraph struct {
	Edges []*hash_map.Chaining
	*Symbol
}

func NewDigraph(size uint) *Digraph {
	edges := make([]*hash_map.Chaining, size)
	for i := range edges {
		edges[i] = hash_map.New()
	}
	return &Digraph{Edges: edges}
}

func (dg *Digraph) NumVertical() uint {
	return uint(len(dg.Edges))
}

func (dg *Digraph) HasVertical(v int) bool {
	return v >= 0 && v < len(dg.Edges)
}

func (dg *Digraph) NumEdge() uint {
	n := uint(0)
	for i := range dg.Edges {
		n += dg.Edges[i].Size()
	}
	return n
}

func (dg *Digraph) AddEdge(src, dst int) error {
	return dg.addWeightedEdge(src, dst, 1)
}

func (dg *Digraph) addWeightedEdge(src, dst int, w float64) error {
	if !dg.HasVertical(src) || !dg.HasVertical(dst) {
		return ErrVerticalNotExist
	}
	if src == dst {
		return ErrSelfLoop
	}
	dg.Edges[src].Put(util.Int(dst), w)
	return nil
}

func (dg *Digraph) DelEdge(src, dst int) {
	if !dg.HasVertical(src) {
		return
	}
	dg.Edges[src].Del(util.Int(dst))
}

func (dg *Digraph) HasEdge(v1, v2 int) bool {
	if !dg.HasVertical(v1) || !dg.HasVertical(v2) {
		return false
	}
	return dg.Edges[v1].Get(util.Int(v2)) != nil
}

func (dg *Digraph) IterateAdj(v int, fn func(int) bool) {
	dg.IterateWAdj(v, func(a int, _ float64) bool {
		return fn(a)
	})
}

func (dg *Digraph) IterateWAdj(v int, fn func(int, float64) bool) {
	if !dg.HasVertical(v) {
		return
	}
	dg.Edges[v].Range(func(key hash_map.Key, val util.T) bool {
		return fn(int(key.(util.Int)), val.(float64))
	})
}

func (dg *Digraph) Adjacent(v int) (adj []int) {
	dg.IterateAdj(v, func(a int) bool {
		adj = append(adj, a)
		return true
	})
	return adj
}

func (dg *Digraph) getWeightMust(from, to int) float64 {
	return dg.Edges[from].Get(util.Int(to)).(float64)
}

func (dg *Digraph) GetWeight(from, to int) (float64, bool) {
	if !dg.HasVertical(from) {
		return 0, false
	}
	w := dg.Edges[from].Get(util.Int(to))
	if w == nil {
		return 0, false
	}
	return w.(float64), true
}

func (dg *Digraph) String() string {
	var i2a func(int) string
	if dg.Symbol != nil {
		i2a = dg.SymbolOf
	} else {
		i2a = strconv.Itoa
	}
	out := ""
	for i := range dg.Edges {
		out += i2a(i) + " :"
		dg.IterateAdj(i, func(j int) bool {
			out += " " + i2a(j)
			return true
		})
		out += "\n"
	}
	out += "\n"
	return out
}

func (dg *Digraph) Reverse() *Digraph {
	rg := NewDigraph(dg.NumVertical())
	for v := 0; v < int(dg.NumVertical()); v++ {
		dg.IterateAdj(v, func(w int) bool {
			rg.AddEdge(w, v)
			return true
		})
	}
	return rg
}

func (dg *Digraph) FindNegativeEdge() (src, dst int) {
	src, dst = -1, -1
	dg.IterateWEdge(func(v int, v2 int, w float64) bool {
		if w < 0 {
			src = v
			dst = v2
			return false
		}
		return true
	})
	return
}

func (dg *Digraph) iterateAdj(v int, fn func(int, int, float64) bool) {
	dg.Edges[v].Range(func(key hash_map.Key, val util.T) bool {
		return fn(v, int(key.(util.Int)), val.(float64))
	})
}

func (dg *Digraph) HasCycle() bool {
	return dg.FindCycle() != nil
}

func (dg *Digraph) FindCycle() *stack.IntStack {
	marks := make([]bool, dg.NumVertical())
	path := stack.NewInt(4)
	for v, m := range marks {
		if !m {
			if dg.DetectCycleDFS(v, marks, path) {
				return path
			}
		}
	}
	return nil
}

func (dg *Digraph) FindCycleFrom(v int) *stack.IntStack {
	marks := make([]bool, dg.NumVertical())
	path := stack.NewInt(4)
	if dg.DetectCycleDFS(v, marks, path) {
		return path
	}
	return nil
}

func ParseCycleInStack(stk *stack.IntStack) []int {
	path := make([]int, 0, stk.Size())
	w := stk.Pop()
	path = append(path, w)
	for {
		v := stk.Pop()
		path = append(path, v)
		if v == w {
			break
		}
	}
	util.ReverseInts(path)
	return path
}

func (dg *Digraph) DetectCycleDFS(v int, marked []bool, path *stack.IntStack) bool {
	path.Push(v)
	found := false
	dg.IterateAdj(v, func(w int) bool {
		if marked[w] {
			return true
		}
		if path.Contains(w) {
			path.Push(w)
			found = true
			return false
		}
		found = dg.DetectCycleDFS(w, marked, path)
		return !found
	})
	marked[v] = true
	if !found {
		path.Pop()
	}
	return found
}

// Topological return a stack that pop vertices in topological order
func (dg *Digraph) Topological() *stack.IntStack {
	if dg.HasCycle() {
		return nil
	}
	order := stack.NewInt(int(dg.NumVertical()))
	dg.IterateVetRDFS(func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func (dg *Digraph) IterateWEdge(fn func(int, int, float64) bool) {
	for src, hm := range dg.Edges {
		goon := true
		hm.Range(func(dst hash_map.Key, v util.T) bool {
			goon = fn(src, int(dst.(util.Int)), v.(float64))
			return goon
		})
		if !goon {
			break
		}
	}
}

func (dg *Digraph) IterateEdge(fn func(int, int) bool) {
	dg.IterateWEdge(func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

func (dg *Digraph) IterateWEdgeFrom(v int, fn func(int, int, float64) bool) {
	dg.IterateVetDFS(v, func(v int) bool {
		goon := true
		dg.IterateWAdj(v, func(a int, w float64) bool {
			goon = fn(v, a, w)
			return goon
		})
		return goon
	})
}

func (dg *Digraph) IterateEdgeFrom(v int, fn func(int, int) bool) {
	dg.IterateWEdgeFrom(v, func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

func (dg *Digraph) IterateVetDFS(src int, fn func(int) bool) {
	dg.iterateUnMarkVetFrom(src, nil, fn)
}

func (dg *Digraph) iterateUnMarkVetFrom(src int, marked []bool, fn func(int) bool) {
	if !dg.HasVertical(src) {
		return
	}
	if len(marked) == 0 {
		marked = make([]bool, dg.NumVertical())
	}
	dg.iterateUnMarkVetDFS(src, marked, fn)
}

func (dg *Digraph) iterateUnMarkVetDFS(v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	if !fn(v) {
		return false
	}
	goon := true // continue DFS or abort
	dg.IterateAdj(v, func(adj int) bool {
		if !marked[adj] {
			if !dg.iterateUnMarkVetDFS(adj, marked, fn) {
				goon = false
			}
		}
		return goon
	})
	return goon
}

func (dg *Digraph) ReachableSlice(src int) []int {
	if !dg.HasVertical(src) {
		return nil
	}
	var arrived []int
	dg.IterateVetDFS(src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

func (dg *Digraph) ReachableBits(src int) []bool {
	if !dg.HasVertical(src) {
		return nil
	}
	marked := make([]bool, dg.NumVertical())
	dg.iterateUnMarkVetDFS(src, marked, func(_ int) bool { return true })
	return marked
}

func (dg *Digraph) IterateVetRDFS(fn func(int) bool) {
	marked := make([]bool, dg.NumVertical())
	for v := range marked {
		if marked[v] {
			continue
		}
		dg.rDFS(v, marked, fn)
	}
}

func (dg *Digraph) IterateRDFSFrom(src int, fn func(int) bool) {
	if !dg.HasVertical(src) {
		return
	}
	marked := make([]bool, dg.NumVertical())
	dg.rDFS(src, marked, fn)
}

func (dg *Digraph) RDFSOrderVertical(src int) *stack.IntStack {
	order := stack.NewInt(int(dg.NumVertical()))
	dg.IterateRDFSFrom(src, func(v int) bool {
		order.Push(v)
		return true
	})
	return order
}

func (dg *Digraph) rDFS(v int, marked []bool, fn func(int) bool) bool {
	marked[v] = true
	goon := true // continue DFS or abort
	dg.IterateAdj(v, func(a int) bool {
		if !marked[a] {
			if !dg.rDFS(a, marked, fn) {
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

func (dg *Digraph) IsBipartiteGraph() bool {
	marks := make([]bool, dg.NumVertical())
	colors := make([]bool, dg.NumVertical())
	for i, m := range marks {
		if m {
			continue
		}
		if !dg.isBipartiteDFS(i, true, colors, marks) {
			return false
		}
	}
	return true
}

func (dg *Digraph) isBipartiteDFS(cur int, color bool, colors []bool, marked []bool) bool {
	isBip := true
	if !marked[cur] {
		marked[cur] = true
		colors[cur] = color
		dg.IterateAdj(cur, func(adj int) bool {
			if !dg.isBipartiteDFS(adj, !color, colors, marked) {
				isBip = false
				return false
			}
			return true
		})
	} else {
		isBip = colors[cur] == color
	}
	return isBip
}

func (dg *Digraph) SameWith(other Digraph) bool {
	if dg.NumVertical() != other.NumVertical() {
		return false
	}
	if dg.NumEdge() != other.NumEdge() {
		return false
	}
	isSame := true
	dg.IterateWEdge(func(from int, to int, w float64) bool {
		w2, ok := other.GetWeight(from, to)
		if !ok || w != w2 {
			isSame = false
			return false
		}
		return true
	})
	return isSame
}

func (dg *Digraph) TotalWeight() float64 {
	var total float64
	dg.IterateWEdge(func(_ int, _ int, w float64) bool {
		total += w
		return true
	})
	return total
}

func (dg *Digraph) Marshal() ([]byte, error) {
	m := make(map[string]map[string]float64)
	for v := 0; v < int(dg.NumVertical()); v++ {
		edges := make(map[string]float64)
		dg.IterateWAdj(v, func(a int, w float64) bool {
			if dg.Symbol == nil {
				edges[strconv.Itoa(a)] = w
			} else {
				edges[dg.vet2syb[a]] = w
			}
			return true
		})
		if dg.Symbol == nil {
			m[strconv.Itoa(v)] = edges
		} else {
			m[dg.vet2syb[v]] = edges
		}
	}
	return yaml.Marshal(m)
}

// SCC is strong connected components of digraph
// vertices in the same component can access each other
type SCC struct {
	locate     []int   // vertical -> componentID
	components [][]int // componentID -> vertices
}

// SCC calculate strong connected components of digraph with kosaraju algorithm
func (dg Digraph) SCC() *SCC {
	scc := &SCC{
		locate: make([]int, dg.NumVertical()),
	}
	marked := make([]bool, dg.NumVertical())
	dg.IterateVetRDFS(func(v int) bool {
		if !marked[v] {
			c := make([]int, 0, 8)
			dg.iterateUnMarkVetFrom(v, marked, func(w int) bool {
				scc.locate[w] = len(scc.components)
				c = append(c, w)
				return true
			})
			scc.components = append(scc.components, c)
		}
		return true
	})
	return scc
}

func (scc *SCC) IsStronglyConn(src, dst int) bool {
	if !scc.hasVertical(src) || !scc.hasVertical(dst) {
		return false
	}
	return scc.locate[src] == scc.locate[dst]
}

func (scc *SCC) GetCompID(v int) int {
	if !scc.hasVertical(v) {
		return -1
	}
	return scc.locate[v]
}

func (scc *SCC) RangeComponent(v int, fn func(int) bool) {
	if !scc.hasVertical(v) {
		return
	}
	for _, w := range scc.components[scc.locate[v]] {
		if !fn(w) {
			break
		}
	}
}

func (scc *SCC) NumComponents() int {
	return len(scc.components)
}

func (scc *SCC) hasVertical(v int) bool {
	return v >= 0 && v < len(scc.locate)
}

type Reachable [][]bool

func (dg Digraph) Reachable() Reachable {
	tc := make(Reachable, dg.NumVertical())
	for v := range tc {
		tc[v] = dg.ReachableBits(v)
	}
	return tc
}

func (tc Reachable) CanReach(src, dst int) bool {
	if !tc.hasVertical(src) || !tc.hasVertical(dst) {
		return false
	}
	return tc[src][dst]
}

func (tc Reachable) Range(v int, fn func(v int) bool) {
	if !tc.hasVertical(v) {
		return
	}
	for w, marked := range tc[v] {
		if marked {
			if !fn(w) {
				break
			}
		}
	}
}

func (tc Reachable) hasVertical(v int) bool {
	return v >= 0 || v < len(tc)
}

type BFS struct {
	src    int
	marked []bool
	edgeTo []int
}

func (dg Digraph) BFS(src int) *BFS {
	if !dg.HasVertical(src) {
		return nil
	}
	bfs := &BFS{
		src:    src,
		marked: make([]bool, dg.NumVertical()),
		edgeTo: make([]int, dg.NumVertical()),
	}
	q := queue.NewIntQ()
	bfs.marked[src] = true
	q.PushBack(src)
	for !q.IsEmpty() {
		edge := q.Front()
		dg.IterateAdj(edge, func(adj int) bool {
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

func (bfs *BFS) ShortestPathTo(dst int) *Path {
	if !bfs.CanReach(dst) {
		return nil
	}
	if dst == bfs.src {
		return nil
	}
	path := &Path{
		stk:      stack.NewInt(2),
		distance: 0,
	}
	for dst != bfs.src {
		path.stk.Push(dst)
		path.distance++
		dst = bfs.edgeTo[dst]
	}
	path.stk.Push(bfs.src)
	return path
}

func (bfs *BFS) checkVertical(v int) bool {
	return v >= 0 && v < len(bfs.marked)
}

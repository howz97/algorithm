package graphs

import (
	"strconv"

	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/basic/stack"
	"github.com/howz97/algorithm/search/hashmap"
	. "github.com/howz97/algorithm/util"
	"gopkg.in/yaml.v2"
)

func NewDigraph(size uint) *Digraph {
	edges := make([]*hashmap.Chaining[Int, float64], size)
	for i := range edges {
		edges[i] = hashmap.New[Int, float64]()
	}
	return &Digraph{edges: edges}
}

type Digraph struct {
	edges []*hashmap.Chaining[Int, float64]
	*Symbol
}

// NumVert get the number of vertices
func (dg *Digraph) NumVert() uint {
	return uint(len(dg.edges))
}

// HasVert indicate whether dg contains vertical v
func (dg *Digraph) HasVert(v int) bool {
	return v >= 0 && v < len(dg.edges)
}

// AddEdge add a new edge
func (dg *Digraph) AddEdge(from, to int) {
	dg.addWeightedEdge(from, to, 1)
}

// NumEdge get the number of edges
func (dg *Digraph) NumEdge() uint {
	n := uint(0)
	for i := range dg.edges {
		n += dg.edges[i].Size()
	}
	return n
}

// HasEdge indicate whether dg contains the edge specified by params
func (dg *Digraph) HasEdge(from, to int) bool {
	if !dg.HasVert(from) || !dg.HasVert(to) {
		return false
	}
	_, ok := dg.edges[from].Get(Int(to))
	return ok
}

func (dg *Digraph) addWeightedEdge(from, to int, w float64) {
	if !dg.HasVert(from) || !dg.HasVert(to) {
		panic(ErrVerticalNotExist)
	}
	if from == to {
		panic(ErrSelfLoop)
	}
	dg.edges[from].Put(Int(to), w)
}

// DelEdge delete an edge
func (dg *Digraph) DelEdge(src, dst int) {
	dg.edges[src].Del(Int(dst))
}

// GetWeight get the weight of edge
// Zero will be returned if edge not exist
func (dg *Digraph) GetWeight(from, to int) float64 {
	w, _ := dg.edges[from].Get(Int(to))
	return w
}

// TotalWeight sum the weight of all edges
func (dg *Digraph) TotalWeight() float64 {
	var total float64
	dg.IterateWEdge(func(_ int, _ int, w float64) bool {
		total += w
		return true
	})
	return total
}

// IterateAdj iterate all adjacent vertices of v
func (dg *Digraph) IterateAdj(v int, fn func(int) bool) {
	dg.IterateWAdj(v, func(a int, _ float64) bool {
		return fn(a)
	})
}

// IterateWAdj iterate all adjacent vertices and weight of v
func (dg *Digraph) IterateWAdj(v int, fn func(int, float64) bool) { // todo: rename
	dg.edges[v].Range(func(key Int, val float64) bool {
		return fn(int(key), val)
	})
}

// Adjacent return a slice contains all adjacent vertices of v
func (dg *Digraph) Adjacent(v int) (adj []int) {
	dg.IterateAdj(v, func(a int) bool {
		adj = append(adj, a)
		return true
	})
	return adj
}

func (dg *Digraph) String() string {
	var i2a func(int) string
	if dg.Symbol != nil {
		i2a = dg.SymbolOf
	} else {
		i2a = strconv.Itoa
	}
	out := ""
	for i := range dg.edges {
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

// Reverse all edges of dg
func (dg *Digraph) Reverse() *Digraph {
	rg := NewDigraph(dg.NumVert())
	dg.IterateWEdge(func(from int, to int, w float64) bool {
		rg.addWeightedEdge(to, from, w)
		return true
	})
	return rg
}

// FindNegativeEdge find a negative edge.
// If here is no negative edge, (-1, -1) will be returned
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

// FindNegativeEdgeFrom find a reachable negative edge from the specified start vertical
// If here is no negative edge, (-1, -1) will be returned
func (dg *Digraph) FindNegativeEdgeFrom(start int) (src int, dst int) {
	src, dst = -1, -1
	dg.IterateWEdgeFrom(start, func(v0 int, v1 int, w float64) bool {
		if w < 0 {
			src = v0
			dst = v1
			return false
		}
		return true
	})
	return
}

// FindCycle find any directed cycle in dg
func (dg *Digraph) FindCycle() *Cycle {
	marks := make([]bool, dg.NumVert())
	path := NewPath()
	for v, m := range marks {
		if !m {
			if dg.detectCycleDFS(v, marks, path) {
				return path.Cycle()
			}
		}
	}
	return nil
}

// FindCycleFrom find any directed cycle from vertical v in dg
// But not include cycle that can not be accessed from v
func (dg *Digraph) FindCycleFrom(v int) *Path {
	marks := make([]bool, dg.NumVert())
	path := NewPath()
	if dg.detectCycleDFS(v, marks, path) {
		return path
	}
	return nil
}

func (dg *Digraph) detectCycleDFS(v int, marked []bool, path *Path) bool {
	found := false
	dg.IterateWAdj(v, func(a int, w float64) bool {
		if marked[a] {
			return true
		}
		if path.ContainsVert(a) {
			path.Push(v, a, w)
			found = true
			return false
		}
		path.Push(v, a, w)
		found = dg.detectCycleDFS(a, marked, path)
		if !found {
			path.Pop()
		}
		return !found
	})
	marked[v] = true
	return found
}

// Topological return a stack that will pop vertices in topological order
func (dg *Digraph) Topological() (order *stack.Stack[int]) {
	if dg.FindCycle() != nil {
		return
	}
	order = stack.New[int](int(dg.NumVert()))
	dg.IterateVetRDFS(func(v int) bool {
		order.Push(v)
		return true
	})
	return
}

// IterateWEdge iterate all edges and their weight in dg
func (dg *Digraph) IterateWEdge(fn func(int, int, float64) bool) {
	for src, hm := range dg.edges {
		goon := true
		hm.Range(func(dst Int, v float64) bool {
			goon = fn(src, int(dst), v)
			return goon
		})
		if !goon {
			break
		}
	}
}

// IterateEdge iterate all edges in dg
func (dg *Digraph) IterateEdge(fn func(int, int) bool) {
	dg.IterateWEdge(func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

// IterateWEdgeFrom iterate all reachable edges and their weight from vertical src
func (dg *Digraph) IterateWEdgeFrom(src int, fn func(int, int, float64) bool) {
	dg.IterateVertDFS(src, func(v int) bool {
		goon := true
		dg.IterateWAdj(v, func(a int, w float64) bool {
			goon = fn(v, a, w)
			return goon
		})
		return goon
	})
}

// IterateEdgeFrom iterate all reachable edges from vertical src
func (dg *Digraph) IterateEdgeFrom(src int, fn func(int, int) bool) {
	dg.IterateWEdgeFrom(src, func(src int, dst int, _ float64) bool {
		return fn(src, dst)
	})
}

// IterateVertDFS iterate all reachable vertices from vertical src in DFS order
func (dg *Digraph) IterateVertDFS(src int, fn func(int) bool) {
	dg.iterateUnMarkVetFrom(src, nil, fn)
}

func (dg *Digraph) iterateUnMarkVetFrom(src int, marked []bool, fn func(int) bool) {
	if !dg.HasVert(src) {
		return
	}
	if len(marked) == 0 {
		marked = make([]bool, dg.NumVert())
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

// ReachableSlice get a slice contains all reachable vertices from src
func (dg *Digraph) ReachableSlice(src int) []int {
	if !dg.HasVert(src) {
		return nil
	}
	var arrived []int
	dg.IterateVertDFS(src, func(v int) bool {
		arrived = append(arrived, v)
		return true
	})
	return arrived
}

// ReachableBits get a bit-map contains all reachable vertices from src
func (dg *Digraph) ReachableBits(src int) []bool {
	marked := make([]bool, dg.NumVert())
	dg.iterateUnMarkVetDFS(src, marked, func(_ int) bool { return true })
	return marked
}

// IterateVetRDFS iterate all vertices in RDFS order
func (dg *Digraph) IterateVetRDFS(fn func(int) bool) {
	marked := make([]bool, dg.NumVert())
	for v := range marked {
		if marked[v] {
			continue
		}
		dg.rDFS(v, marked, fn)
	}
}

// IterateRDFSFrom iterate all reachable vertices from vertical src in RDFS order
func (dg *Digraph) IterateRDFSFrom(src int, fn func(int) bool) {
	if !dg.HasVert(src) {
		return
	}
	marked := make([]bool, dg.NumVert())
	dg.rDFS(src, marked, fn)
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

// IsBipartite check whether dg is a bipartite graph
func (dg *Digraph) IsBipartite() bool {
	marks := make([]bool, dg.NumVert())
	colors := make([]bool, dg.NumVert())
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

// IsSameWith check whether dg is the same with other
func (dg *Digraph) IsSameWith(other Digraph) bool {
	if dg.NumVert() != other.NumVert() {
		return false
	}
	if dg.NumEdge() != other.NumEdge() {
		return false
	}
	isSame := true
	dg.IterateWEdge(func(from int, to int, w float64) bool {
		if w != other.GetWeight(from, to) {
			isSame = false
			return false
		}
		return true
	})
	return isSame
}

// Marshal dg into yaml format
func (dg *Digraph) Marshal() ([]byte, error) {
	m := make(map[string]map[string]float64)
	for v := 0; v < int(dg.NumVert()); v++ {
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
		locate: make([]int, dg.NumVert()),
	}
	marked := make([]bool, dg.NumVert())
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

// IsStronglyConn check whether src is strongly connected to dst
func (scc *SCC) IsStronglyConn(src, dst int) bool {
	return scc.locate[src] == scc.locate[dst]
}

// GetCompID get the strongly connected component ID of vertical v
func (scc *SCC) GetCompID(v int) int {
	return scc.locate[v]
}

// IterateComponent iterate all vertices strongly connected to vertical v (include v)
func (scc *SCC) IterateComponent(v int, fn func(int) bool) {
	for _, w := range scc.components[scc.locate[v]] {
		if !fn(w) {
			break
		}
	}
}

// NumComponents get the number of components
func (scc *SCC) NumComponents() int {
	return len(scc.components)
}

type Reachable [][]bool

// Reachable save all reachable information of dg
func (dg Digraph) Reachable() Reachable {
	tc := make(Reachable, dg.NumVert())
	for v := range tc {
		tc[v] = dg.ReachableBits(v)
	}
	return tc
}

// CanReach check whether src can reach dst
func (tc Reachable) CanReach(src, dst int) bool {
	return tc[src][dst]
}

// Iterate all reachable vertices from src
func (tc Reachable) Iterate(src int, fn func(v int) bool) {
	for w, marked := range tc[src] {
		if marked {
			if !fn(w) {
				break
			}
		}
	}
}

type BFS struct {
	src    int
	marked []bool
	edgeTo []int
}

// BFS save all BFS information from src
func (dg Digraph) BFS(src int) *BFS {
	if !dg.HasVert(src) {
		return nil
	}
	bfs := &BFS{
		src:    src,
		marked: make([]bool, dg.NumVert()),
		edgeTo: make([]int, dg.NumVert()),
	}
	q := queue.NewLinkQ[int]()
	bfs.marked[src] = true
	q.PushBack(src)
	for q.Size() > 0 {
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

// CanReach check whether src can reach dst
func (bfs *BFS) CanReach(dst int) bool {
	return bfs.marked[dst]
}

// ShortestPathTo get the shortest path to dst (ignore weight)
func (bfs *BFS) ShortestPathTo(dst int) *Path {
	if !bfs.CanReach(dst) {
		return nil
	}
	if dst == bfs.src {
		return nil
	}
	path := NewPath()
	for dst != bfs.src {
		path.Push(bfs.edgeTo[dst], dst, 1)
		dst = bfs.edgeTo[dst]
	}
	path.Reverse()
	return path
}

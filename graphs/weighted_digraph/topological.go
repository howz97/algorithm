package weighted_digraph

import "github.com/howz97/algorithm/stack"

func (g EdgeWeightedDigraph) ReversePostOrder() *stack.IntStack {
	marked := make([]bool, g.NumV())
	result := stack.NewInt(g.NumV())
	for i, b := range marked {
		if !b {
			g.reversePostDFS(i, marked, result)
		}
	}
	return result
}

func (g EdgeWeightedDigraph) reversePostDFS(v int, marked []bool, result *stack.IntStack) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		if !marked[e.to] {
			g.reversePostDFS(e.to, marked, result)
		}
	}
	result.Push(v)
}

func (g EdgeWeightedDigraph) TopologicalSort() *stack.IntStack {
	if !g.DetectDirCycle() {
		return g.ReversePostOrder()
	}
	return nil
}

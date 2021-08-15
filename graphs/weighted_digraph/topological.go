package weighted_digraph

import "howz97/algorithm/stack"

func (g EdgeWeightedDigraph) ReversePostOrder() *stack.StackInt {
	marked := make([]bool, g.NumV())
	result := stack.NewStackInt(g.NumV())
	for i, b := range marked {
		if !b {
			g.reversePostDFS(i, marked, result)
		}
	}
	return result
}

func (g EdgeWeightedDigraph) reversePostDFS(v int, marked []bool, result *stack.StackInt) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		if !marked[e.to] {
			g.reversePostDFS(e.to, marked, result)
		}
	}
	result.Push(v)
}

func (g EdgeWeightedDigraph) TopologicalSort() *stack.StackInt {
	if !g.DetectDirCycle() {
		return g.ReversePostOrder()
	}
	return nil
}

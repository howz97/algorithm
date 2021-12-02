package wdigraph

import "github.com/howz97/algorithm/stack"

func (g WDigraph) ReversePostOrder() *stack.IntStack {
	marked := make([]bool, g.NumVertical())
	result := stack.NewInt(g.NumVertical())
	for i, b := range marked {
		if !b {
			g.reversePostDFS(i, marked, result)
		}
	}
	return result
}

func (g WDigraph) reversePostDFS(v int, marked []bool, result *stack.IntStack) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		if !marked[e.to] {
			g.reversePostDFS(e.to, marked, result)
		}
	}
	result.Push(v)
}

func (g WDigraph) TopologicalSort() *stack.IntStack {
	if !g.DetectDirCycle() {
		return g.ReversePostOrder()
	}
	return nil
}

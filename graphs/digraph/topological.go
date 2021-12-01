package digraph

import "github.com/howz97/algorithm/stack"

func ReversePostOrder(g Digraph) *stack.IntStack {
	marked := make([]bool, g.NumVertical())
	result := stack.NewInt(g.NumVertical())
	for i, b := range marked {
		if !b {
			reversePostDFS(g, i, marked, result)
		}
	}
	return result
}

func reversePostDFS(g Digraph, v int, marked []bool, result *stack.IntStack) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, w := range adj {
		if !marked[w] {
			reversePostDFS(g, w, marked, result)
		}
	}
	result.Push(v)
}

func TopologicalSort(g Digraph) *stack.IntStack {
	if !g.HasCycle() {
		return ReversePostOrder(g)
	}
	return nil
}

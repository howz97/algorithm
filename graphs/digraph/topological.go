package digraph

import "github.com/zh1014/algorithm/stack"

func topologicalDFS(g Digraph, v int, marked []bool, result *stack.Stack){
	adj := g.Adjacent(v)
	for _, w := range adj {
		if !marked[w] {
			topologicalDFS(g, w, marked, result)
		}
	}
	result.Push(v)
}

func TopologicalOrder(g Digraph) *stack.Stack {
	marked := make([]bool, g.NumV())
	result := stack.New(g.NumV())
	for i, b := range marked {
		if !b {
			topologicalDFS(g, i, marked, result)
		}
	}
	return result
}

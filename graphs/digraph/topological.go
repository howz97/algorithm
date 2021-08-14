package digraph

import "howz97/algorithm/stack"

func ReversePostOrder(g Digraph) *stack.StackInt {
	marked := make([]bool, g.NumV())
	result := stack.NewStackInt(g.NumV())
	for i, b := range marked {
		if !b {
			reversePostDFS(g, i, marked, result)
		}
	}
	return result
}

func reversePostDFS(g Digraph, v int, marked []bool, result *stack.StackInt) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, w := range adj {
		if !marked[w] {
			reversePostDFS(g, w, marked, result)
		}
	}
	result.Push(v)
}

func TopologicalSort(g Digraph) *stack.StackInt {
	if !DetectDirCycle(g) {
		return ReversePostOrder(g)
	}
	return nil
}

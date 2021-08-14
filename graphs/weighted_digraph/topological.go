package weighted_digraph

import "howz97/algorithm/stack"

func ReversePostOrder(g EdgeWeightedDigraph) *stack.StackInt {
	marked := make([]bool, g.NumV())
	result := stack.NewStackInt(g.NumV())
	for i, b := range marked {
		if !b {
			reversePostDFS(g, i, marked, result)
		}
	}
	return result
}

func reversePostDFS(g EdgeWeightedDigraph, v int, marked []bool, result *stack.StackInt) {
	marked[v] = true
	adj := g.Adjacent(v)
	for _, e := range adj {
		if !marked[e.to] {
			reversePostDFS(g, e.to, marked, result)
		}
	}
	result.Push(v)
}

func TopologicalSort(g EdgeWeightedDigraph) *stack.StackInt {
	if !DetectDirCycle(g) {
		return ReversePostOrder(g)
	}
	return nil
}

package regexp

import (
	"fmt"
	"github.com/zh1014/algorithm/graphs/digraph"
	"github.com/zh1014/algorithm/queue"
	"github.com/zh1014/algorithm/set"
	"github.com/zh1014/algorithm/stack"
)

func IsMatch(pattern, txt string) bool {
	symbolTbl := parsePattern(pattern)
	g := createNFA(symbolTbl)
	reachableStatus := getStartStatus(g)
	for _, r := range txt {
		statusAfterMatch := set.New()
		for !reachableStatus.IsEmpty() {
			status := reachableStatus.RemoveOne()
			if status < len(symbolTbl) && match(symbolTbl[status], r) {
				statusAfterMatch.Add(status + 1)
			}
		}
		if statusAfterMatch.IsEmpty() {
			return false
		}
		getReachableStatus(g, statusAfterMatch, reachableStatus)
	}
	return reachableStatus.Contains(len(symbolTbl))
}

func match(s symbol, r rune) bool {
	return isWildCard(s) || (!s.isPrime && s.r == r)
}

type symbol struct {
	isPrime bool
	r       rune
}

func parsePattern(pattern string) []symbol {
	// this can only transfer \( \) \| \* \. \\
	pttrnRunes := []rune(pattern)
	numRunes := len(pttrnRunes)
	symbolTable := make([]symbol, 0, numRunes)
	for i := 0; i < numRunes; i++ {
		if pttrnRunes[i] == '\\' {
			i++
			if !canBeTransferred(pttrnRunes[i]) {
				panic(fmt.Sprintf("invalid transfer: \\%v", string(pttrnRunes[i])))
			}
			symbolTable = append(symbolTable, symbol{
				isPrime: false,
				r:       pttrnRunes[i],
			})
			continue
		}
		symbolTable = append(symbolTable, symbol{
			isPrime: isPrimeRune(pttrnRunes[i]),
			r:       pttrnRunes[i],
		})
	}
	return symbolTable
}

func isPrimeRune(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '*' || r == '.'
}

func canBeTransferred(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '*' || r == '.' || r == '\\'
}

func getReachableStatus(g digraph.Digraph, src, reachable set.Set) {
	tc := digraph.NewTransitiveClosure(g)
	srcQ := queue.NewIntQ()
	for !src.IsEmpty() {
		srcQ.PushBack(src.RemoveOne())
	}
	reachableQ := tc.ReachableVertices(srcQ)
	for !reachableQ.IsEmpty() {
		reachable.Add(reachableQ.Front())
	}
}

func getStartStatus(g digraph.Digraph) set.Set {
	startStatus := set.New()
	dfs := digraph.NewDFS(g, 0)
	rvQ := dfs.ReachableVertices()
	for !rvQ.IsEmpty() {
		startStatus.Add(rvQ.Front())
	}
	return startStatus
}

func createNFA(symbolTable []symbol) digraph.Digraph {
	tblSize := len(symbolTable)
	g := digraph.NewDigraph(tblSize + 1)
	stck := stack.NewStackInt(tblSize)
	for i, symbl := range symbolTable {
		leftBracket := i
		if symbl.isPrime && (symbl.r == '(' || symbl.r == ')' || symbl.r == '*') {
			g.AddEdge(i, i+1)
		}
		if symbl.isPrime && (symbl.r == '(' || symbl.r == '|') {
			stck.Push(i)
		}
		if symbl.isPrime && symbl.r == ')' {
			allOr := queue.NewIntQ()
			for !stck.IsEmpty() {
				out := stck.Pop()
				if symbolTable[out].r == '|' {
					allOr.PushBack(out)
				}else { // got '('
					leftBracket = out
					break
				}
			}
			for !allOr.IsEmpty() {
				or := allOr.Front()
				g.AddEdge(leftBracket, or+1)
				g.AddEdge(or, i)
			}
		}
		if i+1 < tblSize && isClosure(symbolTable[i+1]) {
			g.AddEdge(leftBracket, i+1)
			g.AddEdge(i+1, leftBracket)
		}
	}
	return g
}

func isWildCard(s symbol) bool {
	return s.isPrime && s.r == '.'
}

func isClosure(s symbol) bool {
	return s.isPrime && s.r == '*'
}

package regexp

import (
	"fmt"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/set"
	"github.com/howz97/algorithm/stack"
	"strconv"
	"strings"
	"unicode/utf8"
)

func IsMatch(pattern, txt string) bool {
	table := makeSymbolTable(compile([]rune(pattern)))
	nfa := makeNFA(table)
	reachableStatus := getStartStatus(nfa)
	for _, r := range txt {
		statusAfterMatch := set.NewIntSet()
		for !reachableStatus.IsEmpty() {
			status := reachableStatus.RemoveOne()
			if status < len(table) && match(table[status], r) {
				statusAfterMatch.Add(status + 1)
			}
		}
		if statusAfterMatch.IsEmpty() {
			return false
		}
		getReachableStatus(nfa, statusAfterMatch, reachableStatus)
	}
	return reachableStatus.Contains(len(table))
}

func match(s symbol, r rune) bool {
	return isWildCard(s) || (!s.isPrime && s.r == r)
}

type symbol struct {
	isPrime bool
	r       rune
}

func makeSymbolTable(compiled []rune) []symbol {
	symbols := make([]symbol, 0, len(compiled))
	for i := 0; i < len(compiled); i++ {
		if compiled[i] == '\\' {
			i++
			if !canBeTransferred(compiled[i]) {
				panic(fmt.Sprintf("invalid transfer: \\%v", string(compiled[i])))
			}
			symbols = append(symbols, symbol{
				isPrime: false,
				r:       compiled[i],
			})
			continue
		}
		symbols = append(symbols, symbol{
			isPrime: isPrimeRune(compiled[i]),
			r:       compiled[i],
		})
	}
	return symbols
}

// compile high-level grammar into low-level grammar
func compile(pattern []rune) []rune {
	compiled := make([]rune, 0, len(pattern)*2)
	lp := 0
	lpStack := stack.NewStackInt(10) // 随意设定
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '\\': // must put \ on top case
			lp = i
			i++
			compiled = append(compiled, '\\', pattern[i])
		case '(':
			lpStack.Push(len(compiled))
			compiled = append(compiled, '(')
		case ')':
			lp = lpStack.Pop()
			compiled = append(compiled, ')')
		case '+':
			// "(regexp)+" -> "(regexp)(regexp)*"
			compiled = append(compiled, compiled[lp:]...)
			compiled = append(compiled, '*')
		case '?':
			// "(regexp)?" -> "(regexp|)"
			lastRegExp := make([]rune, len(compiled[lp:]))
			copy(lastRegExp, compiled[lp:])
			compiled = append(compiled[:lp], '(')
			compiled = append(compiled, lastRegExp...)
			compiled = append(compiled, '|', ')')
		case '{':
			rb := indexRune(pattern[i:], '}')
			if rb < 0 {
				panic(fmt.Sprintf("[surround %v] no corresponding right bracket", i))
			}
			inBrackets := pattern[i+1 : rb+i]
			i += rb
			if len(inBrackets) == 0 {
				panic(fmt.Sprintf("[surround %v] nothing in bracket", i))
			}
			hyphen := indexRune(inBrackets, '-')
			lastRegExp := make([]rune, len(compiled[lp:]))
			copy(lastRegExp, compiled[lp:])
			if hyphen < 0 {
				// repeat regexp for n times
				// example: "(regexp){3}" -> "(regexp)(regexp)(regexp)"
				n, err := strconv.Atoi(string(inBrackets))
				if err != nil {
					panic(fmt.Sprintf("[surround %v] not a number in bracket: %v", i, err.Error()))
				}
				if n < 1 {
					panic(fmt.Sprintf("[surround %v] number in bracket less than 1", i))
				}
				compiled = append(compiled, repeatRunes(lastRegExp, n-1)...)
			} else {
				// repeat regexp for lo-hi times
				// example: "(regexp){1-3}" -> "((regexp)|(regexp)(regexp)|(regexp)(regexp)(regexp))"
				lo, err := strconv.Atoi(string(inBrackets[:hyphen]))
				if err != nil {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket: %v", i, err.Error()))
				}
				if lo < 0 {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket", i))
				}
				hi, err := strconv.Atoi(string(inBrackets[hyphen+1:]))
				if err != nil {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket: %v", i, err.Error()))
				}
				if hi <= lo {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket", i))
				}
				compiled = append(compiled[:lp], '(')
				for j := lo; j <= hi; j++ {
					compiled = append(compiled, repeatRunes(lastRegExp, j)...)
					if j != hi {
						compiled = append(compiled, '|')
					}
				}
				compiled = append(compiled, ')')
			}
		default:
			compiled = append(compiled, pattern[i])
			lp = i
		}
	}
	return compiled
}

// characterSetConv convert character set to multiple OR.
// [abc] => (a|b|c)
// [0-9] => (0|1|2|3|4|5|6|7|8|9)
// [^abc] => complement of (a|b|c)
func characterSetConv(cs []rune) []rune {
	return nil
}

func repeatRunes(runes []rune, count int) []rune {
	return []rune(strings.Repeat(string(runes), count))
}

func indexRune(runes []rune, r rune) int {
	if !utf8.ValidRune(r) {
		panic(fmt.Sprintf("invalid rune: %v", r))
	}
	for i := range runes {
		if runes[i] == r {
			return i
		}
	}
	return -1
}

func isPrimeRune(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '*' || r == '.'
}

func canBeTransferred(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '*' || r == '.' || r == '\\'
}

func getReachableStatus(g digraph.Digraph, src, reachable set.IntSet) {
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

func getStartStatus(g digraph.Digraph) set.IntSet {
	startStatus := set.NewIntSet()
	dfs := digraph.NewDFS(g, 0)
	rvQ := dfs.ReachableVertices()
	for !rvQ.IsEmpty() {
		startStatus.Add(rvQ.Front())
	}
	return startStatus
}

func makeNFA(symbolTable []symbol) digraph.Digraph {
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
				} else { // got '('
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

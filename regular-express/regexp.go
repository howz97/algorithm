package regexp

import (
	"fmt"
	"github.com/zh1014/algorithm/graphs/digraph"
	"github.com/zh1014/algorithm/queue"
	"github.com/zh1014/algorithm/set"
	"github.com/zh1014/algorithm/stack"
	"strconv"
	"strings"
	"unicode/utf8"
)

// IsMatch return whether the txt is match with the pattern
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
	pttrnRunes := compile([]rune(pattern))
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

func compile(pattern []rune) []rune {
	handled := make([]rune, 0, len(pattern)*2)
	lp := 0
	lpStack := stack.NewStackInt(10) // 随意设定
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '\\': // must put \ on top case
			lp = i
			i++
			handled = append(handled, '\\', pattern[i])
		case '(':
			lpStack.Push(len(handled))
			handled = append(handled, '(')
		case ')':
			lp = lpStack.Pop()
			handled = append(handled, ')')
		case '+':
			handled = append(handled, handled[lp:]...)
			handled = append(handled, '*')
		case '?':
			lastRegExp := make([]rune, len(handled[lp:]))
			copy(lastRegExp, handled[lp:])
			handled = append(handled[:lp], '(')
			handled = append(handled, lastRegExp...)
			handled = append(handled, '|', ')')
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
			lastRegExp := make([]rune, len(handled[lp:]))
			copy(lastRegExp, handled[lp:])
			if hyphen < 0 { // a number in brackets: {n}
				n, err := strconv.Atoi(string(inBrackets))
				if err != nil {
					panic(fmt.Sprintf("[surround %v] not a number in bracket: %v", i, err.Error()))
				}
				if n < 1 {
					panic(fmt.Sprintf("[surround %v] number in bracket less than 1", i))
				}
				handled = append(handled, repeatRunes(lastRegExp, n-1)...)
			} else { // a range in brackets: {n-m}
				lowerLimit, err := strconv.Atoi(string(inBrackets[:hyphen]))
				if err != nil {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket: %v", i, err.Error()))
				}
				if lowerLimit < 0 {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket", i))
				}
				upperLimit, err := strconv.Atoi(string(inBrackets[hyphen+1:]))
				if err != nil {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket: %v", i, err.Error()))
				}
				if upperLimit <= lowerLimit {
					panic(fmt.Sprintf("[surround %v] invalid range in bracket", i))
				}
				handled = append(handled[:lp], '(')
				for j := lowerLimit; j <= upperLimit; j++ {
					handled = append(handled, repeatRunes(lastRegExp, j)...)
					if j != upperLimit {
						handled = append(handled, '|')
					}
				}
				handled = append(handled, ')')
			}
		default:
			handled = append(handled, pattern[i])
			lp = i
		}
	}
	return handled
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

package regexp

import (
	"errors"
	"fmt"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/set"
	"github.com/howz97/algorithm/stack"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Regexp struct {
	table []symbol
	nfa   digraph.Digraph
	start *digraph.DFS
}

func (re *Regexp) Match(str string) bool {
	curStatus := re.startStatus()
	for _, r := range str {
		arrived := re.forwardStatus(curStatus, r)
		if arrived.IsEmpty() {
			return false
		}
		curStatus = re.updateCurStatus(arrived)
	}
	return curStatus.Contains(len(re.table))
}

func (re *Regexp) startStatus() set.IntSet {
	if re.start == nil {
		re.start = digraph.NewDFS(re.nfa, 0)
	}
	start := set.NewIntSet()
	rvQ := re.start.ReachableVertices() // todo optimize
	for !rvQ.IsEmpty() {
		start.Add(rvQ.Front())
	}
	return start
}

func (re *Regexp) forwardStatus(curStatus set.IntSet, r rune) set.IntSet {
	arrived := set.NewIntSet()
	for !curStatus.IsEmpty() {
		s := curStatus.RemoveOne()
		if re.table[s].match(r) {
			arrived.Add(s + 1)
		}
	}
	return arrived
}

func (re *Regexp) updateCurStatus(src set.IntSet) set.IntSet {
	// todo optimize
	tc := digraph.NewTransitiveClosure(re.nfa)
	srcQ := queue.NewIntQ()
	for !src.IsEmpty() {
		srcQ.PushBack(src.RemoveOne())
	}
	reachable := set.NewIntSet()
	reachableQ := tc.ReachableVertices(srcQ)
	for !reachableQ.IsEmpty() {
		reachable.Add(reachableQ.Front())
	}
	return reachable
}

func Compile(pattern string) (*Regexp, error) {
	compiled, err := compile([]rune(pattern))
	if err != nil {
		return nil, err
	}
	re := new(Regexp)
	re.table = makeSymbolTable(compiled)
	re.nfa = makeNFA(re.table)
	return re, nil
}

func Match(pattern, str string) (bool, error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.Match(str), nil
}

func makeSymbolTable(compiled []rune) []symbol {
	symbols := make([]symbol, 0, len(compiled))
	for i := 0; i < len(compiled); i++ {
		if compiled[i] == '\\' {
			i++
			symbols = append(symbols, symbol{
				isPrime: false,
				r:       compiled[i],
			})
			continue
		}
		symbols = append(symbols, symbol{
			isPrime: isPrime(compiled[i]),
			r:       compiled[i],
		})
	}
	return symbols
}

// compile high-level grammar into low-level grammar fixme: detect invalid pattern
func compile(pattern []rune) ([]rune, error) {
	compiled := make([]rune, 0, len(pattern)<<1)
	left := 0
	lpStack := stack.NewStackInt(0)
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '\\': // must put \ on top case
			left = i
			i++
			if !isTransferable(pattern[i]) {
				return nil, errors.New(fmt.Sprintf("invalid transfer: \\%v", string(pattern[i])))
			}
			compiled = append(compiled, '\\', pattern[i])
		case '(':
			lpStack.Push(len(compiled))
			compiled = append(compiled, '(')
		case ')':
			left = lpStack.Pop()
			compiled = append(compiled, ')')
		case '+':
			// "(regexp)+" -> "(regexp)(regexp)*"
			compiled = append(compiled, compiled[left:]...)
			compiled = append(compiled, '*')
		case '?':
			// "(regexp)?" -> "(regexp|)"
			lastRegExp := make([]rune, len(compiled[left:]))
			copy(lastRegExp, compiled[left:])
			compiled = append(compiled[:left], '(')
			compiled = append(compiled, lastRegExp...)
			compiled = append(compiled, '|', ')')
		case '{':
			rb := indexRune(pattern[i:], '}')
			if rb < 0 {
				return nil, errors.New(fmt.Sprintf("[surround %v] no corresponding right bracket", i))
			}
			inBrackets := pattern[i+1 : rb+i]
			i += rb
			if len(inBrackets) == 0 {
				return nil, errors.New(fmt.Sprintf("[surround %v] nothing in bracket", i))
			}
			hyphen := indexRune(inBrackets, '-')
			lastRegExp := make([]rune, len(compiled[left:]))
			copy(lastRegExp, compiled[left:])
			if hyphen < 0 {
				// repeat regexp for n times
				// example: "(regexp){3}" -> "(regexp)(regexp)(regexp)"
				n, err := strconv.Atoi(string(inBrackets))
				if err != nil {
					return nil, errors.New(fmt.Sprintf("[surround %v] not a number in bracket: %v", i, err.Error()))
				}
				if n < 1 {
					return nil, errors.New(fmt.Sprintf("[surround %v] number in bracket less than 1", i))
				}
				compiled = append(compiled, repeatRunes(lastRegExp, n-1)...)
			} else {
				// repeat regexp for lo-hi times
				// example: "(regexp){1-3}" -> "((regexp)|(regexp)(regexp)|(regexp)(regexp)(regexp))"
				lo, err := strconv.Atoi(string(inBrackets[:hyphen]))
				if err != nil {
					return nil, errors.New(fmt.Sprintf("[surround %v] invalid range in bracket: %v", i, err.Error()))
				}
				if lo < 0 {
					return nil, errors.New(fmt.Sprintf("[surround %v] invalid range in bracket", i))
				}
				hi, err := strconv.Atoi(string(inBrackets[hyphen+1:]))
				if err != nil {
					return nil, errors.New(fmt.Sprintf("[surround %v] invalid range in bracket: %v", i, err.Error()))
				}
				if hi <= lo {
					return nil, errors.New(fmt.Sprintf("[surround %v] invalid range in bracket", i))
				}
				compiled = append(compiled[:left], '(')
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
			left = i
		}
	}
	return compiled, nil
}

func isPrime(r rune) bool {
	return r == '(' || r == ')' || r == '|' || r == '*' || r == '.'
}

func isTransferable(r rune) bool {
	return isPrime(r) || r == '\\'
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

func makeNFA(table []symbol) digraph.Digraph {
	size := len(table)
	nfa := digraph.NewDigraph(size + 1)
	stk := stack.NewStackInt(size)
	for i, syb := range table {
		left := i
		if syb.isPrime {
			switch syb.r {
			case '(':
				nfa.AddEdge(i, i+1)
				stk.Push(i)
			case ')':
				nfa.AddEdge(i, i+1)
				allOr := queue.NewIntQ()
				for !stk.IsEmpty() {
					out := stk.Pop()
					if table[out].r == '|' {
						allOr.PushBack(out)
					} else {
						// got '('
						left = out
						break
					}
				}
				for !allOr.IsEmpty() {
					or := allOr.Front()
					nfa.AddEdge(left, or+1)
					nfa.AddEdge(or, i)
				}
			case '*':
				nfa.AddEdge(i, i+1)
			case '|':
				stk.Push(i)
			}
		}
		if i+1 < size && table[i+1].isClosure() {
			nfa.AddEdge(left, i+1)
			nfa.AddEdge(i+1, left)
		}
	}
	return nfa
}

type symbol struct {
	isPrime bool
	r       rune
}

func (s symbol) match(r rune) bool {
	return s.equal(r) || s.isWildCard()
}

func (s symbol) isWildCard() bool {
	return s.isPrime && s.r == '.'
}

func (s symbol) equal(r rune) bool {
	return !s.isPrime && s.r == r
}

func (s symbol) isClosure() bool {
	return s.isPrime && s.r == '*'
}

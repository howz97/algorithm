// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package regexp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/graphs"
)

func Match(pattern, str string) (bool, error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.Match(str), nil
}

func Compile(pattern string) (*Regexp, error) {
	compiled, err := compile([]rune(pattern))
	if err != nil {
		return nil, err
	}
	re := new(Regexp)
	re.makeSymbolTable(compiled)
	re.makeNFA()
	return re, nil
}

type Regexp struct {
	table []symbol
	nfa   *graphs.Digraph[int]
	tc    graphs.Reachable
}

func (re *Regexp) Match(str string) bool {
	// lazy init
	if re.tc == nil {
		re.tc = re.nfa.Reachable()
	}

	curStatus := re.startStatus()
	for _, r := range str {
		arrived := re.forwardStatus(curStatus, r)
		if arrived.Len() == 0 {
			return false
		}
		curStatus = re.updateCurStatus(arrived)
	}
	return curStatus.Contains(len(re.table))
}

func (re *Regexp) startStatus() basic.Set[int] {
	start := basic.NewSet[int]()
	re.tc.Iterate(0, func(v graphs.Id) bool {
		start.Add(int(v))
		return true
	})
	return start
}

func (re *Regexp) forwardStatus(curStatus basic.Set[int], r rune) basic.Set[int] {
	arrived := basic.NewSet[int]()
	for curStatus.Len() > 0 {
		s := curStatus.TakeOne()
		if s < len(re.table) && re.table[s].match(r) {
			arrived.Add(s + 1)
		}
	}
	return arrived
}

func (re *Regexp) updateCurStatus(sources basic.Set[int]) basic.Set[int] {
	reachable := basic.NewSet[int]()
	for sources.Len() > 0 {
		vSrc := sources.TakeOne()
		re.tc.Iterate(graphs.Id(vSrc), func(v graphs.Id) bool {
			reachable.Add(int(v))
			return true
		})
	}
	return reachable
}

func (re *Regexp) makeSymbolTable(compiled []rune) {
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
	re.table = symbols

	nfa := graphs.NewDigraph[int](uint(len(symbols) + 1))
	for i := 0; i <= len(symbols); i++ {
		id := nfa.AddVertex(i)
		if id != graphs.Id(i) {
			panic("id != graphs.Id(i)")
		}
	}
	re.nfa = nfa
}

// compile high-level grammar into low-level grammar
func compile(pattern []rune) ([]rune, error) {
	compiled := make([]rune, 0, len(pattern)<<1)
	left := 0
	lpStack := basic.NewStack[int](0)
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '\\': // must put \ on top case
			left = i
			i++
			if !isTransferable(pattern[i]) {
				return nil, fmt.Errorf("invalid transfer: \\%v", string(pattern[i]))
			}
			compiled = append(compiled, '\\', pattern[i])
		case '(':
			lpStack.PushBack(len(compiled))
			compiled = append(compiled, '(')
		case ')':
			if lpStack.Size() <= 0 {
				return nil, errors.New("'(' missing")
			}
			left = lpStack.Back()
			lpStack.PopBack()
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
				return nil, fmt.Errorf("[surround %v] no corresponding right bracket", i)
			}
			inBrackets := pattern[i+1 : rb+i]
			i += rb
			if len(inBrackets) == 0 {
				return nil, fmt.Errorf("[surround %v] nothing in bracket", i)
			}
			hyphen := indexRune(inBrackets, '-')
			lastRegExp := make([]rune, len(compiled[left:]))
			copy(lastRegExp, compiled[left:])
			if hyphen < 0 {
				// repeat regexp for n times
				// example: "(regexp){3}" -> "(regexp)(regexp)(regexp)"
				n, err := strconv.Atoi(string(inBrackets))
				if err != nil {
					return nil, fmt.Errorf("[surround %v] not a number in bracket: %v", i, err.Error())
				}
				if n < 1 {
					return nil, fmt.Errorf("[surround %v] number in bracket less than 1", i)
				}
				compiled = append(compiled, repeatRunes(lastRegExp, n-1)...)
			} else {
				// repeat regexp for lo-hi times
				// example: "(regexp){1-3}" -> "((regexp)|(regexp)(regexp)|(regexp)(regexp)(regexp))"
				lo, err := strconv.Atoi(string(inBrackets[:hyphen]))
				if err != nil {
					return nil, fmt.Errorf("[surround %v] invalid range in bracket: %v", i, err.Error())
				}
				if lo < 0 {
					return nil, fmt.Errorf("[surround %v] invalid range in bracket", i)
				}
				hi, err := strconv.Atoi(string(inBrackets[hyphen+1:]))
				if err != nil {
					return nil, fmt.Errorf("[surround %v] invalid range in bracket: %v", i, err.Error())
				}
				if hi <= lo {
					return nil, fmt.Errorf("[surround %v] invalid range in bracket", i)
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

func (re *Regexp) makeNFA() {
	size := len(re.table)
	stk := basic.NewStack[int](size)
	for i, syb := range re.table {
		left := i
		if syb.isPrime {
			switch syb.r {
			case '(':
				re.nfa.AddEdge(graphs.Id(i), graphs.Id(i+1))
				stk.PushBack(i)
			case ')':
				re.nfa.AddEdge(graphs.Id(i), graphs.Id(i+1))
				allOr := basic.NewList[int]()
				for {
					out := stk.Back()
					stk.PopBack()
					if re.table[out].r == '|' {
						allOr.PushBack(out)
					} else {
						// got '('
						left = out
						break
					}
				}
				for allOr.Size() > 0 {
					or := allOr.Front()
					allOr.PopFront()
					re.nfa.AddEdge(graphs.Id(left), graphs.Id(or+1))
					re.nfa.AddEdge(graphs.Id(or), graphs.Id(i))
				}
			case '*':
				re.nfa.AddEdge(graphs.Id(i), graphs.Id(i+1))
			case '|':
				stk.PushBack(i)
			case '.':
			default:
				panic(fmt.Sprintf("unknown prime %v", syb.r))
			}
		}
		if i+1 < size && re.table[i+1].isClosure() {
			re.nfa.AddEdge(graphs.Id(left), graphs.Id(i+1))
			re.nfa.AddEdge(graphs.Id(i+1), graphs.Id(left))
		}
	}
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

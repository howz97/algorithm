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

package strings

const (
	byteNum = 1 << 8
)

type KMP struct {
	stateCnt int
	// deterministic finite automaton
	dfa [byteNum][]int
}

func NewKMP(pattern string) *KMP {
	kmp := &KMP{}
	kmp.stateCnt = len(pattern)
	for i := range kmp.dfa {
		kmp.dfa[i] = make([]int, kmp.stateCnt)
	}

	// construct dfa
	dfa := kmp.dfa
	dfa[pattern[0]][0] = 1 // dfa[][0] is special

	state := 1
	rs := 0 // restart state
	for state < kmp.stateCnt {
		for i := 0; i < byteNum; i++ {
			dfa[i][state] = dfa[i][rs]
		}
		dfa[pattern[state]][state] = state + 1

		rs = dfa[pattern[state]][rs]
		state++
	}
	return kmp
}

func (kmp *KMP) Index(s string) int {
	if len(s) < kmp.stateCnt {
		return -1
	}
	i, state := 0, 0
	for ; i < len(s) && state < kmp.stateCnt; i++ {
		state = kmp.dfa[s[i]][state]
	}
	if state == kmp.stateCnt {
		return i - state
	}
	return -1
}

func (kmp *KMP) IndexAll(s string) (indices []int) {
	j := 0
	for {
		i := kmp.Index(s)
		if i < 0 {
			break
		}
		indices = append(indices, j+i)
		j = j + i + kmp.stateCnt
		s = s[i+kmp.stateCnt:]
	}
	return
}

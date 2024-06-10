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

type BoyerMoore struct {
	pattern string
	right   [byteNum]int
}

func NewBoyerMoore(pattern string) *BoyerMoore {
	bm := &BoyerMoore{
		pattern: pattern,
	}
	for i := 0; i < len(pattern); i++ {
		bm.right[pattern[i]] = i
	}
	return bm
}

func (bm *BoyerMoore) Index(s string) int {
	for i := 0; i < len(s)-bm.LenP(); {
		j := bm.LenP() - 1
		for ; j >= 0 && s[i+j] == bm.pattern[j]; j-- {
		}
		if j < 0 {
			return i
		}
		skip := j - bm.right[s[i+j]]
		if skip < 1 {
			skip = 1
		}
		i += skip
	}
	return -1
}

func (bm *BoyerMoore) IndexAll(s string) (indices []int) {
	j := 0
	for {
		i := bm.Index(s)
		if i < 0 {
			break
		}
		indices = append(indices, j+i)
		j = j + i + bm.LenP()
		s = s[i+bm.LenP():]
	}
	return
}

func (bm *BoyerMoore) LenP() int {
	return len(bm.pattern)
}

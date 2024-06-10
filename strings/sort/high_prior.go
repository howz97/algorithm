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

package sort

import "github.com/howz97/algorithm/strings/alphabet"

func HighPrior(strings []string) {
	if len(strings) <= 1 {
		return
	}
	aux := make([]string, len(strings))
	highPriorUnicode(strings, aux, 0, len(strings)-1, 0)
}

func highPriorUnicode(strings, aux []string, lo, hi, depth int) {
	if lo >= hi {
		return
	}
	// count[1...257] -> count(missing,0,1...255)
	count := make([]int, 258)
	for i := lo; i <= hi; i++ {
		count[getByte(strings[i], depth)+2]++
	}

	// index[0...256] -> startIndex(missing,0,1...255)
	index := count
	index[0] = lo
	for i := 1; i <= 256; i++ {
		index[i] += index[i-1]
	}

	for i := lo; i <= hi; i++ {
		b := getByte(strings[i], depth)
		aux[index[b+1]] = strings[i]
		// index[0...256] -> nextIndex(missing,0,1...255)
		index[b+1]++
	}
	// write back
	for i := lo; i <= hi; i++ {
		strings[i] = aux[i]
	}
	// index[0...255] -> startIndex(0,1...255)
	for i := 0; i < 256; i++ {
		highPriorUnicode(strings, aux, index[i], index[i+1]-1, depth+1)
	}
}

func getByte(runes string, depth int) int {
	if depth >= len(runes) {
		return -1
	}
	return int(runes[depth])
}

// HighPriorAlp sort strings with a user-defined alphabet
func HighPriorAlp(a alphabet.IAlp, data []string) {
	if len(data) <= 1 {
		return
	}
	runes := make([][]rune, len(data))
	for i := range data {
		runes[i] = []rune(data[i])
	}
	aux := make([][]rune, len(data))
	highPriorSort(a, runes, aux, 0, len(data)-1, 0)
	for i := range runes {
		data[i] = string(runes[i])
	}
}

func highPriorSort(a alphabet.IAlp, runes, aux [][]rune, lo, hi, depth int) {
	if lo >= hi {
		return
	}
	// count[1...R+1] -> count(missing,0,1...R-1)
	count := make([]int, a.R()+2)
	for i := lo; i <= hi; i++ {
		count[toIndex(a, runes[i], depth)+2]++
	}

	// index[0...R] -> startIndex(missing,0,1...R-1)
	index := count
	index[0] = lo
	for i := 1; i <= a.R(); i++ {
		index[i] += index[i-1]
	}

	for i := lo; i <= hi; i++ {
		aux[index[toIndex(a, runes[i], depth)+1]] = runes[i]
		// index[0...R] -> nextIndex(missing,0,1...R-1)
		index[toIndex(a, runes[i], depth)+1]++
	}
	// write back
	for i := lo; i <= hi; i++ {
		runes[i] = aux[i]
	}
	// index[0...R] -> startIndex(0,1...R)
	for i := 0; i < a.R(); i++ {
		highPriorSort(a, runes, aux, index[i], index[i+1]-1, depth+1)
	}
}

func toIndex(a alphabet.IAlp, runes []rune, depth int) rune {
	if depth >= len(runes) {
		return -1
	}
	return a.ToIndex(runes[depth])
}

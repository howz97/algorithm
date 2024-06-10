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

package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

func ReadAllLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func RangeFileLines(filename string, fn func(string) bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if !fn(scan.Text()) {
			break
		}
	}
	return nil
}

func AverageDuration(dur []time.Duration) (avg time.Duration) {
	for _, e := range dur {
		avg += e
	}
	avg /= time.Duration(len(dur))
	return
}

func IndexStringSlice(strings []string, str string) int {
	for i, s := range strings {
		if s == str {
			return i
		}
	}
	return -1
}

func IsRunesMatch(pattern, runes []rune) bool {
	if len(pattern) != len(runes) {
		return false
	}
	for i, p := range pattern {
		if !IsRuneMatch(p, runes[i]) {
			return false
		}
	}
	return true
}

func IsRuneMatch(p, r rune) bool {
	return p == '.' || p == r
}

func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func Max[Ord constraints.Ordered](a, b Ord) Ord {
	if a > b {
		return a
	}
	return b
}

func ExecCost(fn func()) time.Duration {
	t := time.Now()
	fn()
	return time.Since(t)
}

func IndexInt(ints []int, n int) int {
	for i, e := range ints {
		if e == n {
			return i
		}
	}
	return -1
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Reverse[T any](data []T) {
	m := len(data) / 2
	j := len(data) - 1
	for i := 0; i < m; {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
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

type Elems []T

func (e Elems) Len() int {
	return len(e)
}

func (e Elems) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Reversible interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func ReverseSlice(data Reversible) {
	m := data.Len() / 2
	j := data.Len() - 1
	for i := 0; i < m; {
		data.Swap(i, j)
		i++
		j--
	}
}

func ReverseInts(data []int) {
	ReverseSlice(sort.IntSlice(data))
}

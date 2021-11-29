package util

import (
	"io/ioutil"
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

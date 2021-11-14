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

func SimplePatternMatch(pattern, str []rune) bool {
	if len(pattern) != len(str) {
		return false
	}
	for i, r := range pattern {
		if r == '.' {
			continue
		}
		if r != str[i] {
			return false
		}
	}
	return true
}

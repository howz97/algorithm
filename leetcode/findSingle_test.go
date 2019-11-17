package leetcode

import (
	"fmt"
	"testing"
)

func TestFindSingle(t *testing.T) {
	n := findSingle([]int{-555, -555, -555, 234, 999, 456, 999, 234, 234, -1014, 999, 456, 456, 1000, 1000, 1000})
	fmt.Println(n)
}

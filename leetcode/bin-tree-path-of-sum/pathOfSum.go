package pathofsum

import "fmt"

type node struct {
	v     int
	left  *node
	right *node
}

func printPathOfSum(t *node, sum int, pre []int) {
	if t == nil {
		return
	}
	pre = append(pre, t.v)
	sum -= t.v
	if t.left != nil {
		printPathOfSum(t.left, sum, pre)
	}
	if t.right != nil {
		printPathOfSum(t.right, sum, pre)
	}
	if sum == 0 && t.left == nil && t.right == nil {
		fmt.Println(pre)
	}
	pre = pre[:len(pre)-1]
}

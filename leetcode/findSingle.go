package leetcode

// numbers整形数组中除了一个数字唯一，其他都出现了三次，找出唯一的数字
func findSingle(numbers []int) int {
	count := make([]int, 64)
	for _, n := range numbers {
		index := 0
		for n != 0 {
			count[index] += n & 1
			n = n >> 1
			index++
		}
	}
	result := 0
	for i, c := range count {
		if c%3 == 1 {
			result |= 1 << uint(i)
		}
	}
	return result
}

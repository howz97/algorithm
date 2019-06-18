package repeatnumber

/*
在一个长度为 n的数组里的所有数字都在0到 n-1的范围内。
数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。
请找出数组中任意一个重复的数字。
*/

// FindRepeatNo -
func FindRepeatNo(numbers []int, length int) int {
	for i := 0; i < length; i++ {
		if numbers[i] == i {
			continue
		}
		if numbers[numbers[i]] == numbers[i] {
			return numbers[i]
		}
		numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		i--
	}
	return -1
}

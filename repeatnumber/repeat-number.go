package repeatnumber

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

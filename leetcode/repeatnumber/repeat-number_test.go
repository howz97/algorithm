package repeatnumber

import "testing"

func Test_FindRepeatNo(t *testing.T) {
	numbers := []int{4, 2, 8, 3, 7, 2, 9, 3, 4, 1, 0}
	println(FindRepeatNo(numbers, len(numbers)))
}

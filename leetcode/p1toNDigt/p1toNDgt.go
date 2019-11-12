package p1toNDigt

import (
	"errors"
	"strconv"
)

func printR(pre string, n int) {
	if n == 0 {
		println(pre)
		return
	}
	for i := 0; i < 10; i++ {
		printR(pre+strconv.Itoa(i), n-1)
	}
}

func print1toMaxOfNDigits(n int) error {
	if n <= 0 {
		return errors.New("invalid input")
	}
	printR("", n)
	return nil
}

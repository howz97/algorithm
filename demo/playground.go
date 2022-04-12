package main

import (
	"fmt"
)

func main() {
	s := "0123456789"
	s2 := s
	bytes := []byte(s)
	bytes[9] = 'X'
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(string(bytes))
}

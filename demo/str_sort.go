package main

import (
	"fmt"

	"github.com/howz97/algorithm/strings/string_sort"
)

func main() {
	data := []string{
		"ABABC",
		"ABAAA",
		"BBAAA",
		"CA",
		"HUAWEI",
		"AA",
		"张豪",
		"张三",
		"李四",
		"王麻子",
		"李二狗",
	}
	string_sort.Quick3(data)
	fmt.Println(data)
}

package main

import (
	"fmt"

	"github.com/howz97/algorithm/strings/sort"
)

func demo_strSort() {
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
	sort.Quick3(data)
	fmt.Println(data)
}

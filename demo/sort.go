package main

import (
	"fmt"

	"github.com/howz97/algorithm/sort"
)

func main() {
	sli := []int{46, 4, 1, 26, 4, 25}
	sort.QuickSort(sli)
	fmt.Println(sli)
}

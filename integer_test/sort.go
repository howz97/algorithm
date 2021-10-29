package main

import (
	"fmt"
	"github.com/howz97/algorithm/sort"
	stdsort "sort"
)

func main() {
	sli := stdsort.IntSlice{46, 4, 1, 26, 4, 25}
	sort.QuickSort(sli)
	fmt.Println(sli)
}

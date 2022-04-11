package sort

import (
	"fmt"
	"math/rand"
	stdsort "sort"
	"testing"
	"time"
)

func TestContrast(t *testing.T) {
	testPerformance(PopSort[int], "PopSort")
	testPerformance(SelectSort[int], "SelectSort")
	testPerformance(InsertSort[int], "InsertSort")
	testPerformance(ShellSort[int], "ShellSort")
	testPerformance(MergeSort[int], "MergeSort")
	testPerformance(HeapSort[int], "HeapSort")
	testPerformance(QuickSort[int], "QuickSort")

	testPerformance(func(p []int) { stdsort.Sort(stdsort.IntSlice(p)) }, "Go library sort.Ints")
}

const (
	testFreq         = 1
	inputSize        = 10000
	randInputUpLimit = 100000000
	dupInputUpLimit  = 10
)

var (
	inputData = make([]int, inputSize)
)

func testPerformance(sortAlg func(p []int), algName string) {
	performanceRandomInput(sortAlg, algName)
	// performanceDupInput(sortAlg, algName)
	// performanceSortedInput(sortAlg, algName)
	// performanceReverseSortedInput(sortAlg, algName)
	fmt.Println("-------------------------------------")
}

func performanceRandomInput(sortAlg func(p []int), algName string) {
	fmt.Printf("%v : random(%v):\n", algName, inputSize)
	for pass := 0; pass < testFreq; pass++ {
		genRandomData(inputData, randInputUpLimit)
		start := time.Now()
		sortAlg(inputData)
		elapsed := time.Since(start)
		if !stdsort.IntsAreSorted(inputData) {
			panic("failed to sort: " + algName)
		}
		fmt.Print(elapsed.String(), "  ")
	}
	fmt.Println()
}

func performanceDupInput(sortAlg func(p []int), algName string) {
	fmt.Printf("%v : dup(%v):\n", algName, inputSize)
	for pass := 0; pass < testFreq; pass++ {
		genRandomData(inputData, dupInputUpLimit)
		start := time.Now()
		sortAlg(inputData)
		elapsed := time.Since(start)
		if !stdsort.IntsAreSorted(inputData) {
			panic("failed to sort")
		}
		fmt.Print(elapsed.String(), "  ")
	}
	fmt.Println()
}

func performanceSortedInput(sortAlg func(p []int), algName string) {
	fmt.Printf("%v : sorted(%v):\n", algName, inputSize)
	genSortedData(inputData)
	for pass := 0; pass < testFreq; pass++ {
		start := time.Now()
		sortAlg(inputData)
		elapsed := time.Since(start)
		if !stdsort.IntsAreSorted(inputData) {
			panic("failed to sort")
		}
		fmt.Print(elapsed.String(), "  ")
	}
	fmt.Println()
}

func performanceReverseSortedInput(sortAlg func(p []int), algName string) {
	fmt.Printf("%v : reverse(%v):\n", algName, inputSize)
	for pass := 0; pass < testFreq; pass++ {
		genReverseSortedData(inputData)
		start := time.Now()
		sortAlg(inputData)
		elapsed := time.Since(start)
		if !stdsort.IntsAreSorted(inputData) {
			panic("failed to sort")
		}
		fmt.Print(elapsed.String(), "  ")
	}
	fmt.Println()
}

func genRandomData(data []int, uplimit int) {
	rand.Seed(14)
	for i := range data {
		data[i] = rand.Intn(uplimit)
	}
}

func genSortedData(data []int) {
	for i := range data {
		data[i] = i
	}
}

func genReverseSortedData(data []int) {
	n := len(data)
	for i := range data {
		n--
		data[i] = n
	}
}

func Test_cutOff(t *testing.T) {
	//sli := stdsort.IntSlice{9,1,1,4,4,5,6,7,3,2}
	sli := stdsort.IntSlice{9, 8, 7, 4, 4, 5, 6, 7, 3, 2}
	m := cutOff(sli, 0, 9)
	t.Log(m, sli)
}

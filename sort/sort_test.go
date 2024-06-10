// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sort

import (
	"fmt"
	"math/rand"
	stdsort "sort"
	"testing"
	"time"
)

func TestContrast(t *testing.T) {
	testPerformance(Bubble[int], "PopSort")
	testPerformance(Select[int], "SelectSort")
	testPerformance(Insert[int], "InsertSort")
	testPerformance(Shell[int], "ShellSort")
	testPerformance(Merge[int], "MergeSort")
	testPerformance(Heap[int], "HeapSort")
	testPerformance(Quick[int], "QuickSort")

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

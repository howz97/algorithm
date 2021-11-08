package string_sort

import (
	"github.com/howz97/algorithm/alphabet"
	mysort "github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/util"
	"sort"
	"testing"
	"time"
)

//const filename = "./input.txt"
const filename = "../string_search/tale.txt"
const testTimes = 10

func Test_HighPrior(t *testing.T) {
	LoopTest(t, HighPrior, "HighPrior")
}

func Test_HighPriorWithAlphabet(t *testing.T) {
	LoopTest(t, func(data []string) {
		HighPriorWithAlphabet(alphabet.Ascii, data)
	}, "HighPriorWithAlphabet")
}

func Test_Quick3(t *testing.T) {
	LoopTest(t, Quick3, "Quick3")
}

func Test_Quick3WithAlphabet(t *testing.T) {
	LoopTest(t, func(data []string) {
		Quick3WithAlphabet(alphabet.Ascii, data)
	}, "Quick3WithAlphabet")
}

func TestStdSort(t *testing.T) {
	LoopTest(t, sort.Strings, "StdSort")
}

func TestQuickSort(t *testing.T) {
	LoopTest(t, func(data []string) {
		mysort.QuickSort(sort.StringSlice(data))
	}, "StdSort")
}

func LoopTest(t *testing.T, fn func([]string), desc string) {
	var results []time.Duration
	for i := 0; i < testTimes; i++ {
		data := util.ReadAllLines(filename)
		start := time.Now()
		fn(data)
		dur := time.Since(start)
		results = append(results, dur)
		if !sort.StringsAreSorted(data) {
			t.Fatalf("%s sort failed", desc)
		}
	}
	t.Logf("%s results: avg=%v, %v", desc, util.AverageDuration(results), results)
}

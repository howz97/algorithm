package string_sort

import (
	"fmt"
	"github.com/howz97/algorithm/alphabet"
	mysort "github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/util"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"
)

const filename = "./length_rand.txt"
const testTimes = 10
const inputSize = 100000

func TestCreateInput(t *testing.T) {
	CreateInputStrings(t, "./ip.txt", func() string {
		return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	})

	CreateInputStrings(t, "./length5.txt", func() string {
		return RandString(5)
	})

	CreateInputStrings(t, "./length50.txt", func() string {
		return RandString(50)
	})

	CreateInputStrings(t, "./length_rand.txt", func() string {
		return RandString(rand.Intn(100))
	})
}

func RandString(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += string(alphabet.Ascii.Rand())
	}
	return str
}

func CreateInputStrings(t *testing.T, filename string, fn func() string) {
	ipf, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer ipf.Close()
	for i := 0; i < inputSize; i++ {
		_, err = ipf.WriteString(fn() + "\n")
		if err != nil {
			t.Fatal(err)
		}
	}
}

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
	}, "MyQuickSort")
}

func TestCompare(t *testing.T) {
	//Test_HighPrior(t)
	Test_HighPriorWithAlphabet(t)
	Test_Quick3(t)
	Test_Quick3WithAlphabet(t)
	TestQuickSort(t)
	TestStdSort(t)
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

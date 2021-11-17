package string_sort

import (
	"github.com/howz97/algorithm/alphabet"
	mysort "github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/util"
	"os"
	"sort"
	"testing"
	"time"
)

const filename = "./short.txt"
const testTimes = 1
const inputSize = 1000

var alpha = alphabet.NewAlphabetImpl(alphabet.UPPERCASE)

func TestCreateInput(t *testing.T) {
	//CreateInputStrings(t, "./ip.txt", func() string {
	//	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	//})
	//
	//CreateInputStrings(t, "./short.txt", func() string {
	//	return RandString(5)
	//})

	prefix := ""
	for i := 0; i < 300000; i++ {
		prefix += "A"
	}
	CreateInputStrings(t, "./long.txt", func() string {
		return prefix + alpha.RandString(10)
	})

	//CreateInputStrings(t, "./length_rand.txt", func() string {
	//	return RandString(rand.Intn(100))
	//})
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
		HighPriorWithAlphabet(alpha, data)
	}, "HighPriorWithAlphabet")
}

func Test_Quick3(t *testing.T) {
	LoopTest(t, Quick3, "Quick3")
}

func Test_Quick3WithAlphabet(t *testing.T) {
	LoopTest(t, func(data []string) {
		Quick3WithAlphabet(alpha, data)
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
	Test_Quick3WithAlphabet(t)
	Test_Quick3(t)
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

		for i := 1; i < len(data); i++ {
			if data[i] < data[i-1] {
				t.Fatalf("%s sort failed: %s, %s", desc, data[i-1], data[i])
			}
		}
	}
	t.Logf("%s results: avg=%v, %v", desc, util.AverageDuration(results), results)
}

func TestStringCompare(t *testing.T) {
	str1 := alpha.RandString(200000)
	str2 := str1
	b := true

	t0 := time.Now()
	runes1 := []rune(str1)
	runes2 := []rune(str2)
	for i := range runes1 {
		if runes1[i] < runes2[i] {
			b = false
			break
		}
	}
	t.Logf("compare runes %v %v", time.Since(t0), b)

	t0 = time.Now()
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	for i := range bytes1 {
		if bytes1[i] < bytes2[i] {
			b = false
			break
		}
	}
	t.Logf("range bytes %v %v", time.Since(t0), b)

	t0 = time.Now()
	b = str1 > str2
	t.Logf("builtin %v %v", time.Since(t0), b)
}

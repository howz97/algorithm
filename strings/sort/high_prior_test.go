package sort

import (
	"os"
	"sort"
	"testing"
	"time"

	mysort "github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/strings/alphabet"
	"github.com/howz97/algorithm/util"
)

const filename = "./dna.txt"
const testTimes = 1
const inputSize = 3000

var alpha = alphabet.NewAlphabetImpl(alphabet.UPPERCASE)

func TestCreateInput(t *testing.T) {
	//CreateInputStrings(t, "./ip.txt", func() string {
	//	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	//})
	//
	//CreateInputStrings(t, "./short.txt", func() string {
	//	return RandString(5)
	//})

	// alpha = alphabet.NewAlphabetImpl(alphabet.DNA)
	// CreateInputStrings(t, "./dna.txt", func() string {
	// 	return alpha.RandString(1000)
	// })

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

func Test_HighPriorAlp(t *testing.T) {
	LoopTest(t, func(data []string) {
		HighPriorAlp(alpha, data)
	}, "HighPriorAlp")
}

func Test_HighPrior(t *testing.T) {
	LoopTest(t, func(data []string) {
		HighPrior(data)
	}, "HighPrior")
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
		mysort.Quick(data)
	}, "MyQuickSort")
}

func TestCompare(t *testing.T) {
	Test_HighPrior(t)
	Test_HighPriorAlp(t)
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

	runes1 := []rune(str1)
	runes2 := []rune(str2)
	t0 := time.Now()
	for i := range runes1 {
		if runes1[i] < runes2[i] {
			b = false
			break
		}
	}
	t.Logf("compare runes %v %v", time.Since(t0), b)

	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	t0 = time.Now()
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

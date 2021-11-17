package hash_table

import (
	"fmt"
	"github.com/howz97/algorithm/sort"
	"os"
	"strings"
	"testing"
)

func TestStr_HashCode(t *testing.T) {
	file, err := os.Open("/Users/zhanghao/go/src/github.com/zh1014/algorithm/str-search/tale.txt")
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	txt := make([]byte, fileStat.Size())
	_, err = file.Read(txt)
	words := strings.Split(string(txt), " ")
	rang := 97
	count := make([]int, rang)
	for i := range words {
		count[(Str(words[i]).Hash()&0x7fffffffffffffff)%rang]++
	}
	sort.QuickSort(count)
	fmt.Println(count)
}

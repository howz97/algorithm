package main

import (
	"fmt"
	"github.com/howz97/algorithm/str_search"
	"os"
)

func main() {
	file, err := os.Open("../str_search/tale.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	txt := make([]byte, fileStat.Size())
	_, err = file.Read(txt)
	if err != nil {
		panic(err)
	}

	pattern := "It is a far, far better thing that I do, than I have ever done"
	searcher := str_search.NewKMP(pattern)
	//searcher := str_search.NewBM(pattern)
	i := searcher.Index(string(txt))
	//i := str_search.IndexRabinKarp(string(txt), pattern)
	fmt.Println(string(txt[i-50 : i+100]))
}

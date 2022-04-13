package main

import (
	"fmt"
	"io/ioutil"

	"github.com/howz97/algorithm/strings"
)

func main() {
	txt, err := ioutil.ReadFile("../strings/tale.txt")
	if err != nil {
		panic(err)
	}
	pattern := "It is a far, far better thing that I do, than I have ever done"
	searcher := strings.NewKMP(pattern)
	//searcher := strings.NewBM(pattern)
	i := searcher.Index(string(txt))
	//i := strings.IndexRabinKarp(string(txt), pattern)
	fmt.Println(string(txt[i-50 : i+100]))

	/*
		 story, with a tender and a faltering
		voice.

		"It is a far, far better thing that I do, than I have ever done;
		it is a far, far better rest that I
	*/
}

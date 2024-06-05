package main

import (
	"fmt"

	"github.com/howz97/algorithm/strings/compress/lzw"
)

func compress() {
	data := []byte("howhowhowhowhowhowhowhowhow")
	compressed := lzw.Compress(data)
	fmt.Printf("performance %.4f \n", float64(len(compressed))/float64(len(data)))
	lzw.Decompress(compressed)
}

//output: performance 0.8889

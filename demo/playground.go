package main

import (
	"compress/lzw"
	"os"
)

func main() {
	f, _ := os.Open("xxx")
	lzw.NewWriter(f, lzw.LSB, 2)
}

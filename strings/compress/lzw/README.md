[LZW compression](https://pkg.go.dev/github.com/howz97/algorithm/compress/lzw)
```go
func Example() {
	data := []byte("howhowhowhowhowhowhowhowhow")
	compressed := Compress(data)
	fmt.Printf("performance %.4f\n", float64(len(compressed))/float64(len(data)))
	fmt.Println(string(Decompress(compressed)))

	// Output:
	// performance 0.8889
	// howhowhowhowhowhowhowhowhow
}
```
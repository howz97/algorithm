Huffman compression
```go
func Example() {
	data := []byte("zhang how, zhang how, zhang how, zhang how,")
	compressed := Compress(data)
	fmt.Printf("performance %.4f\n", float64(len(compressed))/float64(len(data)))
	de, _ := Decompress(compressed)
	fmt.Println(string(de))

	// Output:
	// performance 0.7674
	// zhang how, zhang how, zhang how, zhang how,
}
```
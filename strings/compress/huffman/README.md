[哈夫曼编码压缩](https://pkg.go.dev/github.com/howz97/algorithm/compress/huffman)

```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/compress/huffman"
)

func main() {
	data := []byte("zhang how, zhang how, zhang how, zhang how,")
	compressed := huffman.Compress(data)
	fmt.Printf("performance %.4f \n", float64(len(compressed))/float64(len(data)))
	huffman.Decompress(compressed)
}

//output:  performance 0.7674
```
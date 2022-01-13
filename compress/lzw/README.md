[LZW压缩](https://pkg.go.dev/github.com/howz97/algorithm/compress/lzw)

```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/compress/lzw"
)

func main() {
	data := []byte("howhowhowhowhowhowhowhowhow")
	compressed := lzw.Compress(data)
	fmt.Printf("performance %.4f \n", float64(len(compressed))/float64(len(data)))
	lzw.Decompress(compressed)
}

//output: performance 0.8889
```
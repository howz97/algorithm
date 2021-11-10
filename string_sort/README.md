字符串排序
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/string_sort"
)

func main() {
	data := []string{
		"ABABC",
		"ABAAA",
		"BBAAA",
		"CA",
		"HUAWEI",
		"AA",
	}
	string_sort.Quick3ASCII(data)
	fmt.Println(data)

	uniData := []string{
		"张豪",
		"张三",
		"李四",
		"王麻子",
		"李二狗",
	}
	string_sort.Quick3(uniData)
	fmt.Println(uniData)
}

/*
[AA ABAAA ABABC BBAAA CA HUAWEI]
[张三 张豪 李二狗 李四 王麻子]
*/
```
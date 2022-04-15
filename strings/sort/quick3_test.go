package sort

import "fmt"

func ExampleQuick3() {
	data := []string{
		"ABABC",
		"ABAAA",
		"BBAAA",
		"CA",
		"HUAWEI",
		"AA",
		"张豪",
		"张三",
		"李四",
		"王麻子",
		"李二狗",
	}
	Quick3(data)
	fmt.Println(data)
	// Output: [AA ABAAA ABABC BBAAA CA HUAWEI 张三 张豪 李二狗 李四 王麻子]
}

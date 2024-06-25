// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphs

import (
	"fmt"
)

func ExampleSymbol() {
	g, err := LoadSymbGraph("../assets/graphs/symbol_graph.yml")
	if err != nil {
		panic(err)
	}
	bfs := g.BFS(g.IdOf("姜文"))
	fmt.Println(bfs.ShortestPathTo(g.IdOf("梁朝伟")).Str())
	fmt.Println(bfs.ShortestPathTo(g.IdOf("宋慧乔")).Str())
	fmt.Println(bfs.ShortestPathTo(g.IdOf("郎雄")).Str())
	fmt.Println(bfs.ShortestPathTo(g.IdOf("周星驰")).Str())
	fmt.Println(bfs.ShortestPathTo(g.IdOf("梁家辉")).Str())

	// Output:
	// [TotalDistance=6] 姜文->《让子弹飞》(1.00) 《让子弹飞》->刘嘉玲(1.00) 刘嘉玲->《阿飞正传》(1.00) 《阿飞正传》->刘德华(1.00) 刘德华->《无间道》(1.00) 《无间道》->梁朝伟(1.00)
	// [TotalDistance=6] 姜文->《让子弹飞》(1.00) 《让子弹飞》->周润发(1.00) 周润发->《卧虎藏龙》(1.00) 《卧虎藏龙》->章子怡(1.00) 章子怡->《一代宗师》(1.00) 《一代宗师》->宋慧乔(1.00)
	// [TotalDistance=4] 姜文->《让子弹飞》(1.00) 《让子弹飞》->周润发(1.00) 周润发->《卧虎藏龙》(1.00) 《卧虎藏龙》->郎雄(1.00)
	// [TotalDistance=6] 姜文->《让子弹飞》(1.00) 《让子弹飞》->刘嘉玲(1.00) 刘嘉玲->《阿飞正传》(1.00) 《阿飞正传》->张国荣(1.00) 张国荣->《家有喜事》(1.00) 《家有喜事》->周星驰(1.00)
	// [TotalDistance=4] 姜文->《让子弹飞》(1.00) 《让子弹飞》->周润发(1.00) 周润发->《赌神2》(1.00) 《赌神2》->梁家辉(1.00)
}

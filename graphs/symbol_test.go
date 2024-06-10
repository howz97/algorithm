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
	g, _ := LoadGraph("./testdata/symbol_graph.yml")
	bfs := g.BFS(g.VetOf("姜文"))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁朝伟")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("宋慧乔")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("郎雄")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("周星驰")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁家辉")).Str(g.Symbol))

	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张学友, 张学友->《东邪西毒》, 《东邪西毒》->梁朝伟,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->章子怡, 章子怡->《一代宗师》, 《一代宗师》->宋慧乔,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->郎雄,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张曼玉, 张曼玉->《家有喜事》, 《家有喜事》->周星驰,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《赌神2》, 《赌神2》->梁家辉,
}

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

func ExampleSymbGraph() {
	g, err := LoadSymbGraph("../assets/graphs/symbol_graph.yml")
	if err != nil {
		panic(err)
	}
	bfs := g.BFS(g.IdOf("姜文"))
	fmt.Println(bfs.ShortestPathTo(g.IdOf("宋慧乔")))
	fmt.Println(bfs.ShortestPathTo(g.IdOf("郎雄")))
	fmt.Println(bfs.ShortestPathTo(g.IdOf("梁家辉")))
	fmt.Println(bfs.ShortestPathTo(g.IdOf("赵本山")))
	fmt.Println(bfs.ShortestPathTo(g.IdOf("赵文瑄")))

	// Output:
	// [Distance=6] 姜文->《让子弹飞》 《让子弹飞》->周润发 周润发->《卧虎藏龙》 《卧虎藏龙》->章子怡 章子怡->《一代宗师》 《一代宗师》->宋慧乔
	// [Distance=4] 姜文->《让子弹飞》 《让子弹飞》->周润发 周润发->《卧虎藏龙》 《卧虎藏龙》->郎雄
	// [Distance=4] 姜文->《让子弹飞》 《让子弹飞》->周润发 周润发->《赌神2》 《赌神2》->梁家辉
	// [Distance=6] 姜文->《让子弹飞》 《让子弹飞》->周润发 周润发->《卧虎藏龙》 《卧虎藏龙》->章子怡 章子怡->《一代宗师》 《一代宗师》->赵本山
	// [Distance=6] 姜文->《让子弹飞》 《让子弹飞》->周润发 周润发->《卧虎藏龙》 《卧虎藏龙》->郎雄 郎雄->《喜宴》 《喜宴》->赵文瑄
}

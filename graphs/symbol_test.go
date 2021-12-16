package graphs

import (
	"fmt"
	"testing"
)

func TestSymbol(t *testing.T) {
	g, err := LoadGraph("./test_data/symbol_graph.yml")
	if err != nil {
		t.Fatal(err)
	}
	bfs := g.BFS(g.VetOf("姜文"))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁朝伟")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("宋慧乔")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("郎雄")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("周星驰")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁家辉")).Str(g.Symbol))

	//Output:
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张学友, 张学友->《东邪西毒》, 《东邪西毒》->梁朝伟,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->章子怡, 章子怡->《一代宗师》, 《一代宗师》->宋慧乔,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->郎雄,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张曼玉, 张曼玉->《家有喜事》, 《家有喜事》->周星驰,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《赌神2》, 《赌神2》->梁家辉,
}

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
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁朝伟")).Symbol(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("宋慧乔")).Symbol(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("郎雄")).Symbol(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("周星驰")).Symbol(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁家辉")).Symbol(g.Symbol))
}

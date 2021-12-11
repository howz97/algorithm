package graphs

import (
	"github.com/howz97/algorithm/util"
	"gopkg.in/yaml.v2"
	"strings"
)

type Symbol struct {
	str2int map[string]int
	int2str []string
}

func NewSymbolGraph() *Symbol {
	return &Symbol{
		str2int: make(map[string]int),
		int2str: nil,
	}
}

func (sg *Symbol) AddEdge(src, dst string) error {
	iSrc, ok := sg.str2int[src]
	if !ok {
		return ErrVerticalNotExist
	}
	iDst, ok := sg.str2int[dst]
	if !ok {
		return ErrVerticalNotExist
	}
	return sg.IGraph.AddEdge(iSrc, iDst)
}

func (sg *Symbol) HasEdge(src, dst string) bool {
	iSrc, ok := sg.str2int[src]
	if !ok {
		return false
	}
	iDst, ok := sg.str2int[dst]
	if !ok {
		return false
	}
	return sg.IGraph.HasEdge(iSrc, iDst)
}

func (sg *Symbol) HasVertical(v string) bool {
	_, ok := sg.str2int[v]
	return ok
}

func (sg *Symbol) Adjacent(v string) []string {
	var adj []string
	sg.RangeAdj(v, func(v string) bool {
		adj = append(adj, v)
		return true
	})
	return adj
}

func (sg *Symbol) RangeAdj(src string, fn func(string) bool) {
	iSrc, ok := sg.str2int[src]
	if !ok {
		return
	}
	sg.IGraph.IterateAdj(iSrc, func(adj int) bool {
		return fn(sg.int2str[adj])
	})
}

func (sg *Symbol) scanVertical(v string) {
	if _, ok := sg.str2int[v]; !ok {
		sg.str2int[v] = len(sg.int2str)
		sg.int2str = append(sg.int2str, v)
	}
}

func (sg *Symbol) Marshal() ([]byte, error) {
	m := make(map[string]map[string]float64)
	for v := 0; v < int(sg.NumVertical()); v++ {
		edges := make(map[string]float64)
		sg.IterateWAdj(v, func(a int, w float64) bool {
			edges[sg.int2str[a]] = w
			return true
		})
		m[sg.int2str[v]] = edges
	}
	return yaml.Marshal(m)
}

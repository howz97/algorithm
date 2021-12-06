package symbol_graph

import (
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/util"
	"strings"
)

type SymbolGraph struct {
	str2int map[string]int
	int2str []string
	Graph   graphs.IGraph
}

func NewSymbolGraph(filename string, scanFn func(string) ([][]string, error), newFn func(int) graphs.IGraph) (*SymbolGraph, error) {
	input, err := scanFn(filename)
	if err != nil {
		return nil, err
	}
	sg := &SymbolGraph{
		str2int: make(map[string]int),
		int2str: make([]string, 0),
	}
	for _, row := range input {
		for _, v := range row {
			sg.scanVertical(v)
		}
	}
	sg.Graph = newFn(len(sg.int2str))
	for _, row := range input {
		if len(row) == 0 {
			continue
		}
		src := row[0]
		for _, dst := range row[1:] {
			sg.AddEdge(src, dst)
		}
	}
	return sg, nil
}

func (sg *SymbolGraph) AddEdge(src, dst string) error {
	iSrc, ok := sg.str2int[src]
	if !ok {
		return graphs.ErrVerticalNotExist
	}
	iDst, ok := sg.str2int[dst]
	if !ok {
		return graphs.ErrVerticalNotExist
	}
	return sg.Graph.AddEdge(iSrc, iDst)
}

func (sg *SymbolGraph) HasEdge(src, dst string) bool {
	iSrc, ok := sg.str2int[src]
	if !ok {
		return false
	}
	iDst, ok := sg.str2int[dst]
	if !ok {
		return false
	}
	return sg.Graph.HasEdge(iSrc, iDst)
}

func (sg *SymbolGraph) HasVertical(v string) bool {
	_, ok := sg.str2int[v]
	return ok
}

func (sg *SymbolGraph) Adjacent(v string) []string {
	var adj []string
	sg.RangeAdj(v, func(v string) bool {
		adj = append(adj, v)
		return true
	})
	return adj
}

func (sg *SymbolGraph) RangeAdj(src string, fn func(string) bool) {
	iSrc, ok := sg.str2int[src]
	if !ok {
		return
	}
	sg.Graph.IterateAdj(iSrc, func(adj int) bool {
		return fn(sg.int2str[adj])
	})
}

func (sg *SymbolGraph) scanVertical(v string) {
	if _, ok := sg.str2int[v]; !ok {
		sg.str2int[v] = len(sg.int2str)
		sg.int2str = append(sg.int2str, v)
	}
}

func ScanInput(filename string) (input [][]string, err error) {
	util.RangeFileLines(filename, func(line string) bool {
		line = strings.TrimSpace(line)
		if line == "" {
			return true
		}
		split := strings.Split(line, ":")
		if len(split) != 2 {
			err = graphs.ErrInputFormat
			return false
		}
		src := strings.TrimSpace(split[0])
		if src == "" {
			err = graphs.ErrInputFormat
			return false
		}
		dsts := strings.Split(split[1], ",")
		if len(dsts) == 0 {
			err = graphs.ErrInputFormat
			return false
		}
		row := []string{src}
		for _, dst := range dsts {
			dst = strings.TrimSpace(dst)
			if dst == "" {
				continue
			}
			row = append(row, dst)
		}
		input = append(input, row)
		return true
	})
	return
}

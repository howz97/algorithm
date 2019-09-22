package graph

import (
	"os"
	"strings"
)

type SymbolGraph struct {
	st  map[string]int
	rst []string
	g   Graph
}

func NewSymbolGraph(filename string) (*SymbolGraph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	dataStr := string(data)
	dataSpltd := strings.Split(dataStr, "\n")
	sg := &SymbolGraph{
		st:  make(map[string]int),
		rst: make([]string, 0),
	}
	for i := range dataSpltd {
		dataSpltd[i] = strings.Trim(dataSpltd[i], " ")
		s := dataSpltd[i]
		if s[len(s)-1] == ':' {
			s = s[:len(s)-1]
		}
		if _, exst := sg.st[s]; !exst {
			sg.rst = append(sg.rst, s)
			sg.st[s] = len(sg.st)
		}
	}
	sg.g = NewGraph(len(sg.st))
	v := 0
	for _, s := range dataSpltd {
		if s[len(s)-1] == ':' {
			v = sg.st[s[:len(s)-1]]
		} else {
			sg.g.AddEdge(v, sg.st[s])
		}
	}
	return sg, nil
}

func (sg *SymbolGraph) Graph() Graph {
	return sg.g
}

func (sg *SymbolGraph) NumV() int {
	return sg.g.NumV()
}

func (sg *SymbolGraph) NumEdge() int {
	return sg.g.NumEdge()
}

func (sg *SymbolGraph) AddEdge(v1, v2 string) error {
	return sg.g.AddEdge(sg.index(v1), sg.index(v2))
}

func (sg *SymbolGraph) HasEdge(v1, v2 string) (bool, error) {
	return sg.g.HasEdge(sg.index(v1), sg.index(v2))
}

func (sg *SymbolGraph) Adjacent(v string) ([]string, error) {
	adjInts, err := sg.g.Adjacent(sg.index(v))
	if err != nil {
		return nil, err
	}
	adj := make([]string, len(adjInts))
	for i := range adj {
		adj[i] = sg.name(adjInts[i])
	}
	return adj, nil
}

func (sg *SymbolGraph) name(i int) string {
	if i < 0 || i >= sg.NumV() {
		panic(errVerticalNotExist)
	}
	return sg.rst[i]
}

func (sg *SymbolGraph) index(n string) int {
	i, exst := sg.st[n]
	if !exst {
		return -1
	}
	return i
}

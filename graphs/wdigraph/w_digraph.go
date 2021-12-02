package wdigraph

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/queue"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
	"os"
	"strconv"
	"strings"
)

var (
	ErrVerticalNotExist = errors.New("vertical not exist")
)

// WDigraph is edge weighted digraph without self loop
type WDigraph struct {
	digraph.Digraph
	weight []*hash_map.Chaining
}

func NewWDigraph(size int) *WDigraph {
	weight := make([]*hash_map.Chaining, size)
	for i := range weight {
		weight[i] = hash_map.New()
	}
	return &WDigraph{
		Digraph: digraph.New(size),
		weight:  weight,
	}
}

//func ImportEWD(filename string) (WDigraph, error) {
//	f, err := os.Open(filename)
//	if err != nil {
//		return nil, err
//	}
//	stat, err := f.Stat()
//	if err != nil {
//		return nil, err
//	}
//	data := make([]byte, stat.Size())
//	_, err = f.Read(data)
//	if err != nil {
//		return nil, err
//	}
//	r := newLineReader(data)
//	headLine, eof := r.nextLine()
//	if eof {
//		panic(fmt.Sprintf("empty file: %v", filename))
//	}
//	headSli := strings.Split(string(headLine), " ")
//	numV, err := strconv.Atoi(headSli[0])
//	if err != nil {
//		return nil, err
//	}
//	ewd := NewWDigraph(numV)
//	for l, eof := r.nextLine(); !eof; l, eof = r.nextLine() {
//		lineSli := strings.Split(string(l), " ")
//		from, err := strconv.Atoi(lineSli[0])
//		if err != nil {
//			return nil, err
//		}
//		to, err := strconv.Atoi(lineSli[1])
//		if err != nil {
//			return nil, err
//		}
//		weight, err := strconv.ParseFloat(lineSli[2], 64)
//		if err != nil {
//			return nil, err
//		}
//		ewd.AddEdge(&Edge{
//			from:   from,
//			to:     to,
//			weight: weight,
//		})
//	}
//	return ewd, nil
//}
//
//type lineReader struct {
//	data []byte
//}
//
//func newLineReader(data []byte) *lineReader {
//	return &lineReader{
//		data: data,
//	}
//}
//
//func (r *lineReader) nextLine() (data []byte, eof bool) {
//	if len(r.data) == 0 {
//		return nil, true
//	}
//	i := bytes.Index(r.data, []byte{'\n'})
//	if i < 0 {
//		data = r.data[:]
//		r.data = nil
//		return data, false
//	}
//	data = r.data[:i]
//	r.data = r.data[i+1:]
//	return data, false
//}

func (g WDigraph) AddEdge(src, dst int, w float64) error {
	err := g.Digraph.AddEdge(src, dst)
	if err != nil {
		return err
	}
	g.weight[src].Put(search.Integer(dst), w)
	return nil
}

// todo: deprecated
func (g WDigraph) Adjacent(v int) []*Edge {
	if !g.HasVertical(v) {
		return nil
	}
	return g[v].traverse()
}

func (g WDigraph) HasNegativeEdge() bool {
	for i := range g {
		adj := g.Adjacent(i)
		for _, e := range adj {
			if e.weight < 0 {
				return true
			}
		}
	}
	return false
}

type Edge struct {
	from, to int
	weight   float64
}

func (e *Edge) From() int {
	return e.from
}

func (e *Edge) To() int {
	return e.to
}

func (e *Edge) GetWeight() float64 {
	return e.weight
}

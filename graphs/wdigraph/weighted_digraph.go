package wdigraph

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/howz97/algorithm/queue"
	"os"
	"strconv"
	"strings"
)

var (
	ErrVerticalNotExist   = errors.New("vertical not exist")
	ErrNotSupportSelfLoop = errors.New("not support self loop")
)

// 不支持自环
type EdgeWeightedDigraph []edgeSet

func NewEWD(numV int) EdgeWeightedDigraph {
	g := make(EdgeWeightedDigraph, numV)
	for i := range g {
		g[i] = NewEdgeSet()
	}
	return g
}

func ImportEWD(filename string) (EdgeWeightedDigraph, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stat.Size())
	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}
	r := newLineReader(data)
	headLine, eof := r.nextLine()
	if eof {
		panic(fmt.Sprintf("empty file: %v", filename))
	}
	headSli := strings.Split(string(headLine), " ")
	numV, err := strconv.Atoi(headSli[0])
	if err != nil {
		return nil, err
	}
	ewd := NewEWD(numV)
	for l, eof := r.nextLine(); !eof; l, eof = r.nextLine() {
		lineSli := strings.Split(string(l), " ")
		from, err := strconv.Atoi(lineSli[0])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(lineSli[1])
		if err != nil {
			return nil, err
		}
		weight, err := strconv.ParseFloat(lineSli[2], 64)
		if err != nil {
			return nil, err
		}
		ewd.AddEdge(&Edge{
			from:   from,
			to:     to,
			weight: weight,
		})
	}
	return ewd, nil
}

type lineReader struct {
	data []byte
}

func newLineReader(data []byte) *lineReader {
	return &lineReader{
		data: data,
	}
}

func (r *lineReader) nextLine() (data []byte, eof bool) {
	if len(r.data) == 0 {
		return nil, true
	}
	i := bytes.Index(r.data, []byte{'\n'})
	if i < 0 {
		data = r.data[:]
		r.data = nil
		return data, false
	}
	data = r.data[:i]
	r.data = r.data[i+1:]
	return data, false
}

func (g EdgeWeightedDigraph) NumV() int {
	return len(g)
}

func (g EdgeWeightedDigraph) NumE() int {
	nume := 0
	for i := range g {
		nume += g[i].len()
	}
	return nume
}

func (g EdgeWeightedDigraph) AddEdge(e *Edge) {
	v := e.from
	w := e.to
	if !g.HasV(v) || !g.HasV(w) {
		panic(ErrVerticalNotExist)
	}
	if v == w {
		panic(ErrNotSupportSelfLoop)
	}
	g[v].add(e)
}

func (g EdgeWeightedDigraph) Adjacent(v int) []*Edge {
	if !g.HasV(v) {
		panic(ErrVerticalNotExist)
	}
	return g[v].traverse()
}

func (g EdgeWeightedDigraph) Edges() *queue.Queen {
	edges := queue.NewQueen(g.NumE())
	for i := range g {
		adj := g.Adjacent(i)
		for _, e := range adj {
			edges.PushBack(e)
		}
	}
	return edges
}

func (g EdgeWeightedDigraph) HasV(v int) bool {
	return v >= 0 && v < g.NumV()
}

func (g EdgeWeightedDigraph) HasNegativeEdge() bool {
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

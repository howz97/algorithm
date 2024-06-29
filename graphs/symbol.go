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
	"os"

	"gopkg.in/yaml.v2"
)

func LoadSymbDigraph(filename string) (*SymbDigraph, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var m map[string][]string
	err = yaml.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	graph := NewDigraph[string](uint(len(m)))
	symbols := Populate(graph, m)
	return &SymbDigraph{
		Digraph: graph,
		symbols: symbols,
	}, nil
}

func LoadSymbGraph(filename string) (*SymbGraph, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var m map[string][]string
	err = yaml.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	graph := NewGraph[string](uint(len(m)))
	symbols := Populate(graph, m)
	return &SymbGraph{
		Graph:   graph,
		symbols: symbols,
	}, nil
}

func LoadSymbWDigraph(filename string) (*SymbWDigraph, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var m map[string]map[string]float64
	err = yaml.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	graph := NewWDigraph[string](uint(len(m)))
	symbols := WPopulate(graph, m)
	return &SymbWDigraph{
		WDigraph: graph,
		symbols:  symbols,
	}, nil
}

func LoadSymbWGraph(filename string) (*SymbWGraph, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var m map[string]map[string]float64
	err = yaml.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	graph := NewWGraph[string](uint(len(m)))
	symbols := WPopulate(graph, m)
	return &SymbWGraph{
		WGraph:  graph,
		symbols: symbols,
	}, nil
}

func NewSymbDigraph(cap uint) *SymbDigraph {
	return &SymbDigraph{
		Digraph: NewDigraph[string](cap),
		symbols: make(map[string]Id, cap),
	}
}

func NewSymbGraph(cap uint) *SymbGraph {
	return &SymbGraph{
		Graph:   NewGraph[string](cap),
		symbols: make(map[string]Id, cap),
	}
}

type SymbDigraph struct {
	*Digraph[string]
	symbols map[string]Id
}

func (sg *SymbDigraph) AddVertex(v string) {
	if _, ok := sg.symbols[v]; !ok {
		sg.symbols[v] = sg.Digraph.AddVertex(v)
	}
}

func (sg *SymbDigraph) AddEdge(src, dst string) error {
	srcId, ok := sg.symbols[src]
	if !ok {
		srcId = sg.Digraph.AddVertex(src)
	}
	dstId, ok := sg.symbols[dst]
	if !ok {
		dstId = sg.Digraph.AddVertex(dst)
	}
	sg.Digraph.AddEdge(srcId, dstId)
	return nil
}

func (sg *SymbDigraph) SymbOf(v Id) string {
	return sg.Digraph.vertices[v]
}

func (sg *SymbDigraph) IdOf(s string) Id {
	return sg.symbols[s]
}

type SymbGraph struct {
	*Graph[string]
	symbols map[string]Id
}

func (sg *SymbGraph) AddVertex(v string) {
	if _, ok := sg.symbols[v]; !ok {
		sg.symbols[v] = sg.Graph.AddVertex(v)
	}
}

func (sg *SymbGraph) AddEdge(src, dst string) error {
	srcId, ok := sg.symbols[src]
	if !ok {
		srcId = sg.Graph.AddVertex(src)
	}
	dstId, ok := sg.symbols[dst]
	if !ok {
		dstId = sg.Graph.AddVertex(dst)
	}
	sg.Graph.AddEdge(srcId, dstId)
	return nil
}

func (sg *SymbGraph) SymbOf(v Id) string {
	return sg.Graph.vertices[v]
}

func (sg *SymbGraph) IdOf(s string) Id {
	return sg.symbols[s]
}

type SymbWGraph struct {
	*WGraph[string]
	symbols map[string]Id
}

func (sg *SymbWGraph) AddVertex(v string) {
	if _, ok := sg.symbols[v]; !ok {
		sg.symbols[v] = sg.WGraph.AddVertex(v)
	}
}

func (sg *SymbWGraph) AddEdge(src, dst string, w float64) error {
	srcId, ok := sg.symbols[src]
	if !ok {
		srcId = sg.WGraph.AddVertex(src)
	}
	dstId, ok := sg.symbols[dst]
	if !ok {
		dstId = sg.WGraph.AddVertex(dst)
	}
	sg.WGraph.AddEdge(srcId, dstId, w)
	return nil
}

func (sg *SymbWGraph) SymbOf(v Id) string {
	return sg.WGraph.vertices[v]
}

func (sg *SymbWGraph) IdOf(s string) Id {
	return sg.symbols[s]
}

type SymbWDigraph struct {
	*WDigraph[string]
	symbols map[string]Id
}

func (sg *SymbWDigraph) AddVertex(v string) {
	if _, ok := sg.symbols[v]; !ok {
		sg.symbols[v] = sg.WDigraph.AddVertex(v)
	}
}

func (sg *SymbWDigraph) AddEdge(src, dst string, w float64) error {
	srcId, ok := sg.symbols[src]
	if !ok {
		srcId = sg.WDigraph.AddVertex(src)
	}
	dstId, ok := sg.symbols[dst]
	if !ok {
		dstId = sg.WDigraph.AddVertex(dst)
	}
	sg.WDigraph.AddEdge(srcId, dstId, w)
	return nil
}

func (sg *SymbWDigraph) SymbOf(v Id) string {
	return sg.WDigraph.vertices[v]
}

func (sg *SymbWDigraph) IdOf(s string) Id {
	return sg.symbols[s]
}

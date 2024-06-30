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

const (
	verticalOverflow = "vertical overflow"
)

type UnionFind struct {
	id    []Id
	numcc int
}

func NewUnionFind(numV int) *UnionFind {
	uf := &UnionFind{
		id:    make([]Id, numV),
		numcc: numV,
	}
	for i := range uf.id {
		uf.id[i] = Id(i)
	}
	return uf
}

func (uf *UnionFind) NumConnectedComponent() int {
	return uf.numcc
}

func (uf *UnionFind) Find(v Id) Id {
	if v < 0 || int(v) >= len(uf.id) {
		panic(verticalOverflow)
	}
	return uf.id[v]
}

func (uf *UnionFind) IsConnected(v1, v2 Id) bool {
	if v1 < 0 || int(v1) >= len(uf.id) || v2 < 0 || int(v2) >= len(uf.id) {
		panic(verticalOverflow)
	}
	return uf.id[v1] == uf.id[v2]
}

func (uf *UnionFind) Union(v1, v2 Id) {
	if v1 < 0 || int(v1) >= len(uf.id) || v2 < 0 || int(v2) >= len(uf.id) {
		panic(verticalOverflow)
	}
	if uf.IsConnected(v1, v2) {
		return
	}
	v2id := uf.id[v2]
	for i := range uf.id {
		if uf.id[i] == v2id {
			uf.id[i] = uf.id[v1]
		}
	}
	uf.numcc--
}

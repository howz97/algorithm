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
	"errors"
	"fmt"
	"math"
)

var (
	ErrInvalidVertex = errors.New("invalid vertex")
	ErrInvalidEdge   = errors.New("invalid edge")
	ErrInvalidYaml   = errors.New("invalid yaml file")
)

const (
	DistanceZero    Weight = 0
	DistanceDefault Weight = 1
	DistanceMax     Weight = math.MaxInt
)

type Weight int

type Id int

func (i Id) Hash() uintptr {
	return uintptr(i)
}

type ErrCycle []Id

func (c ErrCycle) Error() string {
	str := "cycle path exists: "
	for i, v := range c {
		if i > 0 {
			str += " -> "
		}
		str += fmt.Sprint(v)
	}
	return str
}

type edge struct {
	from, to Id
	weight   Weight
}

type IGraph[V any] interface {
	AddVertex(v V) Id
	AddEdge(Id, Id) error
}

type IWGraph[V any] interface {
	AddVertex(v V) Id
	AddEdge(Id, Id, Weight) error
}

func Populate[T comparable](g IGraph[T], m map[T][]T) map[T]Id {
	symbols := make(map[T]Id, len(m))
	for fr, edges := range m {
		for _, to := range edges {
			frId, ok := symbols[fr]
			if !ok {
				frId = g.AddVertex(fr)
				symbols[fr] = frId
			}
			toId, ok := symbols[to]
			if !ok {
				toId = g.AddVertex(to)
				symbols[to] = toId
			}
			g.AddEdge(frId, toId)
		}
	}
	return symbols
}

func WPopulate[T comparable](g IWGraph[T], m map[T]map[T]Weight) map[T]Id {
	symbols := make(map[T]Id, len(m))
	for fr, edges := range m {
		for to, w := range edges {
			frId, ok := symbols[fr]
			if !ok {
				frId = g.AddVertex(fr)
				symbols[fr] = frId
			}
			toId, ok := symbols[to]
			if !ok {
				toId = g.AddVertex(to)
				symbols[to] = toId
			}
			g.AddEdge(frId, toId, w)
		}
	}
	return symbols
}

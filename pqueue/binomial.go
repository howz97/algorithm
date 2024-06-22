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

package pqueue

import "cmp"

const (
	isNil  = 0
	notNil = 1
)

// Binomial is a binomial queue
type Binomial[O cmp.Ordered] struct {
	size  int
	trees []*bNode[O]
}

// NewBinomial return a binomial queue with default capacity
func NewBinomial[O cmp.Ordered]() *Binomial[O] {
	return &Binomial[O]{
		trees: make([]*bNode[O], 8),
	}
}

// Merge bq1 to bq. ErrExceedCap returned when merge would exceed capacity
func (b *Binomial[O]) Merge(b2 *Binomial[O]) {
	if len(b2.trees) > len(b.trees) {
		*b, *b2 = *b2, *b
	}
	b.size += b2.size
	n := len(b.trees)
	var carry *bNode[O]
	for i := 0; i < n; i++ {
		switch carry.isNil()<<2 + b2.isNil(i)<<1 + b.isNil(i) {
		case 2: // 010
			b.trees[i] = b2.trees[i]
		case 3: // 011
			carry = bMerge(b.trees[i], b2.trees[i])
			b.trees[i] = nil
		case 4: // 100
			b.trees[i] = carry
			carry = nil
		case 5: // 101
			carry = bMerge(carry, b.trees[i])
			b.trees[i] = nil
		case 6: // 110
			fallthrough
		case 7: // 111
			carry = bMerge(carry, b2.trees[i])
		default: // 000, 001
		}
	}
	if carry != nil {
		b.trees = append(b.trees, carry)
	}
}

func (b *Binomial[O]) isNil(i int) int {
	if i >= len(b.trees) {
		return isNil
	}
	return b.trees[i].isNil()
}

func (b *Binomial[O]) Push(p O) {
	b.Merge(&Binomial[O]{
		size:  1,
		trees: []*bNode[O]{{p: p}},
	})
}

func (b *Binomial[O]) Pop() O {
	index := 0 // index of node to pop
	for ; index < len(b.trees); index++ {
		if b.trees[index] != nil {
			break
		}
	}
	for i := index + 1; i < len(b.trees); i++ {
		if b.trees[i] != nil && b.trees[i].p < b.trees[index].p {
			index = i
		}
	}
	// remove tree at index
	popNode := b.trees[index]
	b.trees[index] = nil
	b.size -= 1 << uint(index)
	// trees left by popNode become a new binomial
	trees := popNode.son
	b2 := &Binomial[O]{
		trees: make([]*bNode[O], index),
	}
	for i := index - 1; i >= 0; i-- {
		b2.trees[i] = trees
		sibling := trees.sibling
		trees.sibling = nil
		trees = sibling
	}
	b2.size = 1<<uint(index) - 1
	// merge b2 back
	b.Merge(b2)
	return popNode.p
}

func (b *Binomial[O]) Top() O {
	index := 0 // index of node to pop
	for ; index < len(b.trees); index++ {
		if b.trees[index] != nil {
			break
		}
	}
	for i := index + 1; i < len(b.trees); i++ {
		if b.trees[i] != nil && b.trees[i].p < b.trees[index].p {
			index = i
		}
	}
	return b.trees[index].p
}

// Size get the current size of this binomial queue
func (b *Binomial[O]) Size() int {
	return b.size
}

type bNode[O cmp.Ordered] struct {
	p       O
	sibling *bNode[O]
	son     *bNode[O]
}

func (n *bNode[O]) isNil() int {
	if n == nil {
		return isNil
	}
	return notNil
}

// both a and b MUST not be nil
func bMerge[O cmp.Ordered](a, b *bNode[O]) *bNode[O] {
	if a.p > b.p {
		*a, *b = *b, *a
	}
	b.sibling = a.son
	a.son = b
	return a
}

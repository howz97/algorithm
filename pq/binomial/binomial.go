package binomial

import . "github.com/howz97/algorithm/util"

const (
	isNil  = 0
	notNil = 1
)

// Binomial is a binomial queue
type Binomial struct {
	size  int
	trees []*node
}

// New return a binomial queue with default capacity
func New() *Binomial {
	return &Binomial{
		trees: make([]*node, 8),
	}
}

// Merge bq1 to bq. ErrExceedCap returned when merge would exceed capacity
func (b *Binomial) Merge(b2 *Binomial) {
	if len(b2.trees) > len(b.trees) {
		*b, *b2 = *b2, *b
	}
	b.size += b2.size
	n := len(b.trees)
	var carry *node
	for i := 0; i < n; i++ {
		switch carry.isNil()<<2 + b2.isNil(i)<<1 + b.isNil(i) {
		case 2: // 010
			b.trees[i] = b2.trees[i]
		case 3: // 011
			carry = merge(b.trees[i], b2.trees[i])
			b.trees[i] = nil
		case 4: // 100
			b.trees[i] = carry
			carry = nil
		case 5: // 101
			carry = merge(carry, b.trees[i])
			b.trees[i] = nil
		case 6: // 110
			fallthrough
		case 7: // 111
			carry = merge(carry, b2.trees[i])
		default: // 000, 001
		}
	}
	if carry != nil {
		b.trees = append(b.trees, carry)
	}
}

func (b *Binomial) isNil(i int) int {
	if i >= len(b.trees) {
		return isNil
	}
	return b.trees[i].isNil()
}

func (b *Binomial) Push(p Comparable) {
	b.Merge(&Binomial{
		size:  1,
		trees: []*node{{p: p}},
	})
}

func (b *Binomial) Pop() Comparable {
	index := 0 // index of node to pop
	for ; index < len(b.trees); index++ {
		if b.trees[index] != nil {
			break
		}
	}
	for i := index + 1; i < len(b.trees); i++ {
		if b.trees[i] != nil && b.trees[i].Cmp(b.trees[index]) == Less {
			index = i
		}
	}
	// remove tree at index
	popNode := b.trees[index]
	b.trees[index] = nil
	b.size -= 1 << uint(index)
	// trees left by popNode become a new binomial
	trees := popNode.son
	b2 := &Binomial{
		trees: make([]*node, index),
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

// Size get the current size of this binomial queue
func (b *Binomial) Size() int {
	return b.size
}

type node struct {
	p       Comparable
	sibling *node
	son     *node
}

func (n *node) isNil() int {
	if n == nil {
		return isNil
	}
	return notNil
}

func (n *node) Cmp(other *node) Result {
	return n.p.Cmp(other.p)
}

// both a and b MUST not be nil
func merge(a, b *node) *node {
	if a.Cmp(b) == More {
		*a, *b = *b, *a
	}
	b.sibling = a.son
	a.son = b
	return a
}

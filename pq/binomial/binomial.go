package binomial

import (
	"errors"
)

const (
	defaultMaxTrees = 16
)

var (
	// ErrExceedCap returned when merge would exceed capacity
	ErrExceedCap = errors.New("merge would exceed capacity")

	// ErrEmptyBQ means BQ has been empty
	ErrEmptyBQ = errors.New("delete but BQ has been empty")
)

// BQ is a binomial queue
type BQ struct {
	currentSize int
	trees       []*node
}

// New return a binomial queue with default capacity
func New() *BQ {
	return &BQ{
		trees: make([]*node, defaultMaxTrees),
	}
}

// NewWithMaxTrees return a binomial queue that capacity equal to 2^(maxTrees+1)-1
func NewWithMaxTrees(maxTrees int) *BQ {
	return &BQ{
		trees: make([]*node, maxTrees),
	}
}

// Merge bq1 to bq. ErrExceedCap returned when merge would exceed capacity
func (bq *BQ) Merge(bq1 *BQ) error {
	if bq.currentSize+bq1.currentSize > bq.Cap() {
		return ErrExceedCap
	}
	bq.currentSize += bq1.currentSize
	var carry *node
	n := max(bq.MaxTrees(), bq1.MaxTrees())
	for i := 0; i <= n; i++ {
		switch notNil(carry)*4 + bq1.hasTree(i)*2 + bq.hasTree(i) {
		case 2: // 010
			bq.trees[i] = bq1.trees[i]
		case 3: // 011
			carry = merge(bq.trees[i], bq1.trees[i])
			bq.trees[i] = nil
		case 4: // 100
			bq.trees[i] = carry
			carry = nil
		case 5: // 101
			carry = merge(carry, bq.trees[i])
			bq.trees[i] = nil
		case 6: // 110
			fallthrough
		case 7: // 111
			carry = merge(carry, bq1.trees[i])
		default: // 000, 001
		}
	}
	return nil
}

// Insert k into binomial queue. ErrExceedCap returned when BQ has been full
func (bq *BQ) Insert(k int) error {
	bq1 := NewWithMaxTrees(1)
	bq1.trees[0] = &node{
		k: k,
	}
	bq1.currentSize = 1
	return bq.Merge(bq1)
}

// DelMin delete and return the min key. ErrEmptyBQ returned when BQ has been empty
func (bq *BQ) DelMin() (int, error) {
	if bq.currentSize == 0 {
		return 0, ErrEmptyBQ
	}
	min := 1<<63 - 1 // initialed with biggest int64
	minIdx := -1
	for i := 0; i < bq.MaxTrees(); i++ {
		if bq.trees[i] != nil && bq.trees[i].k < min {
			minIdx = i
			min = bq.trees[i].k
		}
	}
	minTree := bq.trees[minIdx]
	deletedTree := minTree.leftSon
	bq.trees[minIdx] = nil
	bq.currentSize -= 1 << uint(minIdx)
	bq1 := NewWithMaxTrees(minIdx)
	bq1.currentSize = 1<<uint(minIdx) - 1
	for i := minIdx - 1; i >= 0; i-- {
		bq1.trees[i] = deletedTree
		deletedTree = deletedTree.nextSibling
		bq1.trees[i].nextSibling = nil
	}
	bq.Merge(bq1)
	return min, nil
}

// IsEmpty tell us weather BQ is empty
func (bq *BQ) IsEmpty() bool {
	return bq.currentSize == 0
}

// Size return the current size of the binomial queue
func (bq *BQ) Size() int {
	return bq.currentSize
}

// Cap return the capacity of the binomail queue
func (bq *BQ) Cap() int {
	maxTrees := uint(len(bq.trees))
	return (1 << maxTrees) - 1
}

// MaxTrees is the upper limit of the number of trees
func (bq *BQ) MaxTrees() int {
	return len(bq.trees)
}

type node struct {
	k           int
	nextSibling *node
	leftSon     *node
}

func merge(r1, r2 *node) *node {
	// both r1 and r2 can not be nil
	if r1.k > r2.k {
		return merge(r2, r1)
	}
	r2.nextSibling = r1.leftSon
	r1.leftSon = r2
	return r1
}

func (bq *BQ) hasTree(i int) int {
	if i >= bq.MaxTrees() || bq.trees[i] == nil {
		return 0
	}
	return 1
}

func notNil(n *node) int {
	if n == nil {
		return 0
	}
	return 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

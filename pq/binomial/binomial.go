package binomial

const (
	defaultMaxTrees = 8

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
		trees: make([]*node, defaultMaxTrees),
	}
}

// Merge bq1 to bq. ErrExceedCap returned when merge would exceed capacity
func (b *Binomial) Merge(other *Binomial) {
	if len(other.trees) > len(b.trees) {
		*b, *other = *other, *b
	}
	b.size += other.size
	n := len(b.trees)
	var carry *node
	for i := 0; i < n; i++ {
		switch carry.isNil()<<2 + other.isNil(i)<<1 + b.isNil(i) {
		case 2: // 010
			b.trees[i] = other.trees[i]
		case 3: // 011
			carry = merge(b.trees[i], other.trees[i])
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
			carry = merge(carry, other.trees[i])
		default: // 000, 001
		}
	}
	if carry != nil {
		b.trees = append(b.trees, carry)
	}
}

// Insert k into binomial queue. ErrExceedCap returned when Binomial has been full
func (b *Binomial) Insert(k int) {
	bq1 := New()
	bq1.trees[0] = &node{
		k: k,
	}
	bq1.size = 1
	b.Merge(bq1)
}

// DelMin delete and return the min key. ErrEmptyBQ returned when Binomial has been empty
func (b *Binomial) DelMin() (int, error) {
	min := 1<<63 - 1 // initialed with biggest int64
	minIdx := -1
	for i := 0; i < len(b.trees); i++ {
		if b.trees[i] != nil && b.trees[i].k < min {
			minIdx = i
			min = b.trees[i].k
		}
	}
	minTree := b.trees[minIdx]
	deletedTree := minTree.leftSon
	b.trees[minIdx] = nil
	b.size -= 1 << uint(minIdx)
	bq1 := New()
	bq1.size = 1<<uint(minIdx) - 1
	for i := minIdx - 1; i >= 0; i-- {
		bq1.trees[i] = deletedTree
		deletedTree = deletedTree.nextSibling
		bq1.trees[i].nextSibling = nil
	}
	b.Merge(bq1)
	return min, nil
}

// Size return the current size of the binomial queue
func (b *Binomial) Size() int {
	return b.size
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

func (b *Binomial) isNil(i int) int {
	if i >= len(b.trees) {
		return isNil
	}
	return b.trees[i].isNil()
}

func (n *node) isNil() int {
	if n == nil {
		return isNil
	}
	return notNil
}

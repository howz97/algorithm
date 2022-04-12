package hashmap

import (
	"fmt"
)

const (
	maxLoadFactor = 8
	minLoadFactor = 1

	MinSizeTbl = 1
)

// ComparableHasher is a type constraint that matches all
// comparable types with a Hash method.
type ComparableHasher interface {
	comparable
	Hash() uintptr
}

type Chaining[CH ComparableHasher, T any] struct {
	Num uint
	tbl table[CH, T]
}

func New[CH ComparableHasher, T any](args ...uint) *Chaining[CH, T] {
	if len(args) > 1 {
		panic("too many arguments")
	}
	var size uint
	if len(args) == 1 {
		size = args[0]
	}
	if size < MinSizeTbl {
		size = MinSizeTbl
	}
	return &Chaining[CH, T]{
		tbl: make(table[CH, T], size),
	}
}

func (c *Chaining[CH, T]) Get(k CH) (T, bool) {
	return c.tbl.get(k)
}

func (c *Chaining[CH, T]) Put(k CH, v T) {
	exist := c.tbl.put(k, v)
	if !exist {
		c.Num++
		if c.loadFactor() >= maxLoadFactor {
			c.expand()
		}
	}
}

func (c *Chaining[CH, T]) Del(k CH) {
	exist := c.tbl.del(k)
	if exist {
		c.Num--
		if c.loadFactor() < minLoadFactor {
			c.shrink()
		}
	}
}

func (c *Chaining[CH, T]) Clean() {
	c.Num = 0
	c.tbl = make(table[CH, T], MinSizeTbl)
}

func (c *Chaining[CH, T]) numBuckets() uint {
	return c.tbl.size()
}

func (c *Chaining[CH, T]) Size() uint {
	return c.Num
}

func (c *Chaining[CH, T]) loadFactor() uint {
	return c.Size() / c.numBuckets()
}

func (c *Chaining[CH, T]) expand() {
	newTbl := make([]bucket[CH, T], c.numBuckets()*2)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *Chaining[CH, T]) shrink() {
	size := c.numBuckets() >> 1
	if size < MinSizeTbl {
		return
	}
	newTbl := make([]bucket[CH, T], size)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *Chaining[CH, T]) Range(fn func(key CH, val T) bool) {
	c.tbl.Range(fn)
}

func (c *Chaining[CH, T]) String() string {
	str := fmt.Sprintf("size=%d\n", c.Size())
	for i, b := range c.tbl {
		str += fmt.Sprintf("bucket%d: ", i)
		n := b.head
		for n != nil {
			str += fmt.Sprintf("(%v:%v) -> ", n.k, n.v)
			n = n.next
		}
		str += "nil\n"
	}
	return str
}

func tblMove[CH ComparableHasher, T any](src table[CH, T], dst table[CH, T]) {
	for _, b := range src {
		n := b.head
		for n != nil {
			dst.put(n.k, n.v)
			n = n.next
		}
	}
}

type table[CH ComparableHasher, T any] []bucket[CH, T]

func (t table[CH, T]) put(k CH, v T) bool {
	return t[uint(k.Hash())%t.size()].put(k, v)
}

func (t table[CH, T]) get(k CH) (T, bool) {
	return t[uint(k.Hash())%t.size()].get(k)
}

func (t table[CH, T]) del(k CH) bool {
	return t[uint(k.Hash())%t.size()].del(k)
}

func (t table[CH, T]) size() uint {
	return uint(len(t))
}

func (t table[Cmp, T]) Range(fn func(key Cmp, val T) bool) {
	for _, bkt := range t {
		nd := bkt.head
		for nd != nil {
			if goOn := fn(nd.k, nd.v); !goOn {
				return
			}
			nd = nd.next
		}
	}
}

type node[Cmp comparable, T any] struct {
	k    Cmp
	v    T
	next *node[Cmp, T]
}

type bucket[Cmp comparable, T any] struct {
	head *node[Cmp, T]
}

func (b *bucket[Cmp, T]) put(k Cmp, v T) bool {
	if b.head == nil {
		b.head = &node[Cmp, T]{
			k: k,
			v: v,
		}
		return false
	}
	if b.head.k == k {
		b.head.v = v
		return true
	}
	pre := b.head
	n := b.head.next
	for n != nil {
		if n.k == k {
			n.v = v
			return true
		}
		pre = n
		n = n.next
	}
	pre.next = &node[Cmp, T]{
		k: k,
		v: v,
	}
	return false
}

func (b *bucket[Cmp, T]) get(k Cmp) (T, bool) {
	n := b.head
	for n != nil {
		if n.k == k {
			return n.v, true
		}
		n = n.next
	}
	var v T
	return v, false
}

func (b *bucket[Cmp, T]) del(k Cmp) bool {
	if b.head == nil {
		return false
	}
	if b.head.k == k {
		b.head = b.head.next
		return true
	}
	pre := b.head
	n := b.head.next
	for n != nil {
		if n.k == k {
			pre.next = n.next
			return true
		}
		pre = n
		n = n.next
	}
	return false
}

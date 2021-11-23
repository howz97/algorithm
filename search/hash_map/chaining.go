package hash_map

import (
	"fmt"
	"github.com/howz97/algorithm/search"
)

const (
	maxLoadFactor = 8
	minLoadFactor = 1

	MinSizeTbl = 1
)

type Key interface {
	Hash() uint
	search.Cmp
}

type Chaining struct {
	Num uint
	tbl table
}

func New(args ...uint) *Chaining {
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
	return &Chaining{
		tbl: make(table, size),
	}
}

func (c *Chaining) Get(k Key) search.T {
	return c.tbl.get(k)
}

func (c *Chaining) Put(k Key, v search.T) {
	exist := c.tbl.put(k, v)
	if !exist {
		c.Num++
		if c.loadFactor() >= maxLoadFactor {
			c.expand()
		}
	}
}

func (c *Chaining) Del(k Key) {
	exist := c.tbl.del(k)
	if exist {
		c.Num--
		if c.loadFactor() < minLoadFactor {
			c.shrink()
		}
	}
}

func (c *Chaining) Clean() {
	c.Num = 0
	c.tbl = make(table, MinSizeTbl)
}

func (c *Chaining) numBuckets() uint {
	return c.tbl.size()
}

func (c *Chaining) Size() uint {
	return c.Num
}

func (c *Chaining) loadFactor() uint {
	return c.Size() / c.numBuckets()
}

func (c *Chaining) expand() {
	newTbl := make([]bucket, c.numBuckets()*2)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *Chaining) shrink() {
	size := c.numBuckets() >> 1
	if size < MinSizeTbl {
		return
	}
	newTbl := make([]bucket, size)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *Chaining) Range(fn func(key Key, val search.T) bool) {
	c.tbl.Range(fn)
}

func (c *Chaining) String() string {
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

func tblMove(src table, dst table) {
	for _, b := range src {
		n := b.head
		for n != nil {
			dst.put(n.k, n.v)
			n = n.next
		}
	}
}

type table []bucket

func (t table) put(k Key, v search.T) bool {
	return t[k.Hash()%t.size()].put(k, v)
}

func (t table) get(k Key) search.T {
	return t[k.Hash()%t.size()].get(k)
}

func (t table) del(k Key) bool {
	return t[k.Hash()%t.size()].del(k)
}

func (t table) size() uint {
	return uint(len(t))
}

func (t table) Range(fn func(key Key, val search.T) bool) {
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

type node struct {
	k    Key
	v    search.T
	next *node
}

type bucket struct {
	head *node
}

func (b *bucket) put(k Key, v search.T) bool {
	if b.head == nil {
		b.head = &node{
			k: k,
			v: v,
		}
		return false
	}
	if b.head.k.Cmp(k) == search.Equal {
		b.head.v = v
		return true
	}
	pre := b.head
	n := b.head.next
	for n != nil {
		if n.k.Cmp(k) == search.Equal {
			n.v = v
			return true
		}
		pre = n
		n = n.next
	}
	pre.next = &node{
		k: k,
		v: v,
	}
	return false
}

func (b *bucket) get(k Key) search.T {
	n := b.head
	for n != nil {
		if n.k.Cmp(k) == search.Equal {
			return n.v
		}
		n = n.next
	}
	return nil
}

func (b *bucket) del(k Key) bool {
	if b.head == nil {
		return false
	}
	if b.head.k.Cmp(k) == search.Equal {
		b.head = b.head.next
		return true
	}
	pre := b.head
	n := b.head.next
	for n != nil {
		if n.k.Cmp(k) == search.Equal {
			pre.next = n.next
			return true
		}
		pre = n
		n = n.next
	}
	return false
}

package hash_map

import (
	"github.com/howz97/algorithm/search"
)

const (
	maxLoadFactor = 8
	minLoadFactor = 1
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
	if size == 0 {
		size = 1
	}
	return &Chaining{
		tbl: makeTable(size),
	}
}

func (c *Chaining) Get(k Key) search.T {
	return c.tbl.get(k)
}

func (c *Chaining) Put(k Key, v search.T) {
	if v == nil {
		c.Del(k)
		return
	}
	if c.LoadFactor() >= maxLoadFactor {
		c.expand()
	}
	c.tbl.put(k, v)
	c.Num++
}

func (c *Chaining) Range(fn func(key Key, val search.T) bool) {
	for _, bkt := range c.tbl {
		nd := bkt.head
		for nd != nil {
			if goOn := fn(nd.k, nd.v); !goOn {
				return
			}
			nd = nd.next
		}
	}
}

func (c *Chaining) Del(k Key) {
	if c.tbl.delete(k) {
		c.Num--
		if c.LoadFactor() < minLoadFactor {
			c.shrink()
		}
	}
}

func (c *Chaining) Clean() {
	c.Num = 0
	c.tbl = makeTable(1)
}

func (c *Chaining) TblSize() uint {
	return c.tbl.size()
}

func (c *Chaining) Size() uint {
	return c.Num
}

func (c *Chaining) LoadFactor() uint {
	return c.Size() / c.TblSize()
}

func (c *Chaining) expand() {
	newTbl := make([]bucket, c.TblSize()*2)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *Chaining) shrink() {
	if c.TblSize() == 1 {
		return
	}
	newTbl := make([]bucket, c.TblSize()/2)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
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

func makeTable(size uint) table {
	return make(table, size)
}

func (t table) put(k Key, v search.T) {
	t[k.Hash()%t.size()].put(k, v)
}

func (t table) get(k Key) search.T {
	return t[k.Hash()%t.size()].get(k)
}

func (t table) delete(k Key) bool {
	return t[k.Hash()%t.size()].delete(k)
}

func (t table) size() uint {
	return uint(len(t))
}

type bucket struct {
	head *node
}

func (b *bucket) put(k Key, v search.T) {
	b.head = b.head.put(k, v)
}

func (b *bucket) get(k Key) search.T {
	return b.head.get(k)
}

func (b *bucket) delete(k Key) bool {
	var deleted bool
	b.head, deleted = b.head.delete(k)
	return deleted
}

type node struct {
	k    Key
	v    search.T
	next *node
}

func (n *node) put(k Key, v search.T) *node {
	if n == nil {
		return &node{
			k: k,
			v: v,
		}
	}
	if k.Cmp(n.k) == search.Equal {
		n.v = v
	} else {
		n.next = n.next.put(k, v)
	}
	return n
}

func (n *node) get(k Key) search.T {
	if n == nil {
		return nil
	}
	if k.Cmp(n.k) == search.Equal {
		return n.v
	}
	return n.next.get(k)
}

func (n *node) delete(k Key) (*node, bool) {
	if n == nil {
		return nil, false
	}
	if k.Cmp(n.k) == search.Equal {
		return n.next, true
	}
	var deleted bool
	n.next, deleted = n.next.delete(k)
	return n, deleted
}

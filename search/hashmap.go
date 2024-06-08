package search

import (
	"fmt"
)

const (
	maxLoadFactor = 8
	minLoadFactor = 1
	MinSizeTbl    = 1
)

// CmpHash is a type constraint that matches all
// comparable types with a Hash method.
type CmpHash interface {
	comparable
	Hash() uintptr
}

type HashMap[K CmpHash, V any] struct {
	Num uint
	tbl hTable[K, V]
}

func NewHashMap[K CmpHash, V any](args ...uint) *HashMap[K, V] {
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
	return &HashMap[K, V]{
		tbl: make(hTable[K, V], size),
	}
}

func (c *HashMap[K, V]) Get(k K) (V, bool) {
	return c.tbl.get(k)
}

func (c *HashMap[K, V]) Put(k K, v V) {
	exist := c.tbl.put(k, v)
	if !exist {
		c.Num++
		if c.loadFactor() >= maxLoadFactor {
			c.expand()
		}
	}
}

func (c *HashMap[K, V]) Del(k K) {
	exist := c.tbl.del(k)
	if exist {
		c.Num--
		if c.loadFactor() < minLoadFactor {
			c.shrink()
		}
	}
}

func (c *HashMap[K, V]) Clean() {
	c.Num = 0
	c.tbl = make(hTable[K, V], MinSizeTbl)
}

func (c *HashMap[K, V]) numBuckets() uint {
	return c.tbl.size()
}

func (c *HashMap[K, V]) Size() uint {
	return c.Num
}

func (c *HashMap[K, V]) loadFactor() uint {
	return c.Size() / c.numBuckets()
}

func (c *HashMap[K, V]) expand() {
	newTbl := make([]hBucket[K, V], c.numBuckets()*2)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *HashMap[K, V]) shrink() {
	size := c.numBuckets() >> 1
	if size < MinSizeTbl {
		return
	}
	newTbl := make([]hBucket[K, V], size)
	tblMove(c.tbl, newTbl)
	c.tbl = newTbl
}

func (c *HashMap[K, V]) Range(fn func(key K, val V) bool) {
	c.tbl.Range(fn)
}

func (c *HashMap[K, V]) String() string {
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

func tblMove[K CmpHash, V any](src hTable[K, V], dst hTable[K, V]) {
	for _, b := range src {
		n := b.head
		for n != nil {
			dst.put(n.k, n.v)
			n = n.next
		}
	}
}

type hTable[K CmpHash, V any] []hBucket[K, V]

func (t hTable[K, V]) put(k K, v V) bool {
	return t[uint(k.Hash())%t.size()].put(k, v)
}

func (t hTable[K, V]) get(k K) (V, bool) {
	return t[uint(k.Hash())%t.size()].get(k)
}

func (t hTable[K, V]) del(k K) bool {
	return t[uint(k.Hash())%t.size()].del(k)
}

func (t hTable[K, V]) size() uint {
	return uint(len(t))
}

func (t hTable[Cmp, V]) Range(fn func(key Cmp, val V) bool) {
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

type hNode[Cmp comparable, V any] struct {
	k    Cmp
	v    V
	next *hNode[Cmp, V]
}

type hBucket[Cmp comparable, V any] struct {
	head *hNode[Cmp, V]
}

func (b *hBucket[Cmp, V]) put(k Cmp, v V) bool {
	if b.head == nil {
		b.head = &hNode[Cmp, V]{
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
	pre.next = &hNode[Cmp, V]{
		k: k,
		v: v,
	}
	return false
}

func (b *hBucket[Cmp, V]) get(k Cmp) (V, bool) {
	n := b.head
	for n != nil {
		if n.k == k {
			return n.v, true
		}
		n = n.next
	}
	var v V
	return v, false
}

func (b *hBucket[Cmp, V]) del(k Cmp) bool {
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

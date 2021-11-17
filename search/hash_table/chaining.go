package hash_table

const (
	maxLoadFactor = 8
	minLoadFactor = 1
)

type ChainHT struct {
	kvNum int
	tbl   table
}

func New(args ...uint) *ChainHT {
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
	return &ChainHT{
		tbl: makeTable(size),
	}
}

func (ht *ChainHT) Get(k Key) T {
	return ht.tbl.get(k)
}

func (ht *ChainHT) Put(k Key, v T) {
	if v == nil {
		ht.Delete(k)
		return
	}
	if ht.LoadFactor() >= maxLoadFactor {
		ht.expand()
	}
	ht.tbl.put(k, v)
	ht.kvNum++
}

func (ht *ChainHT) Range(fn func(key Key, val T) bool) {
	for _, bkt := range ht.tbl {
		nd := bkt.head
		for nd != nil {
			if goOn := fn(nd.k, nd.v); !goOn {
				return
			}
			nd = nd.next
		}
	}
}

func (ht *ChainHT) Delete(k Key) {
	if ht.tbl.delete(k) {
		ht.kvNum--
		if ht.LoadFactor() < minLoadFactor {
			ht.shrink()
		}
	}
}

func (ht *ChainHT) TblSize() int {
	return ht.tbl.size()
}

func (ht *ChainHT) Size() int {
	return ht.kvNum
}

func (ht *ChainHT) LoadFactor() int {
	return ht.Size() / ht.TblSize()
}

func (ht *ChainHT) expand() {
	newTbl := make([]bucket, ht.TblSize()*2)
	tblMove(ht.tbl, newTbl)
	ht.tbl = newTbl
}

func (ht *ChainHT) shrink() {
	if ht.TblSize() == 1 {
		return
	}
	newTbl := make([]bucket, ht.TblSize()/2)
	tblMove(ht.tbl, newTbl)
	ht.tbl = newTbl
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

func (t table) put(k Key, v T) {
	t[k.Hash()%t.size()].put(k, v)
}

func (t table) get(k Key) T {
	return t[k.Hash()%t.size()].get(k)
}

func (t table) delete(k Key) bool {
	return t[k.Hash()%t.size()].delete(k)
}

func (t table) size() int {
	return len(t)
}

type bucket struct {
	head *node
}

func (b *bucket) put(k Key, v T) {
	b.head = b.head.put(k, v)
}

func (b *bucket) get(k Key) T {
	return b.head.get(k)
}

func (b *bucket) delete(k Key) bool {
	var deleted bool
	b.head, deleted = b.head.delete(k)
	return deleted
}

type node struct {
	k    Key
	v    T
	next *node
}

func (n *node) put(k Key, v T) *node {
	if n == nil {
		return &node{
			k: k,
			v: v,
		}
	}
	if k.Equal(n.k) {
		n.v = v
	} else {
		n.next = n.next.put(k, v)
	}
	return n
}

func (n *node) get(k Key) T {
	if n == nil {
		return nil
	}
	if k.Equal(n.k) {
		return n.v
	}
	return n.next.get(k)
}

func (n *node) delete(k Key) (*node, bool) {
	if n == nil {
		return nil, false
	}
	if k.Equal(n.k) {
		return n.next, true
	}
	var deleted bool
	n.next, deleted = n.next.delete(k)
	return n, deleted
}

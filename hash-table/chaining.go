package hashtable

const (
	maxBucketAvrgSize = 4
)

type ChainHT struct {
	kvNum   int
	tblSize int
	tbl     []bucket
}

func New() *ChainHT {
	return &ChainHT{
		tblSize: 1,
		tbl:     make([]bucket, 1),
	}
}

func (t *ChainHT) Put(k Key, v interface{}) {
	t.tbl[k.HashCode()%t.tblSize].put(k, v)
	t.kvNum++
	//if t.BucketAvrgSize() >= maxBucketAvrgSize {
	//	for i :=0; i< t.tblSize;i++ {
	//		t.tbl = append(t.tbl, bucket{})
	//	}
	//	t.tblSize *= 2
	//}
}

func (t *ChainHT) Get(k Key) interface{} {

}

func (t *ChainHT) BucketAvrgSize() int {
	return t.kvNum/t.tblSize
}

type bucket struct {
	head *node
}

func (b *bucket) put(k Key, v interface{}) {
	b.head = b.head.put(k, v)
}

func (b *bucket) get(k Key) interface{} {
	return b.head.get(k)
}

type node struct {
	k    Key
	v    interface{}
	next *node
}

func (n *node) put(k Key, v interface{}) *node {
	if n == nil {
		return &node{
			k:    k,
			v:    v,
		}
	}
	if k.Equal(n.k) {
		n.v = v
	}else {
		n.next = n.next.put(k, v)
	}
	return n
}

func (n *node) get(k Key) interface{} {
	if n == nil {
		return nil
	}
	if k.Equal(n.k) {
		return n.v
	}
	return n.next.get(k)
}

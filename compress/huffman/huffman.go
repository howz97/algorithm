package huffman

import (
	"github.com/howz97/algorithm/pq/heap"
	. "github.com/howz97/algorithm/search"
	. "github.com/howz97/algorithm/util"
)

type node struct {
	isLeaf      bool
	b           byte
	cnt         uint
	left, right *node
}

func (n *node) IsNil() bool {
	return n == nil
}

func (n *node) Val() T {
	return nil
}

func (n *node) Left() ITraversal {
	return n.left
}

func (n *node) Right() ITraversal {
	return n.right
}

func (n *node) makeTable(code []bool, table [256][]bool) {
	if n.isLeaf {
		table[n.b] = make([]bool, len(code))
		copy(table[n.b], code)
		return
	}
	n.left.makeTable(append(code, false), table)
	n.right.makeTable(append(code, true), table)
}

func newNode(br *bitReader) (n *node) {
	n = &node{}
	if br.ReadBit() {
		n.isLeaf = true
		n.b = byte(br.ReadBits(8))
	} else {
		n.left = newNode(br)
		n.right = newNode(br)
	}
	return
}

func Compress(data []byte) []byte {
	bw, table := compile(data)
	bw.WriteUint32(uint32(len(data)))
	for _, b := range data {
		bw.WriteBits(table[b])
	}
	bw.Close()
	return bw.output
}

func compile(data []byte) (*bitWriter, [256][]bool) {
	bw := new(bitWriter)
	huffmanTree := genHuffmanTree(data)
	var table [256][]bool
	huffmanTree.makeTable(make([]bool, 0, 256), table)
	// Encode huffman tree
	PreOrder(huffmanTree, func(t ITraversal) bool {
		n := t.(*node)
		if n.isLeaf {
			bw.WriteBit(true)
			bw.WriteByte(n.b)
		} else {
			bw.WriteBit(false)
		}
		return true
	})
	return bw, table
}

func genHuffmanTree(data []byte) (huffmanTree *node) {
	var stat [256]uint
	for _, b := range data {
		stat[b]++
	}
	pq := heap.New(256)
	for b, cnt := range stat {
		if cnt > 0 {
			pq.Push(Int(cnt), &node{
				isLeaf: true,
				b:      byte(b),
				cnt:    cnt,
			})
		}
	}
	for pq.Size() > 1 {
		n1 := pq.Pop().(*node)
		n2 := pq.Pop().(*node)
		cnt := n1.cnt + n2.cnt
		pq.Push(Int(cnt), &node{
			isLeaf: false,
			cnt:    cnt,
			left:   n1,
			right:  n2,
		})
	}
	return pq.Pop().(*node)
}

func Decompress(data []byte) []byte {
	br := newBitReader(data)
	huffmanTree := newNode(br)
	size := br.ReadBits(32)
	output := make([]byte, 0, size)
	for len(output) < size {
		nd := huffmanTree
		for !nd.isLeaf {
			if br.ReadBit() {
				nd = nd.right
			} else {
				nd = nd.left
			}
		}
		output = append(output, nd.b)
	}
	return output
}

package huffman

import (
	"strconv"

	"github.com/howz97/algorithm/pq"
	"github.com/howz97/algorithm/search"
)

type node struct {
	isLeaf      bool
	b           byte
	cnt         int
	left, right *node
}

func (n *node) IsNil() bool {
	return n == nil
}

func (n *node) String() string {
	str := ""
	if n.isLeaf {
		str = "(" + string([]byte{n.b}) + ")"
	}
	return str + strconv.Itoa(n.cnt)
}

func (n *node) Left() search.ITraversal {
	return n.left
}

func (n *node) Right() search.ITraversal {
	return n.right
}

func (n *node) makeTable(code []bool, table [][]bool) {
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
		// non-leaf node MUST have two kids in a huffman tree
		n.left = newNode(br)
		n.right = newNode(br)
	}
	return
}

// Compress data using huffman algorithm
func Compress(data []byte) []byte {
	bw, table := compile(data)
	bw.WriteUint32(uint32(len(data)))
	for _, b := range data {
		bw.WriteBits(table[b])
	}
	bw.Flush()
	return bw.output
}

func compile(data []byte) (*bitWriter, [256][]bool) {
	bw := new(bitWriter)
	huffmanTree := genHuffmanTree(data)
	var table [256][]bool
	huffmanTree.makeTable(make([]bool, 0, 256), table[:])
	// Encode huffman tree
	search.PreOrder(huffmanTree, func(t search.ITraversal) bool {
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
	var stat [256]int
	for _, b := range data {
		stat[b]++
	}
	pq := pq.NewPaired[int, *node](256)
	for b, cnt := range stat {
		if cnt > 0 {
			pq.PushPair(cnt, &node{
				isLeaf: true,
				b:      byte(b),
				cnt:    cnt,
			})
		}
	}
	for pq.Size() > 1 {
		n1 := pq.Pop()
		n2 := pq.Pop()
		cnt := n1.cnt + n2.cnt
		pq.PushPair(cnt, &node{
			isLeaf: false,
			cnt:    cnt,
			left:   n1,
			right:  n2,
		})
	}
	return pq.Pop()
}

// Decompress data compressed by huffman algorithm
func Decompress(data []byte) ([]byte, error) {
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
		if br.Err() != nil {
			break
		}
		output = append(output, nd.b)
	}
	return output, br.Err()
}

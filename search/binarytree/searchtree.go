package binarytree

import (
	. "github.com/howz97/algorithm/search"
	"time"
)

type BinaryTree struct {
	root *node
}

func New() *BinaryTree {
	return new(BinaryTree)
}

func (st *BinaryTree) Insert(key Cmp, val T) {
	st.root = st.root.insert(key, val)
}

func (st *BinaryTree) Find(key Cmp) T {
	n := st.root.find(key)
	if n == nil {
		return nil
	}
	if n.value == nil {
		st.Delete(n.key) // 删除value为nil的节点
		return nil
	}
	return n.value
}

func (st *BinaryTree) FindMin() T {
	return st.root.findMin().value
}

func (st *BinaryTree) FindMax() T {
	return st.root.findMax().value
}

func (st *BinaryTree) Delete(key Cmp) {
	st.root = st.root.delete(key)
}

type node struct {
	value    T
	key      Cmp
	leftSon  *node
	rightSon *node
}

func (n *node) insert(k Cmp, v T) *node {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n
	}
	switch k.Cmp(n.key) {
	case Less:
		n.leftSon = n.leftSon.insert(k, v)
	case More:
		n.rightSon = n.rightSon.insert(k, v)
	default:
		n.value = v
	}
	return n
}

func (n *node) find(k Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case Less:
		n = n.leftSon.find(k)
	case More:
		n = n.rightSon.find(k)
	}
	return n
}

func (n *node) findMin() *node {
	if n == nil {
		return nil
	}
	for n.leftSon != nil {
		n = n.leftSon
	}
	return n
}

func (n *node) findMax() *node {
	if n == nil {
		return nil
	}
	for n.rightSon != nil {
		n = n.rightSon
	}
	return n
}

func (n *node) delete(k Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case Less:
		n.leftSon = n.leftSon.delete(k)
	case More:
		n.rightSon = n.rightSon.delete(k)
	default:
		if n.leftSon == nil {
			return n.rightSon
		}
		if n.rightSon == nil {
			return n.leftSon
		}
		deleted := n
		// to make it randomly
		if time.Now().UnixNano()&1 == 1 {
			n = n.leftSon.findMax()
			n.leftSon = deleted.leftSon.delete(n.key)
			n.rightSon = deleted.rightSon
		} else {
			n = n.rightSon.findMin()
			n.rightSon = deleted.rightSon.delete(n.key)
			n.leftSon = deleted.leftSon
		}
	}
	return n
}

/* 开始写成了这样... 搞的太复杂了! 要学会使用递归
func DeleteMin(st *Node) *Node {
	var min *Node
	var pOfMin *Node
	if st == nil {
		// do nothing and return nil
	} else {
		min = st
		for min.leftSon != nil {
			pOfMin = min
			min = min.leftSon
		}
		if min.rightSon != nil {
			pOfMin.leftSon = min.rightSon
		} else {
			pOfMin.leftSon = nil
		}
	}
	return min
}

func DeleteMax(st *Node) *Node {
	var max *Node
	var pOfMax *Node
	if st == nil {
		// do nothing and return nil
	} else {
		max = st
		for max.rightSon != nil {
			pOfMax = max
			max = max.rightSon
		}
		if max.leftSon != nil {
			pOfMax.rightSon = max.leftSon
		} else {
			pOfMax.rightSon = nil
		}
	}
	return max
}

func Delete(st *Node, id int) *Node {
	return delete(st, id, nil, false)
}

func delete(st *Node, id int, p *Node, leftOfP bool) *Node {
	var deleted *Node
	if st == nil {
		// do nothing and return nil
	} else if id < st.Elem.ID() {
		deleted = delete(st.leftSon, id, st, true)
	} else if id > st.Elem.ID() {
		deleted = delete(st.rightSon, id, st, false)
	} else { // delete st node
		if st.leftSon == nil && st.rightSon == nil {
			// 要删除的节点是叶子节点
			if p != nil { // 有父节点
				if leftOfP {
					p.leftSon = nil
				} else {
					p.rightSon = nil
				}
				deleted = st
			} else { // 无父节点
				deleted = st
				st.Elem = nil // 因为st是调用者持有的指针的副本，没法修改调用者持有的指针，只好这样做
			}
		} else if st.leftSon != nil {
			// 要删除的节点有左儿子
			if p != nil { // 有父节点
				if leftOfP {
					p.leftSon = st.leftSon
				} else {
					p.rightSon = st.leftSon
				}
				st.leftSon = nil
				deleted = st
			} else { // 无父节点
				deleted = new(Node)
				*deleted = *st
				*st = *st.leftSon
			}
		} else if st.rightSon != nil {
			// 要删除的节点有右儿子
			if p != nil { // 有父节点
				if leftOfP {
					p.leftSon = st.rightSon
				} else {
					p.rightSon = st.rightSon
				}
				st.rightSon = nil
				deleted = st
			} else { // 无父节点
				deleted = new(Node)
				*deleted = *st
				*st = *st.rightSon
			}
		} else {
			// 要删除的节点有左 右儿子
			if p != nil { // 有父节点
				replace := DeleteMin(st.rightSon)
				replace.leftSon = st.leftSon
				replace.rightSon = st.rightSon
				if leftOfP {
					p.leftSon = replace
				} else {
					p.rightSon = replace
				}
				st.leftSon = nil
				st.rightSon = nil
				deleted = st
			} else { // 无父节点
				replace := DeleteMin(st.rightSon)
				replace.leftSon = st.leftSon
				replace.rightSon = st.rightSon
				deleted = new(Node)
				*deleted = *st
				*st = *replace
			}
		}
	}
	return deleted
}
*/

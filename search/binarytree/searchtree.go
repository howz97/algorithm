package binarytree

import (
	"github.com/howz97/algorithm/search"
	"time"
)

type BinaryTree struct {
	*node
}

func New() *BinaryTree {
	return new(BinaryTree)
}

func (st *BinaryTree) Put(key search.Cmp, val search.T) {
	st.node = st.node.insert(key, val)
}

func (st *BinaryTree) Get(key search.Cmp) search.T {
	n := st.node.find(key)
	if n == nil {
		return nil
	}
	if n.value == nil {
		st.Del(n.key) // 删除value为nil的节点
		return nil
	}
	return n.value
}

func (st *BinaryTree) GetMin() search.T {
	return st.node.findMin().value
}

func (st *BinaryTree) GetMax() search.T {
	return st.node.findMax().value
}

func (st *BinaryTree) Del(key search.Cmp) {
	st.node = st.node.delete(key)
}

func (st *BinaryTree) Clean() {
	st.node = nil
}

func (st *BinaryTree) GetITraversal() search.ITraversal {
	return st.node
}

type node struct {
	value search.T
	key   search.Cmp
	left  *node
	right *node
}

func (n *node) insert(k search.Cmp, v search.T) *node {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n
	}
	switch k.Cmp(n.key) {
	case search.Less:
		n.left = n.left.insert(k, v)
	case search.More:
		n.right = n.right.insert(k, v)
	default:
		n.value = v
	}
	return n
}

func (n *node) find(k search.Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case search.Less:
		n = n.left.find(k)
	case search.More:
		n = n.right.find(k)
	}
	return n
}

func (n *node) findMin() *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node) findMax() *node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node) delete(k search.Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case search.Less:
		n.left = n.left.delete(k)
	case search.More:
		n.right = n.right.delete(k)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		deleted := n
		// to make it randomly
		if time.Now().UnixNano()&1 == 1 {
			n = n.left.findMax()
			n.left = deleted.left.delete(n.key)
			n.right = deleted.right
		} else {
			n = n.right.findMin()
			n.right = deleted.right.delete(n.key)
			n.left = deleted.left
		}
	}
	return n
}

func (n *node) Left() search.ITraversal {
	return n.left
}

func (n *node) Right() search.ITraversal {
	return n.right
}

func (n *node) IsNil() bool {
	return n == nil
}

func (n *node) Key() search.Cmp {
	return n.key
}

func (n *node) Val() search.T {
	return n.value
}

/* 开始写成了这样... 搞的太复杂了! 要学会使用递归
func DeleteMin(st *Node) *Node {
	var min *Node
	var pOfMin *Node
	if st == nil {
		// do nothing and return nil
	} else {
		min = st
		for min.left != nil {
			pOfMin = min
			min = min.left
		}
		if min.right != nil {
			pOfMin.left = min.right
		} else {
			pOfMin.left = nil
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
		for max.right != nil {
			pOfMax = max
			max = max.right
		}
		if max.left != nil {
			pOfMax.right = max.left
		} else {
			pOfMax.right = nil
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
		deleted = delete(st.left, id, st, true)
	} else if id > st.Elem.ID() {
		deleted = delete(st.right, id, st, false)
	} else { // delete st node
		if st.left == nil && st.right == nil {
			// 要删除的节点是叶子节点
			if p != nil { // 有父节点
				if leftOfP {
					p.left = nil
				} else {
					p.right = nil
				}
				deleted = st
			} else { // 无父节点
				deleted = st
				st.Elem = nil // 因为st是调用者持有的指针的副本，没法修改调用者持有的指针，只好这样做
			}
		} else if st.left != nil {
			// 要删除的节点有左儿子
			if p != nil { // 有父节点
				if leftOfP {
					p.left = st.left
				} else {
					p.right = st.left
				}
				st.left = nil
				deleted = st
			} else { // 无父节点
				deleted = new(Node)
				*deleted = *st
				*st = *st.left
			}
		} else if st.right != nil {
			// 要删除的节点有右儿子
			if p != nil { // 有父节点
				if leftOfP {
					p.left = st.right
				} else {
					p.right = st.right
				}
				st.right = nil
				deleted = st
			} else { // 无父节点
				deleted = new(Node)
				*deleted = *st
				*st = *st.right
			}
		} else {
			// 要删除的节点有左 右儿子
			if p != nil { // 有父节点
				replace := DeleteMin(st.right)
				replace.left = st.left
				replace.right = st.right
				if leftOfP {
					p.left = replace
				} else {
					p.right = replace
				}
				st.left = nil
				st.right = nil
				deleted = st
			} else { // 无父节点
				replace := DeleteMin(st.right)
				replace.left = st.left
				replace.right = st.right
				deleted = new(Node)
				*deleted = *st
				*st = *replace
			}
		}
	}
	return deleted
}
*/

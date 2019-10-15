package searchtree

import (
	"fmt"
	"time"
)

// SearchTree -
type SearchTree struct {
	root *node
}

// New return an empty dictionary
func New() *SearchTree {
	return new(SearchTree)
}

// Insert insert a k-v pair into this dictionary
// if value is nil, is is a virtual deletion operation
func (st *SearchTree) Insert(key int, value interface{}) {
	st.root = st.root.insert(key, value)
}

// Find -
func (st *SearchTree) Find(key int) interface{} {
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

// FindMin -
func (st *SearchTree) FindMin() interface{} {
	return st.root.findMin().value
}

// FindMax -
func (st *SearchTree) FindMax() interface{} {
	return st.root.findMax().value
}

// Delete -
func (st *SearchTree) Delete(key int) {
	st.root = st.root.delete(key)
}

type node struct {
	value    interface{}
	key      int
	leftSon  *node
	rightSon *node
}

func (n *node) insert(k int, v interface{}) *node {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n
	}
	if k < n.key {
		n.leftSon = n.leftSon.insert(k, v)
		return n
	}
	if k > n.key {
		n.rightSon = n.rightSon.insert(k, v)
		return n
	}
	n.value = v
	return n
}

func (n *node) find(key int) *node {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.leftSon.find(key)
	}
	if key > n.key {
		return n.rightSon.find(key)
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

func (n *node) delete(key int) *node {
	if n == nil {
		fmt.Printf("avltree: delete not existed key(%v)", key)
		return nil
	}
	if key < n.key {
		n.leftSon = n.leftSon.delete(key)
		return n
	}
	if key > n.key {
		n.rightSon = n.rightSon.delete(key)
		return n
	}

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
	return n
}

/* 开始写成了这样... 做的太复杂了
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

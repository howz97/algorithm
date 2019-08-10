package searchtree

type element interface {
	ID() int
}

// Node -
type Node struct {
	Elem     element
	leftSon  *Node
	rightSon *Node
}

// NewSearchTree new an empty search tree
func NewSearchTree(e element) *Node {
	return &Node{
		Elem: e,
	}
}

// Insert insert an element to search tree
func Insert(st *Node, e element) *Node {
	if st == nil {
		st = new(Node)
		st.Elem = e
	} else if e.ID() == st.Elem.ID() {
		st.Elem = e
	} else if e.ID() < st.Elem.ID() {
		st.leftSon = Insert(st.leftSon, e)
	} else {
		st.rightSon = Insert(st.rightSon, e)
	}
	return st
}

// Find find the element by id in search tree
func Find(st *Node, id int) *Node {
	var position *Node
	if st == nil {
		// do nothing and return nil
	} else if id == st.Elem.ID() {
		position = st
	} else if id < st.Elem.ID() {
		position = Find(st.leftSon, id)
	} else {
		position = Find(st.rightSon, id)
	}
	return position
}

// FindMin -
func FindMin(st *Node) *Node {
	var min *Node
	if st == nil {
		// do nothing and return nil
	} else {
		min = st
		for min.leftSon != nil {
			min = min.leftSon
		}
	}
	return min
}

// FindMax -
func FindMax(st *Node) *Node {
	var max *Node
	if st == nil {
		// do nothing and return nil
	} else {
		max = st
		for max.rightSon != nil {
			max = max.rightSon
		}
	}
	return max
}

// DeleteMin -
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

// DeleteMax -
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

// Delete -
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

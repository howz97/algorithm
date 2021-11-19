package search

const (
	Equal Result = iota
	Less
	More
)

type Result int

type T interface{}

type Cmp interface {
	Cmp(other Cmp) Result
}

type Searcher interface {
	Insert(key Cmp, value T)
	Find(k Cmp) T
	Delete(key Cmp)
	//Size() uint
}

type IBinaryTree interface {
	Left() IBinaryTree
	Right() IBinaryTree
}

type BinaryNode struct {
	Key   Cmp
	Val   T
	Left  *BinaryNode
	Right *BinaryNode
}

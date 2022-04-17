package redblack

import (
	"fmt"
	"math/rand"
	"testing"
)

func Example() {
	rb := New[int, int]()
	for i := 0; i < 20; i++ {
		rb.Put(i, i)
	}
	v, ok := rb.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v)\n", rb.Size(), v, ok)
	rb.Print()

	for i := 0; i < 10; i++ {
		rb.Del(i)
	}
	v, ok = rb.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v)\n", rb.Size(), v, ok)
	rb.Print()
	fmt.Println("traversal in order:")
	rb.InOrder(func(n *Node[int, int]) bool {
		fmt.Printf("%v,", n.String())
		return true
	})

	// Output (Why not match?):
	// Size=20 Get(5)=(5,true)
	//               (7)
	//               / \
	//              /   \
	//             /     \
	//            /       \
	//           /         \
	//          /           \
	//         /             \
	//      red[3]         red[11]
	//       / \             / \
	//      /   \           /   \
	//     /     \         /     \
	//   (1)     (5)      /       \
	//   / \     / \    (9)      (15)
	// (0) (2) (4) (6)  / \       / \
	//                 /   \     /   \
	//               (8)  (10)  /     \
	//                         /       \
	//                     red[13]   red[17]
	//                       / \       / \
	//                      /   \     /   \
	//                    (12) (14) (16) (18)
	//                                      \
	//                                    red[19]
	// Size=10 Get(5)=(0,false)
	//         (15)
	//          / \
	//         /   \
	//        /     \
	//       /       \
	//     (11)     (17)
	//     / \       / \
	//    /   \     /   \
	//   /     \  (16) (18)
	// (10)  red[13]      \
	//         / \      red[19]
	//        /   \
	//      (12) (14)
	// traversal in order:
	// (10),(11),(12),red[13],(14),(15),(16),(17),(18),red[19],
}

func TestReadBlack(t *testing.T) {
	rb := New[int, int]()
	verify := make(map[int]int)
	BulkDelete(verify, rb, 100)
	BulkInsert(verify, rb, 100)
	BulkDelete(verify, rb, 1000)
	BulkInsert(verify, rb, 500)
	BulkDelete(verify, rb, 100)
	for i := 0; i < 100; i++ {
		BulkInsert(verify, rb, rand.Intn(1000))
		VerifyResult(t, verify, rb)
		BulkDelete(verify, rb, rand.Intn(1000))
		VerifyResult(t, verify, rb)
	}
}

func BulkInsert(verify map[int]int, rb *Tree[int, int], cnt int) {
	for i := 0; i < cnt; i++ {
		k := rand.Int()
		rb.Put(k, k)
		verify[k] = k
	}
}

func BulkDelete(verify map[int]int, rb *Tree[int, int], cnt int) {
	for i := 0; i < cnt; i++ {
		k := rand.Int()
		rb.Del(k)
		delete(verify, k)
	}
}

func VerifyResult(t *testing.T, verify map[int]int, rb *Tree[int, int]) {
	for k, v := range verify {
		vGot, _ := rb.Get(k)
		if vGot != v {
			t.Fatalf("key %v has wrong value %v, should be %v", k, vGot, v)
		}
	}
	if uint(len(verify)) != rb.Size() {
		t.Fatalf("size not equal %d != %d", len(verify), rb.Size())
	}
	CheckValid(t, rb)
}

func CheckValid(t *testing.T, rb *Tree[int, int]) {
	if rb.root.color == red {
		t.Fatalf("root is red")
	}
	cnt := uint(0)
	blackHeight := -1
	rb.InOrder(func(n *Node[int, int]) bool {
		cnt++
		if n.left == rb.null && n.right == rb.null {
			// this is leaf
			bh := 0
			if n.color == black {
				bh++
			}
			c := n.color
			n = n.p
			for n != rb.null {
				if n.color == red && c == red {
					panic("red next to red")
				}
				c = n.color
				if n.color == black {
					bh++
				}
				n = n.p
			}
			if blackHeight < 0 {
				blackHeight = bh
			} else if bh != blackHeight {
				panic("black height not euqal")
			}
		}
		return true
	})
	if cnt != rb.size {
		panic("size not match")
	}
}

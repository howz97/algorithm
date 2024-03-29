ReadBlack search tree
```go
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
```
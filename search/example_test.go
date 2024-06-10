// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"fmt"
	"strconv"

	"github.com/howz97/algorithm/util"
)

func ExampleAVL() {
	toStr := func(nd *AvlNode[int, string]) string { return nd.Value() }
	avl := NewAVL[int, string]()
	for i := 0; i < 20; i++ {
		avl.Put(i, strconv.Itoa(i))
	}
	v, ok := avl.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v) \n", avl.Size(), v, ok)
	PrintTree(avl.Root(), toStr)

	for i := 0; i < 10; i++ {
		avl.Del(i)
	}
	v, ok = avl.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v) \n", avl.Size(), v, ok)
	PrintTree(avl.Root(), toStr)

	fmt.Println("traversal in order:")
	Inorder(avl.Root(), func(t *AvlNode[int, string]) bool {
		fmt.Printf("%v,", t.value)
		return true
	})

	// Size=20 Get(5)=(5,true)
	//            7
	//           / \
	//          /   \
	//         /     \
	//        /       \
	//       /         \
	//      3          15
	//     / \         / \
	//    /   \       /   \
	//   1     5     11   17
	//  / \   / \   / \   / \
	// 0   2 4   6 /   \ 16 18
	//            9    13     \
	//           / \   / \    19
	//          8  10 12 14
	// Size=10 Get(5)=(0,false)
	//     15
	//     / \
	//    /   \
	//   11   17
	//  / \   / \
	// 10 13 16 18
	//    / \     \
	//   12 14    19
	// traversal in order:
	// 10,11,12,13,14,15,16,17,18,19,
}

func ExampleRBTree() {
	toStr := func(nd *RbNode[int, string]) string { return nd.Value() }
	rb := NewRBTree[int, string]()
	for i := 0; i < 20; i++ {
		rb.Put(i, strconv.Itoa(i))
	}
	v, ok := rb.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v)\n", rb.Size(), v, ok)
	PrintTree(rb.Root(), toStr)

	for i := 0; i < 10; i++ {
		rb.Del(i)
	}
	v, ok = rb.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v)\n", rb.Size(), v, ok)
	PrintTree(rb.Root(), toStr)
	fmt.Println("traversal in order:")
	Inorder(rb.Root(), func(n *RbNode[int, string]) bool {
		fmt.Printf("%v,", n.value)
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

func ExampleHashMap() {
	hm := NewHashMap[util.Str, string]()
	hm.Put(util.Str("a"), "A")
	hm.Put(util.Str("b"), "B")
	hm.Put(util.Str("c"), "C")
	hm.Put(util.Str("d"), "D")
	hm.Put(util.Str("e"), "E")
	hm.Put(util.Str("f"), "F")
	hm.Put(util.Str("g"), "G")
	hm.Put(util.Str("h"), "H")
	hm.Put(util.Str("i"), "I")
	fmt.Println(hm.String())

	fmt.Println("delete (d/f/g/x) ...")
	hm.Del(util.Str("d"))
	hm.Del(util.Str("f"))
	hm.Del(util.Str("g"))
	hm.Del(util.Str("x"))
	fmt.Println(hm.String())

	fmt.Println("delete all ...")
	hm.Range(func(key util.Str, _ string) bool {
		hm.Del(key)
		return true
	})
	fmt.Println(hm.String())

	// size=9
	// bucket0: (b:B) -> (d:D) -> (f:F) -> (h:H) -> nil
	// bucket1: (a:A) -> (c:C) -> (e:E) -> (g:G) -> (i:I) -> nil

	// delete (d/f/g/x) ...
	// size=6
	// bucket0: (b:B) -> (h:H) -> nil
	// bucket1: (a:A) -> (c:C) -> (e:E) -> (i:I) -> nil

	// delete all ...
	// size=0
	// bucket0: nil
}

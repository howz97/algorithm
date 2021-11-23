AVL树(平衡二叉树)
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
)

func main() {
	avl := avltree.New()
	for i := 0; i < 20; i++ {
		avl.Put(search.Integer(i), i)
	}
	fmt.Printf("Size=%d Get(5)=%v \n", avl.Size(), avl.Get(search.Integer(5)))
	search.PrintBinaryTree(avl)

	for i := 0; i < 10; i++ {
		avl.Del(search.Integer(i))
	}
	fmt.Printf("Size=%d Get(5)=%v \n", avl.Size(), avl.Get(search.Integer(5)))
	search.PrintBinaryTree(avl)

	fmt.Println("traversal in order:")
	search.InOrder(avl, func(t search.ITraversal) bool {
		fmt.Printf("%v,", t.Val())
		return true
	})
}

/*
Size=20 Get(5)=5 
           7             
          / \            
         /   \           
        /     \          
       /       \         
      /         \        
     3          15       
    / \         / \      
   /   \       /   \     
  1     5     11   17    
 / \   / \   / \   / \   
0   2 4   6 /   \ 16 18  
           9    13     \ 
          / \   / \    19
         8  10 12 14     
Size=10 Get(5)=<nil> 
    15       
    / \      
   /   \     
  11   17    
 / \   / \   
10 13 16 18  
   / \     \ 
  12 14    19
traversal in order:
10,11,12,13,14,15,16,17,18,19,%  
*/
```
package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	var bt BinaryTree

	bt.Insert(15)
	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(5)
	bt.Insert(12)
	bt.Insert(17)
	bt.Insert(25)

	fmt.Println("Traversals")
	fmt.Println("In order:")
	bt.TraverseInOrder()
	fmt.Println("Pre order:")
	bt.TraversePreOrder()
	fmt.Println("Post order:")
	bt.TraversePostOrder()
	fmt.Println()

	fmt.Printf("Found: %v", bt.Search(25))
	fmt.Println()

	bt.Delete(15)
	fmt.Println("In order:")
	bt.TraverseInOrder()
	// bt.Delete(25)
	// fmt.Println("2 - In order:")
	// bt.TraverseInOrder()

}

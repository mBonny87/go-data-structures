package tree

import "fmt"

func (bt *BinaryTree) Insert(v int) {
	bt.Root = bt.insertRecursive(v, bt.Root)
}

func (bt *BinaryTree) insertRecursive(v int, n *Node) *Node {
	if n == nil {
		return &Node{Value: v}
	}

	if v < n.Value {
		n.Left = bt.insertRecursive(v, n.Left)
	} else {
		n.Right = bt.insertRecursive(v, n.Right)
	}

	return n
}

func (bt *BinaryTree) TraverseInOrder() {
	if bt.Root == nil {
		fmt.Println("Empty Binary tree")
	}
	bt.traverseInOrderRecursive(bt.Root)
}

func (bt *BinaryTree) traverseInOrderRecursive(n *Node) {
	if n.Left != nil {
		bt.traverseInOrderRecursive(n.Left)
	}
	fmt.Println(n.Value)
	if n.Right != nil {
		bt.traverseInOrderRecursive(n.Right)
	}
}

func (bt *BinaryTree) TraversePreOrder() {
	if bt.Root == nil {
		fmt.Println("Empty Binary tree")
	}
	bt.traversePreOrderRecursive(bt.Root)
}

func (bt *BinaryTree) traversePreOrderRecursive(n *Node) {
	fmt.Println(n.Value)
	if n.Left != nil {
		bt.traversePreOrderRecursive(n.Left)
	}
	if n.Right != nil {
		bt.traversePreOrderRecursive(n.Right)
	}
}

func (bt *BinaryTree) TraversePostOrder() {
	if bt.Root == nil {
		fmt.Println("Empty Binary tree")
	}
	bt.traversePostOrderRecursive(bt.Root)
}

func (bt *BinaryTree) traversePostOrderRecursive(n *Node) {
	if n.Left != nil {
		bt.traversePostOrderRecursive(n.Left)
	}
	if n.Right != nil {
		bt.traversePostOrderRecursive(n.Right)
	}
	fmt.Println(n.Value)
}

func (bt *BinaryTree) Search(v int) bool {
	return bt.searchRecursive(v, bt.Root)
}

func (bt *BinaryTree) searchRecursive(v int, n *Node) bool {
	if n.Value == v {
		return true
	}
	found := false
	if n.Left != nil {
		found = bt.searchRecursive(v, n.Left)
		if found {
			return found
		}
	}
	if n.Right != nil {
		found = bt.searchRecursive(v, n.Right)
		if found {
			return found
		}
	}
	return false
}

func (bt *BinaryTree) Delete(v int) bool {
	if bt.Root == nil {
		return false
	}
	if bt.Root.Value == v {
		var successor *Node
		if bt.Root.Right != nil {
			// find the In Order Successor
			successor = bt.FindInOrderSuccessor(bt.Root)
			if bt.Root.Left != nil {
				successor.Left = bt.Root.Left
				successor.Right = bt.Root.Right
				bt.Root = successor
				return true
			}
		} else {
			successor = bt.FindInOrderSuccessor(bt.Root)
			bt.Root = successor
			return true
		}
	}
	return bt.deleteRecursive(v, bt.Root)
}

func (bt *BinaryTree) deleteRecursive(v int, n *Node) bool {
	if n.Left != nil {
		if n.Left.Value == v {
			// ok, it is a leaf
			if n.Left.Left == nil && n.Left.Right == nil {
				n.Left = nil
				return true
			}
			var successor *Node
			if n.Left.Right != nil {
				successor = bt.FindInOrderSuccessor(n.Left) // if the statement above the latter is true, it's obvious we'll have a successor
				if n.Left.Left != nil {                     // and we set the new pointer of the successor to the left node, if it exists
					successor.Left = n.Left.Left
					successor.Right = n.Left.Right
					n.Left = successor
					return true
				}
			} else {
				successor = bt.FindInOrderSuccessor(n.Left)
				n.Left = successor
				return true
			}
		} else {
			bt.deleteRecursive(v, n.Left)
		}
	}
	if n.Right != nil {
		if n.Right.Value == v {
			// ok, it is a leaf
			if n.Right.Left == nil && n.Right.Right == nil {
				n.Right = nil
				return true
			}
			var successor *Node
			if n.Right.Right != nil {
				// find the In Order Successor
				successor = bt.FindInOrderSuccessor(n.Right) // if the statement above the latter is true, it's obvious we'll have a successor
				if n.Right.Left != nil {                     // and we set the new pointer of the successor to the left node, if it exists
					successor.Left = n.Right.Left
					successor.Right = n.Right.Right
					n.Right = successor
					return true
				}
			} else {
				successor = bt.FindInOrderSuccessor(n.Right)
				n.Right = successor
				return true
			}
		} else {
			bt.deleteRecursive(v, n.Right)
		}
	}

	return false
}

func (bt *BinaryTree) FindInOrderSuccessor(n *Node) *Node {
	var next *Node
	if n.Right != nil {
		//Check if the next one is the leaf
		next = n.Right
		if n.Right.Right == nil && n.Right.Left == nil {
			n.Right = nil
		}
		return bt.FindInOrderSuccessor(next)
	}
	if n.Left != nil {
		next = n.Left
		if n.Left.Right == nil && n.Left.Left == nil {
			n.Left = nil
		}
		return bt.FindInOrderSuccessor(next)
	}

	return n
}

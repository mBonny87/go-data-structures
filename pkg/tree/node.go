package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

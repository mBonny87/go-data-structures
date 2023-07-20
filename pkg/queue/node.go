package queue

type Node struct {
	Value    int
	Priority int
	Next     *Node
}

package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	queue := &[5]int{}
	//Enqueue
	Enqueue(6, queue)
	Enqueue(4, queue)
	Enqueue(5, queue)
	Enqueue(10, queue)
	Enqueue(10, queue)
	//This is skipper
	Enqueue(10, queue)

	// Dequeue
	Dequeue(queue)

	fmt.Printf("Peeked: %v", Peek(queue))
	fmt.Println()
	fmt.Printf("Empty: %v", IsEmpty(queue))
	fmt.Println()
	fmt.Printf("Size: %v", Size(queue))
	fmt.Println()

	fmt.Println("Current queue:")
	Print(queue)

	//Clear

	fmt.Println("Clearing...")
	Clear(queue)
	fmt.Printf("Peeked: %v", Peek(queue))
	fmt.Println()
	fmt.Printf("Empty: %v", IsEmpty(queue))
	fmt.Println()
	fmt.Printf("Size: %v", Size(queue))
	fmt.Println()

	fmt.Println("Current queue:")
	Print(queue)
}

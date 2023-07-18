package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}

	s.Push(5)
	s.Push(10)
	s.Push(15)
	s.Push(20)

	s.Pop()
	s.Pop()
	fmt.Printf("peeked %v", s.Peek())
	fmt.Println()

	if s.Size() != 2 {
		fmt.Printf("Expected 2 but found %v", s.Peek())
	}

	// s.Pop()
	// s.Pop()

	// if s.Size() != 0 {
	// 	fmt.Printf("Expected 0 but found %v", s.Peek())
	// }

	s.Print()
	fmt.Println("Clearing the Stack...")
	s.Clear()
	fmt.Println("Printing to check if empty...")
	s.Print()
}

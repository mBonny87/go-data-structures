package stack

import "fmt"

type Stack struct {
	Top *Node
}

func (s *Stack) Push(v int) {
	if s.Top == nil {
		s.Top = &Node{Value: v, Next: nil}
		return
	}
	curr := s.Top
	s.Top = &Node{Value: v, Next: curr}
}

func (s *Stack) Pop() {
	if s.Top == nil {
		// already empty
		return
	}
	if s.Top.Next == nil {
		//one element
		s.Top = nil
		return
	}
	s.Top = s.Top.Next
}

func (s *Stack) Print() {
	if s.Top == nil {
		return
	}
	curr := s.Top
	for curr.Next != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	//print last value
	fmt.Println(curr.Value)
}

func (s *Stack) Peek() int {
	if s.Top == nil {
		return 0
	}
	return s.Top.Value
}

func (s *Stack) Size() int {
	if s.Top == nil {
		return 0
	}
	i := 1
	curr := s.Top
	for curr.Next != nil {
		i = i + 1
		curr = curr.Next
	}
	return i
}

func (s *Stack) Clear() {
	s.Top = nil
}

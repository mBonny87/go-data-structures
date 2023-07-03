package main

import (
	"bonnydev87/datastructures/pkg/linkedlist"
)

// func main() {
// 	//1 Declare and initialize the LinkedList struct
// 	ll := linkedlist.SinglyLinkedList{}

// 	//2 Start the timer
// 	start := time.Now()

// 	//3 Add some nodes to the linked list
// 	ll.Insert(5)
// 	ll.Insert(15)
// 	ll.Insert(25)
// 	ll.Insert(35)

// 	//4 Find a value
// 	valueToFind := 45
// 	found := ll.Search(valueToFind)

// 	if found {
// 		fmt.Printf("%v exists\n", valueToFind)
// 	} else {
// 		fmt.Printf("%v does not exist\n", valueToFind)
// 	}

// 	// Delete a node with value 15
// 	ll.Delete(15)

// 	//Print some informations
// 	ll.Print()

// 	fmt.Println("Length:", ll.Length())

// 	elapsed := time.Since(start)

// 	fmt.Println()

// 	fmt.Println("Execution time:", elapsed)

// 	ll.Reverse()

// 	ll.Print()
// }

func main() {
	dll := linkedlist.DoublyLinkedList{}

	dll.InsertAtBeginning(5)
	dll.InsertAtBeginning(10)

	dll.Print()
}

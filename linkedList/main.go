package main

import "fmt"

// Main function
func main() {

	// Appending at the end of the list
	// ll := new(SinglyLinkedList)
	// ll.Append(1)
	// ll.Append(2)
	// ll.Append(3)
	// ll.PrintLL()

	// Inserting elements using InsertAT
	ll := new(SinglyLinkedList)

	ll.InsertAt(1, 0)
	ll.InsertAt(2, 0)
	ll.InsertAt(3, 1)
	ll.Append(4)
	ll.InsertAt(5, 3)
	fmt.Println(ll.Size())
	fmt.Println(ll.ValueAt(3))
	ll.PrintSinglyLL()
}

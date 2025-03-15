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
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.PrintSinglyLL()
	fmt.Println("--------------------")
	ll.Reverse()
	ll.PrintSinglyLL()
}

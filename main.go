package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(data int) {
	// Prepare a node with the data
	temp := Node{
		Data: data,
		Next: nil,
	}

	// Check if the list is empty
	if ll.Head == nil {
		ll.Head = &temp
		return
	}

	// If we have nodes in the list we have to iterate to the last one
	curr := ll.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &temp

}

func (ll *LinkedList) PrintLL() {
	temp := ll.Head
	if temp == nil {
		fmt.Println("The list is empty")
	}
	for temp != nil {
		fmt.Println("Data: ", temp.Data)
		temp = temp.Next
	}
}

// Main function
func main() {
	ll := new(LinkedList)
	ll.Add(1)
	// ll.Add(2)
	// ll.Add(3)
	ll.PrintLL()
}

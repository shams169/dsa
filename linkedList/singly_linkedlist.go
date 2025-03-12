package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *SinglyLinkedList) Append(data int) {
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
	ll.Tail = curr.Next
}

func (ll *SinglyLinkedList) InsertAt(data, index int) error {

	temp := &Node{
		Data: data,
		Next: nil,
	}

	// Check if the index is greater than or equal to 0
	if index < 0 {
		return fmt.Errorf("index %d is invalid", index)
	}

	//Check if the list is empty and index value is greater than 0
	if index != 0 && ll.Head == nil {
		return fmt.Errorf("list is empty cannot do insert at %d ", index)
	}

	if index == 0 && ll.Head == nil {
		ll.Head = temp
		return nil
	}

	if index == 0 && ll.Head != nil {
		temp.Next = ll.Head
		ll.Head = temp
		return nil
	}

	curr := ll.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
		if curr == nil {
			return fmt.Errorf("insufficient number of elements to do insert at index %d", index)
		}
	}

	temp.Next = curr.Next
	curr.Next = temp

	return nil
}

func (ll *SinglyLinkedList) Size() int {
	size := 0
	if ll.Head == nil {
		return size
	}
	temp := ll.Head
	for temp != nil {
		temp = temp.Next
		size++
	}
	return size
}

func (ll *SinglyLinkedList) ValueAt(index int) int {

	//Check if the list is empty
	if ll.Head == nil {
		fmt.Printf("list is empty")
		return 0
	}

	// Check if the index is greater than or equal to 0
	if index < 0 {
		fmt.Printf("index %d is invalid", index)
		return 0
	}

	temp := ll.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
		if temp == nil {
			fmt.Printf("index %d is out of range", index)
			return 0
		}
	}
	return temp.Data

}

func (ll *SinglyLinkedList) PrintSinglyLL() {
	temp := ll.Head
	if temp == nil {
		fmt.Println("The list is empty")
	}
	index := 0
	for temp != nil {
		fmt.Printf("Data at %d: %d\n", index, temp.Data)
		temp = temp.Next
		index++
	}
}

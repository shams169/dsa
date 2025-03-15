package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (dll *DoublyLinkedList) InsertAtHead(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	temp := dll.Head
	temp.Prev = newNode
	newNode.Next = temp
	dll.Head = newNode
}

func (dll *DoublyLinkedList) InsertAtTail(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}

	temp := dll.Tail
	temp.Next = newNode
	newNode.Prev = temp
	dll.Tail = newNode

}

func (dll *DoublyLinkedList) InsertAtIndex(data int, index int) {
	newNode := &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	if index < 0 {
		fmt.Println("invalid index")
		return
	}

	if index == 0 {
		dll.InsertAtHead(data)
		return
	}

	if dll.Head == nil {
		fmt.Println("the list is empty")
		return
	}

	temp := dll.Head

	for i := 0; i < index-1; i++ {
		temp = temp.Next
		if temp == nil {
			fmt.Println("index out of range for the list ")
			return
		}
	}
	newNode.Prev = temp
	newNode.Next = temp.Next
	if temp.Next != nil {
		temp.Next.Prev = newNode
	}

	temp.Next = newNode

}

func (dll *DoublyLinkedList) DeleteAtHead() {
	if dll.Head == nil {
		fmt.Println("the list is empty")
		return
	}

	temp := dll.Head
	dll.Head = temp.Next
	if dll.Head != nil {
		dll.Head.Prev = nil
	}
	temp = nil

}

func (dll *DoublyLinkedList) DeleteAtTail() {
	if dll.Tail == nil {
		fmt.Println("the list is empty")
		return
	}

	temp := dll.Tail
	dll.Tail = temp.Prev
	if dll.Tail != nil {
		dll.Tail.Next = nil
	}
	temp = nil
}

func (dll *DoublyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		fmt.Println("invalid index")
		return
	}

	if index == 0 {
		dll.DeleteAtHead()
		return
	}

	temp := dll.Head

	for i := 0; i < index; i++ {
		temp = temp.Next
		if temp == nil {
			fmt.Println("index out of range for the list ")
			return
		}
	}
	temp2 := temp.Prev
	temp2.Next = temp.Next
	if temp.Next != nil {
		temp.Next.Prev = temp2
	}
	temp = nil

	if dll.Head.Next == nil {
		dll.Tail = dll.Head
	}
}

func (dll *DoublyLinkedList) PrintDLL() {
	if dll.Head == nil {
		fmt.Println("the list is empty")
		return
	}

	temp := dll.Head
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
	fmt.Println("Head: ", dll.Head.Data)
	fmt.Println("Tail: ", dll.Tail.Data)
	fmt.Println("-----------------------")
}

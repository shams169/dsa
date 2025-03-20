package main

import "fmt"

func main() {
	queue := Queue[string]{size: 5}
	fmt.Println("Queue is empty: ", queue.IsEmpty())
	queue.Enqueue("Hello")
	fmt.Println("Queue is not Empty: ", queue.IsEmpty())
	queue.PrintQueue()
}

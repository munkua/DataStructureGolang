package main

import "fmt"

type List struct {

	Head *Node
	Tail *Node

}

type Node struct {

	Element int
	Next *Node

}

func pushNode(value int, listPointer *List) {

	node := &Node{Element : value}
	
	if listPointer.Head == nil {

		listPointer.Head = node
	
	} else {

		listPointer.Tail.Next = node
	
	}
	
	listPointer.Tail = node
	
}

func main() {

	listPointer := &List{}	// listPointer 는 *List 타입
	pushNode(200, listPointer)
	
	fmt.Println(listPointer.Head.Element)
	
	
}
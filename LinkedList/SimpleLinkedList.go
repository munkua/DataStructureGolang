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

func isEmpty(listPointer *List) bool{

	if listPointer.Head == nil {

		fmt.Println("List is empty!")
		return true	

	} else {

	return false
	
	}
	

}

func search(value int, listPointer *List) *Node{

	if isEmpty(listPointer) == false {

		findNode := listPointer.Head
		
		for {
			
			if findNode.Element == value {

				return findNode
			
			}
		
		}
		
	
	}
	
}

func getList() *List {

	listPointer := &List{}
	return listPointer
	
}

func getNode() *Node {
	
	node := &Node{}
	return node
	
}
/*
func insertNode(value int, listPointer *List) {

	nodePointer := getNode(Element : value)
	
	if isEmpty(listPointer) == true {

		listPointer.Head = nodePointer
	
	} else {
		
		for {
			
			
			
		}
	
	}

}
uncomplited code. make List Search function first. */

func push(value int, listPointer *List) {	// To push a Node after the Tail Node:

	node := &Node{Element : value}
	
	if isEmpty(listPointer) == true {

		listPointer.Head = node
	
	} else {

		listPointer.Tail.Next = node
	
	}
	
	listPointer.Tail = node
	
}

func pushToHead(value int, listPointer *List) {

	node := &Node{Element : value}
	
	if isEmpty(listPointer) == true {

		listPointer.Head = node
	
	} else {

		node.Next = listPointer.Head
		listPointer.Head = node
	
	}
}

/* func deleteTail(listPointer *List) {

	if listPointer.Head == nil {

		fmt.Println("List is empty!")
		
	} else {

		listPointer.Tail = nil
		&listPointer.Tail
	
	}

} uncomplited code. */

func deleteHead(listPointer *List) {

	if isEmpty(listPointer) == false {
		
		tempNode := listPointer.Head.Next
		listPointer.Head.Next = nil
		listPointer.Head = tempNode
		
	}
	
}

func main() {

	listPointer := getList()	// listPointer 는 *List 타입
	
	push(200, listPointer)
	fmt.Println(listPointer.Head.Element)
	
	pushToHead(300, listPointer)
	fmt.Println(listPointer.Head.Element)
	
	deleteHead(listPointer)
	fmt.Println(listPointer.Head.Element)
	
	
}
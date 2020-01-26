package main

import "fmt"

type List struct {

	head *Node
	tail *Node

}

type Node struct {

	key interface{}
	link *Node

}

func (list *List)PushToHead(key interface{}) {
	
	node := &Node{key : key}

	if list.head == nil {

		list.head = node
		list.head.link = node
		list.tail = list.head
		list.tail.link = list.head
	
	} else {

		temp1 := list.head
		list.head = node
		list.head.link = temp1
		list.tail = list.head
	
	}

}

func (list *List)PushToTail(key interface{}) {

	node := &Node{key : key}
	
	if list.head == nil {

		list.head = node
		list.head.link = node
		list.tail = node
		list.tail.link = node
	
	} else {
		
		list.tail.link = node
		list.tail = node
		list.tail.link = list.head
		
	}

}

func (list *List)DeleteHead() {



}

func (list *List)DeleteTail() {



}

func (list *List)Insert(key interface{}) {



}

func (list *List)search(key interface{}) {



}

func (list *List)Display() {

	if list.head == nil {

		fmt.Println("list is empty.")

	} else {
		
		searchNode := list.head
		
		for searchNode.link != list.head {

			fmt.Printf("%v -> ", searchNode)
			searchNode = searchNode.link
			
		}
		
		fmt.Printf("%v -> \n", searchNode)
		
	}

}

func main() {

	list := &List{}
	
	list.PushToHead(3)
	list.PushToTail(4)
	list.PushToTail(2)
	
	list.Display()

}
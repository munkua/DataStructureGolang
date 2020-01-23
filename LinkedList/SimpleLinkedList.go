package main

import "fmt"

// List, Node struct 

type List struct {

	Head *Node
	Tail *Node

}	// complete!!

type Node struct {

	Key interface{}
	Next *Node

}	// complete!!

func getList() *List {		// return Empty List

	listPointer := &List{}
	return listPointer
	
}	// complete!!

func getNode() *Node {		// return Empty Node

	nodePointer := &Node{}
	return nodePointer		// 노드포인터의 주소를 반환

}	// complete!!

func (listPointer *List)isEmpty() bool{	// check this List is Empty

	if listPointer.Head == nil {

		fmt.Println("List is empty!")
		return true	

	} else {

	return false
	
	}
	

}	// complete!!

func (listPointer *List)pushToTail(key interface{}) {	// To push a Node after the Tail Node:

	node := &Node{Key : key}
	
	if listPointer.isEmpty() == true {

		listPointer.Head = node
		listPointer.Tail = node
		node.Next = nil
	
	} else {

		listPointer.Tail.Next = node
		listPointer.Tail = node
		listPointer.Tail.Next = nil
	
	}
	
}	// complete!!

func (listPointer *List)pushToHead(key interface{}) {

	node := &Node{Key : key}
	
	if listPointer.isEmpty() == true {

		listPointer.Head = node
		listPointer.Tail = node
		node.Next = nil
	
	} else {		// 공백리스트가 아닐 경우
	
		node.Next = listPointer.Head
		listPointer.Head = node
		
	}
	
}	// complete!!

func (listPointer *List)deleteHead() {

	if listPointer.isEmpty() == true {
		
		fmt.Println("List is empty!")
		
	} else {

		tempNode := listPointer.Head.Next
		listPointer.Head.Next = nil
		listPointer.Head = tempNode
		
	}
	
}	// complete!!

func (listPointer *List)search(key interface{}) *Node{	// search Node from argument key

	cycleNode := listPointer.Head
	
	if listPointer.isEmpty() == true {

		fmt.Println("List is Empty! search fail")
		
	} else if listPointer.Head == listPointer.Tail {

		if listPointer.Head.Key == key {

			searchPointer := listPointer.Head
			return searchPointer
		
		}
	
	} else {
		
		for cycleNode.Next != nil {

			if cycleNode.Key == key {

				return cycleNode
				
			} else {

				cycleNode = cycleNode.Next
				
			}
		
		}
	
	}
	
	return cycleNode
	
}	// complete!!

func (listPointer *List)insertNode(key interface{}, searchPointer *Node) {
	
	newNode := &Node{Key : key}

	if listPointer.isEmpty() == true {

		listPointer.pushToHead(key)
	
	} else if listPointer.Head == listPointer.Tail {

		listPointer.pushToTail(key)
	
	} else {

		if searchPointer.Next == nil {

			listPointer.pushToTail(key)
		
		} else if searchPointer == listPointer.Head {

			temp := listPointer.Head.Next
			listPointer.Head.Next = searchPointer
			searchPointer.Next = temp
			
		} else {

			temp := searchPointer.Next
			searchPointer.Next = newNode
			newNode.Next = temp
		
		}
	
	}

}	// complete!!


func (listPointer *List)deleteTail() {
	
	if listPointer.Head == nil {		// if list is empty?

		fmt.Println("List is empty!")
		
	} else if listPointer.Head.Next == nil {	// if list have only one Node
	
		listPointer.Head = nil
		listPointer.Tail = nil
		
	} else {		// list have two or more Node
		
		currentNode := listPointer.Head.Next	
		previousNode := listPointer.Head

		if currentNode.Next == nil {		// if List have only two Node

			previousNode.Next = nil			
			listPointer.Tail = previousNode
		
		} else {				// if List have two or more Node

			for currentNode.Next != nil {	// search until List Tail

				previousNode = currentNode
				currentNode = currentNode.Next
		
			}
			
			previousNode.Next = nil
			listPointer.Tail = previousNode

		}
		
	
	}

}	// complete!!


func (listPointer *List)printAllNode() {

	if (listPointer.isEmpty() == true) {

		fmt.Println("List is empty!")
	
	} else {
		
		cycleNode := listPointer.Head		// cycleNode는 임시적으로 사용할 순환 노드포인터
		count := 1
		
		for (cycleNode.Next != nil) {		// cycleNode의 Next필드가 nil일 때까지

			fmt.Println(count , "Node is = ", cycleNode.Key)		// 가리키는 노드의 Key필드값을 출력
			cycleNode = cycleNode.Next		// cycleNode를 한칸 앞으로 이동
			count++
		
		}
		
		fmt.Println(count , "Node is = ", cycleNode.Key)		// 가리키는 노드의 Key필드값을 출력


	}

}	// Complete!!

func main() {

	listPointer := getList()
	node := getNode()
	
	listPointer.pushToHead(1)
	listPointer.pushToTail(2)
	listPointer.pushToHead(3)
	
	node = listPointer.search(2)
	listPointer.insertNode(4, node)
	
	listPointer.printAllNode()		// 3124
	
	listPointer.deleteHead()		// 124
	listPointer.deleteTail()		// 12
	
	listPointer.printAllNode()		// 12

}
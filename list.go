package main

import "fmt"


type myNode struct {
	value string
	next *myNode
}

type myLinkedList struct {
	length int
	head *myNode
}

func (l *myLinkedList) len() int {
	return l.length
}

func (L *myLinkedList) append(data string) {
	newMyNode := &myNode{value: data, next: nil }

	if L.head == nil {
		fmt.Println("first time")
		L.head = newMyNode
	}else{
		currentPointer := L.head
		for currentPointer.next != nil {
			currentPointer = currentPointer.next
		}
		currentPointer.next = newMyNode
	}
	L.length++;
}


func main(){
	testingList := myLinkedList{length: 0}
	testingList.append("hola")
	testingList.append("que")
	testingList.append("tal?")
	fmt.Println(testingList.len())
	fmt.Println(testingList.head)
}

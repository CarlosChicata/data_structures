/*
Purpose:
	implement my version of simple linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [x] : add element to end of list
	- display [x] : display element in list
	- popend [] : remove  last element of list
	- appstart [] : add element to first in list
	- popstart [] : remove element to first in list
	- insert [] : add element into list based in position
	- pop [] : remove element in list based in position
*/

package main

import "fmt"

//////////////////////
////// data structures

type myNode struct {
	value string
	next *myNode
}

type myLinkedList struct {
	length int
	head *myNode
}

//////////////////////
////// methods

/// get lenght of list
func (l *myLinkedList) len() int {
	return l.length
}

/// insert element to end of list
func (L *myLinkedList) append(data string) {
	newMyNode := &myNode{value: data, next: nil }

	if L.head == nil {
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

func (L *myLinkedList) display (){
	currentNode := L.head
	index := 1
	var msg string
	for currentNode != nil {
		msg = fmt.Sprintf("position %d : %s", index, currentNode.value)
		fmt.Println(msg)
		currentNode = currentNode.next
		index++
	}	
}


func main(){
	testingList := myLinkedList{length: 0}
	testingList.append("hola")
	testingList.append("que")
	testingList.append("tal?")
	fmt.Println()
	testingList.display()
}

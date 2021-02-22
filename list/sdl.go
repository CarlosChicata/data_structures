/*
Purpose:
	implement my version of double linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [x] : add element to end of list
	- display [x] : display element in list
	- popend [x] : remove  last element of list
	- appstart [ ] : add element to first in list
	- popstart [ ] : remove element to first in list
	- insert [ ] : add element into list based in position if apply
	- pop [ ] : remove element in list based in position
*/
package main

import "fmt"


//////////////////////
////// data structures

type myNode struct {
	value string
	next *myNode
	previous *myNode
}

type myDoubledList struct {
	length int
	head *myNode
}

/// get lenght of list
func (l *myDoubledList) len() int {
	return l.length
}

/// add element to tail of list
func (L *myDoubledList) append(data string){ 
	if L.length == 0 {
		L.head = &myNode{ value: data, next: nil, previous: nil}
	} else{
		currentPointer := L.head
		for currentPointer.next != nil {
			currentPointer = currentPointer.next
		}
		currentPointer.next = &myNode{value: data, previous: currentPointer, next: nil }
	}
	L.length++;
}

/// remove first element in list
func (L *myDoubledList) popend() *myNode {
	var rptaNode *myNode

	if L.length <= 1 {
		rptaNode = L.head
		L.head = nil
		return rptaNode
	}

	currentNode := L.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	if currentNode.previous != nil {
		currentNode.previous.next = nil
		currentNode.previous = nil
	}

	L.length--;
	return rptaNode
}


/// display all elements in list
func (L *myDoubledList) display (){
	currentNode := L.head
	index := 1
	var msg string
	for currentNode != nil {
		msg = fmt.Sprintf("position %d : %s", index, currentNode.value)
		fmt.Println(msg)
		currentNode = currentNode.next
		index++
	}
	fmt.Println("-----end----")
}

/// remove first element in list
func (L *myDoubledList) popstart () *myNode {
	var rptaNode *myNode
	rptaNode = nil

	if L.head == nil {
		L.length = 0
	}else if L.length == 1 {
		rptaNode = L.head
		L.head = nil
		L.length = 0
	}else {
		rptaNode = L.head
		L.head = L.head.next
		L.head.previous.next = nil
		L.head.previous = nil
		rptaNode.next = nil
		rptaNode.previous = nil
	}

	return rptaNode
}


func main(){
	testingList := myDoubledList{length: 0}
	testingList.append("¿")
	testingList.append("Hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.append("?")
	testingList.display()
	fmt.Println(testingList.popstart())
	fmt.Println(testingList.head)
	testingList.display()
}
/*
Purpose:
	implement my version of simple linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [x] : add element to end of list
	- display [x] : display element in list
	- popend [] : remove  last element of list
	- appstart [x] : add element to first in list
	- popstart [x] : remove element to first in list
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

/// display all elements in list
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
	fmt.Println("-----end----")
}

/// add first of list
func (L *myLinkedList) appstart (data string){
	newNode := &myNode{ value: data, next: nil }
	if L.head == nil {
		L.head = newNode
	}else{
		newNode.next = L.head
		L.head = newNode
	}
	L.length++;
}

/// pop first element in list
func (L* myLinkedList) popstart () *myNode{
	var headNode *myNode
	headNode = nil
	if L.head != nil {
		headNode = L.head
		L.head = L.head.next
	}
	L.length--;
	return headNode
}

/// pop last element in list
func (L *myLinkedList) popend () *myNode {
	var rptaNode *myNode
	rptaNode = nil
	if L.length >= 1 {
		rptaNode = L.head
		L.head = nil
	}else{
		nextNode := L.head
		nextnextNode := L.head.next
		for nextnextNode != nil {
			nextNode = nextnextNode
			nextnextNode = nextnextNode.next
		}
		rptaNode = nextnextNode
		nextNode.next = nil
	}
	L.length--;
	return rptaNode
}

func main(){
	testingList := myLinkedList{length: 0}
	testingList.append("hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.append("?")
	fmt.Println(testingList.popend())
	testingList.display()
}

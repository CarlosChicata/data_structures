/*
Purpose:
	implement my version of double linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [x] : add element to end of list
	- display [x] : display element in list
	- popend [x] : remove  last element of list
	- appstart [x] : add element to first in list
	- popstart [x] : remove element to first in list
	- insert [x] : add element into list based in position if apply
	- pop [x] : remove element in list based in position
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
		L.length = 0
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
		L.length--
	}

	return rptaNode
}

/// add new head in list 
func (L *myDoubledList) appstart (data string) {
	if L.head == nil {
		L.length = 1
		L.head = &myNode{ value : data , previous : nil , next : nil }
	}else {
		L.head.previous = &myNode{ value : data , previous : nil , next : L.head }
		L.head = L.head.previous
		L.length++
	}
}

/// insert element in list based in position, never insert in tail list
func (L *myDoubledList) insert (data string, pos int) bool {
	if L.length < pos || pos < 1 {
		return false
	} else if pos  == 1 {
		newNode :=  &myNode{ value: data, next: L.head, previous: nil }
		if L.head != nil {
			L.head.previous = newNode
		}
		L.head = newNode
	}else {
		previousNode := L.head
		for idx := 1; idx < pos - 1; idx++ {
			previousNode = previousNode.next
		}
		newNode := &myNode{ value: data, next: previousNode.next, previous: previousNode }
		previousNode.next = newNode
	}
	L.length++
	return true
}

/// pop element of list based in position
func (L *myDoubledList) pop (pos int) *myNode {
	if L.length < pos || pos < 1 || L.head == nil { // 5
		return nil
	} else if pos == 1 && L.head != nil { // 1 y 2
		myRpta := L.head
		L.head = L.head.next
		myRpta.next = nil

		if L.head.previous != nil {
			L.head.previous = nil
		}

		L.length--
		return myRpta
	}else {
		previousNode := L.head

		for idx := 1; idx < pos - 1 ; idx++ {
			previousNode = previousNode.next
		}

		myRpta := previousNode.next
		previousNode.next = myRpta.next
		myRpta.previous = nil


		if myRpta.next != nil {
			myRpta.next.previous = previousNode
		}
		L.length--
		return myRpta
	}
}


func main(){
	testingList := myDoubledList{length: 0}
	testingList.append("Hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.display()
	testingList.insert("¿", 1)
	testingList.display()
	testingList.pop(3)
	testingList.pop(1)
	testingList.display()
	fmt.Println(testingList.length)
}
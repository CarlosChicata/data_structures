/*
Purpose:
	implement my version of circle double linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [x] : add element to end of list
	- display [x] : display element in list
	- popend [x] : remove  last element of list
	- appstart [x] : add element to first in list
	- popstart [x] : remove element to first in list
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

type myCircleDList struct {
	length int
	head *myNode
}

/// get lenght of list
func (l *myCircleDList) len() int {
	return l.length
}

/// add element to last of list
func (L *myCircleDList) append (data string) {
	if L.head == nil {
		L.head = &myNode{ value: data, next: nil, previous: nil }
		L.head.next = L.head
		L.head.previous = L.head
	}else {
		newNode := &myNode{ value: data, next: L.head, previous: L.head.previous }
		L.head.previous.next = newNode
		L.head.previous = newNode
	}
	L.length++
}

/// display all elements in list
func (L *myCircleDList) display (){
	currentNode := L.head
	index := 1
	var msg string
	for index <= L.length && currentNode != nil {
		msg = fmt.Sprintf("position %d : %s", index, currentNode.value)
		fmt.Println(msg)
		currentNode = currentNode.next
		index++
	}
	fmt.Println("-----end----")
}

/// add element to first of list
func (L *myCircleDList) appstart (data string) {
	if L.head == nil {
		L.head = &myNode{ value: data, next: nil, previous: nil }
		L.head.next = L.head
		L.head.previous = L.head
	}else {
		newNode := &myNode{ value: data, next: L.head, previous: L.head.previous }
		L.head.previous.next = newNode
		L.head.previous = newNode
		L.head = newNode
	}
	L.length++
}

/// pop last element of list
func (L *myCircleDList) popend () *myNode {
	if L.length == 0 || L.head == nil {
		return nil
	}else if L.length == 1 {
		rptaNode := L.head
		L.head = nil
		L.length = 0
		return rptaNode
	}else {
		rptaNode := L.head.previous
		rptaNode.previous.next = L.head
		L.head.previous = rptaNode.previous
		rptaNode.previous = nil
		rptaNode.next = nil
		L.length--
		return rptaNode
	}
}

/// pop start element of list
func (L *myCircleDList) popstart () *myNode {
	if L.length == 0 || L.head == nil {
		return nil
	}else if L.length == 1 {
		rptaNode := L.head
		rptaNode.next = nil
		rptaNode.previous = nil
		L.head = nil
		L.length = 0
		return rptaNode
	}else {
		rptaNode := L.head
		rptaNode.next.previous = rptaNode.previous
		rptaNode.previous.next = rptaNode.next
		L.head = rptaNode.next
		rptaNode.previous = nil
		rptaNode.next = nil
		L.length--
		return rptaNode
	}
}

func (L *myCircleDList) pop (pos int) *myNode {
	if L.length < pos || pos < 1 || L.head == nil {
		return nil
	} else if pos == 1 && L.head != nil {
		rptaNode := L.popstart()
		return rptaNode
	} else {
		previousNode := L.head

		for idx := 1; idx < pos ; idx++ {
			previousNode = previousNode.next
		}

		myRpta := previousNode
		myRpta.previous.next = myRpta.next
		myRpta.next.previous = myRpta.previous
		myRpta.previous = nil
		myRpta.next = nil
		L.length--
		return myRpta
	}
}


func main(){
	testingList := myCircleDList{length: 0}
	testingList.append("Hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.append("?")
	testingList.display()
	testingList.appstart("¿")
	testingList.display()
	testingList.pop(1)
	testingList.pop(3)
	testingList.display()
}
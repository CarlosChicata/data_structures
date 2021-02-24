/*
Purpose:
	implement my version of circle simple linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [x] : add element to end of list
	- display [ ] : display element in list
	- popend [ ] : remove  last element of list
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

type myCircleList struct {
	length int
	head *myNode
}

/// get lenght of list
func (l *myCircleList) len() int {
	return l.length
}

/// add element to last of list
func (L *myCircleList) append (data string) {
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
func (L *myCircleList) display (){
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

func main(){
	testingList := myCircleList{length: 0}
	testingList.append("Hola")
	testingList.append("que")
	testingList.display()
	testingList.append("tal")
	testingList.display()
}
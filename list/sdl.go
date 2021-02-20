/*
Purpose:
	implement my version of double linked list.

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

type myDoubledList struct {
	length int
	head *myNode
}

/// get lenght of list
func (l *myDoubledList) len() int {
	return l.length
}

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

func main(){
	testingList := myDoubledList{length: 0}
	testingList.append("¿")
	testingList.append("hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.append("?")
	testingList.display()
}
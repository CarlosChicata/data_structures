/*
Purpose:
	implement my version of simple circle linked list.

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
		L.head = &myNode{ value: data, next: nil }
		L.head.next = L.head
	}else {
		previousNode := L.head

		for previousNode.next != L.head {
			previousNode = previousNode.next
		}
		previousNode.next = &myNode{ value: data, next: L.head}
	}
	L.length++
}

/// add element to start of list
func (L *myCircleList) appstart (data string) {
	if L.head == nil {
		L.head = &myNode{ value: data, next: nil }
		L.head.next = L.head
	}else {
		newNode := &myNode{ value: data, next: L.head}
		previousNode := L.head

		for previousNode.next != L.head {
			previousNode = previousNode.next
		}
		previousNode.next = newNode
		L.head = newNode
	}
	L.length++
}

/// display all elements in list
func (L *myCircleList) display (){
	currentNode := L.head
	index := 1
	var msg string
	for idx := 1; idx <= L.length; idx++ {
		msg = fmt.Sprintf("position %d : %s", index, currentNode.value)
		fmt.Println(msg)
		currentNode = currentNode.next
		index++
	}
	fmt.Println("-----end----\n")
}

/// remove last element in list
func (L *myCircleList) popend () *myNode {
	if L.length == 0 || L.head == nil {
		return nil
	} else if L.length == 1 || L.head == L.head.next {
		newNode := L.head
		newNode.next = nil
		L.head = nil
		L.length = 0
		return newNode
	} else { 
		previousNode := L.head

		for idx := 1 ; idx < L.length ; idx++ {
			previousNode = previousNode.next
		}
		newNode := previousNode.next
		L.length--
		previousNode.next = L.head
		return newNode
	}
}

/// pop first element in list
func (L* myCircleList) popstart () *myNode{
	if L.length == 0 || L.head == nil {
		return nil
	} else if L.length == 1 || L.head == L.head.next {
		newNode := L.head
		newNode.next = nil
		L.head = nil
		L.length = 0
		return newNode
	} else { 
		LastNode := L.head
		newNode := L.head

		for idx := 1 ; idx <= L.length ; idx++ {
			LastNode = LastNode.next
		}

		LastNode.next = L.head.next
		L.head = L.head.next
		newNode.next = nil
		L.length--
		return newNode
	}
}

/// pop element of list based in position
func (L *myCircleList) pop (pos int) *myNode {
	if L.length < pos || pos < 1 || L.head == nil {
		return nil
	} else if pos == 1 && L.head != nil {
		rptaNode := L.popstart()
		return rptaNode
	} else {
		previousNode := L.head

		for idx := 1; idx < pos - 1; idx++ {
			previousNode = previousNode.next
		}
		fmt.Println(L.head)
		fmt.Println(previousNode)
		myRpta := previousNode.next
		previousNode.next = myRpta.next
		myRpta.next = nil
		L.length--
		return myRpta
	}
}

/// pop element of list based in position - not insert to last position
func (L *myCircleList) insert (data string, pos int) bool {
	if L.length < pos || pos < 1 || L.head == nil {
		return false
	} else if pos == 1 {
		L.appstart(data)
		return true
	} else {
		previousNode := L.head

		for idx := 1; idx < pos -1 ; idx++ {
			previousNode = previousNode.next
		}

		myRpta := &myNode{ value: data, next: previousNode.next }
		previousNode.next = myRpta
		L.length++
		return true
	}
}

func main(){
	testingList := myCircleList{length: 0}
	testingList.appstart("hola")
	testingList.popend()
	testingList.append("hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.append("?")
	testingList.append("?")
	testingList.appstart("¿")
	testingList.popend()
	testingList.display()
	testingList.pop(5)
	testingList.insert("amigo" , 2)
	testingList.insert("?" , 5)
	fmt.Println(testingList.length)
	
	testingList.display()
}
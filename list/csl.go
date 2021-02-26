/*
Purpose:
	implement my version of simple circle linked list.

Version: 1.0

Methods:
	- len [x] : get length of list
	- append [ ] : add element to end of list
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
		newNode := &myNode{ value: data, next: L.head }
		previousNode = L.head

		for idx := 1; idx <= L.length ; idx++ {
			previousNode = previousNode.next
		}
		previousNode.next = newMyNode
	}
	L.length++
}

func main(){
	testingList := myCircleList{length: 0}
	fmt.Println(testingList.length)
}
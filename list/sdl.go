/*
Purpose:
	implement my version of double linked list.

Version: 1.0

Methods:
	- len [ ] : get length of list
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


func main(){
	testingList := myDoubledList{length: 0}
	fmt.Println(testingList.len())
}
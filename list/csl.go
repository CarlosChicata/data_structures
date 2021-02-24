/*
Purpose:
	implement my version of circle simple linked list.

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

type circleList struct {
	length int
	head *myNode
}



func main(){

}
/*
Purpose:
	implement my version of unique elements in set.

Version: 1.0

Methods:
	- add [ ] : add element in set if not exists
	- remove [ ] : remove element in set if exists
	- intersec [ ] : all elements will intersect in two sets
	- belong [ ] : element belong this set
	- difference [ ] : all elemnet will not interset in two sets
	- len [ ] : count all elements of universe set
*/

package main

import "fmt"

type uniqeSet {
	
}




func main(){
	testingList := myLinkedList{length: 0}
	testingList.append("hola")
	testingList.append("que")
	testingList.append("tal")
	testingList.append("?")
	testingList.display()
	testingList.insert("¿", 1)
	testingList.display()
	testingList.pop(5)
	testingList.display()
}

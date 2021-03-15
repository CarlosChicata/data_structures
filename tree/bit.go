/*
Purpose:
	implement my version of duplicated elements in set.

Version: 1.0

Methods:
	- add [ ] : add element in tree
	- remove [ ] : remove element in tree if exists
	- find [ ] : find element in this tree
	- len [ ] : count all elements of tree
	- display [ ] : display all element in tree
*/

package main

import "fmt"


type myNode struct {
	value string
	left *myNode
	right *myNode
}

type myBinaryTree struct {
	length int
	head *myNode
	comparing func(string, string) bool
}


func main(){
	
}
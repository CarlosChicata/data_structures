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

func (T *myBinaryTree) add (value string) {
	currentNode := T.head
	
	for currentNode != nil {
		if T.comparing(value, currentNode.value) {
			currentNode = currentNode.right
		}else{
			currentNode = currentNode.left
		}
	}

	T.length++
	currentNode = &myNode{value: value, left: nil, right: nil}
}



func main(){
	testingTree := myBinaryTree{length: 0, comparing: func(a string, b string) bool { return a < b}}
	fmt.Println(testingTree)
	testingTree.add("hola")
	testingTree.add("que")
	testingTree.add("tal")
	testingTree.add("?")
}
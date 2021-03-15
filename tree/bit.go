/*
Purpose:
	implement my version of duplicated elements in set.

Version: 1.0

Methods:
	- add [x] : add element in tree
	- remove [ ] : remove element in tree if exists
	- find [ ] : find element in this tree
	- len [ ] : count all elements of tree
	- display [x] : display all element in tree
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

/// add element in tree
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

/// display tree
func (T *myBinaryTree) display(){
	currentNode := T.head
	var queue [] *myNode

	if currentNode != nil {
		queue = append(queue, currentNode)
	}

	for len(queue) != 0 {
		fmt.Println("value is %s: ", currentNode.value, " left: ", currentNode.left, " and right: ", currentNode.right)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}

		currentNode = queue[0]
		queue = append(queue[:0], queue[1:]...)
	}

}


func main(){
	testingTree := myBinaryTree{length: 0, comparing: func(a string, b string) bool { return a < b}}
	fmt.Println(testingTree)
	testingTree.add("hola")
	testingTree.add("que")
	testingTree.add("tal")
	testingTree.add("?")
	testingTree.display()
}
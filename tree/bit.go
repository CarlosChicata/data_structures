/*
Purpose:
	implement my version of simple binary tree.

Version: 1.0

Methods:
	- add [x] : add element in tree
	- remove [ ] : remove element in tree if exists
	- find [x] : find element in this tree
	- len [x] : count all elements of tree
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
	if T.head == nil {
		T.length = 1
		T.head = &myNode{value: value, left: nil, right: nil}
	} else {
		currentNode := T.head
		var flag bool

		for currentNode != nil {

			if T.comparing(value, currentNode.value) {
				flag = true
				if currentNode.right == nil {
					break
				} else {
					currentNode = currentNode.right
				}
			}else{
				flag = false
				if currentNode.left == nil {
					break
				} else {
					currentNode = currentNode.left
				}
			}

		}

		T.length++
		if flag {
			currentNode.right = &myNode{value: value, left: nil, right: nil}
		} else {
			currentNode.left = &myNode{value: value, left: nil, right: nil}	
		}
	}
}

/// display tree
func (T *myBinaryTree) display(){
	currentNode := T.head
	var queue [] *myNode

	if currentNode != nil {
		queue = append(queue, currentNode)
	}

	for len(queue) != 0 {
		fmt.Println("value is '", currentNode.value, "' left: ", currentNode.left, " and right: ", currentNode.right)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}

		queue = append(queue[:0], queue[1:]...)

		if len(queue) != 0 {
			currentNode = queue[0]
		}
	}

}

/// length of node in tree
func (T *myBinaryTree) len() int {
	return T.length
}

/// find value in tree
func (T *myBinaryTree) find(value string) bool {
	currentNode := T.head
	rptaFlag := false

	for currentNode != nil {
		if currentNode.value == value {
			rptaFlag = true
			break
		} else {
			if T.comparing(value, currentNode.value) {
				currentNode = currentNode.right
			}else{
				currentNode = currentNode.left
			}
		}
	}
	return rptaFlag
}

func main(){
	testingTree := myBinaryTree{length: 0, comparing: func(a string, b string) bool { return a < b}}
	testingTree.add("hola")
	testingTree.add("que")
	testingTree.add("tal")
	fmt.Println("---------------------")
	testingTree.display()
	fmt.Println(testingTree.len())
	fmt.Println(testingTree.find("kha"))
}
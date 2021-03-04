/*
Purpose:
	implement my version of disjoint-union set.

Version: 1.0

Methods:
	- add [ ] : add element in universe if not exists
	- belong [ ] : element belong this universe
	- len [ ] : count all elements of universe set
	- display [ ] : display all element in universe set
	- union [ ] : merge two set
	- parentIn [ ] : get parent of set
	- sizeIn [ ] : get size of set
	- sizeSet [ ] : get number of set in universe.
*/

package main

import "fmt"


type DisjointUnion struct {
	universe map[string]string
}


func (D *DisjointUnion) len () int {
	return len(D.universe)
}

func main(){
	testingDUS := DisjointUnion{}
	testingDUS.add("hola")
	testingDUS.add("que")
	testingDUS.add("tal")
	testingDUS.add("?")
	testingDUS.add("?")
	testingDUS.add("?")
	fmt.Println(testingDUS.len());
}
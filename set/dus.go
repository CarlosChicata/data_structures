/*
Purpose:
	implement my version of disjoint-union set.

Version: 1.0

Methods:
	- add [x] : add element in universe if not exists
	- belong [ ] : element belong this universe
	- len [x] : count all elements of universe set
	- display [ ] : display all element in universe set
	- union [ ] : merge two set
	- parentIn [ ] : get parent of set
	- sizeIn [ ] : get size of set
	- sizeSet [ ] : get number of set in universe.

Internal Methods:
	- preparing [x]: preparing map fields in Disjoint Union
*/

package main

import "fmt"


type DisjointUnion struct {
	universe map[string]string
}

func (D *DisjointUnion) preparing () {
	D.universe = make(map[string]string)
}

func (D *DisjointUnion) add (value string) bool {
	var rptaFlag bool

	if _, ok := D.universe[value]; ok {
		rptaFlag = true
	}else{
		D.universe[value] = value
		rptaFlag = true
	}

	return rptaFlag
}

func (D *DisjointUnion) len () int {
	return len(D.universe)
}

func (D *DisjointUnion) belong (value string) bool {
	rptaFlag := false
	if _, ok := D.universe[value]; ok {
		rptaFlag = true
	}
	return rptaFlag
}

func main(){
	testingDUS := DisjointUnion{}
	testingDUS.preparing()
	testingDUS.add("hola")
	testingDUS.add("que")
	testingDUS.add("tal")
	testingDUS.add("?")
	testingDUS.add("?")
	testingDUS.add("?")
	fmt.Println(testingDUS.len());
}
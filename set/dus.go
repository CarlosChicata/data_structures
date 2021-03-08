/*
Purpose:
	implement my version of disjoint-union set.

Version: 1.0

Methods:
	- add [x] : add element in universe if not exists
	- belong [x] : element belong this universe
	- len [x] : count all elements of universe set
	- display [ ] : display all element in universe set
	- union [ ] : merge two set: they are two variants:
		- unionBySize [ ] : merge using size
		- unionByRank [ ] : merge using rank
	- parentIn [x] : get parent of set
	- sizeIn [ ] : get size of set
	- sizeSet [ ] : get number of set in universe.

Internal Methods:
	- preparing [x]: preparing map fields in Disjoint Union
*/

package main

import "fmt"


type DisjointUnion struct {
	universe map[string]string
	size_sets map[string]int
}

/// preparing maps of disjoint-union set
func (D *DisjointUnion) preparing () {
	D.universe = make(map[string]string)
	D.size_sets = make(map[string]int)
}

/// add value of disjoint-union set if it is not exists
func (D *DisjointUnion) add (value string) bool {
	var rptaFlag bool

	if _, ok := D.universe[value]; ok {
		rptaFlag = true
	}else{
		D.universe[value] = value
		D.size_sets[value] = 1
		rptaFlag = true
	}

	return rptaFlag
}

/// length of disjoint-union set
func (D *DisjointUnion) len () int {
	return len(D.universe)
}

/// if value belong this disjoint-union set
func (D *DisjointUnion) belong (value string) bool {
	rptaFlag := false
	if _, ok := D.universe[value]; ok {
		rptaFlag = true
	}
	return rptaFlag
}

/// find parent of value
func (D *DisjointUnion) parentIn (value string) (string, bool) {
	if D.belong(value) == false {
		return value, false
	}
	
	valueMap := D.universe[value]
	
	for value != valueMap {
		value = D.universe[value]
		valueMap = D.universe[valueMap]
		fmt.Println(value, valueMap)
	}

	return valueMap, true
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
	fmt.Println(testingDUS.parentIn("hola"))
}
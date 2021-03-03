/*
Purpose:
	implement my version of duplicated elements in set.

Version: 1.0

Methods:
	- add [x] : add element in set if not exists
	- remove [x] : remove element in set if exists
	- intersect [x] : all elements will intersect in two sets
	- belong [x] : element belong this set
	- difference [x] : all elemnet will not interset in two sets
	- len [x] : count all elements of universe set
	- display [x] : display all element in universe set
*/

package main

import "fmt"


type duplicatedSet struct {
	universe []string
}

/// verify if element inside set
func (S *duplicatedSet) belong (element string) bool {
	isBelong := false
	for _, value := range S.universe {
		if element == value {
			isBelong = true
			break;
		}
	}
	return isBelong
}

/// display all element in set
func (S *duplicatedSet) display () {
	var msg string
	for _, value := range S.universe {
		msg = fmt.Sprintf("value is: %s", value)
		fmt.Println(msg)
	}
	fmt.Println("------END------")
}

/// add element in set if not exists
func (S *duplicatedSet) add (element string) bool {
	S.universe = append(S.universe, element)
	return true
}

/// length of set
func (S *duplicatedSet) len() int {
	return len(S.universe)
}

/// remove element in set if exists
func (S *duplicatedSet) remove(element string) bool {
	rptaFlag := false
	var newUniverse []string

	for _, value := range S.universe {
		if value == element && rptaFlag == false {
			rptaFlag = true
		}else if rptaFlag == false {
			newUniverse = append(newUniverse, value)			
		}else{
			newUniverse = append(newUniverse, value)
		}
	}

	S.universe = newUniverse
	return rptaFlag
}

/// get new uniqueSet is intersection of two set
func (S *duplicatedSet) intersect (otherSet duplicatedSet) duplicatedSet {
	var rptaNewSet []string
	
	for _, value := range S.universe {
		if otherSet.belong(value) {
			rptaNewSet = append(rptaNewSet, value)
		}
	}
	return duplicatedSet{ universe: rptaNewSet }
}

/// get a uniqueSet with element belong host but not belong of guess
func (S *duplicatedSet) difference (otherSet duplicatedSet) duplicatedSet {
	var rptaNewSet []string

	for _, value := range S.universe {
		if !otherSet.belong(value) {
			rptaNewSet = append(rptaNewSet, value)
		}
	}

	return duplicatedSet{ universe: rptaNewSet }
}

func main(){
	testingSet := duplicatedSet{}
	testingSet.add("¿")
	testingSet.add("hola")
	testingSet.add("que")
	testingSet.add("tal")
	testingSet.add("?")
	testingSet.add("?")
	testingSet.add("?")
	testingSet.display()
	testingSet2 := duplicatedSet{}
	testingSet2.add("?")
	testingSet2.add("hola")
	testingSet2.add("gaaaaaaaaa")
	testingSet2.display()
	testingSet3 := testingSet.intersect(testingSet2)
	testingSet3.display()
	testingSet4 := testingSet.difference(testingSet2)
	testingSet4.display()
}
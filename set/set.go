/*
Purpose:
	implement my version of unique elements in set.

Version: 1.0

Methods:
	- add [x] : add element in set if not exists
	- remove [x] : remove element in set if exists
	- intersect [x] : all elements will intersect in two sets
	- belong [x] : element belong this set
	- difference [ ] : all elemnet will not interset in two sets
	- len [x] : count all elements of universe set
	- display [x] : display all element in universe set
*/

package main

import "fmt"


type uniqueSet struct {
	universe []string
}

/// verify if element inside set
func (S *uniqueSet) belong (element string) bool {
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
func (S *uniqueSet) display () {
	var msg string
	for _, value := range S.universe {
		msg = fmt.Sprintf("value is: %s", value)
		fmt.Println(msg)
	}
	fmt.Println("------END------")
}

/// add element in set if not exists
func (S *uniqueSet) add (element string) bool {
	var isBelong bool;
	isBelong = S.belong(element)

	if isBelong {
		isBelong = false
	}else{
		S.universe = append(S.universe, element)
		isBelong = true
	}
	return isBelong
}

/// length of set
func (S *uniqueSet) len() int {
	return len(S.universe)
}

/// remove element in set if exists
func (S *uniqueSet) remove(element string) bool {
	rptaFlag := false
	var newUniverse []string
	for _, value := range S.universe {
		if value == element {
			rptaFlag = true
		}else{
			newUniverse = append(newUniverse, value)
		}
	}
	S.universe = newUniverse
	return rptaFlag
}

/// get new uniqueSet is intersection of two set
func (S *uniqueSet) intersect (otherSet uniqueSet) uniqueSet {
	var rptaNewSet []string
	
	for _, value := range S.universe {
		if otherSet.belong(value) {
			rptaNewSet = append(rptaNewSet, value)
		}
	}
	return uniqueSet{ universe: rptaNewSet }
}

func (S *uniqueSet) difference (otherSet uniqueSet) uniqueSet {
	var rptaNewSet []string

	for _, value := range S.universe {
		if !otherSet.belong(value) {
			rptaNewSet = append(rptaNewSet, value)
		}
	}

	return uniqueSet{ universe: rptaNewSet }
}

func main(){
	testingSet := uniqueSet{}
	testingSet.add("¿")
	testingSet.add("hola")
	testingSet.add("que")
	testingSet.add("tal")
	testingSet.add("?")
	testingSet.display()
	testingSet2 := uniqueSet{}
	testingSet2.add("que")
	testingSet2.add("hace")
	testingSet2.add("?")
	testingSet2.display()
	var testingSet3 uniqueSet
	testingSet3 = testingSet.intersect(testingSet2)
	testingSet3.display()
	testingSet4 := testingSet.difference(testingSet2)
	testingSet4.display()
}

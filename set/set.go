/*
Purpose:
	implement my version of unique elements in set.

Version: 1.0

Methods:
	- add [x] : add element in set if not exists
	- remove [ ] : remove element in set if exists
	- intersec [ ] : all elements will intersect in two sets
	- belong [x] : element belong this set
	- difference [ ] : all elemnet will not interset in two sets
	- len [ ] : count all elements of universe set
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


func main(){
	testingSet := uniqueSet{}
	testingSet.add("hola")
	testingSet.add("que")
	testingSet.add("tal")
	testingSet.add("?")
	testingSet.display()
	testingSet.add("¿")
	testingSet.display()
}

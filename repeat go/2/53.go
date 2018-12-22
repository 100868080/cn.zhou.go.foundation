package main

import (
	"fmt"
)

func main() {
	var a [3]int             //array of 3 integers
	fmt.Println(a[0])        //print the first element
	fmt.Println(a[len(a)-1]) //print the last element,a[2]
	//print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)

	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], "\n", q) //"0"
	qu := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) //"[3]int"
	fmt.Println(qu)
	qp := [3]int{1, 2, 3}
	qp := [4]int{1, 2, 3, 4} //compile error:cannot assign [4]int to [3]int
	fmt.Println(qp)
}

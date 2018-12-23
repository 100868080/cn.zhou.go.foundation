//reverse reverses a slice of ints in place
package main

import (
	"fmt"
)

func main() {

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) //"[5 4 3 2 1 0]"
	sk := []int{0, 1, 2, 3, 4, 5}
	//rotate s left by two positions.
	reverse(sk[:2])
	reverse(sk[2:])
	reverse(sk)
	fmt.Println(s) //"[2 3  4 5 0 1]"

}

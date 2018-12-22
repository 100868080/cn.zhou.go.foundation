package main

import (
	"fmt"
)

func nonempty2(strings []string) []string {
	out := strings[:0] //zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

/*stack=append(stack,v)  //push v
top:=stack[len(stack)-1]  //top of stack
stack=stack[:len(stack)-1]  //pop
*/
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) //"[5 6 8 9]"
}

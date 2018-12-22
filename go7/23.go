package main

import (
	"fmt"
)

func main() {
	f1(3, 1, "go", "3.156", "language", 'a', false, 'F', 3.654879)
}
func f1(a ...interface{}) {
	var num = make([]int, 0, 9)
	var str = make([]string, 0, 6)
	var ch = make([]int32, 0, 6)
	var other = make([]interface{}, 0, 6)
	for _, a := range a {
		switch v := a.(type) {
		case int:
			num = append(num, v)
		case string:
			str = append(str, v)
		case int32:
			ch = append(ch, v)
		default:
			other = append(other, v)
		}
	}
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(ch)
	fmt.Println(other)
}

package main

import (
	"fmt"
)

func main() {
	var b int = 66565
	f2(&b)
	fmt.Println(b)
}
func f2(a *int) {
	*a += 2456999

	fmt.Printf("%d\n", *a)
}

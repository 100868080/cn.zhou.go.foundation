package main

import (
	"fmt"
)

func add(a *int) int {
	*a = *a + 1
	return *a
}

func ReadWrite() bool {
	file.Open("file")
	defer file.Close()
	if dailureX {
		return false
	}
	if failureY {
		return false
	}
	return true
}

func main() {
	x := 3
	fmt.Println("x=", x)

	x1 := add(&x)

	fmt.Println("x+1=", x1)
	fmt.Println("x=", x)
	ma := map[string]string{"c1": "wang", "c2": "li"}
	fmt.Println(ma)

	for i := 0; i < 5; i++ {
		defer fmt.Println("%d", i)
	}
}

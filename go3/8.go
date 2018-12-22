package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	if a > 3 {
		fmt.Println("ok")
	} else {
		fmt.Println("ko")
	}
	fmt.Println("over")

}

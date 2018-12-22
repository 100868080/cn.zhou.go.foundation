package main

import (
	"fmt"
)

func test() {
	fmt.Println("GO...")
}
func main() {
	for i := 0; i < 10; i++ {
		go test()
	}
}

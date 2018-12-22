package main

import (
	"fmt"
)

func test() {
	fmt.Println("Go...")
}

func main() {
	for i := 0; i < 10; i++ {
		go test()
	}
}

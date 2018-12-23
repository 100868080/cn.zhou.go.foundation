package main

import (
	"fmt"
)

func main() {
	f(3)
}

func f(x int) {
	fmt.printf("f(%d)\n", x+0/x)
	defer fmt.printf("defer %d\n", x)
	f(x - 1)
}

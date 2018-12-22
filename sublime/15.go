package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type ele struct {
	Human
	spen  string
	phone string
}

func main() {
	bob := ele{Human{"bob", 32, "8-456-8546-456-xxx"}, "deth", "564-121"}
	fmt.Println(bob.phone)
	fmt.Println(bob.Human.phone)

}

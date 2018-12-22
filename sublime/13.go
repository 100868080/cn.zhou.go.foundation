package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	spen string
}

func main() {
	mark := Student{Human{"mark", 12, 66}, "computer games"}
	fmt.Println(mark.Human)
	fmt.Println(mark.spen)
	fmt.Println(mark)

	mark.spen = "AI"
	fmt.Println(mark.spen)

	mark.age = 166

	mark.weight += 231
	fmt.Println(mark)

	mark.Human = Human{"wang", 654, 22}
	mark.Human.age -= 9
	fmt.Println(mark.Human, mark.Human.age)

}

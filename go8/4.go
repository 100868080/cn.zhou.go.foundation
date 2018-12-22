package main

import (
	"fmt"
)

type user struct {
	name string
	id   int
}

func main() {
	u2 := user{id: 456}
	u1 := user{"limti", 123}
	fmt.Println(u1, u2)
	var u3 *user = new(user)
	*u3 = u1
	var u4 user
	u4 = u2
	fmt.Println(u4, u4 == u2, u1 != u2)
	fmt.Println(u3)
	u5 := new(user)
	u5 = u3
	fmt.Println(u5)
}

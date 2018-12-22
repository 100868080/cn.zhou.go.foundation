package main

import (
	"fmt"
)

type User struct {
	id   int
	Name string
	age  int
}

func (u *User) Say() {
	fmt.Println("I,m %s nice to!\n", u.Name)
}
func (u User) old() int {
	return u.age
}
func main() {
	User.Say()
}

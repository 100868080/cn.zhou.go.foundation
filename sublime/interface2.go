package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}
func (h Human) Sing(lyrics string) {
	fmt.Println("la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s.call me %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"mike", 25, "564-89564-xxx"}, "Mit", 0.00}
	paul := Student{Human{"paul", 26, "123-98-xxx"}, "haryvard", 100}
	sam := Employee{Human{"sam", 56, "564-78-x"}, "golang inc", 1000}
	tom := Employee{Human{"tom", 123, "7+89+-6132"}, "thing ltd.", 5000}
	var i Men
	i = mike
	fmt.Println("this is mike,a student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("this is tom,an employee:")
	i.SayHi()
	i.Sing("born to be wild")

	fmt.Println("let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
	/*
		var a interface{}
		var i int = 5
		s := "Hello world!"
		a = i
		a = s*/
}

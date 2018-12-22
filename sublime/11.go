package main

import (
	. "fmt"
	_ "os"
)

type person struct {
	id   int
	name string
	age  int
}

//var p person

func main() {
	Println("the color is blue!")
	/*p.name = "alise"
	p.id = 12354
	p.age = 213
	println(p)*/
	p := person{12036, "tom", 25}
	Println(p)
	P := person{name: "tomi", id: 1764, age: 62}
	Println(P)
}

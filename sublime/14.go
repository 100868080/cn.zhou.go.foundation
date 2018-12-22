package main

import (
	. "fmt"
)

type skills []string
type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	skills
	int
	Spen string
}

func main() {
	jane := Student{Human: Human{"jane", 23, 456}, Spen: "player"}
	Println(jane)
	Println(jane.Human, jane.Spen, jane.age)
	jane.skills = []string{"today"}
	Println(jane.skills)
	jane.skills = append(jane.skills, "physics", "golang", "math")
	Println(jane.skills)
	jane.int = 34
	Println(jane.int)

}

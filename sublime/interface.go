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

func (h *Human) sayhi() {
	fmt.Printf("hi,I'm %s ,you can call me on %s\n", h.name, h.phone)
}

func (h *Human) sing(ly string) {
	fmt.Println("la la,la la la...", ly)

}
func (h *Human) guzzle(beerstein string) {
	fmt.Println("guzzle guzzle...", beerstein)
}
func (e *Employee) sayhi() {
	fmt.Printf("hi,I'm %s ,I work at %s.call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) borrowmoney(amount float32) {
	s.loan += amount
}
func (e *Employee) spendsalary(amount float32) {
	e.money -= amount
}

type men interface {
	sayhi()
	sing(ly string)
	guzzle(beerstein string)
}

type youngchap interface {
	sayhi()
	sing(song string)
	borrowmoney(amount float32)
}

type elderlygent interface {
	sayhi()
	sing(song string)
	spendsalary(amount float32)
}

func main() {

}

package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "0" + h.name + "-" + strconv.Itoa(h.age) + "years-@" + h.phone + "0"
}

func main() {
	bob := Human{"bob", 39, "002-5456-xxx"}
	fmt.Println("this huaman is:", bob)
	//Println("the biggest one is", boxes.biggestcolor().String())
}

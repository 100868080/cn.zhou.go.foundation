package main

import "fmt"

type student struct {
	height, width float64
}

type stu struct {
	radius float64
}

func (r student) say() float64 {
	return r.width * r.height
}

func (st stu) say() float64 {
	return st.radius * st.radius * 3.14
}

func main() {

	jk := student{56.3, 223}
	m := stu{32}

	fmt.Println(jk.say(), m.say())
}

/*

package main

import (
	"fmt"
)

type re struct {
	width, height float64
}

func area(r re) float64 {
	return r.width * r.height
}

func main() {
	r := re{456, 23}
	fmt.Println(area(r))
}*/

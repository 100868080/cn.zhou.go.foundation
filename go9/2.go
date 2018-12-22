//接口的定义与赋值
package main

import (
	"fmt"
)

type People struct {
	Name string
}
type Teacher struct {
	People
	Depar string
}
type Student struct {
	People
	School string
}

func (p People) SayHi() {
	fmt.Printf("I,m %s .\n", p.Name)
}
func (t Teacher) SayHi() {
	fmt.Printf("my nmae is %s.I,m work in %s.\n", t.Name, t.Depar)
}
func (s Student) SayHi() {
	fmt.Printf("I,m %s I,m study int %s\n", s.Name, s.School)
}
func (s Student) Study() {
	fmt.Printf("I,m learning golang in %s.\n", s.School)
}

type Speaker interface {
	SayHi()
}
type Leaner interface {
	SayHi()
	Study()
}

func main() {
	people := People{"zhang san"}
	teacher := Teacher{People{"zheng zhi"}, "computer science"}
	stu := Student{People{"li ming"}, "yale universty"}
	var is Speaker
	is = people
	is.SayHi()
	is = teacher
	is.SayHi()
	is = stu
	is.SayHi()
	var il Leaner
	il = stu
	il.Study()
	fmt.Println()
	ix := make([]Speaker, 3)
	ix[0], ix[1], ix[2] = people, teacher, stu
	for _, value := range ix {
		value.SayHi()
	}
}

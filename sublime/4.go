package main

import (
	"errors"
	"fmt"
)

func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	m := ",world"
	a := s + m
	b := s2 + m
	fmt.Println(a, b)

	fmt.Printf("%s\n", s2)
	s = "c" + s[1:]
	fmt.Println(s)
	k := `helle
112  66 65 5   2 32 2 4 4 4 7 7 8 9 9
               world!`
	fmt.Println(k)

	err := errors.New("emit macho dwarf:elf header  corrupted")
	if err != nil {
		fmt.Println(err)
	}
	h, d := true, false
	fmt.Println(h, d)
}

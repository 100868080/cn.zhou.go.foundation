package main

import (
	"fmt"
)

func main() {
	type s []string
	type people struct {
		name string
		id   int
	}
	type teacher struct {
		s
		int
		people
		feil bool
	}

	t := teacher{}
	t.feil = false
	t.int = 131
	t.s = append(t.s, "asda", "adadaffa", "qweopqoiweioqiop", "qqwweoppa")
	t.name = "adsdq"
	t.id = 131
	fmt.Println(t)
}

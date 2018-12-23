package main

import (
	"fmt"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
	var apples int32 = 1
	var oranges int16 = 2
	var compote = apples + int32(oranges)
	fmt.Println(compote)

	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	o := 0666
	fmt.Printf("%d %[1]o\v %#[1]o\n", o)
	fmt.Println()
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

}

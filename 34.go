package main

import (
	"fmt"
)

func main() {
	f := 3.141 //a floa64
	i := int(f)
	fmt.Println(f, i) //"3.141 3"
	f = 1.99
	fmt.Println(int(f)) //"1

	o := 0666
	fmt.Printf("%d %[1]o\n", o) //"438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]x\n", x)
	// Output:
	//3735928559 deadbeef 0xdeadbeef 0XDEADCEEF

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   //"97 a 'a' "
	fmt.Printf("%d %[1]c %[1]q\n", unicode) //"22269 国 ‘国’ "
	fmt.Printf("%d %[1]q\n", newline)

}

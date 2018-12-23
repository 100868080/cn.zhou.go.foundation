package main

import (
	"fmt"
)

func main() {
	s := "hello,world!"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])
	sh := "hello,world"
	fmt.Println(len(sh))
	nfc := "6564654898"
	fmt.Println(len(nfc))
	fmt.Println(s[:3] + "\tkild")
	fmt.Println("hello,world!")
}

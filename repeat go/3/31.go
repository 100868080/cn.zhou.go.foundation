package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	if x := f(); x == 0 {
		fmt.Println(x)

	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	fmt.Println(x, y)

	if f, err := os.Open(fname); err != nil {
		return err
	}
	f.ReadByte()
	f.Close()
}

var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("os.Get failed:%v", err)
	}
}

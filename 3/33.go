package main

import (
	"fmt"
	"log"
	"os"
)

func f() {}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
	//fmt.Println(h)

	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
}

var cwd string

func init() {
	var err error

	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal("os.Getwd failed:%v", err)
	}
}

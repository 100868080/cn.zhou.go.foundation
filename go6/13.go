package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("Hello,world!")
	c := "Hdlr!"
	fmt.Println(string(bytes.Trim(s, c)))
}

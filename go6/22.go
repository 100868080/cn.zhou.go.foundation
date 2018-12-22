package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("中华人民共和国")
	var buf [3]byte
	e := bytes.NewReader(a)
	e.Read(buf[:])
	err := e.UnreadRune()
	fmt.Println(err)
	c, _, _ := e.ReadRune()
	fmt.Println(string(c))
	err = e.UnreadRune()
	fmt.Println(err)
	c, _, _ = e.ReadRune()
	fmt.Println(string(c))
}

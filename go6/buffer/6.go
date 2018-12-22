package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*b := []byte("hello ha ba gou!")
	a := bytes.NewBuffer(b)
	var p [8]byte
	n, err := a.Read(p[:])
	fmt.Println(string(p[:]), n, err)*/

	buf := []byte("lang gou!")
	a := bytes.NewBuffer(buf)
	c, err := a.ReadByte()
	fmt.Println(string(c), err)
}

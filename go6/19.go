package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*a := []byte("go lang")
	e := bytes.NewReader(a)
	c, err := e.ReadByte()
	fmt.Println(string(c), err)*/
	a := []byte("抓人民")
	f := bytes.NewReader(a)
	c, z, err := f.ReadRune()
	fmt.Println(string(c), z, err)
}

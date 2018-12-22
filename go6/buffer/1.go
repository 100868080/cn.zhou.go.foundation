package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*p := []byte("Hello,world!")
	b := bytes.NewBuffer(nil)
	n, err := b.Write(p)
	fmt.Println(string(b.Bytes()), n, err)*/

	var c1, c2 byte = 'G', 'o'
	a := bytes.NewBuffer(nil)
	err := a.WriteByte(c1)
	a.WriteByte(c2)
	fmt.Println(string(a.Bytes()), err)
}

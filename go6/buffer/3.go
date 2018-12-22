package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x, y byte = 'G', 'g'
	a := bytes.NewBuffer(nil)
	err := a.WriteByte(x)
	a.WriteByte(y)
	fmt.Println(string(a.Bytes()), err)
}

package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("the close, is gone!")
	a := bytes.NewBuffer(buf)
	var d byte = ','
	fmt.Println("wei zhi xing :")
	fmt.Println(string(a.Bytes()), a.Len())
	a.ReadString(d)
	fmt.Println("zhi xing:")
	fmt.Println(string(a.Bytes()), a.Len())
}

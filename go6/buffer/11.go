package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("golang is a beautiful language,i like it!")
	a := bytes.NewBuffer(buf)
	var d byte = ','
	fmt.Println("wei cao zuo buffer:")
	fmt.Println(string(a.Bytes()), a.Len())
	a.ReadString(d)
	fmt.Println("zhi xing cao zuo:")
	fmt.Println(string(a.Bytes()), a.Len())
}

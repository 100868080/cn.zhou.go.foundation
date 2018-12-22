package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*buf := []byte("kil z/hwu  adj?ka")
	t := bytes.NewBuffer(buf)
	var d byte = '/'
	line, err := t.ReadBytes(d)
	fmt.Println(string(line), err)*/

	buf := []byte("中华人民共和国")
	a := bytes.NewBuffer(buf)
	b := bytes.NewBuffer(nil)
	n, err := a.ReadFrom(a)
	fmt.Println(string(b.Bytes()), n, err)
}

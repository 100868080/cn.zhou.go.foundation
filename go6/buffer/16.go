package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("中华人,民共和国")
	b := bytes.NewBuffer(buf)
	var d byte = ','
	n, err := b.ReadBytes(d)
	fmt.Println(string(n), err)
	b.ReadRune()
	b.ReadRune()
	b.ReadRune()
	b.UnreadRune()
	fmt.Println(string(b.Bytes()))
}

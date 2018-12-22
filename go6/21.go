package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("golang")
	var buf [6]byte
	q := bytes.NewReader(a)
	err := q.UnreadByte()
	fmt.Println(err)
	n, _ := q.Read(buf[:])
	fmt.Println(string(buf[:n]), n)
	err = q.UnreadByte()
	fmt.Println(err)
	n, _ = q.Read(buf[:])
	fmt.Println(string(buf[:n]), n)
}

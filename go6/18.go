package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("golang")
	var buf [3]byte
	r := bytes.NewReader(a)
	n, err := r.ReadAt(buf[:], 2)
	fmt.Println(string(buf[:n]), n, err)
	n, err = r.ReadAt(buf[:], 3)
	fmt.Println(string(buf[:n]), n, err)
	n, err = r.ReadAt(buf[:], 4)
	fmt.Println(string(buf[:n]), n, err)
}

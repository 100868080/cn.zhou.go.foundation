package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("go lang")
	var buf [3]byte
	r := bytes.NewReader(a)
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

}

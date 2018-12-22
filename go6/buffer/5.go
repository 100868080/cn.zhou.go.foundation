package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("asdaggaasfas")
	t := bytes.NewBuffer(buf)
	a := bytes.NewBuffer(nil)
	n, err := t.WriteTo(a)
	fmt.Println(string(a.Bytes()), n, err)
}

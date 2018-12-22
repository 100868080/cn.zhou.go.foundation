package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("Go la ng")
	e := bytes.NewReader(a)
	fmt.Println(e.Len())
}

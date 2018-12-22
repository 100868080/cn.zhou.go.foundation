package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("size ,i like i>t.")
	a := bytes.NewBuffer(buf)
	var d byte = '>'
	line, err := a.ReadString(d)
	fmt.Println(line, err)
}

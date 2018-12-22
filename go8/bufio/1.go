package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("go lang")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)
}

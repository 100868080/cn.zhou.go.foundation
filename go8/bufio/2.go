package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("hello,world!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
}

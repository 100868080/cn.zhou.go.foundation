package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("zhong hua ren ming gong")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:]), n, err)
}

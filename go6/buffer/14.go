package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("Hello,world!")
	a := bytes.NewBuffer(buf)
	fmt.Println("wei zhi xing:", string(a.Bytes()))
	a.Truncate(1)
	fmt.Println("zhi xing:", string(a.Bytes()))
	a.Reset()
	fmt.Println("zhi xing reset:", string(a.Bytes()))
}

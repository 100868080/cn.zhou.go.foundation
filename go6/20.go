package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("中华人民共和国")
	b := bytes.NewReader(a)
	b.Seek(3, 0)
	c, v, err := b.ReadRune()
	fmt.Println(string(c), v, err)

	b.Seek(3, 1)
	c, v, err = b.ReadRune()
	fmt.Println(string(c), v, err)

	b.Seek(-9, 2)
	c, v, err = b.ReadRune()
	fmt.Println(string(c), v, err)
}

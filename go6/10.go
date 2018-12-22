package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("abc")
	b := []byte("ABC")
	s := []byte("GOLANG")
	t := []byte("golang")
	fmt.Println(bytes.Compare(a, t))
	fmt.Println(bytes.Equal(a, s))
	fmt.Println(bytes.EqualFold(a, b))
}

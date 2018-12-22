package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*var z, x rune = 's', 'k'
	a := bytes.NewBuffer(nil)
	a.WriteRune(z)
	a.WriteRune(x)
	fmt.Println(string(a.Bytes()))*/
	s := "你好 时间!"
	a := bytes.NewBuffer(nil)
	n, err := a.WriteString(s)
	fmt.Println(string(a.Bytes()), n, err)
}

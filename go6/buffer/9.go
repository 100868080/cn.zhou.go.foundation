package main

import (
	"bytes"
	"fmt"
)

func main() {
	//var r ,
	buf := []byte("adasdaasdasasdasdaasd中华人民共和国asdadasdaasdasdqweqweqwrqwreqweqwrttrdghd")
	g := bytes.NewBuffer(buf)
	r, d, err := g.ReadRune()
	fmt.Println(string(r), d, err)
}

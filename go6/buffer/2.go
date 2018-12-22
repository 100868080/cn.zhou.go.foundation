package main

import (
	"bytes"
	"fmt"
)

func main() {
	q := []byte("adsagga a   asaad   asda")
	r := bytes.NewBuffer(nil)
	n, err := r.Write(q)
	fmt.Println(string(r.Bytes()), n, err)
}

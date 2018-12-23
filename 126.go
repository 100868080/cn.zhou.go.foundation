package io

import (
	"fmt"
)

type Write interface {
	Write(p []byte) (n int, err error)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
var c ByteCounter
c.Write([] byte("hello"))

fmt.Println(c)
c=0
var name ="Dolly"

fmt.Fprintf(%c,"hello,%s",name)
fmt.Println(c)
package main

import (
	"fmt"
)

type ByteSize float64

const (
	_           = iota
	kb ByteSize = 1 << (10 * iota)
	m
	g
	t
	e
)

func main() {
	fmt.Println(kb, m, g, t, e)
}

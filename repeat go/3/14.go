package main

import (
	"fmt"
	//"os"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g。F=%g 。C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g 。F=%g 。C\n", boilingF, fToC(boilingF))

	type fg string
	var t fg
	t = "a"
	fmt.Println(t)
	var s string
	var r int
	var z bool
	fmt.Println(s, r, z)

	var i, j, k int
	fmt.Println(i, j, k)
	var jf, jk, jl = true, 23.15, "floa"
	fmt.Println(jf, jk, jl)
	var name []string
	fmt.Println(name)
	//var fh, err = os.Open(int)
	//fmt.Println(fh)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

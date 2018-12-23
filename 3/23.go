package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f)
	fmt.Println(c == Celsius(f))

	C := FToC(212.0)
	fmt.Println(C.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

}

func (c Celsius) String() string { return fmt.Sprintf("%goC", c) }

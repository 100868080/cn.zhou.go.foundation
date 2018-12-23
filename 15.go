//Ftoc prints two Fahrenheit-to-Calsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g。f=%g。c\n", freezingF, fToc(freezingF)) //"32。F=0。c"
	fmt.Printf("%g。F=%g。c\n", boilingF, fToc(boilingF))   //"212。F=100。C"
}
func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}

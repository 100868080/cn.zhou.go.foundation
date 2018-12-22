package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "January" /* ...*/, 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     //["April "May" "June"]
	fmt.Println(summer) //["June" "July" "August"]
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Println(summer[:20])    //panic :out of range
	endlessSummer := summer[:5] //extend a slice (within capacity)
	fmt.Println(endlessSummer)  //"[June July August September October]"
}

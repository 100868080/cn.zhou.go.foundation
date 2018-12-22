package main

import (
	"fmt"
)

func main() {
	type Currency int
	const (
		USD Currency = iota //美元
		EUR                 //欧元
		GBP                 //英镑
		RMB                 //人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) //"3 ￥"

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) //"true false false"
	d := [3]int{1, 2}
	fmt.Println(a == d) //compile erroe:cannot compare [2]int==[3]int
}

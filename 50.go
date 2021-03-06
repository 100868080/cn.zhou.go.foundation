package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))             //"123 123"
	fmt.Println(strconv.FormatInt(int64(x), 2)) //"1111011"
	s := fmt.Sprintf("x=%b", x)                 //"x=1111011"
	fmt.Println(s)
	l, err := strconv.Atoi("123")             //x is an int
	k, err := strconv.ParseInt("123", 10, 64) //base 10,up to 64 bits

}

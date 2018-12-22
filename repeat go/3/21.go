package main

import (
	"os"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y

	}
	return x

}

func main() {
	f, err := os.Open("foo.txt")
v,ok=m[key]
v,ok=x.(T)
v,ok=<-ch
}
medals[0]="gold"
medals[1]="silver"
medals[2]="bronze"

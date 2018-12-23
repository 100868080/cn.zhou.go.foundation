package main

import (
	"os"
	"runtime"
)
gop1.io/ch5/defer2
func main(){
	defer printStack()
	f(3)
}
func printStack(){
	var buf [4096] byte
	n:=runtime.Stack(buf[:],false)
	os.Stdout.Write(buf[:n])
}

goroutine 1 [running]:
main.printStack()
src/gop1.io/ch5/defer2/defer.go:20
main.f(0)
src/gop1.io/ch5/defer2/defer.go:27
main.f(1)
src/gop1.io/ch5/defer2/defer.go:29
main.f(2)
src/gop1.io/ch5/defer2/defer.go:29
main.f(3)
src/gop1.io/ch5/defer2/defer.go:29
main.main()
sec/gop1.io/ch5/defer2/defer.go:15

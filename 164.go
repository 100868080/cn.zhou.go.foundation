package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j <- tick
	}
	launch()
}
abort:=make(chan struct{})
go func(){
	os.Stdin.Read(make([]byte,1))
	abort<-struct{} {}
}()

select{
	case <-ch1:
	//...
	case x:=<-ch2:
	//...use x...
	case ch3<-y:
	//...
	default:
	//...
}
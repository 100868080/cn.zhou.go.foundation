package main

import (
	"time"
	"fmt"
)
fmt.Println("Commencing countdown. Parse return to abort.")
select{
	case <-time.After(10*time.Second):
	//Do nothing.
	case <-abort:
	fmt.Println("Launch aborted!")
	return
}
launch()
}
ch:=make(chan int,1)
for i:=0;i<10;i++{
	select{
		case x:=<-ch:
		fmt.Println(x)
		case ch<-i:
	}
}
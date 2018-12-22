package main

import (
	"time"
	"fmt"
)
func main(){
	fmt.Println("Commencing countdown.Parse return to abort.")
tick:=time.Tick(1*time.Second)
for countdown:=10;countdown>0;countdown---{
	fmt.Println(countdown)
	select{
		case <-tick:
		//Do thing.
		case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
}
launch()

}
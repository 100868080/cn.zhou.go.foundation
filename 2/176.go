package main

import (
	"fmt"
	"sync"
)
var mu sync.RWMutex
var balance int
func Balance()int{
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
var x,y int
go func(){
	x=1 
	fmt.Print("y:",y,"")
	
}()

go func(){
	y=1
	fmt.Print("x:",x,"")
}()
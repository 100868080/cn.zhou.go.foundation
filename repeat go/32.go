package main

import (
	"fmt"
)
var u uint8=255
fmt.Println(u,u+1,u*u)  //"255 0 1"
var i int8=127
fmt.Println(i,i+1,i*i)  //"127 -128 1"

var x uint8=1<<1|1<<5
var y uint8=1<<1|1<<2

fmt.Printf("%0.8b\n",x)      //"00100010",the set{1,5}

fmt.Printf("%0.8b\n",y)      //"00000110",the set{1,2}
fmt.Printf("%0.8b\n",x&y)   //"000000010",the intersection{1}
fmt.Printf("%0.8b\n",x|y)   //"001000110",the union{1,2,5}
fmt.Printf("%0.8b\n",x^y)   //"00100100",the symmetric {2,5}
fmt.Printf("%0.8b\n",x&^y)    //"0010000",the difference{5}
for i:=uint(0);i<8;i++{
	if x&(1<<i)!=0{//membership test
	fmt.Println(i)   //"1","5"
	}
}
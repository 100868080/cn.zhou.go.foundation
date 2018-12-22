package main

import (
	"fmt"
)

func main() {
	summer := []int{}
	if summer == nil { /*...*/
	}
	var s []int    //len(s)==0,s==nil
	s = nil        //len(s)==0,s==nil
	s = []int(nil) //len(s）==0，s==nil
	s = []int{}    //len(s)==0,s!=nil
	fmt.Println(s)
}

package main

import (
	"fmt"
)

func main() {
	var t = [5]int{1, 2, 3, 60}
	//fmt.Println(t[3] * t[2])
	var h int
	h = max(t[1], t[3])
	fmt.Println("%d",h)
}
func max(t[1],t[3] int)int{
	var g int
	if t[1]>t[3]{
		g=t[1]
	}else{
	g=t[3]
}
return g
}
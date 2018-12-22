package main

import (
	"fmt"
)
func preOrder(n*Node){
	if n!=nil{
		fmt.Printf("%v",n.Data)
		preOrder(n.Left)
		preOrder(n.Right)
}
func InOrder (n*Node){
	if n!=nil{
		preOrder(n.Left)
		fmt.Printf("%v",n.Data)
	}
}
func postOrder(n*Node){
	preOrder(n.Left)
	preOrder(n.Right)
	fmt.Printf("%v",n.Data)
}

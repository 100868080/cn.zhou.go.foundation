package main

import (
	"fmt"
)

func printBT(n *Node) {
	if n != nil {
		fmt.Printf("%v", v.Data)
		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			printBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			printBT(n.Right)
			fmt.Printf(")")
		}
	}
}

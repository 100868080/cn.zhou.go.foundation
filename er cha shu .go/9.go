package main

import (
	"fmt"
)

func main() {
	//创建二叉树
	root := NewNode(nil, nil)
	root.SetData("root node")
	a := NewNode(nil, nil)
	a.SetData("left node")
	al := NewNode(nil, nil)
	al.SetData(100)
	ar := NewNode(nil, nil)
	ar.SetData(3.14)
	a.Left = al
	a.Right = ar
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.Left = a
	root.Right = b
	//使用operater接口实现对二叉树的基本操作
	var it Operater
	it = root
	it.PrintBT()
	fmt.Println()
	fmt.Println("The depths of Btreeis :", it.Dapth())
	fmt.Println("the leaf counts of the btree is:", it.LeafCount())
}

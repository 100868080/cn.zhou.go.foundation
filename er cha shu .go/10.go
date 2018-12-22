//二叉树的遍历
package main

import (
	"fmt"
)
//创建二叉树
root:=NewNode(nil,nil)
root.SetNode("root node")
a:=NewNode(nil,nil)
a.SetData("left node")
al:=NewNode(nil,nil)
al.SetData(100)
ar:=NewNode(nil,nil)
ar.SetData(3.14)
a.Left=al
a.Right=ar
b:=NewNode(nil,nil)
b.SetData("right node")
root.Left=a
root.Right=b
//使用order接口实现对二叉树的遍历
var it Order
it=root
it.PreOrder()  //先序遍历
fmt.Println()
it.InOrder()  //中序遍历
fmt.Println()
it.PostOrder()  //厚序遍历
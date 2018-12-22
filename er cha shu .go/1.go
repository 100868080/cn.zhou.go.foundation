package main

import (
	"link"
)

func main() {
	var head *link.Node
	stu1 := link.Node{link.Student{100, "Liming"}, nil}
	stu2 := link.Node{link.Student{101, "彰显"}, nil}
	stu3 := link.Node{link.Student{102, "李宁"}, nil}
	stu4 := link.Node{link.Student{103, "当挖"}, nil}
	//创建新链表
	head = head.Creat()
	head = stu1.Insert(head)
	head = stu2.Insert(head)
	head = stu3.Insert(head)
	head = stu4.Insert(head)
	//输出链表
	head.printlink()
	//删除节点
	head = stu3.dealete(head)
	head.printlink
}

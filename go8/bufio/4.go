package main

import (
	"link"
)

func main() {
	var head *line.node
	s1 := link.node{link.student{100, "li ming"}, nil}
	s2 := link.node{link.student{101, "zhangqiong"}, nil}
	s3 := link.node{link.student{102, "zhao xiao"}, nil}
	head = head.creat()
	head = s1.insert(head)
	head = s2.insert(head)
	head.printlink()
	head = s3.delete(head)
	head.printlink()
}

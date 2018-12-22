//链表综合操作
package main

import (
	"fmt"
)

type student struct {
	id   int
	name string
}
type node struct {
	student
	next *node
}

func (head *node) creat() *node {
	head = nil
	return head
}
func (p *node) printlink() {
	for p != nil {
		fmt.Printf("%d,%s\n", p.id, p.name)
		p = p.next
	}
}
func (newnode *node) insert(head *node) *node {
	var p0, p1, p2 *node
	p0 = newnode
	p1 = head
	if head == nil {
		head = p0
		p0.next = nil
	} else {
		for (p0.id > p1.id) && p1.next != nil {
			p2 = p1
			p1 = p1.next
		}
		if p0.id <= p1.id {
			if head == p1 {
				head = p0
			} else {
				p2.next = p0
				p0.next = p0
			}
		} else {
			p1.next = p0
			p0.next = nil
		}
	}
	return head
}
func (delnode *node) delete(head *node) *node {
	var p1, p2 *node
	if head == nil {
		fmt.Println("list nil!")
		goto end
	}
	p1 = head
	for delnode.id != p1.student.id && p1.next != nil {
		p2 = p1
		p1 = p1.next
	}
	if delnode.id == p1.student.id {
		if p1 == head {
			head = p1.next
		} else {
			p2.next = p1.next
		}
		fmt.Printf("delete %d\n", delnode.id)
	} else {
		fmt.Printf("node %d not been found!\n", delnode.id)
	}
end:
	return head
}

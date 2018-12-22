package main

import (
	"fmt"
)

func main() {
	var usr = struct {
		id       int
		name     string
		password string
	}{name: "张三"}
	fmt.Println(usr)
	var usr2 = new(struct {
		id   int
		name string
	})
	usr2.id = 132
	fmt.Println(usr2)
}

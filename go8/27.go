package main

import (
	"fmt"
	"mystruct"
)

func main() {
	usr := new(mystruct.User)
	usr.Id = 100
	usr.Name = "zhangsan"
	usr.Sayhi()
	fmt.Println(usr)
}

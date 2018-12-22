package main

import (
	"crypto/rand"
	"fmt"
)

func Test(ch chan int) {
	ch <- rand.Int() //向channel中写入一个随机数
	fmt.Println("Go...")
}
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Test(chs[i]) //启动10个goroutine
	}
	for _, ch := range chs {
		value := <-ch //阻塞等待退出信号
		fmt.Println(value)
	}
}

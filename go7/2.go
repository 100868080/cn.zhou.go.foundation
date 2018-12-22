package main

import (
	"fmt"
	"time"
)

func main() {
	//获取时间戳
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006年1月2日"))
	fmt.Println(t.Weekday().String())
	fmt.Println(t.ISOWeek())
}

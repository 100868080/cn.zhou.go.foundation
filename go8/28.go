//字段标签:
package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int "hong cold:"
	Name string
	Sex  bool
}

func main() {
	u := user{100, "zhangsan", false}
	//使用typeof函数获取对象的类型
	t := reflect.TypeOf(u)
	//使用value（）获取对象值
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s(%s=%v)\n", f.Tag, f.Name, v.Field(i).Interface())
	}
}

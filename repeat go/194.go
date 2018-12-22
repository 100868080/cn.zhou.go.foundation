package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}


t:=reflect.TypeOf(3)
fmt.Println(t.String())
fmt.Println(t)

v:=reflect.ValueOf(3)
fmt.Println(v)
fmt.Printf("%v\n",v)

fmt.Println(v.String())
v:=reflect.ValueOf(3)
x:=v.Interface()
i:=x.(int)
fmt.Printf("%d\n",i)

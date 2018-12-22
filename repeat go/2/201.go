package main

import (
	"fmt"
	"reflect"
)

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumField(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s %s\n", t, t.Method(i).Name, strins.TrimPrefix(methType.String(), "func"))
	}
}

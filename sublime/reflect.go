package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	tag := t.Elem().Field(0).Tag
	name := v.Elem().Field(0).String()
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type :", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

}

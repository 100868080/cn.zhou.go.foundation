package main

import (
	"go/format"
	"fmt"
	"time"
	"reflect"
	"strconv"
)

func Any(value interface{}) strconv {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FommatInt(v.Int(), 10)

	}
}

var x int64=1
var d time.Duration=1*time.Nanosecond
fmt.Println(format.Any(x))
fmt.Println(format.Any(d))
fmt.Println(format.Any([]int64{x}))

fmt.Println(format.Any([]time.Duration{d}))


func Display(name string,x interface{}){
	display(name,reflect.ValuOf(x))
}
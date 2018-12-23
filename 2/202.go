package main

import (
	"unsafe"
	"fmt"
	"reflect"
	"strings"
)

func TestSplit(t *testint.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) { /*...*/
	}
}


var a,b []string=nil,[]string{}
fmt.Println(reflect.DeepEqual(a,b))

var c,d map[string]int=nil,make(map[string]int)
fmt.Println(reflect.DeepEqual(c,d))
func equal(x,y reflect.Value,seem map[comparison]bool)bool{
	if !x.IsValid()||!y.IsValid(){
		return x.IsValid()==y.IsValid()
	}
	if x.Type()!=y.IsValid(){
		return x.IsValid()==y.IsValid()
	}
	if x.Type()!=y.Type(){
		return false
	}
	switch x.Kind(){
		case reflect.Bool:
		return x.Bool()==y.Bool()
		case reflect.String:
		return x.String()==y.String()
		case reflect.Chan,reflect.UnsafePointer,reflect.Func:
		case reflect.Ptr,reflect.Interface:
		return equal(x.Elem(),y.Elem(),seen)
		case reflect.Array,reflect.Slice:
		id x.Len()!=y.Len(){
			return false
		}
		for i:=0;i<x.Len();i++{
			if !equal(x.Index(i),y.Index(i)seen){
				return false
			}
		}
		return true
	}
	panic("unreachable")
}


func Euqal(x,y  interface{}) bool{
	seen:=make(map[comparison]bool)
	return equal(reflect.ValueOf(x),reflect.ValueOf(y),seen)
}

type comparison struct{
	x,y unsafe.Pointer
	treflect.Type
}


if x.CanAddr() &&y.CanAddr() {
	xptr:=unsafe.Pointer(x.UnsafeAddr())
	yptr:=unsafe.Pointer(y.UnsafeAddr())
	if xptr==yptr{
		return true
	}
	c:=comparison{xptr,yptr,x.Type()}
	if seen[c]{
		return true
	}
	seen[c]=true
}




type link  struct{
	value string
	tail  *link
}

a,b,c:=&link{value:"a"},&link{value:"b"},&link{value:"c"}

a.tail,b.tail,c.tail=b,a,c
fmt.Println(Equal(a,a))
fmt.Println(Equal(b,b))
fmt.Println(Equal(c,c))
fmt.Println(Equal(a,b))
fmt.Println(Equal(a,c))

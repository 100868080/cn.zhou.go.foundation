package main
import "fmt"

x:=1
p:=&x  //p,of type*int,point to x
fmt.Println(*p)  //"1"
*p=2
fmt.Println(x)   //"2"

var x,y int
fmt.Println(&x==&x,&x==&y,&x==nil)   //true false false

var p=f()
func f()*int{
	v:=1
	return &v
}


func incr(p*int) int{
	*p++   //非常重要 ：只是增加p指向的变量的值，并不改变p指针!!!
	return *p
}
v:=1
incr(&v)  //side effect:v is now 2
fmt.Println(incr(&v))   //"3"(and v is 3)
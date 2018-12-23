package main

import (
	"io"
	"os"
)
import(
	"fmt"
)
x=1     //命名变量的赋值
*p=true   //通过指针间接赋值
person.name="bob"   //结构字段赋值
count[x]=count[x]*scale  //数组、slice或map的元素赋值

count[x]*=scale
  
v:=1
v++  //等价方式 v=v+1;v变成2
v--   //等价方式v=v-1;v变成1

x,y=y,x
a[i],a[j]=a[j],a[i]


func gcd(x,y int)int{
	for y!=0{
		x,y=y,x%y
	}
	return x
}
func fib(n int)int{
	x,y:=0,1
	for i:=0;i<n;i++{
		x,y=y,x+y
	}
	return x
}


f, err=os.Open("foo.txt")  //function call retutns two values

v,ok=m[key]  //map lookup
 v,ok=x.(T)   //type assertion
v,ok=<-ch     //channel receive


v=m[key]   //map查找，失败时返回零值
v=x.(T)    //type断言，失败时panic异常
v=<-ch    //管道接收，失败时返回零值（阻塞不算是失败

_,ok=m[key]   //map返回2个值
_,ok=mm[""],false   //map返回1个值
_=mm[""]         //map返回1个值


_,err=io.Copy(dst,src)  //丢弃字节数
_,ok=x.(T)     //只检测类型，忽略具体值



medals:=[]string{"gold","silver","bronze"}

medals[o]="gold"
medals[1]="silver"
medals[2]="bronze"
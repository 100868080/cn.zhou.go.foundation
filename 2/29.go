package main

import (
	"os"
	"fmt"
)
if x:=f();x==0{
	fmt.Println(x)
}else if y:=g(x);x==y{
	fmt.Println(x,y)
}else{
	fmt.Println(x,y)
}
fmt.Println(x,y)    //compile error:x and y are not visible here
 
if f,err:=os.Open(fname);err!=nil{//compile error:unnsed:f return err}
f.ReadByte()
f.Close()
 
f,err:=os.Open(fname)
if err!=nil{
	return err
}
f.ReadByte()
f.Close()

if f,err:=os.Open(fname);err!=nil{
	return err
}else{
	//f and err are visible here too
	f.ReadByte()
	f.Close()
}


var cwd string
func init(){
	cwd,err:=os.Getwd()  //compile error:unused:cwd
	if err!=nil{
		log.Fatalf("os.Getwd failed:%v",err)
	}
}
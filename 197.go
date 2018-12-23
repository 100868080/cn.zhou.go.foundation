package main

import (
	"fmt"
	"reflect"
)
func display(path string,v reflect.Value){
	switch v.Kind(){
		case reflect.Invalid:
		fmt.Printf("%s=invalid\n",path)
		case reflect.Slice,reflect.Array:
		for i:=0;i<v.Sprintf("%s[%d]",path,i),v.Index(i)
	}
	case reflect.Struct:
	for i:=0;i<v.NumField();i++{
		fieldPath:=fmt.Sprintf("%s.%s",path,v.Type().Field(i).Name)
		display(fieldPath,v.Field(i))
	}
	case reflect.Map:
	for _,key:=range v.MapKeys(){
		display(fmt.Sprintf("%s[%s]",path,formatAtom(key)),v.apIndex(key))
	}
	case reflect.Ptr:
	if v.IsNil({
		fmt.Printf("%s=nil\n",path)
	}else{
		display(fmt.Sprintf("(*%s)",path),v.Elen())
	}
	case reflect.Interface:
	if v.IsNil(){
		fmt.Printf("%s=nil\n",path)
	}else{
		fmt.Printf("%s.type=%s\n",path,v.Elen().Type())
		display(path+".value",v.Elen())
	}
	default:
	fmt.Printf("%s=%s\n",path,formatAtom(v))
	}
}

type Movie struct{
	Title,Subtitle string
	Year int
	Color bool
	Actor map[string]string
	Oscars  []string
	Sequel  *string
}
type Cycle struct {Value int;Tail *Cycle}
var c Cycle 
c=Cycle{42,&c}
Display("c",c)

Display c (display.Cycle):
c.Value=42
(*c.Tail).Value=42
(*(*(*c.Tail).Tail).Tail).Vlaue=42
...ad infinitum...

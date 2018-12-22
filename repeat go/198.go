package main

import (
	"encoding/json"
	"os"
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Ptr:
	return encode(buf,v.Elen())
	case reflect.Map:
	buf.WriteByte('')
	for i,key:=range v.MapKeys(){
		if i>0{
			buf.WriteByte('')
		}
		buf.WriteByte('(')
		if err:=encode(buf,key);err!=nil{
			return err
		}
		buf.WriteByte('')
		if err:=encode(buf,v.MapIndex(key));err!=nil{
			return err
		}
		buf.WriteByte(')')
	}
	buf.WriteByte(')')
	default:
	return fmt.Errorf("unspported types:%s",v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte,error){
	var err:=encode(&buf,reflect.ValueOf(v));err!=nil{
		return nil,err
	}
	return buf.Bytes(),nil
}


fmt.Println(a.CanAddr())   使用ＣａｎＡｄｄｒ方法判断是否可以被取地址

x:=2
d:=reflect .ValueOf(&x).Elem()
px:=d.Addr().Interface().(*int)
*px=3
fmt.Println(x)

d.set(reflect.ValueOf(4))
fmt.Println(x)

d.set(reflect.ValueOf(int64(5)))

x:=2
b:=reflect.ValueOf(x)
b:=Set(reflect.ValueOf(3))

d:=reflect.ValueOf(&x).Elem()
d.SetInt(x)


x:=1
rx:=reflect.ValueOf(&x).Elem()
rx.SetInt(2)
rx.Set(reflect.ValueOf(3))
rx.SetString("hello")
rx.Set(reflect.ValueOf("hello"))

var y interface{}
ry:=reflect.ValueOf(&y).Elem()
ry.SetInt(2)
ry.Set(reflect.ValueOf(3))
ry.SetString("hello")

ry.Set(reflect.ValueOf("hello"))

stdout:=reflect.ValueOf(os.Stdout).Elem()
fmt.Println(stdout.Type())
fd:=stdout.FieldByName("fd")
fmt.Println(fd.Int())
fd.SetInt(2)

data:=[]byte{/*...*/}
var movie Movie 
err:=json.Unmarshal(data,&movie)


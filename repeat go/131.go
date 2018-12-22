package  main

import (
	"fmt"
	"time"
	"bytes"
	"os"
	"io"
)
var w io.Writer
w=os.Stdout
w=new(bytes.Buffer)
w=nil

var w io.Writer
os.Stdout.Write([] byte("helo"))
w:=new(bytes.Buffer)

var x interface{} =time.Now()

var x interface{} =[] int{1,2,3}
fmt.Println(x==x)

var w io.Writer
fmt.Printf("%T\n",w)
w=os.Stdout
fmt.Printf("%T\n",w)
w=new(bytes.Buffer)
fmt.Printf("%T\n",w)

const debug=true
func main(){
	var buf *bytes.Buffer
	if debug{
		buf=new(bytes.Buffer)
	}
	f(buf)
	if debug{
		//..use buf...
	}
}

func f(out io.Write){
	if out!=nil{
		out.Write([]byte("done!\n"))
	}
}

if out!=nil{
	out.Write([]byte("done!\n"))
}

var buf io.Writer
if debug{
	buf=new(bytes.Buffer)
}
f(buf)
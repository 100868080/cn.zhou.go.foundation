package main
import (
	"os"
	"fmt"
	"crypto/rand"
)
var s string
fmt.Println(s)

var i,j,k int   //int,int,int
var b,f,s=true,2.3,"four"   //bool,float64,string

var f,err=os.Open(name) //os.Open return a file and an error

anim:=gif.GIF{loopCount:nframes}
freq:=rand.Float64()*3.0
t:=0.0

i:=100  //a int
var boiling float64=100 //a int
var names []string
var err error
var p point

f,err:=os.Open(name)
if err!=nil{
	return err
}
// ...use f..
f.Close()


in,err:=os.Open(infile)
//...
out,err:=os.Create(outfile)

f,err:=os.Open(infile)
//...
f,err:=os.Create(outfile)   //compile error:no new variables
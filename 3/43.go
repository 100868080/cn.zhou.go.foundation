package main 
import (
	"unicode/utf8"
)
func main() {
	s:="Hello,shijie"
	fmt.println(len(s))
	fmt.println(utf8.RuneCountIntString(s))
}

for i:=0;i<len(s);{
	r,size:=utf8.DecodeRuneIntString(s[i:])
	fmt.printf("%d\t%c\",i,r)
i+=size
for i,r:=range "Hello,shijie"{
	fmt.printf("%d\t%q\t%d\n",i,r,r)
}

n:=0
for _,_=range s{
	n++
}

n:=0
for range s{
	n++
}
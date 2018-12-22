package regexp

import (
	"fmt"
)
func Compile(expr string)(*Regexp,error){/* ... */}
func MustCompile(expr string)*Regexp{
	re,err:Compile(expr string) *Regexp{
		re,err:=Compile(expr)
		if err!=nil{
			panic(err)
		}
		return re 
	}
}
var httpSchemeRE=regexp.MustCompile(`^http?:`)  
func main(){
	f(3)
}
func f(x int){
	fmt.Printf("f(%d)\n",x+0/x)
	defer fmt.Printf("defer %d\n",x)
	f(x-1)
}
panic:runtime error:integer divide by zero
main.f(0)
src/gopl.io/ch5/defer1/defer.go:14
main.f(1)
src/gop1.io/ch5/defer1/defer.go:16
main.f(2)
src/gop1.io/ch5/defer1/defer.go:16
main.f(3)
src/gop1.io/ch5/defer1/defer.go:16
main.main()
src/gop1.io/ch5/defer1/defer.go:10
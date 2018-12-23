package ioutil

import (
	"fmt"
	"log"
	"time"
	"sync"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}
var mu sync.Mutex
var m=make(map[string]int)
func lookup(key string)int{
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func bigSlowOperation(){
	defer trace("bigSlowOperation")()   //don't forget the 
	extra parentheses
	time.Sleep(10*time.econd)
	operation by sleeping
}
func trace(msg string)func() {
	sart:=time.Now()
	log.Printf("enter %s",msg)
	return func(){
		log.printf("exit %s (%s)",msg,time.Since(start))
	}
}

func double (x int)int{
	return x+x
}

func double(x int)(result int){
	defer func()  {fmt.printf("double(%d)=%d\n",x,result)}()
	return x+x
}

_=double(4)

//Output:
//"double(c)=8"

func triple(x int)(result int){
	defer func() {result+=x}()
	return double(x)
}
fmt.println(triple(4))   //12
package memo

import (
	"sync"
	"fmt"
	"log"
	"time"
)

type Memo struct {
	f     Func
	cache map[string]result
}
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
m:=memo.New(httpGetBody)
for url:=range incomingURLs(){
	start:=time.Now()
	value,err:=m.Get(url)
	if err!=nil{
		log.Print(err)
	}
	fmt.Printf("%s,%s,%d bytes\n",url,time.Since(start),len(value.([]byte)))
}

m:=memo.New(httpGetBody)
var n sync.WaitGroup
for url:=range incomingURLs(){
	n.Add(1)
	go func(url string){
		start:=time.Now()
		value,err:=m.Get(url)
		if err!=nil{
			log.Print(err)
		}
		fmt.Printf("%s,%s,%d bytes\n",url,time.Since(start),len(value.([]byte)))
		n.Done()
	}(url)
}
n.Wait()

func (memo *Memo) Get(key string) (interface{},error){
	res,ok:=memo.cache(key)
	if !ok{
		res.value,res.err=memo.f(key)
		memo.cache[key]=res
	}
	return res.value,res.err
}

type Memo struct{
	f Func 
	mu sync.Mutex
	cache map[string]result
}
func (memo *Memo)Get(key string) (value interface{},err error) {
	res,ok:=memo.cache[key] if !ok{
		res.value,res.err=memo.f(key)
		memo.cache[key]=res
		memo.mu.Lock()
		res,ok:=memo.cache[key]
		if !ok{
			res.value,res.err=memo.f(key)
			memo.cache[key]=res
		}
		memo.mu.Unlock()
		return res.value,res.err 
	}
}

func (memo *Memo)Get(key string) (value interface{},err error){
	memo.mu.Lock()
	res,ok:=memo.cache[key]
	memo.mu.Unlock()
	if !ok{
		res.value,res.err=memo.f(key)
		memo.mu.Lock()
		memo.cache[key]=res
		memo.mu.Unlock()
	}
	return res.value,res.err
}

type entry struct{
	res result
	ready chan struct{}
}

func New(f Func )*Memo{
	return &Memo{f:f,cache:make(map [string]*entry)}
}
type Memo string{
	f Func
	mu sync.Mutex
	cache map[string]*entry
}
func (memo *Memo)Get(key string)(value interface{},err error){
	memo.mu.Lock()
	e:=memo.cache[key]
	if e==nil{
		e=&entry{ready:make(chan struct{})}
		memo.cache[key]=e
		memo.mu.Unlock()
		e.res.value,e.res.err=memo.f(key)
		close(e.ready)
	}else{
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value,e.res.err
}

type Func func(key string) (interface{},error)
type result struct{
	value interface{}
	err error
}

type entry struct{
	res result
	ready chan struct{}
}

 type request struct{
	 key string
	response chan <-result
}

type Memo struct{requests chan request}
func New(f Func) *Memp{
	memo:=&Memo{requests:make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo)Get(key string) (interface{},error)  {
	response:=make(chan result)
	memo.request<-request{key,response}
	res:=<-response
	return res.value,res.err
}

func (memo *Memo) close()  {close(memo.requests)}

func (memo *Memo)server(f Func){
	cache:=make(map[string]*entry)
	for req:=range memo.requests{
		e:=cache[req.key]
		if e==nil{
			e=&entry{ready:make(chan struct{})}
			cache[req.key]=e
			go e.call(f,req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func,key string){
	e.res.value,e.res.err=f(key)
	close(e.ready)
}

func(e *entry)deliver(response chan <-result){
	<-e.ready
	response <-e.res
}

for {
	go fmt.Print(0)
	fmt.Print(1)
}
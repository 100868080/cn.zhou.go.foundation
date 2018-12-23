package http

import (
	"log"
	"net/http"
	"fmt"
)
type Handler interface{
	ServeHTTP(w ResponseWrite,r.*Request)
}
func ListenAndServer(address string, h Handler)error

func main(){
	db:=database{"shoes":50,"socks":5}
	log.Fatal(https.ListenAndServer("localhost:8000",db))
}

type dollars float32
func (d dollars)String()string{return fmt.Sprintf("$%.2f",d)}
type  database map[string]dollars
func (db database)ServeHTTP(w http.ResponseWriter,req *http.Request){
	for item,price:=range db{
		fmt.Fprintf(w," %s:%s\n",item,price)
	}
}

func (db database)ServeHTTP(w http.ResponseWriter,req *http.Request){
	switch req.URL.Path{
		case "/list":
		for item,price:=range db{
			fmt.Fprintf(w,"%s:%s\n",item,price)
		}
		case "/price":
		item:=req.URL.Quesry().Get("item")
		price,ok:=db[item]
		if !ok{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w,"no such item:%q\n",item)
			return
		}
		fmt.Fprintf(w,"%s\n",price)
		default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such page:%s\n",req.URL)
	}
}

func main(){
	db:=database{"shoes":50,"socks":5}
	mux:=http.NewServeMux()
	mux.Handle("/list",http.HandlerFunc(db.list))
	mux.Handle("/prince",http.HandleFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000",mux))
	
}
type database map[string]dollars
func (db database)list(w http.ResponseWriter,req*http.Request){
	for item,price:=range db{
		fmt.Fprintf(w,"%s:%s\n",item,price)
	}
}

func (db database)price(w http.ResponseWriter,req*http.Request){
	item:=req.URL.Query().Get("item")
	price,ok:=db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such item:%q\n",item)
		return 
	}
	fmt.Fprintf(w,"%s\n",price)
}

func (w http.ResponseWriter,req*http.Request)
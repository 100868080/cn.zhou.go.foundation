package http

import (
	"log"
	"net/http"
)

type HandlerFunc func(w RrsponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
}
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	htpp.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

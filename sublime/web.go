package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func say(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("sad", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "")

	}
	fmt.Fprintf(w, "heelooas100012do!")
}

func main() {
	http.HandleFunc("/", say)
	err := http.ListenAndServe(":100", nil)
	if err != nil {
		log.Fatal("list:", err)
	}
}

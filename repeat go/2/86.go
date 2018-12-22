package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if err != nil {
		fmt.Printf(os.Stderr, "findlinks1:%v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

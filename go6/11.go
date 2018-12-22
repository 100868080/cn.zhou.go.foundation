package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("I'm a student to day")
	for _, f := range bytes.Fields(s) {
		fmt.Printf("%q\n", f)
	}

}

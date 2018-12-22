package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var golang string = "the goler is a wonder!"
	var num int
	buf := make([]byte, 1024)
	f, _ := os.Open("golang")
	defer f.Close()
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]), n)
	for _, sentence := range bytes.FieldsFunc(buf[:n], f1) {
		for _, words := range bytes.Fields(sentence) {
			num++
			fmt.Printf("%q", words)
		}
	}
	fmt.Println()
	fmt.Println(f.Name(), num, "ge dan ci")
}
func f1(a rune) bool {
	if a == ',' || a == '.' || a == '!' {
		return true
	} else {
		return false
	}

}

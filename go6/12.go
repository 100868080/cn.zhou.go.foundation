package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("黄瓜,南瓜，西红柿,茄子,sudo apt-get install language-pack-kde-zh-hans-base")
	sep := []byte(",")
	n := 2
	for _, f := range bytes.SplitN(s, sep, n) {
		fmt.Printf("%q\n", f)
	}
	fmt.Println()
	for _, f := range bytes.SplitN(s, sep, 0) {
		fmt.Printf("%q\n", f)
	}
	fmt.Println()
	for _, f := range bytes.SplitN(s, sep, -1) {
		fmt.Printf("%q\n", f)
	}
}

//sudo apt-get install language-pack-kde-zh-hans-base

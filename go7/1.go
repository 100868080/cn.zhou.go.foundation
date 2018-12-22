package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int = 3, 2
	f, err := f1(a, b)
	fmt.Println(f, err)
}
func f1(a, b int) (ret float32, err error) {
	if b == 0 {
		err = errors.New("overflow")
		return
	} else {
		return float32(a) / float32(b), nil
	}

}

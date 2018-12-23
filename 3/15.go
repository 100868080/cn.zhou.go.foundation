package main

import (
	"fmt"
	/*"crypto/rand"
	"fmt"
	"image/gif"
	"os"*/)

func main() {

	/*anim := gif.GIF{LoopCount: nframes}
	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Println(anim, freq, t)

	*/

	/*f, err := os.Open(name)
	if err != nil {
		return err
	}
	*/
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	x = 3 * (*p)
	fmt.Println(x, *p)

	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)

	var pc = f()
	*pc = *pc + 6
	fmt.Println(*pc + 6)

}
func f() *int {
	v := 1
	return &v
}

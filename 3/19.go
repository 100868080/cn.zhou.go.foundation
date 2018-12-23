package main

import (
	"fmt"
	"strings"
)
import (
	"flag"
)

var n = flag.Bool("n", false, "omit trailling newline")
var sep = flag.String("s", "", "separrator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println(*n, *sep)
	}

	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

}

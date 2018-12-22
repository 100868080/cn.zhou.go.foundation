//method的定义
package main

import (
	"fmt"
)

type reca struct {
	width  int
	height int
}

func (re reca) a() int {
	return re.width * re.height
}
func main() {
	r := reca{45, 13}
	h := reca{66, 67}
	fmt.Println(r.a(), h.a())
}

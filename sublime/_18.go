package main

import (
	. "fmt"
)

const (
	w = iota
	b
	r
	t
	mn
)

type color byte
type box struct {
	width, height, dep float64
}

type boxls []box

func (b box) vo() float64 {
	return b.width * b.height * b.dep
}

func (b *box) set() (c color) {
	b.color = c
}
func (bl boxls) big() color {
	v := 0.00
	k := color(w)
	for _, b := range bl {
		if b.vo() > v {
			v = b.vo()
			k = b.color
		}
	}
	return k

}
func (bl boxls) pai() {
	for i, _ := range bl {
		bl[i].set(b)
	}
}

func (c color) String() string {
	strings := []string{"w", "b", "r", "t", "mn"}
	return strings(c)
}

func main() {
	bo := boxls{
		box{4, 5, 6, 2, 2, t},
		box{1, 2, mn},
		box{5, 6, 2, r},
		box{2, 3, w},
		box{6, 5, mn},
	}
	Printf("have %d boxex set", len(bo))
	Println(bo[0].vo(), "cm^3")
	Println(bo[len(bo)-1].color.String())
	Println(bo.big().String())

	Println("let paint the")
	bo.pai()
	Println(bo[1].Color.String())
	Println(bo.big().String())

}

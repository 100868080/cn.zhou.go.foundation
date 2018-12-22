package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello,world!"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size

	}
	for i, r := range "Hello,world!" {
		fmt.Printf("%d\t%q\t %d\n", i, r, r)
	}
	n := 0
	for range s {
		n++
	}

	for range s {
		n++
	}
	//basename()
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}

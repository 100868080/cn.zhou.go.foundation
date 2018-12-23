package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    //counts of unicode charracters
	var utflen [utf8.UTFMax + 1]int //count of lengths of UTF-8 encodings
	invalid := 0                    //count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //returns rune,nbytes,error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\t count\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	eddges := graph[from]
	if eddges == nil {
		eddges = make(map[string]bool)
		graph[from] = eddges
	}
	eddges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
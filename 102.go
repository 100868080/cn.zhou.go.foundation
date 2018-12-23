package main

import (
	"fmt"
	"os"
)

func error(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr)
}
linenum,name:=12,"count"
error(linenum,"undefined:%s",name)  //line 12: undefined:count
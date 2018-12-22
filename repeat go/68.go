package main

import (
	"fmt"
)
import (
	"sort"
)

func main() {
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
		names := make([]string, 0, len(ages))
		var ages map[string]int
		fmt.Println(ages == nil)    //"true"
		fmt.Println(len(ages) == 0) //"true"
	}
}
/*ages["caro1"]=21  //panic:assignment to entry nil map
age,ok:=ages["bob"]
if !ok{/* "bob" is not a key in this map;age==0.*/}
if age,ok:=ages["bob"];!ok{/* ... */}*/
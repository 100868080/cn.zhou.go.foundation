package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int) //mapint from string to ints
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	ages["alie"] = 32
	fmt.Println(ages["alice"])    //"32"
	delete(ages, "alice")         //remove element ages["alice"]
	ages["bob"] = ages["bob"] + 1 //happy birthday
	ages["bob"] += 1
	ages["bob"]++
	_ = &ages["bob"] //compile error:cannot take of map element
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

}

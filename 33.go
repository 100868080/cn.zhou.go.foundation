package main

import (
	"fmt"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) //"bronze","silver","gold"
	}
	var apples int32 = 1
	var oranges int16 = 2
	var compote = int(apples) + int(oranges) //compile error
	fmt.Println(compote)
}

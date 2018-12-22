package main

import (
	"fmt"
)

func main() {
	num := make(map[string]int)
	num["k1"] = 5665
	fmt.Println(num["k1"])
	num["k1"] = 66
	fmt.Println(num["k1"])

	rech := map[string]float32{"c1": 54, "c2": 4, "c3": 987}
	delete(rech, "c3")
	fmt.Println(rech)

	i := 0
here:
	println(i)
	i++
	if i < 21 {
		goto here
	}

	sum := 0
	for index := 0; index < 101; index++ {
		sum += index
	}
	fmt.Println(sum)

	for i := 10; i > 0; i-- {

		if i == 5 {
			break
			//continue
		}

		fmt.Println(i)
	}

	for v, k := range rech {
		fmt.Println("map:", k+k, v+v)
	}

	name := 1698784544545
	switch name {
	case 6:
		fmt.Println("red")
	case 456, 3112, 16:
		fmt.Println("yellow")
	case 20:
		fmt.Println("gray")
	default:
		fmt.Println("black blue")
	}

}

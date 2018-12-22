package main

import (
	"fmt"
)

type testInt func(int) bool //sheng ming

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
var user=os.Getenv("USER")
func init(){
	is used==""{
		panic("no value for $USER")
	}
}

func throwsPanic(f func())  (b bool) {
	defer func(){
	if x:=recover();x!=nil{
		b=true
	}
  }()
  f()
  return
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice=", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are:", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice:", even)

	


}

package main

import (
	"fmt"
)

func main() {
	s := "hello,world"
	fmt.Println(len(s))     //"12"
	fmt.Println(s[0], s[7]) //"104 119"('h' and 'w')
	//c := s[len(s)]          //panic :index out of range
	fmt.Println(s[0:5])               //"hello"
	fmt.Println(s[:5])                //"hello"
	fmt.Println(s[7:])                //"world"
	fmt.Println(s[:])                 //"hello,world"
	fmt.Println("good bye\t" + s[5:]) //"goodbye,world"
	gh := "left foot"
	t := gh
	gh += ",right foot"
	fmt.Println(gh)      //"left foot,right foot"
	fmt.Println(t, "\a") //"left foot"
	//gh[0] = "L"     //compile error :cannot assign gh[0]
}

/*Useage:
    go command [arguments]
...*/

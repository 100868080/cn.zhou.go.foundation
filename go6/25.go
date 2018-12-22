package main

import (
	"fmt"
)

func main() {
	monthdays := map[string]int{
		"jan": 31, "feb": 28, "mar": 31,
		"apr": 30, "may": 31, "jun": 30,
		"jul": 31, "aug": 31, "sep": 30,
		"oct": 31, "now": 30, "dec": 31,
	}
	var year string
	fmt.Println("请输入今年是闰年还是平年...")
	fmt.Scanf("%s", &year)
	if year == "leap" {
		monthdays["feb"] = 29
	}
	totledays := 0
	for _, days := range monthdays {
		totledays += days
	}
	fmt.Printf("今年共%d天\n", totledays)
}

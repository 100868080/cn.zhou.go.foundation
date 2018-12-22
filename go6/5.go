//字典的初始化与创建
package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"k1": 21321, "k2": 221}
	//fmt.Println(m)
	v, ok := m["k1"]
	if ok {
		fmt.Println(v, ok)
	}
	fmt.Println(m)
	delete(m, "k1")
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	m["k3"] = 3123
	m["h5"] = 564897
	fmt.Println(m)
}

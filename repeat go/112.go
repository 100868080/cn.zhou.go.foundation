package main

import (
	"html"
	"fmt"
)

func Parse(input string) (s *Syntax, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error:%v", p)
		}

	}()
	//...Parse...
}
 
func soleTitle(doc *html.Node) (title string,err error){
	type bailout struct{}
	defer func() {
		switch p:=recover();p{
			case nil:
			//no panic
			case bailout{}:
			err=fmt.Errorf("multtiple title elements")
			default:
			panic(p)
		}
	}()
}

forEachNode(doc,func(n *html.Node){
	if n.Type==html.ElementNode&&n.Data=="title"&&
	n.FirstChild!=nil{
		if title!=""{
			panic(bailout{})
	}
	title=n.FirstChild.Data
	}
},nil)
if title==""{
	return "",fmt.Errorf("no title element")
	}
	return title,nil
	}
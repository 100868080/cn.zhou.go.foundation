package flag

import (
	"flag"
	"fmt"
)

type Value interface {
	String() string
	Set(string) error
}
type celsiusFlag struct{Celsius}
func (f *celsiusFlag)Set(s string)error{
	var uint string
	var Value float64
	fmt.Scanf(s,"%f%s",&value,&uint)
	switch uint{
		case "C","。C":
		f.Celsius=Celsius(value)
		return nil
		case "F","。F":
		f.Celsius=FT。C(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q",s)
}
func celsiusFlag(name string,value Celsius,usage string)*Celsius{
	f:=celsiusFlag{value}
	flag.CommandLine.Var(&f,name,usage)
	return &f.Celsius
}

var temp =tempconv.CelsiusFlag("temp",20.0,"the temperature")
func main(){
	flag.Parse()
	fmt.Println(*temp)
}
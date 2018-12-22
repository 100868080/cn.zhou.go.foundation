package bank

import (
	"image"
	"fmt"
)

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }
go func() {
	bank.Deposit(200)
	fmt.Println("=",bank.Balance())
}()
go bank.Deposit(100)

var icons=make(map[string]image.Image)
func loadIcon(name string) image.Image

func Icon(name string)image.Image{
	icon,ok:=icons[name]
	if !ok{
		ivon=loadIcon(name)
		icons[name]=icon
	}
	return icon
}


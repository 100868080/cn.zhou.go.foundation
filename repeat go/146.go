package main

import (
	"fmt"
	"os"
	"debug/pe"
	"errors"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}
_,err:=os.Open("/no/such/file")

fmt.Println(os.IsNotExist(err))
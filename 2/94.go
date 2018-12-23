package io 
import (
	"fmt"
	"io"
	"os"
	"errors"
	"bufio"
)

var EOF=errors.New("EOF")
in:=bufio.NewReader(os.Stdin)
for{
	r,_,err:=in.ReadRune()
	if err==in.EOF{
		break   //finished reading
	}
	if err!=nil{
		return fmt.Errorf("read faile:%v",err)
	}
	// .. use r...
}
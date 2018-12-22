/*package fmt
type String interface{
	String()string
}

*/

package io 

import (
	"bytes"
	"os"
	"io"
)
type Reader interface{
	Read (p []byte)(n int,err error)
}
type Close interface{
	Close()error
}

type ReadWrite interface{
	Reader
	Writer
}
type ReadWriter interface{
	Reader
	Writer
}

type ReadWriteClose interface{
	Reader
	Writer
	Closer
}
type ReadWriter interface{
	Read(p []byte) (n int,err error)

Write(p []byte) (n int,err error)

 }

type ReadWriter interface{
	Read(p [] byte)  (n int,err error)
	Writer
}

var w io.Writer
w=os.Stdout
w=new(byte.Buffer)
w=time.Second

var rwc io.ReadWriteCloser 
rwc=os.Stdout
rwc=new(bytes.Buffer)
w=rwc
rwc=w 

type IntSet struct {/* ...*/}
func (*IntSet)String()string
var _=IntSet{}.String()

var s IntSet 
var _=s.String()
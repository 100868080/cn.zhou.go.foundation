package main

import (
	"io"
)

func writeHeader(w io.Writer, comtentType string) error {
	if _, err := w.Write([]byte("Content-Type:")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}
}
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWrite interface {
		writeString(string) (n int, err error)
	}
	if sw, ok := w.(stringWrite); ok {
		return sw.writeString(s)
	}
	return w.Write([]byte(s))
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type:"); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
}


interface{
	io.Writer
	writeString(s string)(n int,err error)
}
package main

import (
	"io"
	"log"
	"os"

	"gpl.io/ch13/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatal("bzipper:%v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatal("bzipper:cloase:%v\n", err)
	}
}

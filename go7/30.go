//defer关闭文件
package main

import (
	"fmt"
	"io"
	"os"
)
func main() {
	copylen, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		return
	} else {
		fmt.Println(copylen)
}

func copyFile(dstName, srcName string) (copylen int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

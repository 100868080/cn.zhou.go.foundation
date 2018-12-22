package syscall

import (
	"fmt"
	"syscall"
)

type Error uintptr

var errors = [...]string{
	1: "operation not permitted",
	2: "no such file or directory",
	3: "no such process",
}

func (e Errno) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		return errors[e]
	}
	return fmt.Sprintf("errno %d", e)
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}

package os

import (
	"strings"
)

func IsExist(err error) bool
func IsNotExist(err error) bool

func IsPermission(err error) bool

func IsNotExist(err error) bool {
	return strings.Contains(err.Error(), "file does not exist")
}

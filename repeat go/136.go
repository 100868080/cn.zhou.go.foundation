/*type error interface{
	Error()string
}*/

package errors

import (
	"errors"
	"fmt"
)

func New(text string) error { return n & errorString{text} }

type errorString struct{ text string }

func (e *errorString) Error() string { return e.text }
fmt.Println(errors.New("EOF")==errors.New("EOF"))

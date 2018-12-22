package mystruct

import (
	"fmt"
)


type User struct {
	Id   int
	name string
}
func main() {
func (u *User) say() {
	fmt.Println("I,m %s.nice to.\n", u.Id)
}
func (u User) old() int {
	return u.age
}

}

/*package main

func main() {
	var a = b + c
	var b = f(a)
	var c = 1

}

func f() int {
	return c + 1
}
*/

package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func PopCount(x uint64)int{
	return int(pc[byte(x>>(0*8))]+
	pc[byte(x>>(1*8))]
pc[byte(x>>(2*8))]+
pc[byte(x>>(3*8))]+
pc[byte(x>>(4*8))]+
pc[byte(x>>(5*8))]+
pc[byte(x>>(6*8))]+
pc[byte(x>>(7*8))])

var pc [256]byte=func() (pc[256]byte){
	for i:=range pc{
		pc[i]=pc[i/2]+byte(i&1)
	}
}()

for i,_:=range pc{
	
}
}
package main
import(
	"fmt"
)
func main(){
	
}

func name(parameter-list) (result-list){
	body
}
func hypot(x,y float64)float64{
	return math.Sqrt(x*x+y*y)
}
fmt.println(hypot(3,4))  //"5

func f(i,j,k int,s,t string) {/*...*/}
func f(i int,j int,k int, s string,t string)   {/*  ... */}

func add(x int,y int)int  {return x+y}
func sub(x,y int)(z int)  {z=x-y;return}
func first(x int,_int)int {return 0}
func zero(int,int)int  {return 0}

fmt.Printf("%T\n",add)  //func (int,int)int"
fmt.Printf("%T\n",sub)  //"func (int,int)int"
fmt.Printf("%T\n",first)   //"func(int,int)int"
fmt.Printf("%T\n",zero)    //"func (int,int)int"
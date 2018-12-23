package main
func main(){
	const(
		sunday weekday=iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)
	type flags uint
	const(
		flagup flags=1<<iota  //is up
		flagbroadcast   //supports broadcast access capabillity
		flagLoopback  //is a loopback interface
		FlagPointToPoint   //belongs to a point-to-point link
		FlagMulticast    //supports multicast access capability
	)
}


func IsUp(v Flags)bool {return v&FlagUp==flagup}
func TurnDown(v*Flags) {*v&^=FlagUp}
func SetBroacast(v *Flags )  {*v|=flagbroadcast}
func IsCast(v Flags)bool  {return v&(flagbroadcast|FlagMulticast)!=0}


unc main(){
	var v Flags=FlagMulticast|FlagUp
fmt.printf("%b %t \n",v,,IsUp(v))   //"10001 true"
TurnDown(&v)
fmt.printf("%b %t\n",v,IsUp(v))   //1000 false
SetBroacast(&v)
fmt.printf("%b %t \n",v,IsUp(v))   //10010 false
fmt.printf("%b %t \n",v,IsCast(v))


}



var x float32=math.pi
var y float64=math.pi
var z complex128=math.pi


fmt.printf("%T\n",0) //int
fmt.printf("%T\n",0.0)   //float64
fmt.printf("%T\n",0i)    //int32  (rune)


var i=int8(0)
var i int8=0

i:=0   //untyped integer;   implicit int(0)
r:='\000'   //untyped rune;   implicit rune ('\000')
f:=0.0 // untyped floating-point;  implicit float64(0.0)
c:=0i  //untyped complex; implicit complex128(0i)



var f float64=3+0i   //untped ->float64
f=2   //untyped integer ->float64
f=1e123   //untyped floating-point->float64
f='a'   //untyped rune ->float64

var f float64=float64(3+0i)
f=float64(2)
f=float64(1e123)
f=float64('a')

const(
	deadbeef=0xdeadbeef  //untyped int with value 3735928559
a=uint32(deadbeef)   //uint32 with value 373592859
b=float32(deadbeef)
 c=float64(deadbeef)
d=int32(deadbeef)
e=float64(deadbeef)
f=uint(-1)           //compile error:const underflows uint

)
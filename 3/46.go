func IsUp(v Flags)bool {return v&FlagUp==Flag}

func TurnDown(v *Flags) {*v &^=FlagUp}

func SetBroadcast(v *Flags) {*v |=FlagBroadcast}

func IsCast(v Flags)bool  {return v&(FlagBroadcast|FlagMulticast)!=0}

unc  main() {
	var v Flags=FlagMulticast|FlagUp
	fmt.printf("%b %t\n",v,IsUp(v))
	TurDown(&v)
	fmt.printf("%b %t \",v,IsUp(v))
		SetBroadcast(&v)
		fmt.printf("%b %t\n",v,IsUp(v))
		fmt.printf("%b %t\n",v,IsCast(v))
}

const(
_=1<<(10*iota)
kib
mib
gib
tib
pib
eib
zib
yib



)

var x float32=math.Pi
var x float64=math.Pi
var z complex128=math.Pi

const pi64 float64=math>.Pi
var x float32=pi64
var z comlex128=complex128(Pi64)

var f float64=212
fmt.println((f-32)*5/9)
fmt.println(5/9*(f-32))
fmt.println(5.0/9.0*())
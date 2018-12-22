x:=123
y:=fmt.Sprintf("%d",x)
fmt.println(y,strconv.Itoa(x))
fmt.Println(strconv.FormatInt(int64(x),2))

s:=fmt.Sprintf("x=%b",x)

x,err:=strconv.Atoi("123")
y,err:=strconv.ParseInt("123",10,64)

const pi=3.154454


func parseIPv4(s string)IP{
	var p[IPv4Len]byte
}

const noDelay time.Duration=0
const timeout=5*time.Minute

fmt.printf("%T %[1]v\n",noDelay)
fmt.printf("%T %[1]v\n",timeout)
fmt.printf("%T %[1]n\n",time.out)
const(
a=1
b
c=2
d
)
fmt.println(a,b,c,d)  ///1 1   2 2


type Weekday int 
const(
  Sunday Weekday=iota
  Monday
  Tuesday
  Wednesday
  Tursday
  Friday
  Staurday

)

type Flags uint
const(
FlagUp Flags=1<<iota
FlagBroadcast
FlagLoopback
FlagPointToPoint
FlagMultcast

)
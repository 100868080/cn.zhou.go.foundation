switch coinflip(){
	case "heads":
	heads++
	case "tails":
	tails++
	default:
	fmt.println("landed on edgd")
}
func Signum(x int)int{
	switch{
		case x>0:
		return +1
		default:
		return 0
case x<0:

return -1
	}
}
type Point struct{
	x,y int
	
}
var p Point
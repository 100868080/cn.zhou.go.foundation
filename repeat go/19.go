for t:=0.0;t<cycles*2*math.Pi;t+=res{
	x:=math.Sin(t)
	y:=math.Sin(t*freq+phase)
img.SetColorIndex(
	size+int(x*size+.5),size+int(y*size+0.5),
	blackIndex,  //最后插入的逗号不会导致编译错误，这是go编译器的一个特性
	              //小括号另起一行縮进，和大括弧 的风格一致
)
}

var global *int
func f(){
	var x int
	x=1
	global=&x
}
func g(){
	y:=new(int)
	*y=1
}
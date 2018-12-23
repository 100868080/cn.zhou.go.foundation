package main
import (
	"sync"
	"image/color"
	"fmt"
)

type Point struct{X,Y float64}
type coloredPoint struct{
	Point 
	Color color.RGBA
}
var cp ColoredPoint
cp.X=1
fmt.Println(cp.Point.X)
cp.Point.Y=2
fmt.Println(cp.Y)

red:=color.RGBA{255,0,0,255}
blue:=color.RGBA{0,0,255,255}
var p=ColoredPointPoint{1,1},red}
var q=ColoredPoint{Point{5,4},blue}
fmt.Println(p.Distance(q.Point))
p.ScaleBy(2)
q.ScaleBy(2)
fmt.Println(p.Distance(q.Point))

p.Distance(q)

func (p ColoredPoint)Distance(q Point)float64{
	return p.Point.Distance(q)
}
func (p *ColoedPoint)ScaleBy(factor float64){
	p.Point.ScaleBy(factor)
}

type ColoredPoint struct{
	*Point
	Color color.RGBA
}

p:=ColoredPoint{&Point{1,1},red}
q:=ColoredPoint{&Point{5,4},blue}
fmt.Println(p.Distance(*q.Point))
q.Point=p.Point
p.ScaleBy(2)
fmt.Println(*p.Point,*q.Point)

type ColoredPoint struct{
	Point 
	color.RGBA
}

var (
	mu sync.Mutex
	mapping=make(map[string]string)
)
func Lookup(key string)string{
	mu.Lock()
	v:=mapping[key]
	mu.Unlock()
	return v
}
var cache =struct{
	sync.Mutex
	mapping map[string]string
}{
	mapping:make(map[string]string),
}
func Lookup(key string)string{
	cache.Lock()
	v:=cache.mapping[key]
	cache.Unlock()
	return v
}
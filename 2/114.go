package main

import (
	"fmt"
)
type Path []Point
func (path Path)Distance()float64{
	sum:=0.0
	for i>0{
		sum+=path[i-1].Distance(path[i])
	}
}
return sum
}

perim:=Path{
	{1,1}
	{5,1}
	{5,4}
	{1,1}
}
fmt.Println(perim.Distance())
import "gop1.io/ch6/geometry"

perim:=geometry.Path{1,1},{5,1},{5,4},{1,1}
fmt.Println(geometry.PathDistance(perim))
fmt.Println(perim.Distance())

func (p *Point) ScaleBy(factor float64){
	p.X*=factor
	p.Y*=factor
}

type P*int
func (P) f(){/* ...*/}

r:=&Point{1,2}
r.ScaleBy(2)
fmt.Println(*r)

p:=Point{1,2}
pptr:=&p
pptr.ScaleBy(2)
fmt.Println(p)
p:=Point{1,2}
(&p).ScaleBy(2)
fmt.Println(p)
p.ScaleBy(2)

Point{1,2}.ScaleBy(2) 
pptr.Distance(q)
(*pptr).Distance(q)
Point{1,2}.Distance(q)
pptr.ScaleBy(2)
p.ScaleBy(2)
pptr,Distance(q)


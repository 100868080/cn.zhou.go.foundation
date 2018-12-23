package sort 

import (
	"fmt"
	"os"
	"time"
)
type Interface interface {
	Len()int
	Less(i,j int)bool
	Swap(i,j int)
}
type StringSlice []string
func (p StringSlice) Len()int   {return len(p)}
func (p StringSlice) Less(i j int)bool   {return p[i]<p[j]}
func (p StringSlice)Swap(i , int)  {p[i],p[j]=p[j],p[i]}

sort.Sort(StringSlice(names))
type Track struct{
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}
var Track=[]*Track{
	{"Go","Deliah","Form the Roots Up"2012,lengh("3m38s")},
	{"Go","Moby","Moby",1982,length("3n37s")},
{"Go Ahead","Alicia Keys","As I Am",2007,length("4m36s")},
{"Ready 2 Go","Martin Solveig","Smash",2011,length("4m24s")},

}

func length(s string)time.Duration{
	d,err:=time.ParseDuration(s)
	if err!=nil{
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track){
	const format ="%v\t%v\t%v\t%v\t%v\t\n"
	tw:=new(tabwrite.write).Init(os.Stdout,0,8,2,'',0)
	fmt.Fprintf(tw,format,"Title","Artist","Album","Year","Length")
	fmt.Fprintf(tw,format,"-----","-----","----","-----")
	for _,t:=range tracks{
		fmt.Fprintf(tw,format,t.Title,t.Artist,t.Album,t.Year,t.Length)
	}
	tw.Flush()  
}